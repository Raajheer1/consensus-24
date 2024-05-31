package models

import (
	"errors"
	"steller-api/pkg/database"
	"time"
)

type Feed struct {
	ID            uint      `gorm:"primary_key" json:"id"`
	AccountID     uint      `json:"account_id"`
	LoanID        uint      `json:"loan_id"`
	GovernanceID  uint      `json:"governance_id"`
	FeedType      string    `json:"feed_type"` // ENUM: [loan, governance, account]
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
	CreatedBy     uint      `json:"created_by"`
	CreatedByUser Account   `gorm:"foreignKey:CreatedBy" json:"created_by_user"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (f *Feed) Validate() error {
	if f.FeedType != "loan" && f.FeedType != "governance" && f.FeedType != "account" {
		return errors.New("invalid feed type")
	}
	if f.FeedType == "loan" {
		if f.LoanID == 0 {
			return errors.New("missing loan id")
		} else {
			loan := Loan{}
			database.DB.Where("id = ?", f.LoanID).First(&loan)
			if loan.ID == 0 {
				return errors.New("loan not found")
			}
		}
	}
	if f.FeedType == "governance" {
		if f.GovernanceID == 0 {
			return errors.New("missing governance id")
		} else {
			gov := GovernanceVote{}
			database.DB.Where("id = ?", f.GovernanceID).First(&gov)
			if gov.ID == 0 {
				return errors.New("governance not found")
			}
		}
	}
	if f.FeedType == "account" {
		if f.AccountID == 0 {
			return errors.New("missing account id")
		} else {
			account := Account{}
			database.DB.Where("id = ?", f.AccountID).First(&account)
			if account.ID == 0 {
				return errors.New("account not found")
			}
		}
	}
	return nil
}

func (f *Feed) Create() error {
	if err := f.Validate(); err != nil {
		return err
	}
	return database.DB.Create(f).Error
}

func GetFeed(id uint, feedType string) ([]Feed, error) {
	var feeds []Feed
	if feedType == "loan" {
		database.DB.Where("loan_id = ?", id).Preload("CreatedByUser").Find(&feeds)
		return feeds, nil
	}
	if feedType == "governance" {
		database.DB.Where("governance_id = ?", id).Preload("CreatedByUser").Find(&feeds)
		return feeds, nil
	}
	if feedType == "account" {
		database.DB.Where("account_id = ?", id).Preload("CreatedByUser").Find(&feeds)
		return feeds, nil
	}

	return nil, errors.New("invalid feed")
}

func (f *Feed) Update() error {
	if err := f.Validate(); err != nil {
		return err
	}
	return database.DB.Save(f).Error
}

func (f *Feed) Delete() error {
	return database.DB.Delete(f).Error
}
