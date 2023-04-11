export interface IResponse {
    status: number;
}

export interface ISuccessResponse<T> extends IResponse {
    data: T;
}

export interface IErrorResponse extends IResponse {
    message: string;
}