import {
  Dialog,
  DialogPanel,
  DialogTitle,
  Description,
} from '@headlessui/react'

const _errorMessage = 'Prøv igen senere eller med en anden fil.'
const _errorTitle = 'Teknisk fejl'

export type ErrorDialogProps = {
  open: boolean
  onClose: () => void
  title?: string
  message?: string
}

export const ErrorDialog = ({
  open,
  onClose,
  title,
  message,
}: ErrorDialogProps) => {
  return (
    <Dialog open={open} onClose={onClose} className="z-50">
      <div className="fixed inset-0 flex w-screen items-center justify-center p-4">
        <DialogPanel className="rounded-md border-2 border-darkblue bg-white p-8 drop-shadow-lg">
          <DialogTitle>{title ?? _errorTitle}</DialogTitle>
          <Description>{message ?? _errorMessage}</Description>
        </DialogPanel>
      </div>
    </Dialog>
  )
}
