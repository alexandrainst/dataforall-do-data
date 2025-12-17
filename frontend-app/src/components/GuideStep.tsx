import type { ReactNode } from 'react'

type GuideStepProps = {
  step: number
  title: string
  icon?: ReactNode
  children?: ReactNode
  className?: string
}

export const GuideStep = ({
  step,
  title,
  icon,
  children,
  className = '',
}: GuideStepProps) => {
  return (
    <div
      className={`relative flex max-w-md items-start gap-4 rounded-xl border-2 border-darkblue bg-white p-6 text-darkblue shadow-sm ${className}`}
      role="group"
      aria-label={`Guide step ${step}: ${title}`}
    >
      {icon !== undefined && (
        <div className="absolute top-4 right-4 text-darkblue">{icon}</div>
      )}

      <div className="flex-shrink-0">
        <div className="flex h-14 w-14 items-center justify-center rounded-full bg-darkblue text-2xl font-bold text-white">
          {step}
        </div>
      </div>

      <div className="pr-9">
        <h3 className="mb-1 text-lg font-semibold">{title}</h3>
        <div className="text-sm text-gray-700">{children}</div>
      </div>
    </div>
  )
}
