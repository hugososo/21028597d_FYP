export type VoteSignature = {
    interactContract: `0x${string}`;
    proposalId: string;
    voter: `0x${string}`;
    selection: number;
    reason: string;
}