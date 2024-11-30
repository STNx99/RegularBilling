"use client"

import { CartesianGrid, Line, LineChart, XAxis } from "recharts"

import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart"
import { Bill } from "@/types/type"
const chartData = [
  { month: "January", price: 186 },
  { month: "February", price: 305 },
  { month: "March", price: 237 },
  { month: "April", price: 73 },
  { month: "May", price: 209 },
  { month: "June", price: 214 },
  { month: "July", price: 100 },
  { month: "August", price: 200 },
  { month: "September", price: 500 },
  { month: "October", price: 50 },
  { month: "November", price: 200 },
  { month: "December", price: 200  },

]

const chartConfig = {
  desktop: {
    label: "Desktop",
    color: "hsl(var(--chart-1))",
  },
} satisfies ChartConfig

export function LineChartComponent(chartData : Bill[]) {
  return (
    <Card className="w-2/3">
      <CardHeader>
        <CardTitle>Line Chart</CardTitle>
        <CardDescription>January - June 2024</CardDescription>
      </CardHeader>
      <CardContent>
        <ChartContainer config={chartConfig}>
          <LineChart
            accessibilityLayer
            data={chartData}
            margin={{
              left: 12,
              right: 12,
            }}
          >
            <CartesianGrid vertical={false} />
            <XAxis
              dataKey="month"
              tickLine={false}
              axisLine={false}
              tickMargin={8}
              tickFormatter={(value) => value.slice(0, 3)}
            />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent hideLabel />}
            />
            <Line
              dataKey="price"
              type="natural"
              stroke="var(--color-desktop)"
              strokeWidth={2}
              dot={false}
            />
          </LineChart>
        </ChartContainer>
      </CardContent>
    </Card>
  )
}
