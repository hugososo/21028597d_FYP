import Button from "components/UI/Button";
import Typography from "components/UI/Typography";
import styles from "./Rewards.module.scss";
import StakeBox from "./StakeBox";

const Reward = () => {
  return (
    <div className={styles.root}>
      <div className={`${styles["flex-col"]} ${styles["top-panel"]}`}>
        <Typography variant="h1">Rewards</Typography>
        <Typography variant="body1" grey>
          Stake $LORE tokens to earn extra rewards, and get the access to the DAO.
        </Typography>
        <Button variant="tertiary">Buy $LORE</Button>
      </div>
      <div className={styles.line}></div>
      <div className={styles["main-content"]}>
        <StakeBox />
      </div>
    </div>
  );
};

export default Reward;
