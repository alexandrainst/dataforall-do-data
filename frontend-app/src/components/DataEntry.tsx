import { Form } from './Form'
import {
  useClimateCompassForm,
  type ClimateCompassFormInput,
} from '../hooks/useClimateCompassForm'
import { Route } from '../routes/_auth.dataEntry.$year'
import { PageTemplate } from './layout/PageTemplate'

export const DataEntry = () => {
  const { year } = Route.useParams()

  const { submit, formInput, isLoading } = useClimateCompassForm(
    year === '0' ? undefined : year //TODO: Workaround -> cannot get optional params working
  )

  return (
    <PageTemplate title="Data">
      <div className="max-w-sm space-y-4">
        <h4>Energiforbrug og udledning</h4>
        {isLoading ? (
          <p>Loading... </p>
        ) : (
          <Form
            formInput={formInput}
            onSubmit={data =>
              submit(data as unknown as ClimateCompassFormInput)
            }
          />
        )}
      </div>
    </PageTemplate>
  )
}
