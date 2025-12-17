import { Link, useNavigate } from '@tanstack/react-router'
import { usePocketBase } from '../context/PocketBaseContext'
import { useState } from 'react'
import { Button, Input } from '@headlessui/react'
import { ArrowRightEndOnRectangleIcon } from '@heroicons/react/24/outline'
import VejleLogo from '../assets/logo.png'
import Logos from '../assets/logos.png'

export const Login = () => {
  const { pb } = usePocketBase()
  const navigate = useNavigate()
  const [errorText, setErrorText] = useState<string | undefined>(undefined)

  const onSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    const email = (document.getElementById('email') as HTMLInputElement).value
    const password = (document.getElementById('password') as HTMLInputElement)
      .value
    login(email, password)
  }

  const login = async (email: string, password: string) => {
    try {
      const authData = await pb
        .collection('users')
        .authWithPassword(email, password, { expand: 'role' })

      // Check if user is verified
      if (!authData.record.verified) {
        pb.authStore.clear()
        setErrorText(
          'Din konto er ikke verificeret. Tjek venligst din email for at verificere din konto.'
        )
        return
      }

      // Municipality users need to have their role assigned manually
      if (authData.record.expand?.role?.type === undefined) {
        setErrorText(
          'Din konto skal aktiveres. Kontakt venligst systemadministrator.'
        )
        return
      }

      setErrorText(undefined)

      if (authData.record.expand?.role?.type === 'MUNICIPALITY') {
        navigate({ to: '/export' })
      } else {
        navigate({
          to:
            localStorage.getItem('hideGuide') === 'true'
              ? '/dashboard'
              : '/guide',
        })
      }
    } catch (error) {
      console.error('Login failed:', error)
      setErrorText(
        'Log ind mislykkedes. Tjek venligst dine oplysninger og prøv igen.'
      )
    }
  }

  return (
    <div className="flex min-h-screen items-center justify-center bg-greyblue p-6">
      <div className="z-10 w-full max-w-sm">
        <div className="mb-8 flex justify-center">
          <img
            src={VejleLogo}
            alt="Vejle Logo"
            className="h-16 w-auto object-contain"
          />
        </div>
        <h4 className="mb-6 text-center font-semibold text-darkblue">
          Do Data - for a{' '}
          <span className="font-semibold text-[#00b2bc]">Greener</span> Future
        </h4>
        <form onSubmit={e => onSubmit(e)} className="flex flex-col gap-2">
          <Input
            type="text"
            placeholder="Email"
            id="email"
            className="rounded border border-darkblue bg-transparent px-3 py-2 text-xl placeholder-gray-400 focus:border-transparent focus:ring-2 focus:outline-none"
          />
          <Input
            type="password"
            placeholder="Adgangskode"
            id="password"
            className="rounded border border-darkblue bg-transparent px-3 py-2 text-xl placeholder-gray-400 focus:border-transparent focus:ring-2 focus:outline-none"
          />
          <Button
            type="submit"
            className="group flex cursor-pointer items-center justify-center gap-2 rounded bg-sky px-4 py-2 text-xl text-gray-800 shadow-xs shadow-gray-800 hover:brightness-110 focus:ring-2 focus:ring-black focus:brightness-120 focus:outline-none active:border-none active:shadow-none"
          >
            <ArrowRightEndOnRectangleIcon className="h-5 w-5 transition-transform duration-200 group-hover:translate-x-1" />
            Log ind
          </Button>
          <div className="hover:text-gray-1000 mt-4 text-center text-sm text-gray-600 underline">
            <Link to="/register">Ny bruger?</Link>
          </div>
          {errorText !== undefined && (
            <div className="mt-2 rounded p-2 text-sm text-sadred">
              {errorText}
            </div>
          )}
        </form>
      </div>

      {/* Triangle in top right corner */}
      <div
        className="absolute top-0 right-0 h-44 w-200 bg-sky"
        style={{
          clipPath: 'polygon(100% 0%, 100% 100%, 0% 0%)',
        }}
      ></div>

      {/* Logos image inside triangle */}
      <img
        src={Logos}
        alt="Logos"
        className="absolute top-0 right-0 mt-3 mr-3 h-[100px]"
      />

      {/* Soft wave shape left - extends full width */}
      <svg
        className="absolute bottom-0 left-0 h-58 max-h-[220px] w-full opacity-70"
        viewBox="0 0 1000 250"
        preserveAspectRatio="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          d="M0,80 Q100,60 200,90 T400,70 T600,100 T800,80 T1000,110 L1000,250 L0,250 Z"
          fill="currentColor"
          className="text-sky"
        />
      </svg>

      {/* Soft wave shape right - extends full width */}
      <svg
        className="absolute right-0 bottom-0 h-42 max-h-[180px] w-full opacity-70"
        viewBox="0 0 1000 150"
        preserveAspectRatio="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          d="M0,100 Q125,40 250,70 T500,30 T750,50 T1000,20 L1000,150 L0,150 Z"
          fill="currentColor"
          className="text-lightblue"
        />
      </svg>

      <div className="absolute bottom-4">
        Vejle Kommune & Alexandra Instituttet 2025
      </div>
    </div>
  )
}
