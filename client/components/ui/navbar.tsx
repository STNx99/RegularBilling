"use client";
import { useRouter } from "next/navigation";
import { useState } from "react";

const Navbar: React.FC = () => {
  // const [mobileMenuOpen, setMobileMenuOpen] = useState(false);
  const [isMenuOpen, setIsMenuOpen] = useState(false);
  const router = useRouter();
  const toggleMenu = () => {
    setIsMenuOpen((prev) => !prev);
  };
  const backendUrl = "http://localhost:8080/v1";

  const handleLogOut = async (): Promise<void> => {
    try {
      const response = await fetch(`${backendUrl}/user/logout`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include",
      });
      if (!response.ok) {
        throw new Error(`HTTP status ${response.status}`);
      }

      router.push("/login");
    } catch (error) {
      console.log("Error : ", error);
      alert("Loi dang xuat");
    }
  };

  return (
    <nav className="bg-white-500 w-full rounded-lg border-2">
      <div className="relative flex h-16 items-center justify-between">
        <div className="flex items-center justify-center">
          <div className="ml-auto shrink-0 items-center">
            <img
              className="h-8 ml-5"
              src="https://tailwindui.com/plus/img/logos/mark.svg?color=indigo&shade=500"
              alt="Logo"
            />
          </div>

          <div className="sm:block sm:ml-6">
            <div className="flex space-x-4">
              <a
                href="../layout"
                className="text-gray-500 hover:bg-blue-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium"
                aria-current="page"
              >
                Home
              </a>
              <a
                href="#"
                className="text-gray-500 hover:bg-blue-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium"
              >
                Dashboard
              </a>
              <a
                href="#"
                className="text-gray-500 hover:bg-blue-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium"
              >
                Team
              </a>
              <a
                href="../login"
                className="text-gray-500 hover:bg-blue-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium"
              >
                Login
              </a>
            </div>
          </div>
        </div>

        <div className="absolute inset-y-0 right-0 flex items-center pr-2 sm:static sm:inset-auto sm:ml-6 sm:pr-0">
          <button className="relative rounded-full p-1 text-gray-800 hover:text-white focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800">
            <span className="sr-only">View notifications</span>
            <svg
              className="w-6 h-6"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
              aria-hidden="true"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="M14.857 17.082a23.848 23.848 0 0 0 5.454-1.31A8.967 8.967 0 0 1 18 9.75V9A6 6 0 0 0 6 9v.75a8.967 8.967 0 0 1-2.312 6.022c1.733.64 3.56 1.085 5.455 1.31m5.714 0a24.255 24.255 0 0 1-5.714 0m5.714 0a3 3 0 1 1-5.714 0"
              />
            </svg>
          </button>

          <div className="relative ml-3">
            <button 
              className="flex text-sm rounded-full bg-blue-300 p-1 mr-5"
              onClick={toggleMenu}>
              <img
                className="h-8 w-8  rounded-full"
                src="https://cdn-icons-png.flaticon.com/512/1698/1698743.png"
                alt="User"
              />
            </button>
            {isMenuOpen && (
              <div className="absolute right-0 mt-2 w-40 bg-white rounded-lg shadow-lg z-10">
              <button 
                className="block w-full px-4 py-2 text-left text-gray-700 hover:bg-blue-100"
                onClick={() => alert("Cài đặt")}
              >
                Cài đặt
              </button>
              <button
                className="block w-full px-4 py-2 text-left text-gray-700 hover:bg-blue-100"
                onClick={() => handleLogOut()}
              >
                Logout
              </button>
        </div>
      )}
          </div>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
