import { INotification } from "features/notificaiton/providers/NotificationProvider";

export interface INotificationWithTimestamp extends INotification {
  timestamp: number;
  endTime: number;
}
