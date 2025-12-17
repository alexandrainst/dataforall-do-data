type MaskChartProps = {
  title: string
  data: MaskChartData[]
  className?: string
}

type MaskChartData = {
  percentage: number
  year: string
  fillColor?: string
  backgroundColor?: string
}

const renderInnerSvg = (
  fillColor: string | undefined,
  clipPath: string | undefined
) => {
  return (
    <g fill={fillColor} clipPath={clipPath}>
      <circle cx="3326.4" cy="3345.28" r="2083.56" />
      <path d="M3424.61,188.099c-0,-51.392 -41.724,-93.115 -93.116,-93.115c-51.391,-0 -93.115,41.723 -93.115,93.115l0,690.397c0,51.392 41.724,93.116 93.115,93.116c51.392,-0 93.116,-41.724 93.116,-93.116l-0,-690.397Z" />
      <path d="M3417.19,5781.43c-0,-51.392 -41.724,-93.115 -93.116,-93.115c-51.391,-0 -93.115,41.723 -93.115,93.115l0,690.397c0,51.392 41.724,93.116 93.115,93.116c51.392,-0 93.116,-41.724 93.116,-93.116l-0,-690.397Z" />
      <path d="M185.923,3233.13c-51.391,-0 -93.115,41.723 -93.115,93.115c0,51.392 41.724,93.115 93.115,93.115l690.398,0c51.391,0 93.115,-41.723 93.115,-93.115c-0,-51.392 -41.724,-93.115 -93.115,-93.115l-690.398,-0Z" />
      <path d="M5779.25,3240.56c-51.391,-0 -93.115,41.724 -93.115,93.115c0,51.392 41.724,93.115 93.115,93.115l690.398,0c51.391,0 93.115,-41.723 93.115,-93.115c-0,-51.391 -41.724,-93.115 -93.115,-93.115l-690.398,-0Z" />
      <path d="M1037.69,5483.12c-36.339,36.34 -36.339,95.346 0,131.685c36.34,36.34 95.346,36.34 131.685,0l488.185,-488.184c36.339,-36.34 36.339,-95.346 -0,-131.685c-36.34,-36.339 -95.346,-36.339 -131.685,-0l-488.185,488.184Z" />
      <path d="M4998.01,1533.3c-36.339,36.339 -36.339,95.345 0,131.684c36.34,36.34 95.346,36.34 131.685,0l488.184,-488.184c36.34,-36.339 36.34,-95.345 0,-131.685c-36.339,-36.339 -95.345,-36.339 -131.684,0l-488.185,488.185Z" />
      <path d="M5480.95,5620.06c36.339,36.34 95.345,36.34 131.685,0c36.339,-36.339 36.339,-95.345 -0,-131.685l-488.185,-488.184c-36.339,-36.34 -95.345,-36.34 -131.685,-0c-36.339,36.339 -36.339,95.345 0,131.685l488.185,488.184Z" />
      <path d="M1531.12,1659.73c36.339,36.339 95.345,36.339 131.685,-0c36.339,-36.34 36.339,-95.346 -0,-131.685l-488.185,-488.185c-36.339,-36.339 -95.345,-36.339 -131.685,0c-36.339,36.34 -36.339,95.346 0,131.685l488.185,488.185Z" />
    </g>
  )
}

export const MaskChart = ({ data, title, className = '' }: MaskChartProps) => {
  const displayData = data.slice(-3)

  const renderDataPoint = (dp: MaskChartData, index: number) => {
    const {
      percentage,
      year,
      fillColor = '#fdb913',
      backgroundColor = '#ebebeb',
    } = dp
    const fillHeight = percentage
    const yPosition = 100 - fillHeight

    const size = 360 - displayData.length * 60

    return (
      <div className="relative flex flex-col items-center">
        {/* Content taken from sun.svg in assets */}
        <svg
          width={size}
          height={size}
          viewBox="0 0 6667 6667"
          xmlns="http://www.w3.org/2000/svg"
          style={{ overflow: 'visible' }}
        >
          <defs>
            {/* Define clip path that reveals the fill color from bottom up */}
            <clipPath id={`fillClip-${year}-${index}`}>
              <rect
                x="0"
                y={`${yPosition}%`}
                width="100%"
                height={`${fillHeight}%`}
              />
            </clipPath>
          </defs>

          {/* Background sun (gray) */}
          {renderInnerSvg(backgroundColor, undefined)}

          {/* Filled sun (colored) - clipped to show percentage */}
          {renderInnerSvg(fillColor, `url(#fillClip-${year}-${index})`)}
        </svg>

        {/* Percentage text - centered in the SVG */}
        <div className="absolute inset-0 -mt-10 flex items-center justify-center">
          <div className="text-center">
            <p className="text-3xl font-bold text-darkblue">{percentage}%</p>
          </div>
        </div>

        {/* Year text below */}
        <div className="mt-4 text-center">
          <p className="text-xl font-bold text-darkblue">{year}</p>
        </div>
      </div>
    )
  }

  return (
    <div className="w-full">
      <h3 className="mb-4 text-xl font-semibold text-darkblue">{title}</h3>
      <div
        className={`flex h-[400px] w-full flex-row items-center justify-center gap-4 ${className}`}
      >
        {displayData.map((dataPoint, index) => (
          <div key={index}>{renderDataPoint(dataPoint, index)}</div>
        ))}
      </div>
    </div>
  )
}
