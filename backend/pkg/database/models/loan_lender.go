package models

import (
	"errors"
	"steller-api/pkg/database"
	"time"
)

type LoanLender struct {
	ID         uint      `gorm:"primary_key" json:"id"`
	LoanID     uint      `json:"loan_id"`
	Loan       Loan      `json:"loan"`
	LenderID   uint      `json:"lender_id"`
	Lender     Account   `json:"lender"`
	LoanAmount float64   `json:"loan_amount"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (ll *LoanLender) Validate() error {
	if ll.LoanAmount <= 0 {
		return errors.New("loan amount is required")
	}
	return nil
}

func (ll *LoanLender) Create() error {
	if err := ll.Validate(); err != nil {
		return err
	}
	return database.DB.Create(ll).Error
}

func (ll *LoanLender) Get() error {
	return database.DB.Preload("Loan").Preload("Lender").First(ll).Error
}

func GetLoanLendersByLenderID(lenderID uint) ([]LoanLender, error) {
	var loanLenders []LoanLender
	if err := database.DB.Where("lender_id = ?", lenderID).Preload("Loan").Preload("Loan.Borrower").Preload("Loan.GovernanceVotes").Preload("Loan.GovernanceVotes.Votes").Find(&loanLenders).Error; err != nil {

		return nil, err
	}

	return loanLenders, nil
}

func GetLoanLendersByLoanID(loanID uint) ([]LoanLender, error) {
	var loanLenders []LoanLender
	if err := database.DB.Where("loan_id = ?", loanID).Preload("Loan").Preload("Lender").Find(&loanLenders).Error; err != nil {
		return nil, err
	}

	return loanLenders, nil
}

func GetLoanLenderByLenderIDAndLoanID(lenderID, loanID uint) (*LoanLender, error) {
	var loanLender LoanLender
	if err := database.DB.Where("lender_id = ? AND loan_id = ?", lenderID, loanID).Preload("Loan").Preload("Lender").First(&loanLender).Error; err != nil {
		return nil, err
	}

	return &loanLender, nil
}

func (ll *LoanLender) Update() error {
	if err := ll.Validate(); err != nil {
		return err
	}
	return database.DB.Save(ll).Error
}

func (ll *LoanLender) Delete() error {
	return database.DB.Delete(ll).Error
}
