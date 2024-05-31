package models

import (
	"steller-api/pkg/database"
	"time"
)

type Vote struct {
	ID               uint      `gorm:"primary_key" json:"id"`
	GovernanceVoteID uint      `json:"governance_vote_id"`
	LenderID         uint      `json:"lender_id"`
	Approves         bool      `json:"approves" gorm:"not null;"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (v *Vote) Create() error {
	return database.DB.Create(v).Error
}

func (v *Vote) Get() error {
	return database.DB.First(v).Error
}

func GetVotesByGovernanceVoteID(governanceVoteID uint) ([]Vote, error) {
	var votes []Vote
	if err := database.DB.Where("governance_vote_id = ?", governanceVoteID).Find(&votes).Error; err != nil {
		return nil, err
	}

	return votes, nil
}

func (v *Vote) Update() error {
	return database.DB.Save(v).Error
}

func (v *Vote) Delete() error {
	return database.DB.Delete(v).Error
}
