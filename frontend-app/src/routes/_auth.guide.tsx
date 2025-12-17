import { createFileRoute } from '@tanstack/react-router'
import { Guide } from '../components/Guide'

export const Route = createFileRoute('/_auth/guide')({
  component: Guide,
})
