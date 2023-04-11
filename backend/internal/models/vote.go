package models

type Vote struct {
	ProposalId string `json:"proposal_id" gorm:"primaryKey"`
	Voter      string `json:"voter" gorm:"primaryKey"`
	Selection  bool   `json:"selection"`
	Reason     string `json:"reason"`
}
