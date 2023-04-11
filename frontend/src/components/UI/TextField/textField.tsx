import {
  ChangeEvent,
  cloneElement,
  forwardRef,
  FocusEvent,
  HTMLInputTypeAttribute,
  KeyboardEvent,
  MouseEvent,
  PropsWithChildren,
  ReactElement,
} from "react";

import styles from "./textField.module.scss";
import Icon from "../Icon";

interface ITextFieldProps {
  value?: number | string;
  onChange?: (e: ChangeEvent<HTMLInputElement>) => void;
  onFocus?: (e: FocusEvent<HTMLInputElement>) => void;
  onBlur?: (e: FocusEvent<HTMLInputElement>) => void;
  onKeyDown?: (e: KeyboardEvent<HTMLInputElement>) => void;
  onClick?: (e: MouseEvent<HTMLInputElement>) => void;
  type?: HTMLInputTypeAttribute;
  pattern?: string;
  name?: string;
  label?: string;
  required?: boolean;
  placeholder?: string;
  startIcon?: ReactElement;
  endIcon?: ReactElement;
  disabled?: boolean;
  autoFocus?: boolean;
  maxLength?: number;
  errorMessage?: string | null;
  successMessage?: string | null;
  className?: string;
  takeupFullSpace?: boolean;
  isDuplicated?: boolean;
  readOnly?: boolean;
  hideErrorText?: boolean;
  accept?: string;
}

const TextField = forwardRef<HTMLInputElement, PropsWithChildren<ITextFieldProps>>(
  (
    {
      value,
      onChange = undefined,
      onFocus,
      onBlur,
      onKeyDown,
      onClick,
      type = "text",
      name = "",
      pattern,
      label,
      required,
      placeholder,
      startIcon,
      endIcon,
      disabled,
      autoFocus,
      maxLength,
      errorMessage,
      successMessage,
      className,
      children,
      takeupFullSpace,
      isDuplicated,
      readOnly,
      hideErrorText,
      accept,
    },
    ref
  ) => {
    return (
      <div
        onClick={onClick}
        className={`${styles["text-field"]} ${className || ""}  ${takeupFullSpace ? styles["full-space"] : ""}`}
        role="group"
      >
        {label && (
          <div className={styles.label}>
            <div>{label}</div>
            {required && <span>*</span>}
          </div>
        )}
        <div className={styles.field}>
          {(startIcon &&
            cloneElement(startIcon, {
              ...startIcon.props,
              className: `${styles["start-icon"]} ${startIcon.props.className || ""}`,
              role: "img",
            })) ||
            null}
          {onChange || !value ? (
            <input
              data-testid="textFieldInput"
              ref={ref}
              type={type}
              value={value}
              pattern={pattern}
              onChange={onChange}
              onFocus={onFocus}
              onBlur={onBlur}
              onKeyDown={onKeyDown}
              placeholder={placeholder}
              disabled={disabled}
              autoFocus={autoFocus}
              maxLength={maxLength}
              className={`${styles.input} 
                        ${errorMessage ? styles["error-border"] : ""}
                      `}
              readOnly={readOnly}
              // enable the best practice for autofill browser data
              name={name}
              // disenable autofill reference: https://stackoverflow.com/questions/12374442/chrome-ignores-autocomplete-off
              autoComplete="nope"
              accept={accept}
            />
          ) : (
            <input
              data-testid="textFieldInput"
              ref={ref}
              type={type}
              pattern={pattern}
              onFocus={onFocus}
              onBlur={onBlur}
              onKeyDown={onKeyDown}
              placeholder={placeholder}
              disabled={disabled}
              autoFocus={autoFocus}
              maxLength={maxLength}
              className={`${styles.input} 
                      ${errorMessage ? styles["error-border"] : ""}
                    `}
              readOnly={readOnly}
              // enable the best practice for autofill browser data
              name={name}
              // disenable autofill reference: https://stackoverflow.com/questions/12374442/chrome-ignores-autocomplete-off
              autoComplete="nope"
              accept={accept}
            />
          )}
          {(endIcon &&
            cloneElement(endIcon, {
              ...endIcon.props,
              className: `${styles["end-icon"]} ${endIcon.props.className || ""}`,
              role: "img",
            })) ||
            null}
        </div>
        {children}
        {errorMessage &&
          !hideErrorText &&
          (errorMessage !== "This field is required" && isDuplicated ? (
            <div className={styles["message-wrapper"]}>
              <div className={`${styles.message} ${styles.error}`}>{styledErrorMsg(errorMessage, isDuplicated)}</div>
            </div>
          ) : (
            <div className={styles["message-wrapper"]}>
              <div className={`${styles.icon} ${styles.error}`}>
                <Icon variant="error" className={styles["icon"]} />
              </div>
              <div className={`${styles.message} ${styles.error}`}>{styledErrorMsg(errorMessage)}</div>
            </div>
          ))}
        {successMessage && (
          <div className={styles["message-wrapper"]}>
            <div className={`${styles.icon} ${styles.success}`}>
              <Icon variant="success" className={styles["icon"]} />
            </div>
            <span className={`${styles.message} ${styles.success}`}>{successMessage}</span>
          </div>
        )}
      </div>
    );
  }
);

const styledErrorMsg = (errorMessage: string, isDuplicated?: boolean) => {
  if (errorMessage.includes("allowed:")) {
    const results = errorMessage.split("allowed: ");
    const errTags = results[1].split("").map((str) => (
      <div key={str} className={`${styles["error-tag-wrapper"]} ${styles["tag-end"]}`}>
        <span className={styles["error-tag"]}>{str}</span>
      </div>
    ));
    return (
      <div className={styles["styled-error"]}>
        {results[0]}allowed: {errTags}
      </div>
    );
  }

  if (errorMessage.includes("is/are not allowed")) {
    const results = errorMessage.split(" is/are not allowed");
    const errTags = results[0].split("").map((str) => (
      <div key={str} className={`${styles["error-tag-wrapper"]} ${styles["tag-front"]}`}>
        <span className={styles["error-tag"]}>{str === " " ? "space" : str}</span>
      </div>
    ));
    return <div className={styles["styled-error"]}>{errTags}is/are not allowed</div>;
  }

  if (errorMessage !== "This field is required" && isDuplicated) {
    return <div className={styles["styled-error"]}></div>;
  }

  return errorMessage;
};

export default TextField;
