import { createContext, useContext } from 'react'
import type { ReactNode } from 'react'
import PocketBase from 'pocketbase'

// TODO: Url for prod
const pb = new PocketBase('http://localhost:8080')
pb.autoCancellation(false)

export type PocketBaseContextType = {
  pb: PocketBase
}

const PocketBaseContext = createContext<PocketBaseContextType | undefined>(
  undefined
)

export const PocketBaseProvider = ({ children }: { children: ReactNode }) => {
  return (
    <PocketBaseContext.Provider value={{ pb }}>
      {children}
    </PocketBaseContext.Provider>
  )
}

export const usePocketBase = () => {
  const context = useContext(PocketBaseContext)
  if (!context) {
    throw new Error('usePocketBase must be used within a PocketBaseProvider')
  }
  return context
}
