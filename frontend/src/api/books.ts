import { Book } from "models/Book";
import { ServiceInstance } from "./instance";
import { ISuccessResponse } from "models/Response";

export interface IBookApi {
    UploadBook: (
        formData: FormData,
        controller?: AbortController
    ) => Promise<string>;
    GetBooks: (
        controller?: AbortController
    ) => Promise<Book[]>;
    GetBookByAddress: (
        address: string,
        controller?: AbortController
    ) => Promise<Book>;
    DownloadBook: (
        address: string,
        controller?: AbortController
    ) => Promise<Blob | MediaSource>;
}

export const BookApi: IBookApi = {
    UploadBook: async (
        formData: FormData,
        controller?: AbortController
    ) => {
        try {
            const response = await ServiceInstance.post<ISuccessResponse<string>>(
                "/books/",
                formData,
                {
                    headers: {
                        'Content-Type': 'multipart/form-data'
                    },
                    signal: controller?.signal,
                },
            );
            return response.data.data;
        } catch (err) {
            throw err;
        }
    },
    GetBooks: async (
        controller?: AbortController
    ) => {
        try {
            const response = await ServiceInstance.get<ISuccessResponse<Book[]>>(
                `/books/`
            );
            return response.data.data;
        } catch (err) {
            throw err;
        }
    },
    GetBookByAddress: async (
        address: string,
        controller?: AbortController
    ) => {
        try {
            const response = await ServiceInstance.get<ISuccessResponse<Book>>(
                `/books/${address}/`,
            );
            return response.data.data;
        } catch (err) {
            throw err;
        }
    },
    DownloadBook: async (
        address: string,
        controller?: AbortController
    ) => {
        try {
            const response = await ServiceInstance.get<ISuccessResponse<Blob | MediaSource>>(
                `/books/${address}/file`,
                {
                    responseType: 'blob',
                    signal: controller?.signal,
                },
            );
            const href = URL.createObjectURL(response.data.data);
            const link = document.createElement('a');
            link.href = href;
            link.setAttribute('download', 'file.pdf');
            document.body.appendChild(link);
            link.click();
            document.body.removeChild(link);
            URL.revokeObjectURL(href);
            return response.data.data;
        } catch (err) {
            throw err;
        }
    },
};

export default BookApi;
