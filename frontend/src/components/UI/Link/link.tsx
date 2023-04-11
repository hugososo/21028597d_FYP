import { ReactNode } from "react";
import { Link as RouterLink } from "react-router-dom";
import styles from "./link.module.scss";

interface ILinkProps {
  children?: ReactNode;
  to: string;
  className?: string;
  isNewTabOpen?: boolean;
  underline?: boolean;
}

const Link = ({ to, children, className, isNewTabOpen, underline }: ILinkProps) => (
  <RouterLink
    to={to}
    target={isNewTabOpen ? "_blank" : "_self"}
    className={`${styles.root} ${underline ? styles.underline : ""} ${className || ""}`}
    role="link"
  >
    {children}
  </RouterLink>
);

export default Link;
