import { useNavigate } from '@tanstack/react-router'
import { DDButton } from './controls/DDButton'
import { GuideStep } from './GuideStep'
import {
  ArrowRightCircleIcon,
  ArrowUpOnSquareIcon,
  ArrowDownOnSquareIcon,
  ChartBarIcon,
  CheckIcon,
} from '@heroicons/react/24/outline'
import { Checkbox, Field, Label } from '@headlessui/react'
import { useState } from 'react'
import { PageTemplate } from './layout/PageTemplate'
import { useQuery } from '@tanstack/react-query'
import { usePocketBase } from '../context/PocketBaseContext'

export const Guide = () => {
  const { pb } = usePocketBase()
  const navigate = useNavigate()
  const [hideGuide, setHideGuide] = useState(
    localStorage.getItem('hideGuide') === 'true'
  )

  const userRecord = pb.authStore.record

  const { data: user, isLoading: isLoadingUser } = useQuery({
    queryKey: ['userProfile', userRecord?.id],
    queryFn: async () => {
      if (userRecord?.id === undefined) {
        throw new Error('No user ID found')
      }
      const record = await pb.collection('users').getOne(userRecord.id, {
        expand: 'organization',
      })
      return record
    },
    enabled: userRecord?.id !== undefined,
  })

  const isUserWithoutCategory =
    user?.expand?.organization?.category === undefined ||
    user?.expand?.organization?.category === ''

  return (
    <PageTemplate title="Vejledning">
      <div className="mb-8 max-w-md space-y-4">
        <p>
          Følg trinene nedenfor for at komme i gang med at indlæse dine data og
          udforske visualiseringer og indsigter.
        </p>
      </div>

      <div className="space-y-4">
        <GuideStep
          step={1}
          title="Eksportér dine data fra Klimakompasset"
          icon={<ArrowDownOnSquareIcon className="h-[32px] w-[32px]" />}
        >
          Eksportér dine data fra Klimakompasset til en Excel-fil. Data skal
          være for et helt år.
        </GuideStep>

        <GuideStep
          step={2}
          title="Indlæs dine data her i Do Data"
          icon={<ArrowUpOnSquareIcon className="h-[32px] w-[32px]" />}
        >
          På siden "Data" kan du indlæse de Excel-filer du har eksporteret fra
          Klimakompasset.
        </GuideStep>

        <GuideStep
          step={3}
          title="Udforsk visualiseringer og indsigter"
          icon={<ChartBarIcon className="h-[32px] w-[32px]" />}
        >
          Når dine data er indlæst, kan du gå til siden "Dashboard" for at se
          forskellige visualiseringer og indsigter om det indlæste data.
        </GuideStep>
      </div>
      <DDButton
        variant="primary"
        size="md"
        disabled={isLoadingUser}
        onClick={() => {
          if (isUserWithoutCategory) {
            navigate({ to: '/profile' })
            return
          }
          navigate({ to: '/dataEntries' })
        }}
        className="group mt-4 flex items-center"
      >
        Gå til indlæsning af data
        <ArrowRightCircleIcon className="ml-2 h-6 w-6 transition-transform group-hover:translate-x-1" />
      </DDButton>
      <Field className="mt-4 flex cursor-pointer items-center gap-2">
        <Checkbox
          checked={!hideGuide}
          onChange={() => {
            localStorage.setItem('hideGuide', String(!hideGuide))
            setHideGuide(!hideGuide)
          }}
          className="group block rounded border bg-white data-checked:bg-darkblue"
        >
          <CheckIcon className="invisible h-4 w-4 stroke-2 text-white group-data-checked:visible" />
        </Checkbox>

        <Label className="cursor-pointer">Vis vejledning ved log ind</Label>
      </Field>
    </PageTemplate>
  )
}
