package loan

import (
	"errors"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"net/http"
	"steller-api/pkg/database/models"
	"steller-api/pkg/utils"
)

type Request struct {
	BorrowerID         uint    `json:"borrower_id" validate:"required"`
	GoalAmount         float64 `json:"goal_amount" validate:"required"`
	NumberOfPayments   int     `json:"number_of_payments" validate:"required"`
	PaymentSchedule    string  `json:"payment_schedule" validate:"required"`
	InterestRate       float64 `json:"interest_rate" validate:"required"`
	Title              string  `json:"title" validate:"required"`
	Description        string  `json:"description" validate:"required"`
	ImageURL           string  `json:"image_url" validate:"required"`
	LoanTokenAssetCode string  `json:"loan_token_asset_code"`
}

func (req *Request) Validate() error {
	if err := validator.New().Struct(req); err != nil {
		return err
	}

	return nil
}

func (req *Request) Bind(r *http.Request) error {
	return nil
}

type Response struct {
	*models.Loan
}

func NewLoanResponse(loan *models.Loan) *Response {
	resp := &Response{Loan: loan}

	return resp
}

func (res *Response) Render(w http.ResponseWriter, r *http.Request) error {
	if res.Loan == nil {
		return errors.New("missing required loan")
	}
	return nil
}

func NewLoanListResponse(loans []models.Loan) []render.Renderer {
	list := []render.Renderer{}
	for idx := range loans {
		list = append(list, NewLoanResponse(&loans[idx]))
	}
	return list
}

func GetInactiveLoans(w http.ResponseWriter, r *http.Request) {
	loans, err := models.GetInactiveLoans()
	if err != nil {
		utils.Render(w, r, utils.ErrInternalServerWithErr(err))
		return
	}

	if err := render.RenderList(w, r, NewLoanListResponse(loans)); err != nil {
		utils.Render(w, r, utils.ErrRender(err))
		return
	}
}

func CreateLoan(w http.ResponseWriter, r *http.Request) {
	req := &Request{}
	if err := render.Bind(r, req); err != nil {
		utils.Render(w, r, utils.ErrBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		utils.Render(w, r, utils.ErrBadRequestWithErr(err))
		return
	}

	// Verify borrower exists
	borrower := &models.Account{ID: req.BorrowerID}
	if err := borrower.Get(); err != nil {
		utils.Render(w, r, utils.ErrBadRequestWithErr(err))
		return
	}

	loanTokenAssetCode, err := gonanoid.Generate("1234567890ABCDEFHIJKLMNOPQRSTUVWXYZ", 56)
	if err != nil {
		utils.Render(w, r, utils.ErrInternalServerWithErr(err))
		return
	}

	loan := &models.Loan{
		BorrowerID:         req.BorrowerID,
		GoalAmount:         req.GoalAmount,
		AmountRaised:       0,
		NumberOfPayments:   req.NumberOfPayments,
		PaymentSchedule:    req.PaymentSchedule,
		InterestRate:       req.InterestRate,
		Title:              req.Title,
		Description:        req.Description,
		ImageURL:           req.ImageURL,
		LoanTokenAssetCode: loanTokenAssetCode,
	}

	if err := loan.Create(); err != nil {
		utils.Render(w, r, utils.ErrInternalServerWithErr(err))
		return
	}

	utils.Render(w, r, NewLoanResponse(loan))
}

func GetLoans(w http.ResponseWriter, r *http.Request) {
	loans, err := models.GetLoans()
	if err != nil {
		utils.Render(w, r, utils.ErrInternalServerWithErr(err))
		return
	}

	if err := render.RenderList(w, r, NewLoanListResponse(loans)); err != nil {
		utils.Render(w, r, utils.ErrRender(err))
		return
	}
}
