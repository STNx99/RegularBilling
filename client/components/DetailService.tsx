import {Table, TableHeader, TableColumn, TableBody, TableRow, TableCell} from "@nextui-org/react";

export default function DetailService() {
  return (
    <Table className="text-left">
      <TableHeader className="font-bold">
        <TableColumn>NAME SERVICE</TableColumn>
        <TableColumn>DATE REGISTERED</TableColumn>
        <TableColumn>PRICE</TableColumn>
      </TableHeader>
      <TableBody>
        <TableRow key="1">
          <TableCell>Youtube Premium</TableCell>
          <TableCell>11/11/2024</TableCell>
          <TableCell>79000</TableCell>
        </TableRow>
        <TableRow key="2">
          <TableCell>Spotify</TableCell>
          <TableCell>02/11/2024</TableCell>
          <TableCell>29500</TableCell>
        </TableRow>
        <TableRow key="3">
          <TableCell>Netflixr</TableCell>
          <TableCell>04/11/2024</TableCell>
          <TableCell>79000</TableCell>
        </TableRow>
        <TableRow key="4">
          <TableCell>OpenAI</TableCell>
          <TableCell>01/11/2024</TableCell>
          <TableCell>250000</TableCell>
        </TableRow>
      </TableBody>
    </Table>
  );
}