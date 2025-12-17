import { useCallback, useEffect, useMemo, useState } from "react"
import { usePocketBase } from "../context/PocketBaseContext"
import { useNavigate } from "@tanstack/react-router"
import type { FormInput } from "../components/Form"
import { isPositiveNumber } from "../components/validators"
import { useEntries, type ViewEntry } from "./useEntries"

export type ClimateCompassFormData = {
  electricityCO2Emission?: ViewEntry
  electricityConsumptionNonRenewable?: ViewEntry
  electricityConsumptionRenewable?: ViewEntry
  fuelCO2Emission?: ViewEntry
  fuelConsumptionRenewable?: ViewEntry
  fuelConsumptionNonRenewable?: ViewEntry
  heatCO2Emission?: ViewEntry
  heatConsumptionNonRenewable?: ViewEntry
  heatConsumptionRenewable?: ViewEntry
  year: string
}

export type ClimateCompassFormInput = {
  electricityCO2Emission?: number
  electricityConsumptionNonRenewable?: number
  electricityConsumptionRenewable?: number
  fuelCO2Emission?: number
  fuelConsumptionRenewable?: number
  fuelConsumptionNonRenewable?: number
  heatCO2Emission?: number
  heatConsumptionNonRenewable?: number
  heatConsumptionRenewable?: number
  year: string
}

const _emission_type = "CO2 Emission"
const _renewable_type = "Energy Consumption Renewable"
const _nonRenewable_type = "Energy Consumption Non-renewable"
const _electricityDomain = "El"
const _fuelDomain = "Gas"
const _heatDomain = "Varme"

export const useClimateCompassForm = (year?: string) => {
  const { pb } = usePocketBase()
  const navigate = useNavigate()

  const { data: entries, createOrUpdate: createOrUpdateEntries, isLoading: isEntriesLoading } = useEntries({
    filter: `year = ${year}`
  }, year !== undefined)

  const [formData, setFormData] = useState<ClimateCompassFormData>()

  const [isLoading, setIsLoading] = useState(true)

  const submitForm = useCallback(
    async (data: ClimateCompassFormInput) => {
      const auth = pb.authStore.record
      const units = await pb.collection('units').getFullList()
      const unitMWh = units.find(unit => unit.name === 'MWh')?.id
      const unitTon = units.find(unit => unit.name === 'MWh')?.id
      const domains = await pb.collection('domains').getFullList()
      const electricityDomain = domains.find(domain => domain.name === 'El')?.id
      const heatDomain = domains.find(domain => domain.name === 'Varme')?.id
      const fuelDomain = domains.find(domain => domain.name === 'Gas')?.id
      const types = await pb.collection('types').getFullList()
      const renewableType = types.find(type => type.name === 'Energy Consumption Renewable')?.id
      const nonRenewableType = types.find(type => type.name === 'Energy Consumption Non-renewable')?.id
      const co2EmissionType = types.find(type => type.name === 'CO2 Emission')?.id
      const timeAggregations = await pb
        .collection('time_aggregations')
        .getFullList()
      const yearAggregation = timeAggregations.find(agg => agg.name === 'Year')?.id

      if (auth != null) {
        createOrUpdateEntries([
          /* Electricity renewable */
          {
            id: formData?.electricityConsumptionRenewable?.id,
            value: data.electricityConsumptionRenewable,
            organization: auth.organization,
            aggregation: yearAggregation,
            domain: electricityDomain,
            unit: unitMWh,
            year: data.year,
            type: renewableType
          },
          /* Electricity non-renewable */
          {
            id: formData?.electricityConsumptionNonRenewable?.id,
            value: data.electricityConsumptionNonRenewable,
            organization: auth.organization,
            aggregation: yearAggregation,
            domain: electricityDomain,
            unit: unitMWh,
            year: data.year,
            type: nonRenewableType
          },
          /* heat renewable */
          {
            id: formData?.heatConsumptionRenewable?.id,
            value: data.heatConsumptionRenewable,
            organization: auth.organization,
            aggregation: yearAggregation,
            domain: heatDomain,
            unit: unitMWh,
            year: data.year,
            type: renewableType
          },
          /* heat non-renewable */
          {
            id: formData?.heatConsumptionNonRenewable?.id,
            value: data.heatConsumptionNonRenewable,
            organization: auth.organization,
            aggregation: yearAggregation,
            domain: heatDomain,
            unit: unitMWh,
            year: data.year,
            type: nonRenewableType
          },
          /* fuel renewable */
          {
            id: formData?.fuelConsumptionRenewable?.id,
            value: data.fuelConsumptionRenewable,
            organization: auth.organization,
            aggregation: yearAggregation,
            domain: fuelDomain,
            unit: unitMWh,
            year: data.year,
            type: renewableType
          },
          /* fuel non-renewable */
          {
            id: formData?.fuelConsumptionNonRenewable?.id,
            value: data.fuelConsumptionNonRenewable,
            organization: auth.organization,
            aggregation: yearAggregation,
            domain: fuelDomain,
            unit: unitMWh,
            year: data.year,
            type: nonRenewableType
          },
          /* electricity co2Emission */
          {
            id: formData?.electricityCO2Emission?.id,
            value: data.electricityCO2Emission,
            organization: auth.organization,
            aggregation: yearAggregation,
            domain: electricityDomain,
            unit: unitTon,
            year: data.year,
            type: co2EmissionType
          },
          /* fuel co2Emission */
          {
            id: formData?.fuelCO2Emission?.id,
            value: data.fuelCO2Emission,
            organization: auth.organization,
            aggregation: yearAggregation,
            domain: fuelDomain,
            unit: unitTon,
            year: data.year,
            type: co2EmissionType
          },
          /* heat co2Emission */
          {
            id: formData?.heatCO2Emission?.id,
            value: data.heatCO2Emission,
            organization: auth.organization,
            aggregation: yearAggregation,
            domain: heatDomain,
            unit: unitTon,
            year: data.year,
            type: co2EmissionType
          }
        ])

        if (year === undefined) {
          navigate({ to: `/dataEntry/$year`, params: { year: data.year } })
        }
      }
    },
    [pb, formData, year, navigate, createOrUpdateEntries]
  )

  useEffect(() => {
    if (year !== undefined && entries != undefined) {
      const currentFormData: ClimateCompassFormData = {
        year: year,
        electricityCO2Emission: entries.find(entry => entry.typeName === _emission_type && entry.domainName === _electricityDomain),
        fuelCO2Emission: entries.find(entry => entry.typeName === _emission_type && entry.domainName === _fuelDomain),
        heatCO2Emission: entries.find(entry => entry.typeName === _emission_type && entry.domainName === _heatDomain),
        electricityConsumptionRenewable: entries.find(entry => entry.typeName === _renewable_type && entry.domainName === _electricityDomain),
        electricityConsumptionNonRenewable: entries.find(entry => entry.typeName === _nonRenewable_type && entry.domainName === _electricityDomain),
        fuelConsumptionRenewable: entries.find(entry => entry.typeName === _renewable_type && entry.domainName === _fuelDomain),
        fuelConsumptionNonRenewable: entries.find(entry => entry.typeName === _nonRenewable_type && entry.domainName === _fuelDomain),
        heatConsumptionRenewable: entries.find(entry => entry.typeName === _renewable_type && entry.domainName === _heatDomain),
        heatConsumptionNonRenewable: entries.find(entry => entry.typeName === _nonRenewable_type && entry.domainName === _heatDomain)
      }
      setFormData(currentFormData)
      setIsLoading(false)
    } else {
      setIsLoading(false)
    }
  }, [year, entries])


  const formInput: FormInput = useMemo(() => {
    return {
      inputs: [
        {
          name: 'year',
          type: 'text',
          label: 'År',
          placeholder: 'Indtast år',
          value: year,
          defaultValue: ''
        },
        {
          name: 'electricityConsumptionRenewable',
          type: 'text',
          label: 'Elforbrug (vedvarende) - MWh',
          placeholder: 'Indtast forbrug',
          value: formData?.electricityConsumptionRenewable
            ? `${formData.electricityConsumptionRenewable.value}`
            : undefined,
          defaultValue: "",
          registerOptions: {
            validate: {
              isPositiveNumber,
            },
            required: 'Forbrug er påkrævet',
          },
          // TODO: Introduce this for all values?
          //accompanyingSelect: {
          //  name: 'electricityUnit',
          //  label: 'Enhed',
          //  options: [
          //    { id: 'kwh', name: 'kWh' },
          //    { id: 'mwh', name: 'MWh' },
          //  ],
          //  placeholder: 'Vælg',
          //  registerOptions: {
          //    required: 'Påkrævet',
          //  },
          //  defaultValue: year ? 'kwh' : undefined,
          //},
        },
        {
          name: 'electricityConsumptionNonRenewable',
          type: 'text',
          label: 'Elforbrug (ikke-vedvarende)',
          placeholder: 'Indtast forbrug',
          value: formData?.electricityConsumptionNonRenewable
            ? String(formData.electricityConsumptionNonRenewable.value)
            : undefined,
          defaultValue: '',
          registerOptions: {
            validate: {
              isPositiveNumber,
            },
            required: 'Forbrug er påkrævet',
          },
        },
        {
          name: 'heatConsumptionRenewable',
          type: 'text',
          label: 'Fjenvarmeforbrug (vedvarende)',
          placeholder: 'Indtast forbrug',
          value: formData?.heatConsumptionRenewable
            ? String(formData.heatConsumptionRenewable.value)
            : undefined,
          defaultValue: '',
          registerOptions: {
            validate: {
              isPositiveNumber,
            },
            required: 'Forbrug er påkrævet',
          },
        },
        {
          name: 'heatConsumptionNonRenewable',
          type: 'text',
          label: 'Fjernvarmeforbrug (ikke-vedvarende)',
          placeholder: 'Indtast forbrug',
          value: formData?.heatConsumptionNonRenewable
            ? String(formData.heatConsumptionNonRenewable.value)
            : undefined,
          defaultValue: '',
          registerOptions: {
            validate: {
              isPositiveNumber,
            },
            required: 'Forbrug er påkrævet',
          },
        },
        {
          name: 'fuelConsumptionRenewable',
          type: 'text',
          label: 'Brændselsforbrug (vedvarende)',
          placeholder: 'Indtast brændselsforbrug',
          value: formData?.fuelConsumptionRenewable
            ? String(formData.fuelConsumptionRenewable.value)
            : undefined,
          defaultValue: '',
          registerOptions: {
            validate: {
              isPositiveNumber,
            },
            required: 'Brændselsforbrug er påkrævet',
          },
        },
        {
          name: 'fuelConsumptionNonRenewable',
          type: 'text',
          label: 'Brændselsforbrug (ikke-vedvarende)',
          placeholder: 'Indtast brændselsforbrug',
          value: formData?.fuelConsumptionNonRenewable
            ? String(formData.fuelConsumptionNonRenewable.value)
            : undefined,
          defaultValue: '',
          registerOptions: {
            validate: {
              isPositiveNumber,
            },
            required: 'Brændselsforbrug er påkrævet',
          },
        },
        {
          name: 'electricityCO2Emission',
          type: 'text',
          label: 'Eludledning',
          placeholder: 'Indtast udledning',
          value: formData?.electricityCO2Emission
            ? String(formData.electricityCO2Emission.value)
            : undefined,
          defaultValue: '',
          registerOptions: {
            validate: {
              isPositiveNumber,
            },
            required: 'Udledning er påkrævet',
          },
        },
        {
          name: 'heatCO2Emission',
          type: 'text',
          label: 'Fjernvarmeudledning',
          placeholder: 'Indtast udledning',
          value: formData?.heatCO2Emission
            ? String(formData.heatCO2Emission.value)
            : undefined,
          defaultValue: '',
          registerOptions: {
            validate: {
              isPositiveNumber,
            },
            required: 'Udledning er påkrævet',
          },
        },
        {
          name: 'fuelCO2Emission',
          type: 'text',
          label: 'Brændselsudledning',
          placeholder: 'Indtast udledning',
          value: formData?.fuelCO2Emission
            ? String(formData.fuelCO2Emission.value)
            : undefined,
          defaultValue: '',
          registerOptions: {
            validate: {
              isPositiveNumber,
            },
            required: 'Udledning er påkrævet',
          },
        },
      ],
    }
  }, [formData, year])

  return {
    data: formData,
    formInput: formInput,
    isLoading: isLoading || isEntriesLoading,
    submit: submitForm
  }
}
