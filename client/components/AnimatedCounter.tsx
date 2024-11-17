'use client';

import CountUp from 'react-countup';

const AnimatedCounter = ({ amount }: { amount: number }) => {
  return (
    <div className="w-full">
      <CountUp 
        start={1000}
        end={amount}
        duration={1.5}
      />đ
    </div>
  )
}

export default AnimatedCounter