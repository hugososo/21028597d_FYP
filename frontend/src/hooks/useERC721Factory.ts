import { INotification } from "features/notificaiton/providers/NotificationProvider";
import { ContractInstanceContext } from "providers/ContractInstanceProvider";
import { useCallback, useContext } from "react";
import { catchError } from 'utils/error';
import { ethToWei } from "utils/format";

interface IUseDependencies {
    pushNotification: (_: INotification) => void;
}

const useERC721Factory = ({
    pushNotification,
}: IUseDependencies) => {
    const { ERC721FactoryContractWrite } = useContext(ContractInstanceContext);

    const deployERC721 = useCallback(async ({
        name,
        symbol,
        ipfsUrl,
        cost,
        maxSupply = 1000000000,
        maxMintPerTxn = 1
    }: {
        name: string,
        symbol: string,
        ipfsUrl: string,
        cost: number,
        maxSupply?: number,
        maxMintPerTxn?: number
    }) => {
        try {
            const txn = await ERC721FactoryContractWrite?.deployERC721(name, symbol, ipfsUrl, ethToWei(cost), maxSupply, maxMintPerTxn);
            console.log(txn)
            pushNotification({
                type: "success",
                title: "Book Publish!",
                content: `Txn: ${txn.hash}`
            });
        } catch (err) {
            pushNotification(catchError(err));
        }
    }, [ERC721FactoryContractWrite, pushNotification]);

    return { deployERC721 }
}


export default useERC721Factory;