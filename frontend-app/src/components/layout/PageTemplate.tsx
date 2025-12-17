import type { ReactNode } from 'react'

type PageTemplateProps = {
  title: string
  children: ReactNode
}

export const PageTemplate = ({ title, children }: PageTemplateProps) => {
  return (
    <div className="p-2 sm:p-2 md:p-6">
      <div className="flex">
        <h2 className="mb-6 text-darkblue">{title}</h2>
      </div>
      {children}
    </div>
  )
}
