"use client";

import React, { useState } from "react";
import "@fortawesome/fontawesome-free/css/all.min.css";
import Link from "next/link";
export default function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [showPassword, setShowPassword] = useState(false);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    console.log("Email: ", email);
    console.log("Password: ", password);
  };

  return (
    <div className="bg-gray-50 flex items-center justify-center min-h-screen">
      <div className="bg-white rounded-lg shadow-lg flex max-w-4xl">
        <div className="hidden md:flex flex-col items-center justify-center bg-blue-100 p-10 rounded-l-lg">
          <img
            src="/favicon.ico"
            alt="Regular Billing Illustration"
            className="w-60"
          />
        </div>

        <div className="flex-1 p-8">
          <h2 className="text-2xl font-bold text-blue-600 mb-6 text-center">
            Welcome to Regular Billing
          </h2>

          <form onSubmit={handleSubmit}>
            <div className="mb-4">
              <label className="block text-blue-700 text-xs mb-0.5 ml-0.5">
                Số Điện Thoại or Email
              </label>
              <input
                type="email"
                placeholder="Email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                className="p-1 w-full border-solid border-2 border-gray-300 rounded-md focus:border-blue-500 focus:ring-blue-500"
              />
            </div>

            <div className="mb-6 relative">
              <label className="block text-blue-700 text-xs mb-0.5 ml-0.5">
                Password
              </label>
              <div className="relative">
                <input
                  type={showPassword ? "text" : "password"}
                  placeholder="Password"
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                  className="p-1 w-full border-solid border-2 border-gray-300 rounded-md focus:border-blue-500 focus:ring-blue-500"
                />
                <button
                  type="button"
                  aria-label={
                    showPassword ? "Password Visible" : "Password Invisible."
                  }
                  className="text-black dark:text-white"
                  onClick={() => setShowPassword((prev) => !prev)}
                >
                  {showPassword ? (
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      fill="none"
                      viewBox="0 0 24 24"
                      strokeWidth="1.5"
                      stroke="currentColor"
                      className="w-6 select-none cursor-pointer absolute top-1.5 right-2"
                      tabIndex="-1"
                    >
                      <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        d="M2.036 12.322a1.012 1.012 0 010-.639C3.423 7.51 7.36 4.5 12 4.5c4.638 0 8.573 3.007 9.963 7.178.07.207.07.431 0 .639C20.577 16.49 16.64 19.5 12 19.5c-4.638 0-8.573-3.007-9.963-7.178z"
                      ></path>
                      <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
                      ></path>
                    </svg>
                  ) : (
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      fill="none"
                      viewBox="0 0 24 24"
                      strokeWidth="1.5"
                      stroke="currentColor"
                      className="w-6 select-none cursor-pointer absolute top-1.5 right-2"
                    >
                      <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        d="M3.98 8.223A10.477 10.477 0 001.934 12C3.226 16.338 7.244 19.5 12 19.5c.993 0 1.953-.138 2.863-.395M6.228 6.228A10.45 10.45 0 0112 4.5c4.756 0 8.773 3.162 10.065 7.498a10.523 10.523 0 01-4.293 5.774M6.228 6.228L3 3m3.228 3.228l3.65 3.65m7.894 7.894L21 21m-3.228-3.228l-3.65-3.65m0 0a3 3 0 10-4.243-4.243m4.242 4.242L9.88 9.88"
                      ></path>
                    </svg>
                  )}
                </button>
              </div>
            </div>

            <button
              type="submit"
              className="w-full bg-blue-300 text-white font-bold py-2 rounded-md hover:bg-blue-600"
            >
              Login
            </button>
            <Link href="/signup">
              <p className="text-sm mt-4 text-center text-blue-700 underline">
                You do not have an account ?
              </p>
            </Link>
          </form>

          <div className="text-xs mt-6 text-center text-blue-500">
            <div className="inline-flex items-center">
              <i className="fas fa-home mr-2 mb-1"></i>I agree to abide by
              templatana s Terms of Service and its Privacy Policy
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
