import { ObjectId } from "mongodb";

export type User = {
  UserId: ObjectId;
  UserName: string;
  Email: string;
  Password: string;
  Credits: number;
  CreatedAt?: Date;
  UpdatedAt?: Date;
  Bills: Bill[];
  ServiceIds: Service[];
};

export type Service = {
  ServiceId: ObjectId;
  ServiceName: string;
  Price: number;
  CreatedAt: Date;
  ExpireAt: Date;
};

export type Bill = {
  BillId: ObjectId;
  BillName: string;
  UserId: ObjectId;
  Price: number;
  Paid: boolean;
  Expired: Date;
  CreatedAt: Date;
};
