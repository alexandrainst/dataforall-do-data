import { PageTemplate } from './layout/PageTemplate'
import { BarChart } from './charts/BarChart'
import { useForm } from 'react-hook-form'
import { MaskChart } from './charts/MaskChart'
import { Insights } from './charts/Insights'
import {
  useEntriesDistribution,
  type EntriesDistribution,
} from '../hooks/useEntries'
import { useMemo } from 'react'
import {
  useComparisonCategories,
  useComparisonEntriesDistribution,
} from '../hooks/useComparisonCategories'
import { useInsights } from '../hooks/useInsights'
import { useEntryGroups } from '../hooks/useEntryGroups'

export type DashboardFormValues = {
  myDistributionData: string
  distributionComparisonData: string
  distributionComparisonComparison: string
  emissionComparisonData: string
}

type DataOptions = 'Consumption' | 'CO2 Emission'

const getMyConsumptionDistribution = (
  type: DataOptions,
  entriesDistribution: EntriesDistribution
) => {
  if (!entriesDistribution) return undefined
  return Object.entries(entriesDistribution).map(([year, distribution]) => {
    let electricity = 0,
      heat = 0,
      fuel = 0,
      bar1unit = ''
    switch (type) {
      case 'CO2 Emission': {
        electricity = distribution['CO2 Emission']?.['El'] ?? 0
        heat = distribution['CO2 Emission']?.['Varme'] ?? 0
        fuel = distribution['CO2 Emission']?.['Gas'] ?? 0
        bar1unit = 'ton'
        break
      }
      case 'Consumption': {
        electricity =
          (distribution['Energy Consumption Non-renewable']?.['El'] ?? 0) +
          (distribution['Energy Consumption Renewable']?.['El'] ?? 0)
        heat =
          (distribution['Energy Consumption Non-renewable']?.['Varme'] ?? 0) +
          (distribution['Energy Consumption Renewable']?.['Varme'] ?? 0)
        fuel =
          (distribution['Energy Consumption Non-renewable']?.['Gas'] ?? 0) +
          (distribution['Energy Consumption Renewable']?.['Gas'] ?? 0)
        bar1unit = 'MWh'
        break
      }
    }
    // const total = electricity + heat + fuel
    return {
      year: year,
      bar1unit,
      // bar1percentage1: total ? Math.round((electricity / total) * 100) : 0,
      bar1value1: electricity,
      bar1name1: 'Elektricitet',
      // bar1percentage2: total ? Math.round((heat / total) * 100) : 0,
      bar1value2: heat,
      bar1name2: 'Fjernvarme',
      // bar1percentage3: total ? Math.round((fuel / total) * 100) : 0,
      bar1value3: fuel,
      bar1name3: 'Brændsler',
    }
  })
}

export const Dashboard = () => {
  const { control, watch, setValue } = useForm<DashboardFormValues>({
    defaultValues: {
      myDistributionData: 'Consumption',
      distributionComparisonData: 'Consumption',
      distributionComparisonComparison: '1',
      emissionComparisonData: '1',
    },
  })

  const { data: entriesDistribution } = useEntriesDistribution()

  const [
    myDistributionData,
    distributionComparisonData,
    distributionComparisonComparison,
    emissionComparisonData,
  ] = watch([
    'myDistributionData',
    'distributionComparisonData',
    'distributionComparisonComparison',
    'emissionComparisonData',
  ])

  const { data: insights } = useInsights()

  const { data: entryGroups } = useEntryGroups()

  const myInsights = useMemo(() => {
    if (!insights || !entriesDistribution) return []
    return insights.map(insight => ({
      title: insight.title,
      body: insight.text,
      // TODO: change to relevant chart based on insight.typeName and possibly more?
      chart: (
        <BarChart
          name="myInsightDistribution"
          title="Min fordeling af elektricitet, fjernvarme og brændsler"
          data={
            getMyConsumptionDistribution('CO2 Emission', entriesDistribution) ??
            []
          }
        />
      ),
    }))
  }, [insights, entriesDistribution])

  const { data: comparisonCategories } = useComparisonCategories()

  const { data: comparisonEntriesDistribution } =
    useComparisonEntriesDistribution(distributionComparisonComparison)

  const distributionComparisonOptions = useMemo(() => {
    if (comparisonCategories === undefined) return undefined
    const comparisonCategoryOptions = comparisonCategories.map(category => ({
      id: category.id,
      name: category.name,
    }))
    setValue(
      'distributionComparisonComparison',
      comparisonCategories[0]?.id ?? ''
    )
    return comparisonCategoryOptions
  }, [comparisonCategories, setValue])

  const distributionComparison = useMemo(() => {
    if (!entriesDistribution) return []
    return Object.entries(entriesDistribution).map(([year, distribution]) => {
      let electricity = 0,
        heat = 0,
        fuel = 0,
        electricityComparison = 0,
        heatComparison = 0,
        fuelComparison = 0,
        bar1unit = '',
        bar2unit = ''
      switch (distributionComparisonData) {
        case 'CO2 Emission': {
          electricity = distribution['CO2 Emission']?.['El'] ?? 0
          heat = distribution['CO2 Emission']?.['Varme'] ?? 0
          fuel = distribution['CO2 Emission']?.['Gas'] ?? 0
          electricityComparison =
            comparisonEntriesDistribution?.[Number(year)]?.['CO2 Emission']?.[
              'El'
            ] ?? 0
          heatComparison =
            comparisonEntriesDistribution?.[Number(year)]?.['CO2 Emission']?.[
              'Varme'
            ] ?? 0
          fuelComparison =
            comparisonEntriesDistribution?.[Number(year)]?.['CO2 Emission']?.[
              'Gas'
            ] ?? 0
          bar1unit = 'ton'
          bar2unit = 'ton'
          break
        }
        case 'Consumption': {
          electricity =
            (distribution['Energy Consumption Non-renewable']?.['El'] ?? 0) +
            (distribution['Energy Consumption Renewable']?.['El'] ?? 0)
          heat =
            (distribution['Energy Consumption Non-renewable']?.['Varme'] ?? 0) +
            (distribution['Energy Consumption Renewable']?.['Varme'] ?? 0)
          fuel =
            (distribution['Energy Consumption Non-renewable']?.['Gas'] ?? 0) +
            (distribution['Energy Consumption Renewable']?.['Gas'] ?? 0)
          electricityComparison =
            (comparisonEntriesDistribution?.[Number(year)]?.[
              'Energy Consumption Non-renewable'
            ]?.['El'] ?? 0) +
            (comparisonEntriesDistribution?.[Number(year)]?.[
              'Energy Consumption Renewable'
            ]?.['El'] ?? 0)
          heatComparison =
            (comparisonEntriesDistribution?.[Number(year)]?.[
              'Energy Consumption Non-renewable'
            ]?.['Varme'] ?? 0) +
            (comparisonEntriesDistribution?.[Number(year)]?.[
              'Energy Consumption Renewable'
            ]?.['Varme'] ?? 0)
          fuelComparison =
            (comparisonEntriesDistribution?.[Number(year)]?.[
              'Energy Consumption Non-renewable'
            ]?.['Gas'] ?? 0) +
            (comparisonEntriesDistribution?.[Number(year)]?.[
              'Energy Consumption Renewable'
            ]?.['Gas'] ?? 0)
          bar1unit = 'MWh'
          bar2unit = 'MWh'
          break
        }
      }
      const total = electricity + heat + fuel
      const totalComparison =
        electricityComparison + heatComparison + fuelComparison
      return {
        year: year,
        bar1unit: bar1unit,
        bar2unit: bar2unit,
        bar1percentage1: total ? Math.round((electricity / total) * 100) : 0,
        bar1value1: electricity,
        bar1name1: 'Min: elektricitet',
        bar1percentage2: total ? Math.round((heat / total) * 100) : 0,
        bar1value2: heat,
        bar1name2: 'Min: fjernvarme',
        bar1percentage3: total ? Math.round((fuel / total) * 100) : 0,
        bar1value3: fuel,
        bar1name3: 'Min: brændsler',
        bar2percentage1: totalComparison
          ? Math.round((electricityComparison / totalComparison) * 100)
          : 0,
        bar2value1: electricityComparison,
        bar2name1: 'Sammenligning: elektricitet',
        bar2percentage2: totalComparison
          ? Math.round((heatComparison / totalComparison) * 100)
          : 0,
        bar2value2: heatComparison,
        bar2name2: 'Sammenligning: fjernvarme',
        bar2percentage3: totalComparison
          ? Math.round((fuelComparison / totalComparison) * 100)
          : 0,
        bar2value3: fuelComparison,
        bar2name3: 'Sammenligning: brændsler',
      }
    })
  }, [
    entriesDistribution,
    distributionComparisonData,
    comparisonEntriesDistribution,
  ])

  const distributionComparisonDataOptions = [
    { id: 'Consumption', name: 'Forbrug (MWh)' },
    { id: 'CO2 Emission', name: 'CO2-udledning (ton)' },
  ]

  const myDistributionComparison = useMemo(() => {
    if (!entriesDistribution) return undefined
    return Object.entries(entriesDistribution).map(([year, distribution]) => {
      const electricityEmission = distribution['CO2 Emission']?.['El'] ?? 0
      const heatEmission = distribution['CO2 Emission']?.['Varme'] ?? 0
      const fuelEmission = distribution['CO2 Emission']?.['Gas'] ?? 0
      const bar2unit = 'ton'
      const totalEmission = electricityEmission + heatEmission + fuelEmission
      const electricityConsumption =
        (distribution['Energy Consumption Non-renewable']?.['El'] ?? 0) +
        (distribution['Energy Consumption Renewable']?.['El'] ?? 0)
      const heatConsumption =
        (distribution['Energy Consumption Non-renewable']?.['Varme'] ?? 0) +
        (distribution['Energy Consumption Renewable']?.['Varme'] ?? 0)
      const fuelConsumption =
        (distribution['Energy Consumption Non-renewable']?.['Gas'] ?? 0) +
        (distribution['Energy Consumption Renewable']?.['Gas'] ?? 0)
      const bar1unit = 'MWh'
      const totalConsumption =
        electricityConsumption + heatConsumption + fuelConsumption
      return {
        year: year,
        bar1unit: bar1unit,
        bar2unit: bar2unit,
        bar1percentage1: totalConsumption
          ? Math.round((electricityConsumption / totalConsumption) * 100)
          : 0,
        bar1value1: electricityConsumption,
        bar1name1: 'Forbrug elektricitet',
        bar1percentage2: totalConsumption
          ? Math.round((heatConsumption / totalConsumption) * 100)
          : 0,
        bar1value2: heatConsumption,
        bar1name2: 'Forbrug fjernvarme',
        bar1percentage3: totalConsumption
          ? Math.round((fuelConsumption / totalConsumption) * 100)
          : 0,
        bar1value3: fuelConsumption,
        bar1name3: 'Forbrug brændsler',
        bar2percentage1: totalEmission
          ? Math.round((electricityEmission / totalEmission) * 100)
          : 0,
        bar2value1: electricityEmission,
        bar2name1: 'CO2 elektricitet',
        bar2percentage2: totalEmission
          ? Math.round((heatEmission / totalEmission) * 100)
          : 0,
        bar2value2: heatEmission,
        bar2name2: 'CO2 fjernvarme',
        bar2percentage3: totalEmission
          ? Math.round((fuelEmission / totalEmission) * 100)
          : 0,
        bar2value3: fuelEmission,
        bar2name3: 'CO2 brændsler',
      }
    })
  }, [entriesDistribution])

  // CHART 3
  const emissionComparison = [
    {
      year: '2021',
      bar1unit: 'kg pr. ansat',
      bar2unit: 'kg pr. ansat',
      bar1value1: 2000,
      bar1name1: 'Min virksomhed',
      bar2value1: 2450,
      bar2name1: 'Sammenligning',
    },
    {
      year: '2022',
      bar1unit: 'kg pr. ansat',
      bar2unit: 'kg pr. ansat',
      bar1value1: 2400,
      bar1name1: 'Min virksomhed',
      bar2value1: 3750,
      bar2name1: 'Sammenligning',
    },
  ]

  const emissionComparisonDataOptions = [
    { id: '1', name: 'Pr. ansat (kg)' },
    { id: '2', name: 'Pr. mio. omsætning (kg)' },
    { id: '3', name: 'Pr. m2 (kg)' },
  ]

  const myConsumptionDistribution = useMemo(() => {
    if (!entriesDistribution) return undefined
    return getMyConsumptionDistribution(
      myDistributionData as DataOptions,
      entriesDistribution
    )
  }, [entriesDistribution, myDistributionData])

  const myDistributionDataOptions = [
    { id: 'CO2 Emission', name: 'CO2 (ton)' },
    { id: 'Consumption', name: 'Forbrug (MWh)' },
  ]

  return (
    <PageTemplate title="Dashboard">
      {entryGroups !== undefined && entryGroups.length > 0 ?
      <div className="flex flex-wrap gap-4">
        {myInsights.length > 0 && <Insights insights={myInsights} />}
        <div className="w-full rounded-lg bg-white p-6 shadow-sm lg:w-[calc(50%-.5rem)]">
          <BarChart
            name="myDistribution"
            title="Min fordeling af elektricitet, fjernvarme og brændsler"
            data={myConsumptionDistribution ?? []}
            control={control}
            dataOptions={myDistributionDataOptions}
          />
        </div>
        <div className="w-full rounded-lg bg-white p-6 shadow-sm lg:w-[calc(50%-.5rem)]">
          <BarChart
            name="distributionComparison"
            title="Fordeling af elektricitet, fjernvarme og brændsler"
            data={distributionComparison}
            control={control}
            dataOptions={distributionComparisonDataOptions}
            comparisonOptions={distributionComparisonOptions}
          />
        </div>
        <div className="w-full rounded-lg bg-white p-6 shadow-sm lg:w-[calc(50%-.5rem)]">
          <BarChart
            name="myDistributionComparison"
            title="Mit forbrug sammenlignet med min estimerede CO2-udledning"
            data={myDistributionComparison ?? []}
          />
        </div>
         <div className="w-full rounded-lg bg-white p-6 shadow-sm lg:w-[calc(50%-.5rem)]">
          <div className="relative">
            <div className="absolute right-0 top-0 rounded-bl-lg rounded-tr-lg bg-orange-400 px-3 py-1 text-xs font-semibold text-white">
              Testdata
            </div>
            <BarChart
              name="emissionComparison"
              title="Min estimerede CO2-udledning"
              data={emissionComparison}
              control={control}
              dataOptions={emissionComparisonDataOptions}
            />
          </div>
        </div>
        <div className="w-full rounded-lg bg-white p-6 shadow-sm lg:w-[calc(50%-.5rem)]">
          <div className="relative">
            <div className="absolute right-0 top-0 rounded-bl-lg rounded-tr-lg bg-orange-400 px-3 py-1 text-xs font-semibold text-white">
              Testdata
            </div>
            <MaskChart
              title="Vedvarende energi"
              data={[
                {
                  year: '2020',
                  percentage: 45,
                },
                {
                  year: '2021',
                  percentage: 45,
                },
                {
                  year: '2022',
                  percentage: 65,
                },
                {
                  year: '2023',
                  percentage: 32,
                },
              ]}
            />
          </div>
        </div>
      </div> : <div className=''>Ingen data tilgængelig</div>}
    </PageTemplate>
  )
}
