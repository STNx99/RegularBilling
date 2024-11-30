"use client";

import { Bill } from "@/types/type";
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from "chart.js";
import { Doughnut } from "react-chartjs-2";

ChartJS.register(ArcElement, Tooltip, Legend);

interface DoughNutChartProps {
  doughNutData: Bill[] | undefined;
}

const DoughnutChart: React.FC<DoughNutChartProps> = ({ doughNutData }) => {
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

  const monthTotals = monthNames.map((month, index) => {
    const billsForMonth = doughNutData?.filter((bill) => {
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

  const data = {
    labels: monthTotals.map((item) => item.month),
    datasets: [
      {
        label: "Monthly Expenses",
        data: monthTotals.map((item) => item.totalPrice),
        backgroundColor: monthTotals.map((_, index) => {
          const colors = [
            "#FF5733",
            "#FF8D1A",
            "#FFBD33",
            "#A1D8F9",
            "#5D9C59",
            "#C23A2B",
            "#B25D3A",
            "#715CC1",
            "#C36F91",
            "#3B3B98",
            "#006A6A",
            "#803B6F",
          ];
          return colors[index % colors.length];
        }),
      },
    ],
  };

  return (
    <Doughnut
      data={data}
      options={{
        cutout: "60%",
        plugins: {
          legend: {
            display: false,
            position: "bottom",
          },
          tooltip: {
            callbacks: {
              label: (tooltipItem) => {
                const month = monthTotals[tooltipItem.dataIndex].month;
                const total = tooltipItem.raw as number;
                return `${month}: $${total.toFixed(2)}`;
              },
            },
          },
        },
      }}
    />
  );
};

export default DoughnutChart;
