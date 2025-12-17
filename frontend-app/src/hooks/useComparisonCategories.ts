import { useQuery } from "@tanstack/react-query"
import { usePocketBase } from "../context/PocketBaseContext"
import type { Domain, EntriesDistribution, EntryType } from "./useEntries"

export type ComparisonEntry = {
  category: string
  value: number
  domainName: string
  typeName: string
  unitName: string
  year: number
}

export type ComparisonCategory = {
  id: string
  name: string
}

export const useComparisonCategories = () => {
  const { pb } = usePocketBase()
  return useQuery<ComparisonCategory[]>({
    queryKey: ["comparisonCategories"],
    queryFn: async () => {
      const records = await pb
        .collection("comparison_categories")
        .getFullList<ComparisonCategory>({
          sort: "name",
        })
      return records
    }
  })
}

export const useComparisonEntriesDistribution = (categoryId: string | undefined) => {
  const { pb } = usePocketBase()
  return useQuery<EntriesDistribution, Error>({
    queryKey: ["comparisonEntriesDistribution", categoryId],
    queryFn: async () => {
      const entries = await pb.collection<ComparisonEntry>('view_comparison_entries').getFullList({
        filter: `category="${categoryId}"`,
      })
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
        entriesDistribution[year][entry.typeName as EntryType]![entry.domainName as Domain] = entry.value
      })
      return entriesDistribution
    }
  })
}
