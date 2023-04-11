import { FC, MouseEvent, useEffect, useState } from "react";

import styles from "./flag.module.scss";
import { FlagType } from "./flag.type";

import { EXTRA_LONG_NOTIFICATION_DURATION } from "features/notificaiton/components/AppNotification/AppNotification.constants";
import Icon from "../Icon";

interface IFlagProps {
  type: FlagType;
  title: string;
  content?: string;
  positiveActionText?: string;
  negativeActionText?: string;
  onPositiveCallback?: (e: MouseEvent<HTMLButtonElement>) => void;
  onNegativeCallback?: (e: MouseEvent<HTMLButtonElement>) => void;
  onCloseCallback?: () => void | Promise<void>;
  positiveCountDownRequired?: boolean;
  negativeCountDownRequired?: boolean;
  className?: string;
}

const Flag: FC<IFlagProps> = ({
  type,
  title,
  content,
  positiveActionText,
  negativeActionText,
  onPositiveCallback,
  onNegativeCallback,
  onCloseCallback,
  positiveCountDownRequired,
  negativeCountDownRequired,
  className,
}) => {
  const [isClosed, setIsClosed] = useState(false);
  const initialPositiveCountDown = EXTRA_LONG_NOTIFICATION_DURATION / 1000;
  const initialNegativeCountDown = EXTRA_LONG_NOTIFICATION_DURATION / 1000;
  const [positiveCountDownDuration, setPositiveCountDownDuration] = useState(initialPositiveCountDown);
  const [negativeCountDownDuration, setNegativeCountDownDuration] = useState(initialNegativeCountDown);

  useEffect(() => {
    if (!positiveCountDownRequired) return;

    const interval = setInterval(() => {
      setPositiveCountDownDuration((prev) => (prev > 0 ? prev - 1 : 0));
    }, 1000);
    return () => clearInterval(interval);
  }, [positiveCountDownRequired]);

  useEffect(() => {
    if (!negativeCountDownRequired) return;

    const interval = setInterval(() => {
      setNegativeCountDownDuration((prev) => (prev > 0 ? prev - 1 : 0));
    }, 1000);
    return () => clearInterval(interval);
  }, [negativeCountDownRequired]);

  useEffect(() => {
    if (!onCloseCallback) return;
    if (positiveCountDownDuration === 1 || negativeCountDownDuration === 1) {
      onCloseCallback();
    }
  }, [onCloseCallback, positiveCountDownDuration, negativeCountDownDuration]);

  const onClose = (e: MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    setIsClosed(true);
  };

  return (
    <div
      className={`${className || ""} ${styles.root} ${styles[type]} ${
        positiveCountDownDuration === 1 || negativeCountDownDuration === 1 || isClosed ? styles.hide : ""
      }`}
      role="alert"
    >
      <div className={styles["status-icon"]}>
        <FlagIcon type={type} />
      </div>
      <div className={styles.main}>
        <h3 className={styles.header}>{title}</h3>
        {content || (positiveActionText && onPositiveCallback) || (negativeActionText && onNegativeCallback) ? (
          <div role="group" className={`${styles.message}`}>
            <div role="contentinfo" className={styles.contentinfo}>
              {content}
            </div>
            {(positiveActionText && onPositiveCallback) || (negativeActionText && onNegativeCallback) ? (
              <div className={`${styles.buttons}`}>
                {positiveActionText && onPositiveCallback && (
                  <button
                    className={`${styles.button} ${type === "warning" ? styles["warning-button"] : ""}`}
                    onClick={(e) => {
                      onPositiveCallback(e);
                      onClose(e);
                      setPositiveCountDownDuration(initialPositiveCountDown); //reset
                    }}
                  >
                    {positiveCountDownRequired
                      ? `${positiveActionText} (${positiveCountDownDuration - 1}s)`
                      : positiveActionText}
                  </button>
                )}
                {negativeActionText && onNegativeCallback && (
                  <button
                    className={`${styles.button} ${type === "warning" ? styles["warning-button"] : ""}`}
                    onClick={(e) => {
                      onNegativeCallback(e);
                      onClose(e);
                      setNegativeCountDownDuration(initialNegativeCountDown); //reset
                    }}
                  >
                    {negativeCountDownRequired
                      ? `${negativeActionText} (${negativeCountDownDuration - 1}s)`
                      : negativeActionText}
                  </button>
                )}
              </div>
            ) : null}
          </div>
        ) : null}
      </div>
    </div>
  );
};

const FlagIcon = ({ type }: { type: FlagType }) => {
  return <Icon variant={type} className={styles["status-icon"]} />;
};

export default Flag;
