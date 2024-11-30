// 'use client'

// import { useState } from 'react';
// import DetailService from './DetailService';

// interface ServiceData {
//   serviceName: string;
//   price: number;
//   duration: string;
//   registrationDate: string;
// }

// const AddButton = () => {
//   const [showBox, setShowBox] = useState(false);
//   const [selectedService, setSelectedService] = useState<ServiceData | null>(null); 

//   const handleClick = () => {
//     setShowBox(!showBox);
//   };

//   const handleCloseBox = () => {
//     setShowBox(false);
//   };

//   const handleSelectService = (serviceData: ServiceData) => {
//     setSelectedService(serviceData);
//     setShowBox(false); 
//   };

//   return (
//     <div className="relative flex justify-center items-center">
//       <button
//         onClick={handleClick}
//         className="text-2xl bg-blue-600 text-white border-none rounded px-5 py-2.5 cursor-pointer"
//       >
//         +
//       </button>

//       {showBox && (
//         <div className="absolute mt-2 p-2 bg-gray-100 border border-gray-300 rounded w-80">
//           <button
//             onClick={handleCloseBox}
//             className="absolute top-1 right-2 text-lg text-gray-500 hover:text-gray-800"
//           >
//             x
//           </button>
//           <DetailService onSelect={handleSelectService} />
//         </div>
//       )}

//       {selectedService ? (
//         <div className="mt-4 text-center text-lg">
//           <h3>Dịch vụ bạn đã chọn:</h3>
//           <ul>
//             <li>Tên dịch vụ: {selectedService.serviceName}</li>
//             <li>Giá dịch vụ: {selectedService.price} VND</li>
//             <li>Thời hạn: {selectedService.duration}</li>
//             <li>Ngày đăng ký: {selectedService.registrationDate}</li>
//           </ul>
//         </div>
//       ) : (
//         <div className="mt-4 text-center text-lg">Chưa chọn dịch vụ</div>
//       )}
//     </div>
//   );
// };

// export default AddButton;
