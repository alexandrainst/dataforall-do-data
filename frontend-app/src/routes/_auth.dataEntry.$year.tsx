import { createFileRoute } from '@tanstack/react-router'
import { DataEntry } from '../components/DataEntry'

export const Route = createFileRoute('/_auth/dataEntry/$year')({
  component: DataEntry,
})
