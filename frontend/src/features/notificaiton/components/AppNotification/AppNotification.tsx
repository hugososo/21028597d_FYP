import { useCallback, useContext, useEffect, useState } from "react";

import styles from "./AppNotification.module.scss";
import { INotificationWithTimestamp } from "./AppNotification.type";
import {
  SHORT_NOTIFICATION_DURATION,
  LONG_NOTIFICATION_DURATION,
  EXTRA_LONG_NOTIFICATION_DURATION,
} from "./AppNotification.constants";

import Flag from "components/UI/Flag";
import { NotificationContext } from "features/notificaiton/providers/NotificationProvider";

const AppNotification = () => {
  const { notifications: appNotifications, popNotification } = useContext(NotificationContext);
  const [notifications, setNotifications] = useState<INotificationWithTimestamp[]>([]);

  const removeOldestNotification = useCallback(() => {
    setNotifications((prev) => {
      const notificationsCopy = [...prev];
      notificationsCopy.forEach((notification, index) => {
        if (notification.endTime < Date.now()) {
          notificationsCopy.splice(index, 1);
        }
      });
      return notificationsCopy;
    });
  }, []);

  useEffect(() => {
    const notification = popNotification();
    if (!notification) {
      return;
    }
    const notificationDuration =
      notification.positiveActionText || notification.negativeActionText
        ? EXTRA_LONG_NOTIFICATION_DURATION
        : notification.content
        ? LONG_NOTIFICATION_DURATION
        : SHORT_NOTIFICATION_DURATION;
    const notificationWithTimestamp = {
      timestamp: new Date().getTime(),
      endTime: new Date().getTime() + notificationDuration,
      ...notification,
    };
    setNotifications((prev) => [...prev, notificationWithTimestamp]);
    setTimeout(() => {
      removeOldestNotification();
    }, notificationDuration);
  }, [appNotifications, popNotification, removeOldestNotification]);

  return (
    <div className={styles.root}>
      {notifications.map((notification) => {
        return <Flag className={styles.flag} key={notification.timestamp} type="info" {...notification} />;
      })}
    </div>
  );
};

export default AppNotification;
