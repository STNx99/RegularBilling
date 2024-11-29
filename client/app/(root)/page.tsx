"use client";
import React, { useEffect, useState } from "react";
import HeaderBox from "@/components/HeaderBox";
import TotalBalanceBox from "@/components/TotalBalanceBox";
import Navbar from "@/components/ui/navbar";
import Wallet from "@/components/Wallet";
import "../globals.css";
import { DetailService } from "@/components/DetailService";
import { User } from "@/types/type";

export default function Home()  {
  const [user, setUser] = useState<User | null>(null);
  // const [user, setUser] = useState();
  const [loggedIn, setLoggedIn] = useState<string>("Guest");
  const backendUrl = "http://localhost:8080/v1/user";
  
  const handleGetUserData = async (): Promise<void> => {
    const response = await fetch(`${backendUrl}/find`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include', 
    });
  
    const data = await response.json();
    setUser(data);
  };

  useEffect(() => {
    handleGetUserData();
  },[]);

  useEffect(() => {
    if (user) {
      setLoggedIn(user.userName);
    }
  }, [user]);

  console.log(user);

  const handlePayNow = () => {};
  const handleHistory = () => {};
  return (
    <section className="home">
      <div className="home-content">
        <div className="home-header ">
          <Navbar />
          <div className="ml-8">
            <HeaderBox
              type="greeting"
              title="Welcome"
              user={loggedIn}
              subtext="Access and manage your account"
            />
          </div>

          <div className="bg-white-50 flex justify-center p-8 ">
            <div className="flex justify-between w-full">
              <TotalBalanceBox
                accounts={[]}
                totalBanks={1}
                totalCurrentBalance={0}
              />
              <Wallet
                balance={70000}
                onPayNow={handlePayNow}
                onHistory={handleHistory}
              />
            </div>
          </div>

          <div className="w-100% p-8 bg-gray-50 flex flex-col rounded-lg">
            <h1 className="flex justify-center text-3xl font-bold text-gray-800 py-8 ">
              Your Service (Registered)
            </h1>
            <DetailService/>
          </div>
        </div>
      </div>
    </section>
  );
}
