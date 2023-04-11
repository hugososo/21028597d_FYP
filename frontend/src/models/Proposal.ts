import { BigNumber, utils } from "ethers";

export function mapToProposal(id: string, state: ProposalState, rawProposal: RawProposal) {
    let p: Proposal = newProposal();
    p.id = id;
    p.proposer = rawProposal.proposer;
    p.title = utils.parseBytes32String(rawProposal.title);
    p.description = rawProposal.description;
    p.voteStart = new Date(rawProposal.voteStart[0].toNumber() * 1000);
    p.voteEnd = new Date(rawProposal.voteEnd[0].toNumber() * 1000);
    p.upVoteCount = rawProposal.upVoteCount.toString();
    p.downVoteCount = rawProposal.downVoteCount.toString();
    p.isSet = rawProposal.isSet;
    p.isExecuted = rawProposal.isExecuted;
    p.isCanceled = rawProposal.isCanceled;
    p.state = state;
    return p;
}

function newProposal(): Proposal {
    let p: Proposal = {
        id: "",
        proposer: "",
        title: "",
        description: "",
        voteStart: undefined,
        voteEnd: undefined,
        upVoteCount: "",
        downVoteCount: "",
        isSet: false,
        isExecuted: false,
        isCanceled: false,
        state: undefined,
    };

    return p;
}

export enum ProposalState {
    ACTIVE,
    PENDING,
    DEFEATED,
    SUCCEEDED,
    CANCELLED,
    EXPIRED,
    EXECUTED
}

export type RawProposal = {
    proposer: string;
    title: string;
    description: string;
    voteStart: BigNumber[];
    voteEnd: BigNumber[];
    upVoteCount: BigNumber;
    downVoteCount: BigNumber;
    isSet: boolean;
    isExecuted: boolean;
    isCanceled: boolean;
}

export type Proposal = {
    id: string;
    proposer: string;
    title: string;
    description: string;
    voteStart: Date | undefined;
    voteEnd: Date | undefined;
    upVoteCount: string;
    downVoteCount: string;
    isSet: boolean;
    isExecuted: boolean;
    isCanceled: boolean;
    state: ProposalState | undefined;
}

export function proposalStateToText(state: ProposalState | undefined) {
    switch (state) {
        case ProposalState.ACTIVE: return "Active";
        case ProposalState.CANCELLED: return "Cancelled";
        case ProposalState.DEFEATED: return "Defeated";
        case ProposalState.EXECUTED: return "Executed";
        case ProposalState.EXPIRED: return "Expired";
        case ProposalState.PENDING: return "Pending";
        case ProposalState.SUCCEEDED: return "Succeeded";
        default: return "Loading"
    }
}