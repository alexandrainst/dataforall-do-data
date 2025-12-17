import { createRootRouteWithContext, Outlet } from '@tanstack/react-router'
import type { PocketBaseContextType } from '../context/PocketBaseContext'

type MyRouterContext = {
  auth: PocketBaseContextType
}

export const Route = createRootRouteWithContext<MyRouterContext>()({
  component: () => (
    <>
      <Outlet />
      {/* <TanStackRouterDevtools position="bottom-right" initialIsOpen={false} /> */}
    </>
  ),
})
