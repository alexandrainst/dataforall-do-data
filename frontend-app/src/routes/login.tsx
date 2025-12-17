import { createFileRoute, redirect } from '@tanstack/react-router'
import { Login } from '../components/Login'
import z from 'zod'

// eslint-disable-next-line @typescript-eslint/no-unnecessary-type-assertion
const fallback = '/dashboard' as const

export const Route = createFileRoute('/login')({
  validateSearch: z.object({
    redirect: z.string().optional().catch(''),
  }),
  beforeLoad: ({ context, search }) => {
    if (context.auth.pb.authStore.isValid) {
      throw redirect({ to: search.redirect || fallback })
    }
  },
  component: Login,
})
