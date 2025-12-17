import { DDButton } from '../controls/DDButton'
import { Dialog } from '@headlessui/react'

type DDDialogProps = {
  open: boolean
  onClose: () => void
  title: string
  body: string
  confirmLabel: string
  cancelLabel: string
  onConfirm: () => void
  onCancel: () => void
}

export const DDDialog = ({
  open,
  onClose,
  title,
  body,
  confirmLabel,
  cancelLabel,
  onConfirm,
  onCancel,
}: DDDialogProps) => {
  return (
    <Dialog open={open} onClose={onClose}>
      <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/30 backdrop-blur-xs">
        <div className="rounded-xl bg-white p-6 shadow-sm">
          <h2 className="mb-4 text-lg font-semibold text-darkblue">{title}</h2>
          <p className="mb-6">{body}</p>
          <div className="flex justify-end gap-4">
            <DDButton variant="primary" size="md" onClick={onCancel}>
              {cancelLabel}
            </DDButton>
            <DDButton variant="danger" size="md" onClick={onConfirm}>
              {confirmLabel}
            </DDButton>
          </div>
        </div>
      </div>
    </Dialog>
  )
}
