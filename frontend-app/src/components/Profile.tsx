import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { CheckIcon, TrashIcon } from '@heroicons/react/20/solid'
import { usePocketBase } from '../context/PocketBaseContext'
import { PageTemplate } from './layout/PageTemplate'
import { InputSelect } from './InputSelect'
import { useForm } from 'react-hook-form'
import { DDButton } from './controls/DDButton'
import { useState } from 'react'
import { useNavigate } from '@tanstack/react-router'
import { DDDialog } from './layout/Dialog'

export const Profile = () => {
  const { pb } = usePocketBase()
  const queryClient = useQueryClient()
  const navigate = useNavigate()
  const { control, getValues } = useForm()

  const [isSaved, setSaved] = useState(false)
  const [showDeleteConfirm, setShowDeleteConfirm] = useState(false)

  const userRecord = pb.authStore.record

  const { data: user } = useQuery({
    queryKey: ['userProfile', userRecord?.id],
    queryFn: async () => {
      if (userRecord?.id === undefined) {
        throw new Error('No user ID found')
      }
      const record = await pb.collection('users').getOne(userRecord.id, {
        expand: 'organization',
      })
      return record
    },
    enabled: userRecord?.id !== undefined,
  })

  const { data: organizationCategories } = useQuery({
    queryKey: ['organizationCategories'],
    queryFn: async () => {
      const record = await pb
        .collection('organization_categories')
        .getFullList()
      return record
    },
    enabled: userRecord?.id !== undefined,
  })

  const isUserWithoutCategory =
    user?.expand?.organization?.category === undefined ||
    user?.expand?.organization?.category === ''

  const orgCats = organizationCategories?.map(cat => ({
    id: cat.id,
    name: cat.name,
  }))

  const updateCategoryMutation = useMutation({
    mutationFn: async (categoryId: string) => {
      if (user?.expand?.organization?.id === undefined) {
        throw new Error('No organization ID found')
      }
      return await pb
        .collection('organizations')
        .update(user.expand.organization.id, {
          category: categoryId,
        })
    },
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ['userProfile', userRecord?.id],
      })
      if (isUserWithoutCategory) {
        navigate({ to: '/dataEntries' })
      }
      setSaved(true)
    },
    onError: error => {
      console.error('Failed to update organization category:', error)
      setSaved(false)
    },
  })

  // Copilot creation, no idea if it is correct
  // const deleteAllDataMutation = useMutation({
  //   mutationFn: async () => {
  //     if (user?.expand?.organization?.id === undefined) {
  //       throw new Error('No organization ID found')
  //     }
  //     // Delete all data entries for this organization
  //     const dataEntries = await pb.collection('data_entries').getFullList({
  //       filter: `organization="${user.expand.organization.id}"`,
  //     })

  //     await Promise.all(
  //       dataEntries.map(entry => pb.collection('data_entries').delete(entry.id))
  //     )
  //   },
  //   onSuccess: () => {
  //     queryClient.invalidateQueries()
  //     setShowDeleteConfirm(false)
  //     alert('Alle data er blevet slettet')
  //   },
  //   onError: error => {
  //     console.error('Failed to delete data:', error)
  //     alert('Fejl ved sletning af data')
  //   },
  // })

  return (
    <PageTemplate title={user?.expand?.organization?.name}>
      <div className="mb-8 max-w-md space-y-4">
        {isUserWithoutCategory && (
          <p>
            Før du kan begynde at indlæse data, skal du angive din
            virksomhedstype nedenfor. Dette hjælper os med at tilpasse
            oplevelsen til din organisation.
          </p>
        )}
        <InputSelect
          label="Virksomhedstype"
          initialSelectedId={
            isUserWithoutCategory
              ? undefined
              : user?.expand?.organization?.category
          }
          name="businessType"
          control={control}
          options={orgCats || []}
          onValueChange={() => setSaved(false)}
          placeholder="Vælg virksomhedstype"
        />
      </div>

      <div className="flex gap-4">
        <DDButton
          variant="primary"
          onClick={() => {
            if (
              getValues('businessType') &&
              getValues('businessType') !== user?.expand?.organization?.category
            ) {
              updateCategoryMutation.mutate(getValues('businessType'))
            }
          }}
        >
          <span className="flex items-center gap-2">
            {`Gem${isSaved ? 't' : ''}`}
            <CheckIcon
              className={`size-5 fill-white ${isSaved ? '' : 'hidden'}`}
            />
          </span>
        </DDButton>

        <DDButton variant="danger" onClick={() => setShowDeleteConfirm(true)}>
          <span className="flex items-center gap-2">
            <TrashIcon className="size-5" />
            Slet min virksomhed
          </span>
        </DDButton>
      </div>

      <DDDialog
        open={showDeleteConfirm}
        onClose={() => setShowDeleteConfirm(false)}
        title="Bekræft sletning"
        body="Er du sikker på, at du vil slette alle data for din organisation? Denne handling kan ikke fortrydes."
        confirmLabel="Ja, slet alle data"
        cancelLabel="Annuller"
        onConfirm={() => {
          // deleteAllDataMutation.mutate()
        }}
        onCancel={() => setShowDeleteConfirm(false)}
      />
    </PageTemplate>
  )
}
