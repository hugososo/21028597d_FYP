import { FC, MouseEvent } from "react";

import styles from "./tag.module.scss";

import { ReactComponent as CrossIcon } from "assets/icon_cross_sharp.svg";

interface ITagProps {
  className?: string;
  content: string;
  onRemove?: (e: MouseEvent<SVGSVGElement>) => void;
}

const Tag: FC<ITagProps> = ({ content, className, onRemove }) => {
  return (
    <span className={`${styles.root} ${className || ""}`}>
      {content}
      {onRemove && (
        <CrossIcon role="img" className={styles.icon} onClick={onRemove} />
      )}
    </span>
  );
};

export default Tag;
