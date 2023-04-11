import Button from "components/UI/Button";
import Icon from "components/UI/Icon";
import TextField from "components/UI/TextField";
import Typography from "components/UI/Typography";
import { BigNumber, utils } from "ethers";
import { NotificationContext } from "features/notificaiton/providers/NotificationProvider";
import React, { useContext, useEffect, useState } from "react";
import { useAccount } from "wagmi";
import styles from "./StakeBox.module.scss";
import useStakePool from "hooks/useStakePool";
import useLoreToken from "hooks/useLoreToken";
import { weiToEth } from "utils/format";
import { StakePoolAddress } from "contracts/address";

const Reward = () => {
  const { address, isConnected } = useAccount();
  const [input, setInput] = useState("0");
  const onInputFieldChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    var regexp = new RegExp(/^\d*\.?\d*$/);
    if (regexp.test(value)) {
      setInput(value);
    }
    if (Number(value) > Number(balanceOf)) {
      setInput(balanceOf || "0");
    }
  };
  const { pushNotification } = useContext(NotificationContext);
  const { useCalculateReward, useGetAPY, useGetStakedBlances, stake, claim, withdraw } = useStakePool({
    pushNotification,
  });
  const { useAllowance, approve, getBalanceOf } = useLoreToken({ pushNotification });
  const { data: apy } = useGetAPY();
  const { data: staked } = useGetStakedBlances(address);
  const { data: reward } = useCalculateReward(address);
  const { data: allowance, refetch: refetchAllowance } = useAllowance(address, StakePoolAddress);
  const [balanceOf, setBalanceOf] = useState<string | undefined>(undefined);

  useEffect(() => {
    (async () => {
      const data = await getBalanceOf(address);
      const max = weiToEth((data as BigNumber)?.toString());
      setBalanceOf(max);
    })();
  }, [address, getBalanceOf]);

  return (
    <div className={styles.root}>
      <div className={styles["stake-method"]}>
        <div className={styles["flex-col"]}>
          <Typography variant="h5">Standard Method</Typography>
          <Typography variant="body3" bold grey>
            Earn $LORE
          </Typography>
        </div>
        <div className={styles["flex-col"]}>
          <Typography variant="h4">{(apy as BigNumber)?.toString()}%</Typography>
          <Typography variant="body1" bold grey>
            APY
          </Typography>
        </div>
      </div>
      <div>
        {isConnected ? (
          <>
            <div className={styles["stake-action"]}>
              <Typography variant="body3" bold grey>
                Your Staked $LORE
              </Typography>
              <div className={styles["staked-amount"]}>
                <Icon variant="ether" />
                <Typography>{weiToEth((staked as BigNumber)?.toString())}</Typography>
              </div>
              <div className={styles["stake-input"]}>
                <TextField
                  value={input.toString()}
                  className={styles["input"]}
                  onChange={(e) => onInputFieldChange(e)}
                />
                <Button
                  variant="tertiary"
                  onClick={() => {
                    setInput(balanceOf || "0");
                  }}
                >
                  Max
                </Button>
              </div>
              <div className={styles["button-grp"]}>
                <Button
                  onClick={async () => {
                    if (Number(weiToEth((allowance as BigNumber)?.toString())) === 0) {
                      await approve(StakePoolAddress);
                      refetchAllowance();
                    } else {
                      stake(Number(input));
                    }
                  }}
                >
                  Stake
                </Button>
                <Button variant="tertiary" onClick={() => claim()}>
                  Claim
                </Button>
                <Button variant="success" onClick={() => withdraw(Number(input))}>
                  Withdraw
                </Button>
              </div>
            </div>
            <div className={styles["line"]}></div>
            <div className={styles["rewards-amount"]}>
              <Typography variant="body3" bold grey>
                Rewards
              </Typography>
              <Typography variant="h6">{weiToEth((reward as BigNumber)?.toString())}</Typography>
            </div>
          </>
        ) : (
          <Button>Connect</Button>
        )}
      </div>
    </div>
  );
};

export default Reward;
