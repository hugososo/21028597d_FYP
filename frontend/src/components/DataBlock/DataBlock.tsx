import Icon from "components/UI/Icon";
import { IconVariant } from "components/UI/Icon/icon.type";
import Typography from "components/UI/Typography";
import { PropsWithChildren } from "react";
import styles from "./DataBlock.module.scss";

interface Props extends PropsWithChildren {
  icon: IconVariant;
  data: string | number | undefined;
}

const DataBlock = (props: Props) => {
  return (
    <div className={styles.root}>
      <div className={styles.header}>
        <Icon variant={props.icon}></Icon>
        <Typography variant="h2">{props.data}</Typography>
      </div>
      <div className={styles.content}>
        <Typography variant="body2">{props.children}</Typography>
      </div>
    </div>
  );
};

export default DataBlock;
