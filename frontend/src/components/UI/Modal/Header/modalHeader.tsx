import { FC } from "react";
import Typography from "components/UI/Typography";
import { ReactComponent as Cross } from "assets/icon_cross_large.svg";
import styles from "./modalHeader.module.scss";
interface IModalHeader {
  title: string;
  onClose: () => void;
  headerClass?: string;
}

const ModalHeader: FC<IModalHeader> = ({ title, onClose, headerClass }) => {
  return (
    <div className={`${styles["modal-header"]} ${headerClass ?? ""}`}>
      <Typography variant="h2" className={styles["title"]}>
        {title}
      </Typography>
      <Cross className={styles["close-icon"]} onClick={onClose} />
    </div>
  );
};

export default ModalHeader;
