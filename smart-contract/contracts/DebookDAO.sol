// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/utils/cryptography/EIP712.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/Timers.sol";
import "@openzeppelin/contracts/utils/Context.sol";
import "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";
import "./IStakePool.sol";

contract DebookDAO is EIP712, Context {
    using Timers for Timers.Timestamp;
    //variables
    string private name = "DebookDAO";
    IStakePool public sp;
    uint256 public _votingDelay = 1; // 1 block
    uint256 public _votingPeriod = 50400; // 1 week
    bytes32 public constant VOTE_TYPEHASH =
        keccak256(
            "VoteSignature(address interactContract,uint256 proposalId,address voter,uint8 selection,bytes32 reason)"
        );
    bytes32 public constant PROPOSAL_TYPEHASH =
        keccak256(
            "Proposal(address proposer, bytes32 title, uint256 proposeTime)"
        );
    bytes32 public teamMerkleRoot;

    //enum
    enum ProposalState {
        ACTIVE,
        PENDING,
        DEFEATED,
        SUCCEEDED,
        CANCELLED,
        EXPIRED,
        EXECUTED
    }

    enum VoteState {
        DISAGREE,
        AGREE
    }

    //struct
    struct Voter {
        VoteState selection;
        string reason;
        bool isVoted;
    }

    struct Proposal {
        address proposer;
        bytes32 title;
        string description;
        Timers.Timestamp voteStart;
        Timers.Timestamp voteEnd;
        uint256 upVoteCount;
        uint256 downVoteCount;
        bool isSet;
        bool isExecuted;
        bool isCanceled;
    }

    struct VoteSignature {
        address interactContract;
        uint256 proposalId;
        address voter;
        uint8 selection;
        string reason;
    }

    //event
    event Voted(
        address indexed voter,
        uint256 proposalId,
        VoteState selection,
        string reason
    );

    event ProposalCreated(uint256 proposalId);
    event ProposalCanceled(uint256 proposalId);
    event ProposalExecuted(uint256 proposalId);
    event Console(string title, address message);

    mapping(uint256 => Proposal) private _proposals;
    mapping(uint256 => mapping(address => Voter)) private _voters;
    uint256[] public proposalsList;

    modifier onlyTeam(bytes32[] calldata _merkleProof) {
        bytes32 leaf = keccak256(abi.encodePacked(_msgSender()));
        require(
            MerkleProof.verify(_merkleProof, teamMerkleRoot, leaf),
            "You are not in the Team."
        );
        _;
    }

    modifier onlyProposer(uint256 _proposalId) {
        require(
            _msgSender() == _proposals[_proposalId].proposer,
            "You are not proposer"
        );
        _;
    }

    modifier onlyStaker() {
        require(
            sp.isStaker(_msgSender()),
            "You need to at least stake 2000 $LORE first."
        );
        _;
    }

    constructor(address _sp) EIP712("DebookDAO", "1") {
        sp = IStakePool(_sp);
    }

    function state(uint256 proposalId) public view returns (ProposalState) {
        Proposal storage proposal = _proposals[proposalId];

        if (proposal.isExecuted) {
            return ProposalState.EXECUTED;
        }

        if (proposal.isCanceled) {
            return ProposalState.CANCELLED;
        }

        uint256 voteStart = getProposalVoteStart(proposalId);

        if (voteStart == 0) {
            revert("Governor: unknown proposal id");
        }

        if (voteStart > block.timestamp) {
            return ProposalState.PENDING;
        }

        uint256 deadline = getProposalDeadline(proposalId);

        if (deadline > block.timestamp) {
            return ProposalState.ACTIVE;
        }

        if (proposal.upVoteCount == 0 && proposal.downVoteCount == 0) {
            return ProposalState.EXPIRED;
        }

        if (proposal.upVoteCount >= proposal.downVoteCount) {
            return ProposalState.SUCCEEDED;
        } else {
            return ProposalState.DEFEATED;
        }
    }

    function hashVote(
        address interactContract,
        uint256 proposalId,
        address voter,
        VoteState selection,
        string calldata reason
    ) public pure virtual returns (bytes32) {
        return
            keccak256(
                abi.encode(
                    VOTE_TYPEHASH,
                    interactContract,
                    proposalId,
                    voter,
                    selection,
                    keccak256(bytes(reason))
                )
            );
    }

    function hashProposal(
        address proposer,
        bytes32 title,
        uint256 createdAt
    ) public pure virtual returns (uint256) {
        return
            uint256(
                keccak256(
                    abi.encode(PROPOSAL_TYPEHASH, proposer, title, createdAt)
                )
            );
    }

    function propose(
        bytes32 title,
        string memory description
    ) public onlyStaker returns (uint256) {
        uint256 createdAt = block.timestamp;
        uint256 proposalId = hashProposal(_msgSender(), title, createdAt);
        require(!_proposals[proposalId].isSet, "The proposal is existed");
        Proposal storage proposal = _proposals[proposalId];
        proposal.proposer = _msgSender();
        proposal.title = title;
        proposal.description = description;
        proposal.voteStart.setDeadline(uint64(block.timestamp + 1 minutes));
        proposal.voteEnd.setDeadline(
            uint64(proposal.voteStart.getDeadline() + 1 weeks)
        );
        proposal.isSet = true;
        proposalsList.push(proposalId);
        emit ProposalCreated(proposalId);
        return proposalId;
    }

    function cancel(
        uint256 proposalId
    ) public onlyProposer(proposalId) returns (uint256) {
        ProposalState status = state(proposalId);
        require(
            status != ProposalState.CANCELLED &&
                status != ProposalState.EXPIRED &&
                status != ProposalState.EXECUTED,
            "Governor: proposal not active"
        );
        _proposals[proposalId].isCanceled = true;

        emit ProposalCanceled(proposalId);
        return proposalId;
    }

    function exeute(
        uint256 proposalId,
        bytes32[] calldata merkleProof
    ) public onlyTeam(merkleProof) returns (uint256) {
        ProposalState status = state(proposalId);
        require(
            status == ProposalState.SUCCEEDED,
            "Governor: proposal not successful"
        );
        _proposals[proposalId].isExecuted = true;

        emit ProposalExecuted(proposalId);
        return proposalId;
    }

    function castVoteWithReasonBySig(
        uint256 proposalId,
        VoteState selection,
        string calldata reason,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) public virtual onlyStaker {
        bytes32 domainSeparator = keccak256(
            abi.encode(
                keccak256(
                    "EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"
                ),
                keccak256(bytes(name)),
                keccak256(bytes("1")),
                1337,
                address(this)
            )
        );

        bytes32 structHash = hashVote(
            address(this),
            proposalId,
            _msgSender(),
            selection,
            reason
        );

        bytes32 digest = keccak256(
            abi.encodePacked("\x19\x01", domainSeparator, structHash)
        );

        address voter = ecrecover(digest, v, r, s);
        emit Console("voter", voter);
        require(voter != address(0), "Invalid signature");
        require(_msgSender() == voter, "You are not the origin voter");
        _castVote(proposalId, voter, selection, reason);
    }

    function _castVote(
        uint256 proposalId,
        address voter,
        VoteState selection,
        string memory reason
    ) internal virtual {
        require(
            state(proposalId) == ProposalState.ACTIVE,
            "Governor: vote not currently active"
        );
        _countVote(proposalId, voter, selection, reason);
        emit Voted(voter, proposalId, selection, reason);
    }

    function _countVote(
        uint256 proposalId,
        address voter,
        VoteState selection,
        string memory reason
    ) internal virtual {
        Proposal storage proposal = _proposals[proposalId];
        require(
            !isVoterVoted(proposalId, voter),
            "You have voted this proposal already."
        );
        Voter memory voterObj = Voter(selection, reason, true);
        if (selection == VoteState.AGREE) {
            proposal.upVoteCount += 1;
        } else {
            proposal.downVoteCount += 1;
        }
        _voters[proposalId][voter] = voterObj;
    }

    //getter, setter
    function isVoterVoted(
        uint256 proposalId,
        address voterId
    ) public view returns (bool) {
        return _voters[proposalId][voterId].isVoted;
    }

    function setVotingDelay(uint256 delay) public {
        _votingDelay = delay;
    }

    function setVotingPeriod(uint256 period) public {
        _votingPeriod = period;
    }

    function getProposalVoteStart(
        uint256 proposalId
    ) public view returns (uint256) {
        return _proposals[proposalId].voteStart.getDeadline();
    }

    function getProposalDeadline(
        uint256 proposalId
    ) public view returns (uint256) {
        return _proposals[proposalId].voteEnd.getDeadline();
    }

    function setStakePool(
        address _sp,
        bytes32[] calldata merkleProof
    ) public onlyTeam(merkleProof) {
        sp = IStakePool(_sp);
    }

    function getProposalList() public view returns (uint256[] memory) {
        return proposalsList;
    }

    function getProposal(
        uint256 proposalId
    ) public view returns (Proposal memory) {
        return _proposals[proposalId];
    }
}
