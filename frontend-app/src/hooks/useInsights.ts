import { useQuery } from "@tanstack/react-query"
import { usePocketBase } from "../context/PocketBaseContext"
import type { EntryType } from "./useEntries"

export type Insight = {
  title: string
  text: string
  typeName: EntryType
}

export const useInsights = () => {
  const { pb } = usePocketBase()
  return useQuery<Insight[], Error>({
    queryKey: ["insights"],
    queryFn: async () => {
      const insights = await pb
        .collection("active_recommendations")
        .getFullList<Insight>()
      return insights
    }
  })
}
