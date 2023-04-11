import { FC, MouseEvent, cloneElement, ReactElement } from "react";

import styles from "./button.module.scss";
import Spinner from "components/UI/Spinner";

interface IButtonProps {
  startIcon?: ReactElement;
  endIcon?: ReactElement;
  onlyIcon?: ReactElement;
  variant?: "default" | "primary" | "secondary" | "tertiary" | "warning" | "error" | "success";
  loading?: boolean;
  disabled?: boolean;
  big?: boolean;
  long?: boolean;
  value?: any;
  type?: "submit" | "reset" | "button";
  onClick?: (e: MouseEvent<HTMLButtonElement>) => void;
  className?: string;
  children?: React.ReactNode;
}

/**
 * Component for interacting click behaviors with user
 * @param startIcon {ReactElement} - react element representing icon positioned at the start of the button
 * @param endIcon {ReactElement} - react element representing icon positioned at the end of the button
 * @param onlyIcon {ReactElement} - react element representing icon button and icon only
 * @param variant {"primary" | "secondary" | "tertiary" | "warning" | "error"} - color state of the button(blue/gray/white/yellow/red)
 * @param loading {boolean} - loading state of the button
 * @param disabled {boolean} - disabled state of the button
 * @param long {boolean} - long size variant of the button
 * @param value {any} - value content to be rendered inside the button, overrides children if any
 * @param type {"submit" | "reset" | "button"} - type of button
 * @param onClick {(e: MouseEvent<HTMLButtonElement>) => void} - callback fired on button click
 * @param classname {string} - injectable className
 *
 * @component
 * @example
 * const [loading,setLoading] = useState(false);
 * const onClick = () => {setLoading(prev => !prev)}
 * <Button startIcon={<StarIcon />} value={"Star Icon"} loading={loading} onClick={onClick} variant="error"/>
 *
 */
const Button: FC<IButtonProps> = ({
  startIcon,
  endIcon,
  onlyIcon,
  variant = "primary",
  loading,
  disabled,
  big,
  long,
  value,
  type = "button",
  onClick,
  className,
  children,
}) => {
  return (
    <button
      className={`${className || ""} ${styles.root} ${onlyIcon ? styles.only : ""} ${
        disabled ? styles[`${variant || ""}-disabled`] : ""
      } ${variant && styles[variant]} ${long ? styles.long : ""} ${big ? styles.big : ""}`}
      type={type}
      onClick={onClick}
      disabled={disabled || loading}
    >
      {onlyIcon ? (
        onlyIcon
      ) : (
        <>
          {(startIcon &&
            cloneElement(startIcon, {
              className: `${styles["start-icon"]} ${startIcon.props.className || ""}`,
              role: "img",
            })) ||
            null}
          {value || children}
          {(endIcon &&
            cloneElement(endIcon, {
              className: `${styles["end-icon"]} ${endIcon.props.className || ""}`,
              role: "img",
            })) ||
            null}
        </>
      )}
      {loading && (
        <div role="status" className={styles.loading}>
          <Spinner />
        </div>
      )}
    </button>
  );
};

export default Button;
