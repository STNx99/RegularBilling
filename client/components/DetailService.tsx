'use client'

import React, { useState } from 'react';

interface DetailServiceProps {
  onSelect: (serviceData: ServiceData) => void;
}

interface ServiceData {
  serviceName: string;
  price: number;
  duration: string;
  registrationDate: string;
}

const DetailService: React.FC<DetailServiceProps> = ({ onSelect }) => {
  const [serviceName, setServiceName] = useState<string>('youtube');
  const [price, setPrice] = useState<number>(0);
  const [duration, setDuration] = useState<string>('1 tháng');
  const [registrationDate, setRegistrationDate] = useState<string>('');

  const serviceOptions = ['youtube', 'spotify', 'netflix', 'zingmp3', 'Xbox game pass', 'FPT play'];
  const durationOptions = ['1 tháng', '3 tháng', '6 tháng', '12 tháng'];

  const handleSubmit = () => {
    const serviceData: ServiceData = {
      serviceName,
      price,
      duration,
      registrationDate,
    };
    onSelect(serviceData);
  };

  return (
    <div className="p-4 bg-white border rounded-lg shadow-md w-80">
      <h2 className="text-xl font-semibold text-center mb-4">Thêm dịch vụ</h2>
      <hr className="mb-4" />

      <div className="mb-4">
        <label htmlFor="serviceName" className="block text-gray-700">Chọn tên dịch vụ</label>
        <select
          id="serviceName"
          value={serviceName}
          onChange={(e) => setServiceName(e.target.value)}
          className="w-full p-2 border border-gray-300 rounded mt-1"
        >
          {serviceOptions.map((option, index) => (
            <option key={index} value={option}>{option}</option>
          ))}
        </select>
      </div>

      <div className="mb-4">
        <label htmlFor="price" className="block text-gray-700">Giá dịch vụ</label>
        <input
          id="price"
          type="number"
          value={price}
          onChange={(e) => setPrice(Number(e.target.value))}
          className="w-full p-2 border border-gray-300 rounded mt-1"
          placeholder="Nhập giá dịch vụ"
        />
      </div>

      <div className="mb-4">
        <label htmlFor="duration" className="block text-gray-700">Thời hạn gói đăng ký</label>
        <select
          id="duration"
          value={duration}
          onChange={(e) => setDuration(e.target.value)}
          className="w-full p-2 border border-gray-300 rounded mt-1"
        >
          {durationOptions.map((option, index) => (
            <option key={index} value={option}>{option}</option>
          ))}
        </select>
      </div>

      <div className="mb-4">
        <label htmlFor="registrationDate" className="block text-gray-700">Ngày đăng ký dịch vụ</label>
        <input
          id="registrationDate"
          type="date"
          value={registrationDate}
          onChange={(e) => setRegistrationDate(e.target.value)}
          className="w-full p-2 border border-gray-300 rounded mt-1"
        />
      </div>

      <div className="text-center">
        <button
          onClick={handleSubmit}
          className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
        >
          Thêm dịch vụ
        </button>
      </div>
    </div>
  );
};

export default DetailService;
