import { Book } from "./Book";

export type User = {
    userAddress: string;
    name: string;
    publishedBooks?: Book[];
    boughtBooks?: Book[];
}

export type UserReqDTO = {
    address: string;
}