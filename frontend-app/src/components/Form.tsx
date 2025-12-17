import { useForm, type RegisterOptions } from 'react-hook-form'
import { InputText } from './InputText'
import { InputSelect } from './InputSelect'
import { DDButton } from './controls/DDButton'

type FormInputField = {
  name: string
  type: 'text' | 'select'
  label: string
  defaultValue?: string
  value?: string
  placeholder?: string
  options?: { id: string; name: string }[]
  error?: string
  registerOptions?: RegisterOptions<FormOutput, string>
  secret?: boolean
  hide?: boolean
  accompanyingSelect?: {
    name: string
    label: string
    defaultValue?: string
    placeholder: string
    options: { id: string; name: string }[]
    registerOptions?: RegisterOptions<FormOutput, string>
  }
}

export type FormInput = {
  inputs: FormInputField[]
}

type FormOutput = {
  [key: string]: string
}

type FormProps = {
  formInput: FormInput
  actionLabel?: string
  onSubmit: (data: FormOutput) => void
}

export const Form = ({
  formInput,
  actionLabel = 'Gem',
  onSubmit,
}: FormProps) => {
  const {
    register,
    handleSubmit,
    control,
    formState: { errors },
  } = useForm<FormOutput>({
    values: Object.fromEntries(
      formInput.inputs.map(input => [input.name, input.value ?? ''])
    ),
    defaultValues: Object.fromEntries(
      formInput.inputs.flatMap(input => {
        const entries = [[input.name, input.defaultValue || '']]
        if (input.accompanyingSelect !== undefined) {
          entries.push([
            input.accompanyingSelect.name,
            input.accompanyingSelect.defaultValue || '',
          ])
        }
        return entries
      })
    ),
  })

  return (
    <>
      <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
        {formInput.inputs.map(input => {
          if (input.type === 'text') {
            return (
              <div
                key={input.name}
                className={`${input.hide ? 'hidden' : ''} flex gap-2`}
              >
                <InputText
                  key={input.name}
                  label={input.label}
                  secret={input.secret}
                  placeholder={input.placeholder || ''}
                  error={errors[input.name]?.message}
                  {...register(input.name, {
                    ...input.registerOptions,
                  })}
                />

                {input.accompanyingSelect !== undefined && (
                  <div className="w-35">
                    <InputSelect
                      key={input.accompanyingSelect.name}
                      label={input.accompanyingSelect.label}
                      options={input.accompanyingSelect.options}
                      placeholder={input.accompanyingSelect.placeholder}
                      name={input.accompanyingSelect.name}
                      control={control}
                      error={errors[input.accompanyingSelect.name]}
                      rules={input.accompanyingSelect.registerOptions}
                    />
                  </div>
                )}
              </div>
            )
          } else if (input.type === 'select' && input.options) {
            return (
              <div key={input.name} className="flex gap-2">
                <InputSelect
                  key={input.name}
                  label={input.label}
                  options={input.options}
                  placeholder={input.placeholder}
                  name={input.name}
                  control={control}
                  error={errors[input.name]}
                  rules={input.registerOptions}
                />
                {input.accompanyingSelect !== undefined && (
                  <div className="w-35">
                    <InputSelect
                      key={input.accompanyingSelect.name}
                      label={input.accompanyingSelect.label}
                      options={input.accompanyingSelect.options}
                      placeholder={input.accompanyingSelect.placeholder}
                      name={input.accompanyingSelect.name}
                      control={control}
                      error={errors[input.accompanyingSelect.name]}
                      rules={input.accompanyingSelect.registerOptions}
                    />
                  </div>
                )}
              </div>
            )
          }
          return null
        })}
        <DDButton variant="primary" size="md" type="submit">
          {actionLabel}
        </DDButton>
      </form>
    </>
  )
}
