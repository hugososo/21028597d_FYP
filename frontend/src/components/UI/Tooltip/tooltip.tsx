import React, { CSSProperties, FC, ReactNode, useEffect, useState } from "react";

import styles from "./tooltip.module.scss";

interface ITooltipProps {
  text: string;
  position?: "top" | "bottom" | "left" | "right";
  delay?: number;
  className?: string;
  style?: CSSProperties;
  tooltipClassName?: string;
  children: ReactNode;
}

const Tooltip: FC<ITooltipProps> = ({
  children,
  text,
  position = "top",
  delay = 0,
  className,
  style,
  tooltipClassName,
}) => {
  const [mouseEnter, setMouseEnter] = useState(false);
  const [active, setActive] = useState(false);

  useEffect(() => {
    if (!mouseEnter) return;

    let timeout = setTimeout(() => {
      setActive(true);
    }, delay);

    return () => {
      clearInterval(timeout);
      setActive(false);
    };
  }, [mouseEnter, delay]);

  return (
    <div
      className={`${className || ""} ${styles[position]}`}
      id={styles.tooltip}
      onMouseEnter={() => setMouseEnter(true)}
      onMouseLeave={() => setMouseEnter(false)}
      style={style}
    >
      {children}
      {active && (
        <div
          role="tooltip"
          className={`${styles.label} ${position ? styles[position] : styles.top} ${tooltipClassName || ""}`}
        >
          {text}
        </div>
      )}
    </div>
  );
};

export default Tooltip;
