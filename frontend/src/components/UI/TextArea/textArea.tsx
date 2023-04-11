import { ChangeEvent, forwardRef, FocusEvent, KeyboardEvent, MouseEvent, PropsWithChildren } from "react";

import styles from "./textArea.module.scss";
interface ITextAreaProps {
  value?: string;
  name?: string;
  onChange?: (e: ChangeEvent<HTMLTextAreaElement>) => void;
  onFocus?: (e: FocusEvent<HTMLTextAreaElement>) => void;
  onBlur?: (e: FocusEvent<HTMLTextAreaElement>) => void;
  onKeyDown?: (e: KeyboardEvent<HTMLTextAreaElement>) => void;
  onClick?: (e: MouseEvent<HTMLDivElement>) => void;
  label?: string;
  required?: boolean;
  placeholder?: string;
  disabled?: boolean;
  autoFocus?: boolean;
  maxLength?: number;
  className?: string;
  height?: string;
}

const TextArea = forwardRef<HTMLTextAreaElement, PropsWithChildren<ITextAreaProps>>(
  (
    {
      value,
      name = "",
      onChange,
      onFocus,
      onBlur,
      onKeyDown,
      onClick,
      label,
      required,
      placeholder,
      disabled,
      autoFocus,
      maxLength,
      className,
      children,
      height,
    },
    ref
  ) => {
    return (
      <div onClick={onClick} className={`${styles["text-area-wrapper"]} ${className || ""}`} role="group">
        {label && (
          <div className={styles.label}>
            <div>{label}</div>
            {required && <span>*</span>}
          </div>
        )}
        <div className={styles.field}>
          {onChange || !value ? (
            <textarea
              cols={100}
              rows={10}
              data-testid="textArea"
              ref={ref}
              value={value}
              name={name}
              onChange={onChange}
              onFocus={onFocus}
              onBlur={onBlur}
              onKeyDown={onKeyDown}
              placeholder={placeholder}
              disabled={disabled}
              autoFocus={autoFocus}
              maxLength={maxLength}
              className={`${styles["text-area"]} ${height}`}
              autoComplete="nope"
            />
          ) : (
            <textarea
              cols={100}
              rows={10}
              data-testid="textArea"
              ref={ref}
              name={name}
              onFocus={onFocus}
              onBlur={onBlur}
              onKeyDown={onKeyDown}
              placeholder={placeholder}
              disabled={disabled}
              autoFocus={autoFocus}
              maxLength={maxLength}
              className={`${styles["text-area"]} ${height}`}
              autoComplete="nope"
            />
          )}
        </div>
        {children}
      </div>
    );
  }
);

export default TextArea;
