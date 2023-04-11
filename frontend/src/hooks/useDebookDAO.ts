import { VoteSignature } from 'models/Vote';
import { useChainId, useContractRead, useSignTypedData } from "wagmi";
import DAOContactABI from "contracts/abi/DebookDAO.json";
import { DebookDAOAddress } from "contracts/address";
import { INotification } from "features/notificaiton/providers/NotificationProvider";
import { ContractInstanceContext } from "providers/ContractInstanceProvider";
import { useCallback, useContext } from "react";
import { BigNumber, ethers, utils } from "ethers";
import { ethToWei } from "utils/format";
import { mapToProposal, Proposal, RawProposal } from "models/Proposal";
import { catchError } from 'utils/error';

interface IUseDependencies {
    pushNotification: (_: INotification) => void;
}

const useDebookDAO = ({
    pushNotification,
}: IUseDependencies) => {
    const { DAOContractRead, DAOContractWrite } = useContext(ContractInstanceContext);

    const useGetParameters = () => {
        const res = useContractRead({
            abi: DAOContactABI,
            address: DebookDAOAddress,
            args: [],
            onError: (err) => {
                pushNotification(catchError(err));
            }
        })
    }

    const getProposalList = useCallback(async () => {
        const data = await DAOContractRead?.getProposalList();

        let ids: BigNumber[] = data as BigNumber[];

        let proposals: Proposal[] = [];
        for (let i = ids.length - 1; i >= 0; i--) {
            const state = await DAOContractRead?.state(ids[i]);
            const txn: RawProposal = await DAOContractRead?.getProposal(ids[i].toHexString());
            proposals.push(mapToProposal(ids[i].toHexString(), state, txn));
        }
        return proposals
    }, [DAOContractRead]);

    const propose = useCallback(async (_title: string, description: string) => {
        let title = utils.formatBytes32String(_title);
        try {
            const txn = await DAOContractWrite?.propose(title, description);
            pushNotification({
                type: "success",
                title: "Proposal Created!",
                content: `Txn: ${txn.hash}`
            });
            return txn;
        } catch (err) {
            pushNotification(catchError(err));
        }
    }, [DAOContractWrite, pushNotification]);

    const useSignVote = ({ proposalId,
        voter,
        selection,
        reason, }: Omit<VoteSignature, "interactContract">) => {
        const chainId = useChainId();
        const domain = {
            name: "DebookDAO",
            version: "1",
            chainId: chainId,
            verifyingContract: DebookDAOAddress,
        } as const

        const types = {
            VoteSignature: [
                { name: 'interactContract', type: 'address' },
                { name: 'proposalId', type: 'uint256' },
                { name: 'voter', type: 'address' },
                { name: 'selection', type: 'uint8' },
                { name: 'reason', type: 'bytes32' },
            ],
        }

        const value = {
            interactContract: DebookDAOAddress,
            proposalId: BigNumber.from(proposalId),
            voter: voter,
            selection: selection,
            reason: utils.keccak256(utils.toUtf8Bytes(reason)),
        }

        const res =
            useSignTypedData({
                domain,
                types,
                value
            });

        let sign: { r: string, s: string, v: number } = {
            r: "",
            s: "",
            v: 0,
        };
        if (res && res.data) {
            const signature = res.data.substring(2);
            sign.r = "0x" + signature.substring(0, 64);
            sign.s = "0x" + signature!.substring(64, 128);
            sign.v = parseInt(signature!.substring(128, 130), 16);
        }
        if (res.isSuccess) {
            vote({ proposalId, selection, reason, r: sign.r, s: sign.s, v: sign.v });
            res.reset();
        }
        if (res.isError) {
            res.reset();
            pushNotification(catchError(res.error));
        }

        return { sign, ...res };
    };

    const vote = useCallback(async ({ proposalId,
        selection,
        reason,
        r, s, v }: Omit<VoteSignature, "interactContract" | "voter"> & { r: string, s: string, v: number }) => {
        try {
            const txn = await DAOContractWrite?.castVoteWithReasonBySig(proposalId, selection, reason, v, r, s);
            pushNotification({
                type: "success",
                title: "Voted Successfully!",
                content: `Txn: ${txn.hash}`
            });
        } catch (err) {
            pushNotification(catchError(err));
        }
    }, [DAOContractWrite, pushNotification]);

    return { getProposalList, propose, useSignVote, vote };
}


export default useDebookDAO;