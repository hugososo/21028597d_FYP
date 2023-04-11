import { ChangeEvent, FC } from "react";

import styles from "./checkbox.module.scss";
import { ReactComponent as TickIcon } from "assets/icon_tick.svg";

interface ICheckboxProps {
  checked: boolean;
  onChange: (e: ChangeEvent<HTMLInputElement>) => void;
  text?: string;
  required?: boolean;
  error?: boolean;
  disabled?: boolean;
  className?: string;
}

const Checkbox: FC<ICheckboxProps> = ({
  checked,
  onChange,
  text,
  required,
  error,
  disabled,
  className,
}) => (
  <label
    className={`${styles["checkbox-wrapper"]} ${
      disabled ? styles["checkbox-wrapper-disabled"] : ""
    } ${className || ""}`}
  >
    <div className={styles["input-container"]}>
      <input
        type="checkbox"
        checked={checked}
        disabled={disabled}
        onChange={disabled ? undefined : onChange}
      />
      <div
        className={`${styles.checkbox} ${checked ? styles.checked : ""} ${
          error ? styles.error : ""
        }`}
      >
        <TickIcon
          className={`${styles["tick-icon"]} ${checked ? "" : styles.hide}`}
        />
      </div>
    </div>
    {text && (
      <span className={`${styles.text} ${styles["span-disable-highlight"]}`}>
        {text}
      </span>
    )}
    {required && (
      <span
        className={`${styles["required-star"]} ${styles["span-disable-highlight"]}`}
      >
        *
      </span>
    )}
  </label>
);

export default Checkbox;
