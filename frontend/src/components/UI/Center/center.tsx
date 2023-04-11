import { ReactNode } from "react";
import styles from "./center.module.scss";

interface ICenter {
  children?: ReactNode;
  className?: string;
}

const Center = ({ children, className }: ICenter) => {
  return <div className={`${className || ""} ${styles.center}`}>{children}</div>;
};

export default Center;
