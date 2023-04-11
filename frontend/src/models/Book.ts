export type GetBookReqDTO = {
    address: string;
}

export type UploadBookReqDTO = {
    bookName: string;
    description: string;
    author: string;
    authorCity: string;
    bookGenre: string;
    bookLanguage: string;
    edition: string;
    keyword: string[];
}

export type Book = {
    createdAt: number;
    address: string;
    bookName: string;
    description: string;
    author: string;
    authorCity: string;
    bookGenre: string;
    bookLanguage: string;
    edition: string;
    keywords: string[];
    coverImage: string;
    sourceFile: string;
}