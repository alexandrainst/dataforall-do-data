import { usePocketBase } from "../context/PocketBaseContext"
import { useQuery } from "@tanstack/react-query"

export type EntryGroup = {
  organization: string,
  year: string
}

export const useEntryGroups = () => {
  const { pb } = usePocketBase()
  return useQuery<EntryGroup[], Error>({
    queryKey: ['entryGroups'],
    queryFn: () => pb.collection<EntryGroup>('entry_groups').getFullList()
  })
}

