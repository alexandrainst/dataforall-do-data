import { useState } from 'react'
import { Form } from './Form'
import { PageTemplate } from './layout/PageTemplate'
import { usePocketBase } from '../context/PocketBaseContext'
import { useQuery } from '@tanstack/react-query'
import { useNavigate } from '@tanstack/react-router'
import { useForm, useWatch } from 'react-hook-form'
import { InputSelect } from './InputSelect'

export const Register = () => {
  const navigate = useNavigate()
  const { pb } = usePocketBase()
  const [errorMessage, setErrorMessage] = useState<string | undefined>()
  const [successMessage, setSuccessMessage] = useState<string | undefined>()
  const { control } = useForm()

  const selectedType = useWatch({
    control,
    name: 'type',
  })

  const { data: roles } = useQuery({
    queryKey: ['roles'],
    queryFn: async () => await pb.collection('roles').getFullList(),
  })

  const handleCreate = async (data: {
    type: string
    email: string
    password: string
    passwordConfirm: string
    organization: string
  }) => {
    try {
      let role
      let organizationRecord
      // Municipality users require manual assignment in the backend after creation
      if (data.type === 'MUNICIPALITY') {
        role = undefined
      } else {
        role = roles?.find(role => role.type === data.type)?.id
        if (role === undefined) {
          setErrorMessage(
            'Der skete en teknisk fejl under oprettelsen. Er alle felter udfyldt?'
          )
          return
        }
        // Create organization
        organizationRecord = await pb.collection('organizations').create({
          name: data.organization,
        })
      }

      // Create user with the organization ID
      await pb.collection('users').create({
        email: data.email,
        password: data.password,
        passwordConfirm: data.passwordConfirm,
        organization: organizationRecord?.id,
        role,
      })

      // TODO: Setup SMTP server and enable email verification
      await pb.collection('users').requestVerification(data.email)

      setErrorMessage(undefined)
      setSuccessMessage('Bruger oprettet! Du omdirigeres til login...')
      setTimeout(() => {
        navigate({ to: '/login' })
      }, 2000)
    } catch (error) {
      console.error('Error creating user:', error)
      setErrorMessage('Oprettelse af bruger mislykkedes. Prøv igen.')
    }
  }

  return (
    <div className="min-h-screen bg-greyblue">
      <PageTemplate title="Opret bruger">
        <p className="mb-8 max-w-xl space-y-4">
          Adgangskode skal være mindst 8 tegn. Husk at verificere din email
          efter oprettelse. Kommunebrugere kræver manuel godkendelse af
          systemadministrator.
        </p>

        <div className="max-w-md">
          <div className="mb-4">
            <InputSelect
              control={control}
              name="type"
              label="Brugertype"
              placeholder="Vælg type"
              options={[
                { id: 'ORGANIZATION', name: 'Virksomhed' },
                { id: 'MUNICIPALITY', name: 'Kommune' },
              ]}
            />
          </div>
          <Form
            formInput={{
              inputs: [
                {
                  name: 'email',
                  type: 'text',
                  label: 'Email',
                  placeholder: 'Indtast din email',
                },
                {
                  name: 'organization',
                  type: 'text',
                  label: 'Virksomhedsnavn',
                  hide: selectedType === 'MUNICIPALITY',
                  placeholder: 'Indtast navnet på din virksomhed',
                },
                {
                  name: 'password',
                  type: 'text',
                  label: 'Adgangskode',
                  secret: true,
                  placeholder: 'Indtast din adgangskode',
                },
                {
                  name: 'confirmPassword',
                  type: 'text',
                  secret: true,
                  label: 'Bekræft adgangskode',
                  placeholder: 'Bekræft din adgangskode',
                },
              ],
            }}
            actionLabel="Opret"
            onSubmit={data => {
              if (data.password !== data.confirmPassword) {
                setErrorMessage('Adgangskoderne matcher ikke.')
                return
              }
              setErrorMessage(undefined)
              handleCreate({
                type: selectedType,
                email: data.email,
                password: data.password,
                passwordConfirm: data.confirmPassword,
                organization: data.organization,
              })
            }}
          />
        </div>
        <div className="mt-4">
          {successMessage && <p className="text-green-600">{successMessage}</p>}
          {errorMessage && <p className="text-red-600">{errorMessage}</p>}
        </div>
      </PageTemplate>
    </div>
  )
}
