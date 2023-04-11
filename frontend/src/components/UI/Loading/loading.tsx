import { Triangle } from "react-loader-spinner";
import Spinner from "../Spinner";
import styles from "./loading.module.scss";

interface ILoading {
  className?: string;
}

const Loading = ({ className }: ILoading) => {
  return (
    <Triangle
      height="80"
      width="80"
      color={styles.primaryBlue1}
      ariaLabel="triangle-loading"
      wrapperClass={className || ""}
      visible={true}
    />
  );
};

export default Loading;
