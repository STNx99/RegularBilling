import { TotalBalanceBoxProps } from '@/types';
import AnimatedCounter from './AnimatedCounter';
import DoughnutChart from './DoughnutChart';

const TotalBalanceBox = ({
  months = [],
  totalBanks,
  totalCurrentBalance,
}: TotalBalanceBoxProps) => {
  return (
    <section className="total-balance">
      <div className="total-balance-chart">
        <DoughnutChart doughNutData={months} />
      </div>

      <div className="flex flex-col gap-6">
        <h2 className="header-2">Total Spending This Year</h2>
        <div className="flex flex-col gap-2">
          <p className="total-balance-label">Total Spent This Year</p>

          <div className="total-balance-amount flex-center gap-2">
            <AnimatedCounter amount={totalCurrentBalance || 0} />
          </div>
        </div>
      </div>
    </section>
  );
};

export default TotalBalanceBox