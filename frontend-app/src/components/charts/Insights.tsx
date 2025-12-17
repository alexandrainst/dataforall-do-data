import { LightBulbIcon } from '@heroicons/react/24/outline'
//import { ArrowTopRightOnSquareIcon } from '@heroicons/react/24/outline'
import { DDButton } from '../controls/DDButton'
import { useState, type ReactNode } from 'react'

type Insight = {
  title: string
  body: string
  chart?: ReactNode
}
type InsightsProps = {
  insights: Insight[]
}

export const Insights = ({ insights }: InsightsProps) => {
  const [currentIndex, setCurrentIndex] = useState(0)
  const insight = insights[currentIndex]

  return (
    <div className="w-full rounded-lg bg-frost p-6 text-darkblue shadow-sm lg:w-[100%]">
      <div className="flex flex-col gap-8 lg:flex-row">
        <div className="flex w-[100%] flex-col text-justify lg:w-[45%]">
          <div className="mb-4 flex items-center gap-4">
            <LightBulbIcon className="h-12 w-12" />
            <h4 className="mb-4 font-semibold">{insight.title}</h4>
          </div>
          <div className="flex grow-1 flex-col justify-between">
            <div
              className="w-full"
              dangerouslySetInnerHTML={{ __html: insight.body }}
            />
            <div className="flex gap-2">
              <DDButton
                className="mt-4"
                onClick={() => {
                  if (currentIndex > 0) {
                    setCurrentIndex(currentIndex - 1)
                  }
                }}
                disabled={currentIndex === 0}
              >
                Forrige
              </DDButton>
              <DDButton
                className="mt-4"
                onClick={() => {
                  if (currentIndex < insights.length - 1) {
                    setCurrentIndex(currentIndex + 1)
                  }
                }}
                disabled={currentIndex === insights.length - 1}
              >
                Næste
              </DDButton>
            </div>
          </div>
        </div>
        <div className="w-[100%] rounded-lg bg-white p-6 shadow-sm lg:w-[55%]">
          {insight.chart}
        </div>
      </div>
    </div>
  )
}
