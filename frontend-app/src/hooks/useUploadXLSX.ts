import { usePocketBase } from '../context/PocketBaseContext'
import { useMutation, useQueryClient } from '@tanstack/react-query'

export type useUploadXlSXProps = {
  onError?: (error: Error) => void
  onSuccess?: () => void
}

export const useUploadXlSX = ({ onError, onSuccess }: useUploadXlSXProps) => {
  const { pb } = usePocketBase()
  const queryClient = useQueryClient()

  const { mutateAsync, isError } = useMutation<void, Error, File>({
    mutationFn: async file => {
      /* Create multipart form request and POST to pocketbase endpoint */
      const fileFormData = new FormData()
      fileFormData.append('file', file)
      const request = fetch('/api/upload', {
        method: 'POST',
        headers: {
          Authorization: `Bearer ${pb.authStore.token}`,
        },
        body: fileFormData,
      })
      const response = await request
      if (!response.ok) {
        const errorMessage = await response.json()
        return Promise.reject(new Error(errorMessage))
      }
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['entryGroups'] })
      onSuccess?.()
    },
    onError: onError,
  })

  return {
    uploadXLSX: mutateAsync,
    isError: isError,
  }
}
