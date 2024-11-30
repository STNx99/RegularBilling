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
import { Service } from "@/types/type";
import { useEffect, useState } from "react";

export function DetailService() {
  const backendUrl = "http://localhost:8080/v1/service/";
  const [services, setServices] = useState<Service[]>([]); 

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
        setServices(data);
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
          <TableHead>TEN DICH VU</TableHead>
          <TableHead>NGAY DANG KY</TableHead>
          <TableHead>NGAY HET HAN</TableHead>
          <TableHead className="text-right">GIA CA</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {services.map((service) => (
          <TableRow key={service.ServiceId.toString()}>
            <TableCell className="font-medium">{}</TableCell>
            <TableCell>{service.ServiceName}</TableCell>
            <TableCell>{formattedDateTime(service.CreatedAt)}</TableCell>
            <TableCell>{formattedDateTime(service.ExpireAt)}</TableCell>
            <TableCell className="text-right">{service.Price}</TableCell>
          </TableRow>
        ))}
      </TableBody>
      <TableFooter>
        <TableRow>
          <TableCell colSpan={3}>Total</TableCell>
          <TableCell className="text-right">$2,500.00</TableCell>
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