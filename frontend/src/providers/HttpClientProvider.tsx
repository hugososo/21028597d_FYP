import { createContext, PropsWithChildren } from "react";
import BookApi, { IBookApi } from "api/books";
import UserApi, { IUserApi } from "api/users";

type Props = {
  bookApi?: IBookApi;
  userApi?: IUserApi;
};

type ContextProps = Required<Props>;

const defaultValues = {
  bookApi: BookApi,
  userApi: UserApi,
};

export const HttpClientProviderContext = createContext<ContextProps>(defaultValues);

export const HttpClientProvider = ({ children, bookApi, userApi }: PropsWithChildren<Props>) => {
  return (
    <HttpClientProviderContext.Provider
      value={{
        bookApi: bookApi || BookApi,
        userApi: userApi || UserApi,
      }}
    >
      {children}
    </HttpClientProviderContext.Provider>
  );
};
