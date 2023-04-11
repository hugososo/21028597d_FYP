import { ContractInstanceContext } from './../providers/ContractInstanceProvider';
import { useContractRead, useContractWrite, usePrepareContractWrite } from "wagmi";
import loreTokenABI from "contracts/abi/LoreToken.json";
import { LoreTokenAddress } from "contracts/address";
import { INotification } from "features/notificaiton/providers/NotificationProvider";
import { BigNumber } from "ethers";
import { useCallback, useContext } from "react";

interface IUseDependencies {
    pushNotification: (_: INotification) => void;
}

const useLoreToken = ({
    pushNotification,
}: IUseDependencies) => {
    const { loreTokenContractWrite, loreTokenContractRead } = useContext(ContractInstanceContext);
    const useAllowance = (owner: `0x${string}` | undefined, spender: string) => {
        const res = useContractRead({
            address: LoreTokenAddress,
            abi: loreTokenABI,
            functionName: 'allowance',
            args: [owner, spender],
            enabled: owner !== undefined,
            onSuccess(data) {
                console.log("allowance", "Success", "data", (data as BigNumber).toString());
            },
            onError(err) {
                console.error("allowance", err.message);
            },
        });

        return res;
    }

    const approve = useCallback(async (spender: string) => {
        const txn = await loreTokenContractWrite?.approve(spender, "200000000000000000000000000");
        return txn;
    }, [loreTokenContractWrite])

    const getBalanceOf = useCallback(async (address: `0x${string}` | undefined) => {
        const res = await loreTokenContractRead?.balanceOf(address)
        return res;
    }, [loreTokenContractRead])

    return { useAllowance, approve, getBalanceOf };

}


export default useLoreToken;