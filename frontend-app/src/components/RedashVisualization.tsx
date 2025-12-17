import { useEffect, useState } from 'react'
import { usePocketBase } from '../context/PocketBaseContext'
import type { RedashVisualizationType } from '../enum'

type RedashVisualization = {
  query: number
  visualization: number
  apiKey: string
}

export type RedashVisualizationProps = {
  visualizationName: RedashVisualizationType
}

export const RedashVisualization = ({
  visualizationName,
}: RedashVisualizationProps) => {
  const { pb } = usePocketBase()
  const [visualization, setVisualization] = useState<RedashVisualization>()

  useEffect(() => {
    pb.collection('redash_visualizations')
      .getFirstListItem(`name = "${visualizationName}"`)
      .then(redashVisualization =>
        setVisualization({
          query: redashVisualization.query,
          visualization: redashVisualization.visualization,
          apiKey: redashVisualization.api_key,
        })
      )
  }, [pb, visualizationName])

  return visualization !== undefined ? (
    <iframe
      src={`http://localhost:5001/embed/query/${visualization.query}/visualization/${visualization.visualization}?api_key=${visualization.apiKey}&hide_parameters&hide_header&hide_link&hide_timestamp`}
      width="720"
      height="391"
    ></iframe>
  ) : (
    <p>Loading...</p>
  )
}
