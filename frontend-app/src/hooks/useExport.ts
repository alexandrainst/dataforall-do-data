import { useQuery } from "@tanstack/react-query"
import { usePocketBase } from "../context/PocketBaseContext"

export const useExportEntries = () => {
  const { pb } = usePocketBase()
  return ({
    canExportEntries:
      useQuery<boolean, Error>({
        queryKey: ["exportcheck"],
        queryFn: async () => {
          const request = fetch(`http://localhost:8080/api/exportcheck`, {
            method: 'GET',
            headers: {
              Authorization: `Bearer ${pb.authStore.token}`,
            },
          })
          const response = await request
          if (!response.ok) {
            const errorMessage = await response.json()
            return Promise.reject(new Error(errorMessage))
          }
          return (await response.json()).canExport
        }
      }),
    exportEntries: async (year: string) => {
      const request = fetch(`http://localhost:8080/api/export/${year}`, {
        method: 'GET',
        headers: {
          Authorization: `Bearer ${pb.authStore.token}`,
        },
      })
      const response = await request
      if (!response.ok) {
        const errorMessage = await response.json()
        return Promise.reject(new Error(errorMessage))
      }
      const fileUrl = URL.createObjectURL(await response.blob())
      const downloadLink = document.createElement("a")
      downloadLink.href = fileUrl
      downloadLink.download = `${year}-DoData.csv`
      document.body.appendChild(downloadLink)
      downloadLink.click()
      downloadLink.remove()
    }
  })
}
