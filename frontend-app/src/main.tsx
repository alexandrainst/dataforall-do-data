import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import { RouterProvider, createRouter } from '@tanstack/react-router'
import { PocketBaseProvider, usePocketBase } from './context/PocketBaseContext'
import { routeTree } from './routeTree.gen'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'

const router = createRouter({
  routeTree,
  defaultPreload: 'intent',
  scrollRestoration: true,
  context: {
    auth: undefined!,
  },
})

declare module '@tanstack/react-router' {
  interface Register {
    router: typeof router
  }
}

const App = () => {
  const auth = usePocketBase()
  return <RouterProvider router={router} context={{ auth }} />
}

const queryClient = new QueryClient()

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <PocketBaseProvider>
      <QueryClientProvider client={queryClient}>
        <App />
      </QueryClientProvider>
    </PocketBaseProvider>
  </StrictMode>
)
