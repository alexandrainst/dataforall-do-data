import { Field, Label, Input } from '@headlessui/react'
import type { UseFormRegisterReturn } from 'react-hook-form'

type InputTextProps = {
  label: string
  placeholder: string
  secret?: boolean
  error?: string
} & UseFormRegisterReturn

export const InputText = ({
  label,
  placeholder,
  secret = false,
  error,
  ...registerProps
}: InputTextProps) => {
  return (
    <div className={`w-full`}>
      <Field>
        <Label className="mb-1 block text-sm font-medium text-darkblue">
          {label}
        </Label>
        <Input
          type={secret ? 'password' : 'text'}
          placeholder={placeholder}
          className="w-full rounded border border-darkblue bg-white px-3 py-2 text-darkblue placeholder-gray-400 focus:border-darkblue focus:ring-2 focus:ring-darkblue focus:outline-none"
          {...registerProps}
        />
        {error && <p className="mt-1 text-sm text-red-600">{error}</p>}
      </Field>
    </div>
  )
}
