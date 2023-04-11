import { FC } from "react";

import styles from "./switch.module.scss";

import { ReactComponent as TickIcon } from "assets/icon_tick.svg";
import { ReactComponent as CrossIcon } from "assets/icon_cross.svg";

interface ISwitchProps {
  checked: boolean;
  onChange: () => void;
  disabled?: boolean;
  className?: string;
}

const Switch: FC<ISwitchProps> = ({
  checked,
  onChange,
  disabled,
  className,
}) => (
  <label className={`${styles["toggle-switch"]} ${className || ""}`}>
    <input
      type="checkbox"
      checked={checked}
      onChange={disabled ? undefined : onChange}
      disabled={disabled}
    />
    <span className={styles.switch}>
      {checked ? (
        <TickIcon className={`${styles.icon} ${styles.tick}`} />
      ) : (
        <CrossIcon className={`${styles.icon} ${styles.cross}`} />
      )}
    </span>
  </label>
);

export default Switch;
