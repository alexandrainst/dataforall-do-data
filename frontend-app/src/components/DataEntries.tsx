import { ArrowDownOnSquareIcon } from '@heroicons/react/24/outline'
import { useNavigate } from '@tanstack/react-router'
import { useEntryGroups } from '../hooks/useEntryGroups'
import { UploadButton } from './UploadButton'
import { useState } from 'react'
import { DDButton } from './controls/DDButton'
import { ChartBarIcon, TrashIcon } from '@heroicons/react/24/outline'
import { useEntries } from '../hooks/useEntries'
import { PageTemplate } from './layout/PageTemplate'
import { DDDialog } from './layout/Dialog'

export const DataEntries = () => {
  const navigate = useNavigate()
  const [isDeletingYear, setDeletingYear] = useState<string | undefined>(
    undefined
  )
  const { data: entryGroups } = useEntryGroups()
  const { delete: deleteEntry } = useEntries()

  return (
    <PageTemplate title="Data">
      <p className="mb-8 max-w-xl space-y-4">
        Her kan du se en oversigt over dine data. Dine forbrugsdata bliver
        behandlet fortroligt og sikkert. Ingen andre virksomheder eller kommunen
        kan se dine oplysninger - kun du har adgang til dine visualiseringer og
        analyser. Dine data bruges i aggregeret form til sammenligninger på
        tværs at virksomheder.
      </p>

      <div className="mb-4 max-w-xs">
        <UploadButton
          className={`group ml-auto ${entryGroups?.length === 0 ? 'animate-pulse' : ''}`}
          label="Indlæs Excel-fil fra Klimakompasset"
          icon={
            <ArrowDownOnSquareIcon className="h-6 w-6 transition-transform group-hover:translate-y-[-2px]" />
          }
        />
      </div>
      {entryGroups?.length === 0 && (
        <p className="mb-6 max-w-xl text-gray-500">
          Du har ikke tilføjet noget data endnu. Klik på knappen for at tilføje
          data fra Klimakompasset.
        </p>
      )}

      <div className="max-w-xl space-y-4">
        {entryGroups?.map(entryGroup => {
          return (
            <div
              key={entryGroup.year}
              className="flex w-full items-center justify-between rounded-xl border-2 border-darkblue bg-white p-4 text-darkblue"
            >
              <h4>{entryGroup.year}</h4>
              <div className="flex gap-4">
                <DDButton
                  variant="primary"
                  size="md"
                  className="flex items-center"
                  onClick={() => setDeletingYear(entryGroup?.year)}
                >
                  <TrashIcon className="mr-2 h-5 w-5" />
                  Slet
                </DDButton>
                <DDButton
                  variant="primary"
                  size="md"
                  className="flex items-center"
                  onClick={() => navigate({ to: `/dashboard` })}
                >
                  <ChartBarIcon className="mr-2 h-5 w-5" />
                  Gå til dashboard
                </DDButton>
              </div>
            </div>
          )
        })}
      </div>

      <DDDialog
        open={isDeletingYear !== undefined}
        onClose={() => setDeletingYear(undefined)}
        title={`Bekræft sletning af data for ${isDeletingYear}`}
        body="Er du sikker på, at du vil slette disse data? Denne handling kan ikke fortrydes."
        confirmLabel="Slet"
        cancelLabel="Annuller"
        onConfirm={async () => {
          if (isDeletingYear !== undefined) {
            await deleteEntry({ year: isDeletingYear })
            setDeletingYear(undefined)
          }
        }}
        onCancel={() => setDeletingYear(undefined)}
      />
    </PageTemplate>
  )
}
