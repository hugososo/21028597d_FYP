import { ServiceInstance } from "./instance";
import { ISuccessResponse } from "models/Response";
import { User, UserReqDTO } from 'models/User';

export interface IUserApi {
    CreateUser: (
        address: UserReqDTO,
        controller?: AbortController
    ) => Promise<string>;
    GetUser: (
        address: string,
        controller?: AbortController
    ) => Promise<User>;
    GetUserWithType: (
        address: string,
        type: string,
        controller?: AbortController
    ) => Promise<User>;
}

export const UserApi: IUserApi = {
    CreateUser: async (
        address: UserReqDTO,
        controller?: AbortController
    ) => {
        try {
            const response = await ServiceInstance.post<ISuccessResponse<string>>(
                "/users/",
                address,
                {
                    signal: controller?.signal,
                },
            );
            return response.data.data;
        } catch (err) {
            throw err;
        }
    },
    GetUser: async (
        address: string,
        controller?: AbortController
    ) => {
        try {
            const response = await ServiceInstance.get<ISuccessResponse<User>>(
                `/users/${address}`,
            );
            return response.data.data;
        } catch (err) {
            throw err;
        }
    },
    GetUserWithType: async (
        address: string,
        type: string,
        controller?: AbortController
    ) => {
        try {
            const response = await ServiceInstance.get<ISuccessResponse<User>>(
                `/users/${address}/${type}`,
            );
            return response.data.data;
        } catch (err) {
            throw err;
        }
    },
};

export default UserApi;
