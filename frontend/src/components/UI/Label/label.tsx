import styles from "./label.module.scss";
import { LabelColor } from "./label.type";

interface ILabelProps {
  text: string;
  color?: LabelColor;
  onClick?: () => void;
  className?: string;
}

const Label = ({ text, color = "grey1", onClick, className }: ILabelProps) => {
  return (
    <div
      role="status"
      onClick={onClick}
      className={`${styles.root} ${styles[color]} ${onClick ? styles.pointer : ""} ${className || ""}`}
    >
      {text}
    </div>
  );
};

export default Label;
