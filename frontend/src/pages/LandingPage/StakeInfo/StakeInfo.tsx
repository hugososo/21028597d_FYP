import styles from "./StakeInfo.module.scss";
import { NavigateFunction } from "react-router-dom";
import { ReactComponent as Stake1 } from "assets/images/stake1.svg";
import Typography from "components/UI/Typography";
import Button from "components/UI/Button";

interface IStakeInfoProp {
  navigate: NavigateFunction;
}

const StakeInfo = ({ navigate }: IStakeInfoProp) => {
  return (
    <div className={`${styles["flex-row"]} ${styles["stakeinfo-container"]}`}>
      <div className={styles["left-container"]}>
        <div>
          <Typography variant="h1">Earn up to 46.10% APY</Typography>
          <Typography variant="h1">with $LORE</Typography>
        </div>
        <Typography variant="body1" grey>
          Stake LORE tokens to transfer to xLORE and earn a share of daily trading fees in xLORE, and finally unstake to
          earn more LORE.
        </Typography>
        <Button big variant="primary" onClick={() => navigate("/reward")}>
          Start Staking
        </Button>
      </div>
      <div className={styles["right-container"]}>
        <Stake1 className={styles.img} />
      </div>
    </div>
  );
};

export default StakeInfo;
