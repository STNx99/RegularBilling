import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableFooter,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table"
import { Service, ServiceData } from "@/types/type";
import { useEffect, useState } from "react";

export function DetailService() {
  const backendUrl = "http://localhost:8080/v1/service/";
  const [serviceData, setServiceData] = useState<ServiceData>(); 
  const handleGetService = async(): Promise<void> => {
    try {
      const respones = await fetch(`${backendUrl}`,{
        method: 'GET',
        credentials: 'include'
      });
  
      const data = await respones.json()
      console.log(data);
      if(data){
        console.log(data);
        setServiceData(data);
      }


    } catch (error: unknown) {
      console.log("Error: ", error);
    }
  }

  useEffect(() => {
    handleGetService();
  },[]);

  return (
    <Table>
      <TableCaption>A list of your recent invoices.</TableCaption>
      <TableHeader>
        <TableRow>
          <TableHead className="w-[100px]">STT</TableHead>
          <TableHead>SERVICE NAME</TableHead>
          <TableHead>CREATE DATE</TableHead>
          <TableHead>EXPIRATION DATE</TableHead>
          <TableHead className="text-right">PRICE</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {serviceData?.Services.map((service,index) => (
          <TableRow key={service.ServiceId.toString()}>
            <TableCell className="font-medium">{index+1}</TableCell>
            <TableCell>{service.ServiceName}</TableCell>
            <TableCell>{formattedDateTime(service.CreatedAt)}</TableCell>
            <TableCell>{formattedDateTime(service.ExpireAt)}</TableCell>
            <TableCell className="text-right">{service.Price}</TableCell>
          </TableRow>
        ))}
      </TableBody>
      <TableFooter>
        <TableRow>
          <TableCell colSpan={3} className="font-bold">Service Total</TableCell>
          <TableCell className="text-right">{serviceData?.ServiceTotal}</TableCell>
        </TableRow>
      </TableFooter>
    </Table>
  )
}


function formattedDateTime(date: Date | string) : string {
  const dateObj = typeof date === "string" ?  new Date(date): date;
  return dateObj.toLocaleString("en-Us", {
    day: "2-digit",
    month: "2-digit",
    year: "numeric",
  });
}