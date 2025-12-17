import { Button } from '@headlessui/react'
import { useRef, useState, type ReactNode } from 'react'
import { useUploadXlSX } from '../hooks/useUploadXLSX'
import { ErrorDialog } from './ErrorDialog'

export type UploadButtonProps = {
  label: string
  icon?: ReactNode
  className?: string
}

export const UploadButton = ({ label, icon, className }: UploadButtonProps) => {
  const [showUploadError, setShowUploadError] = useState<{
    show: boolean
    message?: string
  }>({ show: false })

  const { uploadXLSX } = useUploadXlSX({
    onError: error => {
      setShowUploadError({
        show: true,
        message:
          error.message === 'year_error'
            ? 'Filen dækker ikke over et helt år.'
            : undefined,
      })
      if (fileInputRef.current != null) {
        fileInputRef.current.value = ''
      }
    },
    onSuccess: () => {
      if (fileInputRef.current != null) {
        fileInputRef.current.value = ''
      }
    },
  })

  const fileInputRef = useRef<HTMLInputElement>(null)

  const handleChange: React.ChangeEventHandler<HTMLInputElement> = async (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    /* Get file from input element */
    const file = event.target.files?.[0]
    if (file == undefined) {
      console.error('File not attach to html input')
      return
    }
    uploadXLSX(file)
  }

  return (
    <>
      <ErrorDialog
        open={showUploadError.show}
        onClose={() => {
          setShowUploadError({ show: false })
        }}
        message={showUploadError.message}
      />
      <div className={className}>
        <Button
          onClick={() => {
            fileInputRef.current?.click()
          }}
          className="transparent flex cursor-pointer items-center justify-center gap-2 rounded-xl border-2 border-dashed border-darkblue p-4 text-darkblue hover:bg-ocean/60 hover:text-white"
        >
          {label}
          {icon}
        </Button>
        <input
          ref={fileInputRef}
          onChange={handleChange}
          type="file"
          accept=".xlsx"
          hidden
          multiple={false}
        />
      </div>
    </>
  )
}
