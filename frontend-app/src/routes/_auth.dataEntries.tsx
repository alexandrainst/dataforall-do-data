import { createFileRoute } from '@tanstack/react-router'
import { DataEntries } from '../components/DataEntries'

export const Route = createFileRoute('/_auth/dataEntries')({
  component: DataEntries,
})
