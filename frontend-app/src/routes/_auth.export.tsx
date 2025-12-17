import { createFileRoute } from '@tanstack/react-router'
import { Export } from '../components/Export'

export const Route = createFileRoute('/_auth/export')({
  component: Export,
})
