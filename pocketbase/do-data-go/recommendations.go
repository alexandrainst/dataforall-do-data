package main

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

type SituationSelf int

const (
	TrendingUp SituationSelf = iota
	TrendingSteady
	TrendingDown
)

var situationSelfName = map[SituationSelf]string{
	TrendingUp:     "trending_up",
	TrendingSteady: "trend_steady",
	TrendingDown:   "trending_down",
}

func (ss SituationSelf) String() string {
	return situationSelfName[ss]
}

type SituationOther int

const (
	Above SituationOther = iota
	Average
	Below
)

var situationOtherName = map[SituationOther]string{
	Above:   "above",
	Average: "average",
	Below:   "below",
}

func (so SituationOther) String() string {
	return situationOtherName[so]
}

type SituationEvaluator func(a, b *core.Record) bool

func isSelfSituation(situation string) bool {
	return slices.Contains([]string{TrendingUp.String(), TrendingSteady.String(), TrendingDown.String()}, situation)
}

func isOtherSituation(situation string) bool {
	return slices.Contains([]string{Above.String(), Average.String(), Below.String()}, situation)
}

func isTrendUp(threshold float64) SituationEvaluator {
	return func(a, b *core.Record) bool {
		return (a.GetFloat("value") / b.GetFloat("value")) > 1+threshold/100
	}
}

func isTrendDown(threshold float64) SituationEvaluator {
	return func(a, b *core.Record) bool {
		return (a.GetFloat("value") / b.GetFloat("value")) < 1-threshold/100
	}
}

func isTrendSteady(threshold float64) SituationEvaluator {
	return func(a, b *core.Record) bool {
		trendQuotient := a.GetFloat("value") / b.GetFloat("value")
		return 1-threshold/100 < trendQuotient && trendQuotient < 1+threshold/100
	}
}

func deleteRecommendations(app *pocketbase.PocketBase, year int, organizationId string) error {
	hashExp := dbx.HashExp{
		"year":         year,
		"organization": organizationId,
	}
	_, err := app.DB().Delete("recommendations", hashExp).Execute()

	return err
}

func deleteRecommendationStatistics(app *pocketbase.PocketBase, organizationId string) error {
	hashExp := dbx.HashExp{
		"organization": organizationId,
	}
	_, err := app.DB().Delete("recommendation_statistics", hashExp).Execute()

	return err
}

func getOrganizationIdsByCategory(app *pocketbase.PocketBase, organizationCategoryId string) ([]string, error) {
	organizations, err := app.FindAllRecords("organizations",
		dbx.HashExp{"category": organizationCategoryId},
	)
	app.Logger().Debug(fmt.Sprintf("%v", organizations))
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch organizations: %s", err.Error())
	}
	var organizationIds []string
	for _, organization := range organizations {
		organizationIds = append(organizationIds, organization.GetString("id"))
	}
	app.Logger().Debug(fmt.Sprintf("Org ids: %v", organizationIds))
	return organizationIds, nil
}

func sortRecordsByYear(a, b *core.Record) int {
	return cmp.Compare(b.GetDateTime("timestamp").Time().Year(), a.GetDateTime("timestamp").Time().Year())
}

func processSelfRule(app *pocketbase.PocketBase, entries []*core.Record, ruleId, domainId, typeId, situationId, organizationId string, checkSituation SituationEvaluator) error {
	slices.SortFunc(entries, sortRecordsByYear)
	for i := 1; i < len(entries); i++ {
		currentYearValue := entries[i-1].GetFloat("value")
		previousYearValue := entries[i].GetFloat("value")
		isSituaionTriggered := checkSituation(entries[i-1], entries[i])
		app.Logger().Debug(fmt.Sprintf("years: %d-%d, values: %f-%f, rule: %t",
			entries[i-1].GetDateTime("timestamp").Time().Year(),
			entries[i].GetDateTime("timestamp").Time().Year(),
			currentYearValue,
			previousYearValue,
			isSituaionTriggered,
		))
		if isSituaionTriggered {
			recommendation_statistics, err := app.FindCollectionByNameOrId("recommendation_statistics")
			if err != nil {
				return err
			}
			year := entries[i-1].GetDateTime("timestamp").Time().Year()
			recommendation_statistic := core.NewRecord(recommendation_statistics)
			recommendation_statistic.Set("domain", domainId)
			recommendation_statistic.Set("type", typeId)
			recommendation_statistic.Set("situation", situationId)
			recommendation_statistic.Set("rule", ruleId)
			recommendation_statistic.Set("organization", organizationId)
			recommendation_statistic.Set("year", year)
			app.Save(recommendation_statistic)
		}
	}
	return nil
}

func getRelationIdOrAll(app *pocketbase.PocketBase, record *core.Record, relation string, collection string) ([]string, error) {
	relationId := record.GetString(relation)
	if len(relationId) > 0 {
		return []string{relationId}, nil
	} else {
		relations, err := app.FindAllRecords(collection)
		if err != nil {
			return nil, fmt.Errorf("Failed to all organizations. %s ", err.Error())
		}
		var relationIds []string
		for _, relation := range relations {
			relationIds = append(relationIds, relation.GetString("id"))
		}
		return relationIds, nil
	}
}

func ProcessRules(app *pocketbase.PocketBase) error {
	rules, err := app.FindAllRecords("rules", nil)
	if err != nil {
		return fmt.Errorf("Failed to fetch rules")
	}
	app.Logger().Debug(fmt.Sprintf("%d", len(rules)))

	// Process each rule
	for _, rule := range rules {
		app.ExpandRecord(rule, []string{"situation"}, nil)
		situation := rule.ExpandedOne("situation")
		if situation == nil {
			return fmt.Errorf("Situation is nil")
		}
		threshold := rule.GetFloat("thresholdPercent")

		if isSelfSituation(situation.GetString("name")) {
			// Get ids for the organizations for which to process the rule.
			// If the organization category is not set, process all organizations.
			organizationCategory := rule.GetString("organization_category")
			app.Logger().Debug(fmt.Sprintf("Org cat: %s", organizationCategory))
			var organizationIds []string
			if len(organizationCategory) > 0 {
				organizationIds, err = getOrganizationIdsByCategory(app, rule.GetString("organization_category"))
				if err != nil {
					return fmt.Errorf("Failed to organization ids: %s", err.Error())
				}
			} else {
				organizations, err := app.FindAllRecords("organizations")
				if err != nil {
					return fmt.Errorf("Failed to all organizations. %s ", err.Error())
				}
				for _, organization := range organizations {
					organizationIds = append(organizationIds, organization.GetString("id"))
				}
			}

			domainIds, err := getRelationIdOrAll(app, rule, "domain", "domains")
			if err != nil {
				return fmt.Errorf("Failed to fetch domains: %s", err.Error())
			}

			typeIds, err := getRelationIdOrAll(app, rule, "type", "types")
			if err != nil {
				return fmt.Errorf("Failed to fetch domains: %s", err.Error())
			}

			app.Logger().Debug(fmt.Sprintf("Processing rule: %s", situation.GetString("name")))
			for _, organizationId := range organizationIds {
				for _, domainId := range domainIds {
					for _, typeId := range typeIds {
						// Filter entries
						entries, err := app.FindAllRecords("entries",
							dbx.HashExp{"domain": domainId},
							dbx.HashExp{"type": typeId},
							dbx.HashExp{"organization": organizationId},
						)
						if err != nil {
							return fmt.Errorf("Failed to fetch entries: %s", err.Error())
						}
						app.Logger().Debug(fmt.Sprintf("Entries: %v", entries))

						var situationEvaluator SituationEvaluator
						switch situation.GetString("name") {
						case TrendingUp.String():
							situationEvaluator = isTrendUp(threshold)
						case TrendingSteady.String():
							situationEvaluator = isTrendSteady(threshold)
						case TrendingDown.String():
							situationEvaluator = isTrendDown(threshold)

						default:
							return fmt.Errorf("Error processing rule: %s, no handler exists", situation.GetString("name"))
						}
						processSelfRule(app, entries, rule.GetString("id"), domainId, typeId, situation.GetString("id"), organizationId, situationEvaluator)
					}
				}
			}
		} else if isOtherSituation(situation.GetString("name")) {
			// Process rules comparing to other companies
		} else {
			return fmt.Errorf("Unknown situation type")
		}
	}
	return nil
}
