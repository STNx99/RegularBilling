"use client";

import { CartesianGrid, Line, LineChart, XAxis } from "recharts";

import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";
import { Bill } from "@/types/type";

const chartConfig = {
  desktop: {
    label: "Desktop",
    color: "hsl(var(--chart-1))",
  },
} satisfies ChartConfig;

interface LineChartProps {
  chartData: Bill[] | undefined;
}

const LineChartComponent: React.FC<LineChartProps> = ({ chartData }) => {
  const monthNames = [
    "January",
    "February",
    "March",
    "April",
    "May",
    "June",
    "July",
    "August",
    "September",
    "October",
    "November",
    "December",
  ];

  const monthlyData = monthNames.map((month, index) => {
    const billsForMonth = chartData?.filter((bill) => {
      const billDate = new Date(bill.Expired);
      return billDate.getMonth() === index;
    });
    const totalPrice =
      billsForMonth?.reduce((total, bill) => total + bill.Price, 0) || 0;

    return {
      month: month,
      totalPrice: totalPrice,
    };
  });

  return (
    <Card className="w-2/3">
      <CardHeader>
        <CardTitle>Line Chart</CardTitle>
        <CardDescription>January - December 2024</CardDescription>
      </CardHeader>
      <CardContent>
        <ChartContainer config={chartConfig}>
          <LineChart
            accessibilityLayer
            data={monthlyData}
            margin={{
              left: 0,
              right: 20,
            }}
          >
            <CartesianGrid vertical={false} />
            <XAxis
              dataKey="month"
              tickLine={true}
              axisLine={true}
              tickMargin={12}
              tickCount={12}
              tickFormatter={(value) => value}
            />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent hideLabel />}
            />
            <Line
              dataKey="totalPrice"
              type="natural"
              stroke="var(--color-desktop)"
              strokeWidth={3}
              dot={true}
            />
          </LineChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
};

export default LineChartComponent;
