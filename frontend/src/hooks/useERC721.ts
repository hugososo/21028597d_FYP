import { INotification } from "features/notificaiton/providers/NotificationProvider";
import { useCallback } from "react";
import { catchError } from 'utils/error';
import { useSigner, useProvider, useContract } from "wagmi";
import erc721ABI from "contracts/abi/ERC721Token.json";

interface IUseDependencies {
    contractAddress: string;
    pushNotification: (_: INotification) => void;
}

const useERC721 = ({
    contractAddress,
    pushNotification,
}: IUseDependencies) => {
    const { data: signer } = useSigner();
    const provider = useProvider();

    const ERC721ContractWrite = useContract({
        address: contractAddress,
        abi: erc721ABI,
        signerOrProvider: signer,
    });

    const ERC721ContractRead = useContract({
        address: contractAddress,
        abi: erc721ABI,
        signerOrProvider: provider,
    });

    const mint = useCallback(async () => {
        try {
            const txn = await ERC721ContractWrite?.mint();
            pushNotification({
                type: "success",
                title: "Enjoy the book!",
                content: `Txn: ${txn.hash}`
            });
        } catch (err) {
            pushNotification(catchError(err));
        }
    }, [ERC721ContractWrite, pushNotification]);

    const cost = useCallback(async () => {
        try {
            const res = await ERC721ContractRead?.cost();
            return res;
        } catch (err) {
            pushNotification(catchError(err));
        }
    }, [ERC721ContractRead, pushNotification]);

    return { mint, cost }
}

export default useERC721;