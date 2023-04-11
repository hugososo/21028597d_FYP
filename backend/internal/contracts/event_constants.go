package contracts

import "github.com/ethereum/go-ethereum/crypto"

var (
	ProposalCreatedBytes     = []byte("ProposalCreated(uint256)")
	VotedBytes               = []byte("Voted(address,uint256,uint8,string)")
	ProposalCreatedBytesHash = crypto.Keccak256Hash(ProposalCreatedBytes)
	VotedBytesHash           = crypto.Keccak256Hash(VotedBytes)
	BookCreatedBytes         = []byte("BookCreated(address,address,string)")
	BookCreatedBytesHash     = crypto.Keccak256Hash(BookCreatedBytes)
	TransferBytes            = []byte("Transfer(address,address,uint256)")
	TransferBytesHash        = crypto.Keccak256Hash(TransferBytes)
)
