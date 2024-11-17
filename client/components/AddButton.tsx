'use client'

import { useState } from 'react';

const AddButton = () => {

    const [showBox, setShowBox] = useState(false);

    const handleClick = () => {
      setShowBox(!showBox);
    }

    const handleCloseBox = () => {
        setShowBox(false);
      }

  
    return (
        <div className="relative flex justify-center items-center">
          <button 
            onClick={handleClick}
            className="text-2xl bg-blue-600 text-white border-none rounded px-5 py-2.5 cursor-pointer"
          >
            +
          </button>
    
            {showBox && 
            <div className="absolute mt-2 p-2 bg-gray-100 border border-gray-300 rounded w-48">
                <button 
                    onClick={handleCloseBox}
                    className="absolute top-1 right-2 text-lg text-gray-500 hover:text-gray-800"
                >
                x
                </button>
            Box service
            </div>}
        </div>
      );
}

export default AddButton
