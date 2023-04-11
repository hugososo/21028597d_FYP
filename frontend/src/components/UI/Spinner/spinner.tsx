import styles from "./spinner.module.scss";

interface ISpinnerProps {
  className?: string;
}

const Spinner = ({ className }: ISpinnerProps) => {
  return <div role="img" className={`${className || ""} ${styles.loader} ${styles["loader-default"]}`} />;
};

export default Spinner;
