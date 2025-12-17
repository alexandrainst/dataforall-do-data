import { ArrowUpOnSquareIcon } from '@heroicons/react/24/outline'
import { DDButton } from './controls/DDButton'
import { PageTemplate } from './layout/PageTemplate'
import { useEntryGroups } from '../hooks/useEntryGroups'
import { useExportEntries } from '../hooks/useExport'

export const Export = () => {
  const { data: entryGroups } = useEntryGroups()
  const {
    canExportEntries: { data: canExport },
    exportEntries,
  } = useExportEntries()
  return (
    <PageTemplate title="Eksportér data">
      <div className="max-w-xl space-y-4">
        {canExport === false && (
          <div className="rounded-xl border-2 border-red-600 bg-red-100 p-4 text-red-600">
            <h4>Eksport ikke mulig grundet manglende data</h4>
          </div>
        )}
        {canExport &&
          entryGroups?.map(data => {
            return (
              <div
                key={data.year}
                className="flex w-full items-center justify-between rounded-xl border-2 border-darkblue bg-white p-4 text-darkblue"
              >
                <h4>{data.year}</h4>
                <div className="flex gap-4">
                  <DDButton
                    variant="primary"
                    size="md"
                    className="flex items-center"
                    onClick={() => {
                      exportEntries(data.year)
                    }}
                  >
                    <ArrowUpOnSquareIcon className="mr-2 h-5 w-5" />
                    Eksportér
                  </DDButton>
                </div>
              </div>
            )
          })}
      </div>
    </PageTemplate>
  )
}