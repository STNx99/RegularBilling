import React, { useEffect, useState } from "react";

type WalletProps = {
  balance: number;
  onPayNow: () => void;
  onHistory: () => void;
};

const Wallet: React.FC<WalletProps> = ({ balance, onPayNow, onHistory }) => {
  const [formattedBalance, setFormattedBalance] = useState("");

  useEffect(() => {
    setFormattedBalance(balance.toLocaleString() + " VND");
  }, [balance]);

  return (
    <div className="items-center bg-white rounded-lg p-6 shadow-md">
      <div className="flex flex-col items-center ">
        <p className="text-xl font-bold">WALLET</p>
        <p className="text-2xl text-teal-500 font-bold my-2">
          {formattedBalance}
        </p>
        <button
          onClick={onPayNow}
          className="bg-teal-500 text-white px-4 py-2 mb-2 rounded-md  hover:bg-teal-800"
        >
          PAY NOW
        </button>
        <button
          onClick={onHistory}
          className="bg-gray-200 text-gray-700 px-4 py-2  rounded-md  hover:bg-gray-400"
        >
          HISTORY
        </button>
      </div>
    </div>
  );
};

export default Wallet;
