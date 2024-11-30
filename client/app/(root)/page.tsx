"use client";
import React, { useEffect, useState } from "react";
import HeaderBox from "@/components/HeaderBox";
import TotalBalanceBox from "@/components/TotalBalanceBox";
import Navbar from "@/components/ui/navbar";
import Wallet from "@/components/Wallet";
import "../globals.css";
import { DetailService } from "@/components/DetailService";
import { Bill, BillData, User } from "@/types/type";
import { LineChart } from "recharts";
import LineChartComponent from "@/components/LineChart";

export default function Home() {
  const [user, setUser] = useState<User | null>(null);
  const [loggedIn, setLoggedIn] = useState<string>("Guest");
  const backendUrl = "http://localhost:8080/v1";
  const [chartData, setChartData] = useState<BillData>();

  const handleGetUserData = async (): Promise<void> => {
    const response = await fetch(`${backendUrl}/user/find`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
    });

    const data = await response.json();
    setUser(data);
  };

  const handleGetChartData = async (): Promise<void> => {
    const response = await fetch(`${backendUrl}/bill`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
    });

    const data = await response.json();
    setChartData(data);
    console.log('====================================');
    console.log(data);
    console.log('====================================');
  };

  useEffect(() => {
    handleGetUserData();
    handleGetChartData();
  }, []);

  useEffect(() => {
    if (user) {
      setLoggedIn(user?.UserName);
    }
  }, [user]);


  const handlePayNow = () => {};
  const handleHistory = () => {};
  return (
    <section className="home">
      <div className="home-content">
        <div className="home-header ">
          <Navbar />
          {/* <button className="" onClick={() => router.push()}/> */}
          <div className="ml-8">
            <HeaderBox
              type="greeting"
              title="Welcome"
              user={loggedIn}
              subtext="Access and manage your account"
            />
          </div>

          <div className="bg-white-50 flex justify-center p-8 ">
            <div className="flex justify-between w-full gap-10">
              <LineChartComponent chartData={chartData?.Bills || []} />
              <TotalBalanceBox
                months={chartData?.Bills || []}
                totalBanks={1}
                totalCurrentBalance={chartData?.YearTotal}
              />
              {/* <Wallet
                balance={70000}
                onPayNow={handlePayNow}
                onHistory={handleHistory}
              /> */}
            </div>
          </div>

          <div className="w-100% p-8 bg-gray-50 flex flex-col rounded-lg">
            <h1 className="flex justify-center text-3xl font-bold text-gray-800 py-8 ">
              Your Service (Registered)
            </h1>
            <DetailService />
          </div>
        </div>
      </div>
    </section>
  );
}

interface LineChartComponentProps {
  chartData: Bill[]; // Define the type for the chartData prop
}
