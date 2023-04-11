import { FC } from "react";

import styles from "./externalLink.module.scss";

interface IExternalLinkProps {
  to: string;
  isNewTabOpen?: boolean;
  className?: string;
  children: React.ReactNode;
}

const ExternalLink: FC<IExternalLinkProps> = ({ to, children, isNewTabOpen = true, className }) => {
  return (
    <a
      href={to}
      className={`${styles.root} ${className || ""}`}
      target={isNewTabOpen ? "_blank" : "_self"}
      rel="noreferrer"
    >
      {children}
    </a>
  );
};

export default ExternalLink;
