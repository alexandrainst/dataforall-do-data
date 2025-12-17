import { createFileRoute, useNavigate } from '@tanstack/react-router'
import { Link, Outlet, redirect, useRouter } from '@tanstack/react-router'
import { usePocketBase } from '../context/PocketBaseContext'
import { Button } from '@headlessui/react'
import { ArrowRightStartOnRectangleIcon } from '@heroicons/react/20/solid'
import VejleLogo from '../assets/vejle.png'
import { useQuery } from '@tanstack/react-query'
import { useState } from 'react'

const AuthLayout = () => {
  const router = useRouter()
  const navigate = useNavigate()
  const { pb } = usePocketBase()
  const [isMunicipality, setMunicipality] = useState(false)

  const handleLogout = () => {
    pb.authStore.clear()
    router.invalidate().finally(() => {
      navigate({ to: '/' })
    })
  }

  const userRecord = pb.authStore.record

  const { data: user } = useQuery({
    queryKey: ['userProfile', userRecord?.id],
    queryFn: async () => {
      if (userRecord?.id === undefined) {
        throw new Error('No user ID found')
      }
      const record = await pb.collection('users').getOne(userRecord.id, {
        expand: 'organization,role',
      })
      setMunicipality(record.expand?.role?.type === 'MUNICIPALITY')
      return record
    },
    enabled: userRecord?.id !== undefined,
  })

  const isUserWithoutCategory =
    user?.expand?.organization?.category === undefined ||
    user?.expand?.organization?.category === ''

  return (
    <div className="min-h-screen bg-greyblue">
      <div className="fixed top-0 left-0 z-100">
        <div className="relative top-3 -left-16 w-48 origin-center -rotate-45 bg-orange-500 py-1.5 text-center text-sm font-bold tracking-wide text-white uppercase shadow-lg">
          Beta
        </div>
      </div>

      <div className="sticky top-2 z-50 mx-auto w-[98%] max-w-[1920px]">
        <header className="rounded-lg bg-darkblue shadow-md">
          <nav className="px-4">
            <div className="flex h-16 items-center gap-4">
              <div className="flex items-center">
                <img
                  src={VejleLogo}
                  alt="Vejle Logo"
                  className="hidden h-8 sm:block"
                />
                <div className="ml-3 hidden text-xl font-semibold text-white select-none sm:block">
                  Do Data
                </div>
              </div>

              {/* Divider */}
              <div className="mx-2 hidden h-8 w-px bg-white/30 sm:block"></div>

              {isMunicipality ? (
                <div className="flex items-center gap-2 sm:gap-4">
                  <Link
                    to="/export"
                    className="rounded-md border border-darkblue px-2 py-2 text-sm text-white hover:bg-ocean sm:px-3 sm:text-base"
                    activeProps={{
                      className:
                        'bg-ocean border-ocean hover:brightness-120',
                    }}
                  >
                    Eksportér data
                  </Link>
                </div>
              ) : (
                <div className="flex items-center gap-2 sm:gap-4">
                  <Link
                    to="/profile"
                    className="rounded-md border border-darkblue px-2 py-2 text-sm text-white hover:bg-ocean sm:px-3 sm:text-base"
                    activeProps={{
                      className:
                        'bg-ocean border-ocean hover:brightness-120',
                    }}
                  >
                    Min virksomhed
                  </Link>
                  {!isUserWithoutCategory && (
                    <>
                      <Link
                        to="/dataEntries"
                        className="rounded-md border border-darkblue px-2 py-2 text-sm text-white hover:bg-ocean sm:px-3 sm:text-base"
                        activeProps={{
                          className:
                            'bg-ocean border-ocean hover:brightness-120',
                        }}
                      >
                        Data
                      </Link>

                      <Link
                        to="/dashboard"
                        className="rounded-md border border-darkblue px-2 py-2 text-sm text-white hover:bg-ocean sm:px-3 sm:text-base"
                        activeProps={{
                          className:
                            'bg-ocean border-ocean hover:brightness-120',
                        }}
                      >
                        Dashboard
                      </Link>
                    </>
                  )}
                  <Link
                    to="/guide"
                    className="rounded-md border border-darkblue px-2 py-2 text-sm text-white hover:bg-ocean sm:px-3 sm:text-base"
                    activeProps={{
                      className:
                        'bg-ocean border-ocean hover:brightness-120',
                    }}
                  >
                    Vejledning
                  </Link>
                </div>
              )}

              <div className="ml-auto flex items-center">
                <Button
                  type="button"
                  onClick={handleLogout}
                  className="group flex cursor-pointer items-center gap-1 rounded-md border border-white px-2 py-2 text-sm text-white hover:bg-white hover:text-darkblue sm:gap-2 sm:px-4 sm:text-base"
                >
                  <ArrowRightStartOnRectangleIcon className="h-4 w-4 transition-transform duration-200 group-hover:translate-x-1" />
                  <span className="hidden sm:inline">Log ud</span>
                </Button>
              </div>
            </div>
          </nav>
        </header>
      </div>

      <main className="mx-auto mt-4 w-[98%] max-w-[1920px]">
        <Outlet />
      </main>
    </div>
  )
}

export const Route = createFileRoute('/_auth')({
  beforeLoad: ({ context, location }) => {
    if (!context.auth.pb.authStore.isValid) {
      throw redirect({
        to: '/login',
        search: {
          redirect: location.href,
        },
      })
    }
  },
  component: AuthLayout,
})
