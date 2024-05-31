package models

import (
	"steller-api/pkg/database"
	"time"
)

type GovernanceVote struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	LoanID       uint      `json:"loan_id"`
	AmountRaised float64   `json:"amount_raised"`
	VotingActive bool      `json:"voting_active"`
	VotePassed   bool      `json:"vote_passed"`
	UUID         string    `json:"uuid"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Votes        []Vote    `json:"votes" gorm:"foreignKey:GovernanceVoteID"`
	EndAt        time.Time `json:"end_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (gv *GovernanceVote) Create() error {
	return database.DB.Create(gv).Error
}

func (gv *GovernanceVote) Get() error {
	return database.DB.Preload("Votes").First(gv).Error
}

func GetActiveGVs() ([]GovernanceVote, error) {
	var votes []GovernanceVote
	if err := database.DB.Where("voting_active = ?", true).Find(&votes).Error; err != nil {
		return nil, err
	}

	return votes, nil
}

func GetGVsByLoanID(loanID uint) ([]GovernanceVote, error) {
	var votes []GovernanceVote
	if err := database.DB.Where("loan_id = ?", loanID).Find(&votes).Error; err != nil {
		return nil, err
	}

	return votes, nil
}

func (gv *GovernanceVote) Update() error {
	return database.DB.Save(gv).Error
}

func (gv *GovernanceVote) Delete() error {
	return database.DB.Delete(gv).Error
}

func (gv *GovernanceVote) CompletionCheck() error {
	if gv.Votes == nil {
		votes, err := GetVotesByGovernanceVoteID(gv.ID)
		if err != nil {
			return err
		}

		gv.Votes = votes
	}

	if err := gv.UpdateVotePassed(); err != nil {
		return err
	}

	if gv.EndAt.Before(time.Now()) || gv.VotePassed {
		gv.VotingActive = false
	}

	if err := gv.Update(); err != nil {
		return err
	}

	return nil
}

// UpdateVotePassed updates the vote passed status based on the votes
// In order for a vote to pass 75% of the votes must be in favor
// loanHolders who do not vote are assumed to vote against.
func (gv *GovernanceVote) UpdateVotePassed() error {
	loan := &Loan{ID: gv.LoanID}
	if err := loan.Get(); err != nil {
		return err
	}

	totalVotes := loan.AmountRaised
	requiredVotes := totalVotes * 0.75
	votesInFavor := 0.0

	for _, vote := range gv.Votes {
		if vote.Approves {
			for _, loanLender := range loan.Lenders {
				if loanLender.LenderID == vote.LenderID {
					votesInFavor += loanLender.LoanAmount
				}
			}
		}
	}

	if votesInFavor > requiredVotes {
		gv.VotePassed = true

		// Only funded if title & description is empty
		if gv.Title == "" {
			loan.FundedAt = time.Now()
			if err := loan.Update(); err != nil {
				return err
			}

			// TODO - ANDREW ADD CALL TO SMART CONTRACT HERE
		}
	}

	return nil
}
