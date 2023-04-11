import { FC, CSSProperties, ReactNode } from "react";

import styles from "./typography.module.scss";
import { TypographyVariant } from "./typography.type";

interface ITypographyProps {
  variant?: TypographyVariant;
  grey?: boolean;
  bold?: boolean;
  className?: string;
  style?: CSSProperties;
  children: ReactNode;
}

/**
 * Component for displaying different types of text
 * @param variant {TypographyVariant} - different text variants, "body1" | "body2" | "body3" | "h1" | "h2" | "h3" | "h4"
 *
 * @component
 * @example
 * <Typography variant="h1">Hello world </Typography>
 */
const Typography: FC<ITypographyProps> = ({ children, variant = "body1", grey, bold, className, style }) => {
  return (
    <div
      role="complementary"
      className={`${styles.root} ${variant ? styles[variant] : styles.body1} ${className || ""} ${
        grey && styles.grey
      } ${bold && styles.bold}`}
      style={style}
    >
      {children}
    </div>
  );
};

export default Typography;
