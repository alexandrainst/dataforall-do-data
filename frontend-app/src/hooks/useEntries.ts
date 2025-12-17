import { usePocketBase } from '../context/PocketBaseContext'
import type { BatchService, RecordListOptions } from 'pocketbase'
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { chain } from 'lodash'

export type EntriesByYear = {
  year: string
  entries: ViewEntry[]
}

export type ViewEntry = {
  id: string
  value: number
  year: number
  unitName: string
  typeName: string
  domainName: string
  organiztionName: string
  timeAggregation: string
}

export type Entry = {
  id?: string
  value: number | undefined
  organization?: string
  aggregation?: string
  domain?: string
  unit?: string
  year?: string
  type?: string
}

export type EntryType =
  | 'Energy Consumption Non-renewable'
  | 'Energy Consumption Renewable'
  | 'CO2 Emission'

export type Domain = 'El' | 'Varme' | 'Gas'

export type EntriesDistribution = {
  [key in number]: { [key in EntryType]?: { [key in Domain]?: number } }
}

const createOrUpdateEntry = async (batch: BatchService, entry: Entry) => {
  if (entry.id !== undefined) {
    batch.collection('entries').update(entry.id, {
      value: entry.value,
    })
  } else {
    batch.collection('entries').create({
      organization: entry.organization,
      value: entry.value,
      aggregation: entry.aggregation,
      domain: entry.domain,
      unit: entry.unit,
      timestamp: `${entry.year}-01-01`,
      type: entry.type,
    })
  }
}

export const useEntries = (
  options?: RecordListOptions,
  enabled: boolean = false
) => {
  const { pb } = usePocketBase()

  const queryClient = useQueryClient()

  const {
    data: entries,
    isLoading,
    isError,
  } = useQuery<ViewEntry[], Error>({
    queryKey: ['entries'],
    queryFn: () =>
      pb.collection<ViewEntry>('view_entries').getFullList(options),
    enabled: enabled,
  })

  const { mutateAsync: mutateAsyncCreate } = useMutation<void, Error, Entry[]>({
    mutationFn: async entries => {
      const batch = pb.createBatch()
      entries.forEach(entry => createOrUpdateEntry(batch, entry))
      await batch.send()
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['entries'] })
    },
  })

  const {
    mutateAsync: mutateAsyncDelete,
    isPending: isPendingDelete,
    isError: isErrorDelete,
  } = useMutation<void, Error, { year: string }>({
    mutationFn: async ({ year }) => {
      const response = await fetch(`/api/dataentries/${year}`, {
        method: 'DELETE',
        headers: {
          Authorization: `Bearer ${pb.authStore.token}`,
        },
      })

      if (!response.ok) {
        const errorData = await response.json()
        throw new Error(errorData.message || 'Failed to delete data entry')
      }
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['entries'] })
    },
    onError: e => console.error('Error deleting entry', e),
  })

  return {
    data: entries,
    createOrUpdate: mutateAsyncCreate,
    delete: mutateAsyncDelete,
    isErrorDelete,
    isPendingDelete,
    isError: isError,
    isLoading: isLoading,
  }
}

export const useEntriesByYear = () => {
  const { pb } = usePocketBase()
  return useQuery<EntriesByYear[], Error>({
    queryKey: ['entriesByYear'],
    queryFn: async () => {
      const entries = await pb
        .collection<ViewEntry>('view_entries')
        .getFullList()
      const entriesByYear = chain(entries)
        .groupBy('year')
        .map((value, key) => ({ year: key, entries: value }))
        .value()
      return entriesByYear
    },
  })
}

export const useEntriesDistribution = () => {
  const { pb } = usePocketBase()
  return useQuery<EntriesDistribution, Error>({
    queryKey: ['entriesByYear'],
    queryFn: async () => {
      const entries = await pb
        .collection<ViewEntry>('view_entries')
        .getFullList()
      const entriesDistribution: EntriesDistribution = {}
      entries.forEach(entry => {
        const year = entry.year
        // Initialize year if it does not exist
        if (!(year in entriesDistribution)) {
          entriesDistribution[year] = {}
        }
        // Initialize type if it does not exists in year
        if (!(entry.typeName in entriesDistribution[year])) {
          entriesDistribution[year][entry.typeName as EntryType] = {}
        }
        // Initialize domain if it does not exists in type
        entriesDistribution[year][entry.typeName as EntryType]![
          entry.domainName as Domain
        ] = entry.value
      })
      return entriesDistribution
    },
  })
}
