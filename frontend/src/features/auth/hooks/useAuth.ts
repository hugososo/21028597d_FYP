import { INotification } from "features/notificaiton/providers/NotificationProvider";
import { useContext, useEffect, useState } from "react";
import { useDispatch } from "react-redux";

type requiredVerifications = {
  email?: string;
  phone_number?: string;
};
interface IUseAuthDependencies {
  pushNotification: (_: INotification) => void;
}

export default function useAuth({ pushNotification }: IUseAuthDependencies) {
  const dispatch = useDispatch();
  // const { authState } = useAppSelector((root) => root.auth);

}
