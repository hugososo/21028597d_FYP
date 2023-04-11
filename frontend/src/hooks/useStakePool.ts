import { useContractRead } from "wagmi";
import stakePoolABI from "contracts/abi/StakePool.json";
import { StakePoolAddress } from "contracts/address";
import { INotification } from "features/notificaiton/providers/NotificationProvider";
import { ContractInstanceContext } from "providers/ContractInstanceProvider";
import { useContext } from "react";
import { BigNumber, utils } from "ethers";
import { ethToWei } from "utils/format";

interface IUseDependencies {
    pushNotification: (_: INotification) => void;
}

const useStakePool = ({
    pushNotification,
}: IUseDependencies) => {
    const { stakePoolContract } = useContext(ContractInstanceContext);

    const useCalculateReward = (address: `0x${string}` | undefined) => {
        const res = useContractRead({
            address: StakePoolAddress,
            abi: stakePoolABI,
            functionName: 'calculateReward',
            args: [address],
            enabled: address !== undefined
        });

        return res;
    }

    const useGetAPY = () => {
        const res = useContractRead({
            address: StakePoolAddress,
            abi: stakePoolABI,
            functionName: 'rewardRate',
        });

        return res;
    }

    const useGetStakedBlances = (address: `0x${string}` | undefined) => {
        const res = useContractRead({
            address: StakePoolAddress,
            abi: stakePoolABI,
            functionName: 'stakedBalances',
            args: [address],
            enabled: address !== undefined
        });

        return res;
    }

    const stake = async (amount: number) => {
        let formatAmount = ethToWei(amount);
        try {
            const txn = await stakePoolContract?.stake(formatAmount);
            pushNotification({
                type: "success",
                title: "Transaction Sent!",
                content: `Txn: ${txn.hash}`
            })
        } catch (err) {
            if (err instanceof Error) {
                pushNotification({
                    type: "error",
                    title: "Something went wrong!",
                    content: err.message,
                })
            }
        }
    }

    const claim = async () => {
        const txn = await stakePoolContract?.claimReward();
        console.log(txn);
    }

    const withdraw = async (amount: number) => {
        let formatAmount = ethToWei(amount);
        const txn = await stakePoolContract?.withdraw(formatAmount);
        console.log(txn);
    }





    return { useCalculateReward, useGetAPY, useGetStakedBlances, stake, claim, withdraw };

}


export default useStakePool;