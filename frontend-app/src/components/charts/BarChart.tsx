import {
  BarChart as BarChartRC,
  Bar,
  XAxis,
  YAxis,
  Tooltip,
  Legend,
  ResponsiveContainer,
} from 'recharts'
import { InputSelect } from '../InputSelect'
import type { Control } from 'react-hook-form'
// TODO: Fix type of Control<any>
// import type { DashboardFormValues } from '../Dashboard'

export type BarChartDataItem = {
  year: string
  bar1percentage1?: number
  bar1value1?: number
  bar1name1?: string
  bar1percentage2?: number
  bar1value2?: number
  bar1name2?: string
  bar1percentage3?: number
  bar1value3?: number
  bar1name3?: string
  bar2percentage1?: number
  bar2value1?: number
  bar2name1?: string
  bar2percentage2?: number
  bar2value2?: number
  bar2name2?: string
  bar2percentage3?: number
  bar2value3?: number
  bar2name3?: string
  bar1unit?: string
  bar2unit?: string
}

export type BarChartProps = {
  data: BarChartDataItem[]
  name: string
  title: string
  control?: Control<any>
  className?: string
  dataOptions?: { id: string; name: string }[]
  comparisonOptions?: { id: string; name: string }[]
  bar1color1?: string
  bar1color2?: string
  bar1color3?: string
  bar2color1?: string
  bar2color2?: string
  bar2color3?: string
}

const formatToolTipItems = (index: number, entry: any) => {
  const dataKey = entry.dataKey as string
  let percentage: string | undefined
  let value: string | undefined
  let unit: string | undefined

  // Determine which unit to use based on bar1 or bar2
  if (dataKey.startsWith('bar1')) {
    unit = entry.payload.bar1unit
  } else if (dataKey.startsWith('bar2')) {
    unit = entry.payload.bar2unit
  }

  if (dataKey.includes('percentage')) {
    percentage = entry.payload[dataKey]
    value = entry.payload[dataKey.replace('percentage', 'value')]
  } else {
    percentage = undefined
    value = entry.payload[dataKey]
  }

  return (
    <p key={index} style={{ color: entry.color }} className="text-sm">
      {entry.name}:{' '}
      {percentage !== undefined && percentage !== null ? (
        <>
          <b>{percentage}%</b>
          {value !== undefined &&
            value !== null &&
            unit &&
            ` (${value} ${unit})`}
        </>
      ) : (
        <>
          <b>{value}</b>
          {unit && ` ${unit}`}
        </>
      )}
    </p>
  )
}

// Custom tooltip formatter
// TODO: Fix any's
const CustomTooltip = ({
  active,
  payload,
  label,
}: {
  active: boolean
  payload: any[]
  label: string
}) => {
  if (active && payload && payload.length) {
    const bar1Items = payload.filter((entry: any) =>
      entry.dataKey.startsWith('bar1')
    )
    const bar2Items = payload.filter((entry: any) =>
      entry.dataKey.startsWith('bar2')
    )

    return (
      <div className="rounded-lg border border-darkblue bg-white p-3 shadow-lg">
        <p className="mb-2 font-bold text-darkblue">{label}</p>

        {bar1Items.map((entry: any, index: number) => {
          return formatToolTipItems(index, entry)
        })}

        {/* Divider */}
        {bar1Items.length > 0 && bar2Items.length > 0 && (
          <div className="my-2 border-t border-darkblue/20" />
        )}

        {bar2Items.map((entry: any, index: number) => {
          return formatToolTipItems(index, entry)
        })}
      </div>
    )
  }
  return null
}

// Custom legend renderer
const renderCustomLegend = (props: any) => {
  const { payload } = props

  const bar1Items = payload.filter((entry: any) =>
    entry.dataKey.startsWith('bar1')
  )
  const bar2Items = payload.filter((entry: any) =>
    entry.dataKey.startsWith('bar2')
  )

  return (
    <div className="flex flex-col gap-1">
      {/* Bar1 items */}
      {bar1Items.map((entry: any, index: number) => (
        <div key={index} className="flex items-center gap-2">
          <div className="h-3 w-3" style={{ backgroundColor: entry.color }} />
          <span className="text-sm text-darkblue">{entry.value}</span>
        </div>
      ))}

      {/* Divider */}
      {bar1Items.length > 0 && bar2Items.length > 0 && (
        <div className="my-1 border-t border-darkblue/20" />
      )}

      {/* Bar2 items */}
      {bar2Items.map((entry: any, index: number) => (
        <div key={index} className="flex items-center gap-2">
          <div className="h-3 w-3" style={{ backgroundColor: entry.color }} />
          <span className="text-sm text-darkblue">{entry.value}</span>
        </div>
      ))}
    </div>
  )
}

export const BarChart = ({
  data,
  name,
  title,
  control,
  dataOptions,
  comparisonOptions,
  className = '',
  bar1color1 = '#77ACC3',
  bar1color2 = '#6E738D',
  bar1color3 = '#17394F',
  bar2color1 = '#004138',
  bar2color2 = '#6E8D88',
  bar2color3 = '#026557',
}: BarChartProps) => {
  // Color mapping
  const colorMap: Record<string, string> = {
    bar1percentage1: bar1color1,
    bar1percentage2: bar1color2,
    bar1percentage3: bar1color3,
    bar2percentage1: bar2color1,
    bar2percentage2: bar2color2,
    bar2percentage3: bar2color3,
    bar1value1: bar1color1,
    bar1value2: bar1color2,
    bar1value3: bar1color3,
    bar2value1: bar2color1,
    bar2value2: bar2color2,
    bar2value3: bar2color3,
  }

  // Dynamically extract bar configuration from first data item
  const barConfigs = data[0]
    ? Object.keys(data[0])
        .filter(
          key =>
            key.startsWith('bar') &&
            (key.includes('percentage') || key.includes('value'))
        )
        .filter(key => {
          // Only include 'value' keys if there's no corresponding 'percentage' key
          if (key.includes('value')) {
            const percentageKey = key.replace('value', 'percentage')
            return !Object.keys(data[0]).includes(percentageKey)
          }
          return true
        })
        .map(dataKey => {
          const nameKey = dataKey
            .replace('percentage', 'name')
            .replace('value', 'name')
          const stackId = dataKey.startsWith('bar1') ? 'bar1' : 'bar2'
          return {
            dataKey,
            stackId,
            name: data[0][nameKey as keyof BarChartDataItem] as string,
          }
        })
        .sort((a, b) => a.name.localeCompare(b.name))
    : []

  return (
    <div className="w-full">
      <h3 className="mb-4 text-xl font-semibold text-darkblue">{title}</h3>
      <div className={`flex h-[400px] w-full flex-col gap-4 ${className}`}>
        {/* Dropdowns */}
        <div className="flex gap-4">
          {/* Data Dropdown */}
          {control !== undefined && dataOptions !== undefined && (
            <div className="flex flex-col gap-1">
              <InputSelect
                control={control}
                label="Data"
                options={dataOptions.map(option => ({
                  id: option.id,
                  name: option.name,
                }))}
                name={name + 'Data'}
                initialSelectedId={Object.values(dataOptions)[0].id}
              />
            </div>
          )}

          {/* Comparison Dropdown */}
          {control !== undefined &&
            comparisonOptions !== undefined &&
            comparisonOptions[0] !== undefined && (
              <div className="flex flex-col gap-1">
                <InputSelect
                  control={control}
                  label="Sammenlign med"
                  options={comparisonOptions.map(option => ({
                    id: option.id,
                    name: option.name,
                  }))}
                  name={name + 'Comparison'}
                  initialSelectedId={comparisonOptions[0].id}
                />
              </div>
            )}
        </div>

        {/* Chart */}
        <div className="flex-1">
          <ResponsiveContainer width="100%" height="100%">
            <BarChartRC
              data={data}
              margin={{ top: 12, right: 12, left: 0, bottom: 0 }}
            >
              <XAxis
                dataKey="year"
                tick={{ fill: '#004138' }}
                axisLine={false}
              />
              <YAxis hide={true} />
              <Tooltip
                content={
                  <CustomTooltip active={false} payload={[]} label={''} />
                }
              />
              <Legend
                layout="vertical"
                align="right"
                verticalAlign="middle"
                content={renderCustomLegend}
                wrapperStyle={{ paddingLeft: '10px' }}
              />
              {barConfigs.map(config => (
                <Bar
                  key={config.dataKey}
                  dataKey={config.dataKey}
                  stackId={config.stackId}
                  fill={colorMap[config.dataKey]}
                  name={config.name}
                />
              ))}
            </BarChartRC>
          </ResponsiveContainer>
        </div>
      </div>
    </div>
  )
}
