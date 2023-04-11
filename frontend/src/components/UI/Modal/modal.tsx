import { FC, MouseEvent, useCallback, useEffect, useMemo, useRef } from "react";
import { createPortal } from "react-dom";

import styles from "./modal.module.scss";

interface IModalProps {
  size?: "default" | "medium";
  open?: boolean;
  onClose?: () => void;
  className?: string;
  children: React.ReactNode;
}

const Modal: FC<IModalProps> = ({ size = "default", open = false, onClose = () => {}, children, className }) => {
  // lock === true -> don't close modal
  // lock === false -> close modal
  // prevent close modal when mouse down inside modal but mouse up in the grey background area
  const lock = useRef(false);

  useEffect(() => {
    if (open) {
      document.body.style.overflow = "hidden";
    } else {
      document.body.style.overflow = "visible";
    }
    return () => {
      document.body.style.overflow = "visible";
    };
  }, [open]);

  const onClick = useCallback(
    (e: MouseEvent<HTMLDivElement>) => {
      e.preventDefault();
      e.stopPropagation();
      onClose();
    },
    [onClose]
  );
  const modalElement = useMemo(() => {
    if (!open) return null;
    return (
      <div
        role="dialog"
        className={`${styles.root}`}
        onClick={(e) => {
          if (!lock.current) {
            onClick(e);
          }
          // after trigger onClick, then unlock
          lock.current = false;
        }}
      >
        <div
          className={`${className || ""} ${styles.content} ${styles[size]}`}
          onClick={(e) => {
            e.stopPropagation();
          }}
          onMouseUp={() => {
            lock.current = false;
          }}
          onMouseDown={() => {
            lock.current = true;
          }}
        >
          {children}
        </div>
      </div>
    );
  }, [children, className, onClick, open, size]);
  return createPortal(modalElement, document.body);
};

export default Modal;
