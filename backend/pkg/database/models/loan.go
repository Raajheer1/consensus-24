package models

import (
	"errors"
	"steller-api/pkg/database"
	"strings"
	"time"
)

type Loan struct {
	ID                 uint             `gorm:"primary_key" json:"id"`
	BorrowerID         uint             `json:"borrower_id"`
	Borrower           Account          `json:"borrower"`
	GoalAmount         float64          `json:"goal_amount"`
	AmountRaised       float64          `json:"amount_raised"`
	NumberOfPayments   int              `json:"number_of_payments"`
	PaymentSchedule    string           `json:"payment_schedule"`
	InterestRate       float64          `json:"interest_rate"`
	Title              string           `json:"title"`
	Description        string           `json:"description"`
	ImageURL           string           `json:"image_url"`
	Lenders            []LoanLender     `json:"lenders" gorm:"foreignKey:LoanID"`
	GovernanceVotes    []GovernanceVote `json:"governance_votes" gorm:"foreignKey:LoanID"`
	LoanTokenAssetCode string           `json:"loan_token_asset_code" gorm:"unique;not null"`
	FundedAt           time.Time        `json:"funded_at"`
	CreatedAt          time.Time        `json:"created_at"`
	UpdatedAt          time.Time        `json:"updated_at"`
}

func (l *Loan) Validate() error {
	if l.GoalAmount <= 0 {
		return errors.New("invalid goal amount")
	}
	if l.NumberOfPayments <= 0 {
		return errors.New("invalid number of payments")
	}
	if l.InterestRate <= 0 {
		return errors.New("invalid interest rate")
	}
	l.PaymentSchedule = strings.ToLower(l.PaymentSchedule)
	if l.PaymentSchedule != "yearly" && l.PaymentSchedule != "weekly" && l.PaymentSchedule != "monthly" {
		return errors.New("invalid payment schedule")
	}
	return nil
}

func (l *Loan) Create() error {
	if err := l.Validate(); err != nil {
		return err
	}
	return database.DB.Create(l).Error
}

func (l *Loan) Get() error {
	return database.DB.Preload("Lenders").Preload("GovernanceVotes").First(l).Error
}

func GetInactiveLoans() ([]Loan, error) {
	var loans []Loan
	if err := database.DB.Where("funded_at IS NULL OR amount_raised = goal_amount").Preload("Lenders").Preload("GovernanceVotes").Find(&loans).Error; err != nil {
		return nil, err
	}

	return loans, nil
}

func GetLoans() ([]Loan, error) {
	var loans []Loan
	if err := database.DB.Preload("Lenders").Preload("GovernanceVotes").Preload("GovernanceVotes.Votes").Find(&loans).Error; err != nil {
		return nil, err
	}

	return loans, nil
}

func (l *Loan) Update() error {
	if err := l.Validate(); err != nil {
		return err
	}
	return database.DB.Save(l).Error
}

func (l *Loan) Delete() error {
	return database.DB.Delete(l).Error
}
