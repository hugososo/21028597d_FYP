import { createContext, MouseEvent, PropsWithChildren, useCallback, useState } from "react";

import { FlagType } from "components/UI/Flag";
import AppNotification from "../components/AppNotification/AppNotification";

export interface INotification {
  type?: FlagType;
  title: string;
  content?: string;
  positiveActionText?: string;
  negativeActionText?: string;
  onPositiveCallback?: (e: MouseEvent<HTMLButtonElement>) => void;
  onNegativeCallback?: (e: MouseEvent<HTMLButtonElement>) => void;
  onCloseCallback?: () => void | Promise<void>;
  positiveCountDownRequired?: boolean;
  negativeCountDownRequired?: boolean;
}

type NotificationContextType = {
  notifications: INotification[];
  pushNotification: (notification: INotification) => void;
  popNotification: () => INotification | null;
};

const defaultValue: NotificationContextType = {
  notifications: [],
  pushNotification: () => {},
  popNotification: () => {
    return null;
  },
};

export const NotificationContext = createContext(defaultValue);

const NotificationProvider = ({ children }: PropsWithChildren<{}>) => {
  const [notifications, setNotifications] = useState<INotification[]>([]);

  const pushNotification = useCallback((notification: INotification) => {
    setNotifications((prev) => [...prev, notification]);
  }, []);

  const popNotification = useCallback(() => {
    if (!notifications.length) return null;
    const firstNotification: INotification = notifications[0];
    setNotifications([]);
    return firstNotification;
  }, [notifications]);

  return (
    <NotificationContext.Provider value={{ notifications, pushNotification, popNotification }}>
      <AppNotification />
      {children}
    </NotificationContext.Provider>
  );
};

export default NotificationProvider;
