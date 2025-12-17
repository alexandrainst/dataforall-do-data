import {
  Field,
  Label,
  Listbox,
  ListboxButton,
  ListboxOption,
  ListboxOptions,
} from '@headlessui/react'
import { ChevronDownIcon, CheckIcon } from '@heroicons/react/20/solid'
import {
  Controller,
  type Control,
  type FieldError,
  type RegisterOptions,
} from 'react-hook-form'

type InputSelectProps = {
  label: string
  options: { id: string; name: string }[]
  placeholder?: string
  name: string
  control: Control
  initialSelectedId?: string
  rules?:
    | Omit<
        RegisterOptions,
        'valueAsNumber' | 'valueAsDate' | 'setValueAs' | 'disabled'
      >
    | undefined
  error?: FieldError
  onValueChange?: () => void
}

export const InputSelect = ({
  label,
  options,
  placeholder,
  name,
  control,
  initialSelectedId,
  rules,
  error,
  onValueChange = () => {},
}: InputSelectProps) => {
  return (
    <Controller
      name={name}
      control={control}
      rules={rules}
      render={({ field: { onChange, value } }) => {
        if (value === undefined && initialSelectedId !== undefined) {
          onChange(initialSelectedId)
        }
        const selectedOption = options.find(
          option => option.id === (value || '')
        )
        const displayValue = selectedOption ? selectedOption.name : placeholder

        return (
          <div className="w-full">
            <Field>
              <Label className="mb-1 block text-sm font-medium text-darkblue">
                {label}
              </Label>
              <Listbox
                value={value || ''}
                onChange={val => {
                  onChange(val)
                  onValueChange()
                }}
              >
                <ListboxButton className="relative w-full cursor-pointer rounded border border-darkblue bg-white py-2 pr-8 pl-3 text-left text-darkblue focus:border-darkblue focus:ring-2 focus:ring-darkblue focus:outline-none">
                  {displayValue}
                  <ChevronDownIcon className="group pointer-events-none absolute top-2 right-2 size-6 fill-darkblue" />
                </ListboxButton>
                <ListboxOptions className="absolute z-10 mt-1 rounded border border-darkblue bg-white shadow-lg focus:outline-none">
                  {options.map(option => (
                    <ListboxOption
                      key={option.id}
                      value={option.id}
                      className="group relative cursor-pointer rounded py-2 pr-2 pl-10 text-darkblue select-none hover:bg-lightblue hover:text-white data-focus:bg-lightblue data-focus:text-white"
                    >
                      {option.name}
                      <CheckIcon className="invisible absolute top-2.5 left-2.5 size-5 fill-darkblue group-hover:fill-white group-data-focus:fill-white group-data-selected:visible" />
                    </ListboxOption>
                  ))}
                </ListboxOptions>
              </Listbox>
              {error && (
                <p className="mt-1 text-sm text-red-600">{error.message}</p>
              )}
            </Field>
          </div>
        )
      }}
    />
  )
}
