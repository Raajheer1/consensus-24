package loan_lender

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"net/http"
	"steller-api/pkg/database/models"
	"steller-api/pkg/utils"
	"strconv"
)

type Request struct {
	LenderID   uint    `json:"lender_id" validate:"required"`
	LoanAmount float64 `json:"loan_amount" validate:"required"`
}

func (req *Request) Validate() error {
	if req.LoanAmount <= 0 {
		return errors.New("loan amount is required")
	}
	if err := validator.New().Struct(req); err != nil {
		return err
	}

	return nil
}

func (req *Request) Bind(r *http.Request) error {
	return nil
}

type Response struct {
	*models.LoanLender
}

func NewLoanLenderResponse(loanLender *models.LoanLender) *Response {
	resp := &Response{LoanLender: loanLender}

	return resp
}

func (res *Response) Render(w http.ResponseWriter, r *http.Request) error {
	if res.LoanLender == nil {
		return errors.New("missing required loan_lender")
	}
	return nil
}

func NewLoanLenderListResponse(loanLenders []models.LoanLender) []render.Renderer {
	list := []render.Renderer{}
	for idx := range loanLenders {
		list = append(list, NewLoanLenderResponse(&loanLenders[idx]))
	}
	return list
}

func GetByLenderID(w http.ResponseWriter, r *http.Request) {
	lenderID := chi.URLParam(r, "id")

	intLenderID, err := strconv.ParseInt(lenderID, 10, 64)
	if err != nil {
		utils.Render(w, r, utils.ErrBadRequest)
		return
	}

	loanLenders, err := models.GetLoanLendersByLenderID(uint(intLenderID))
	if err != nil {
		utils.Render(w, r, utils.ErrInternalServerWithErr(err))
		return
	}

	if err := render.RenderList(w, r, NewLoanLenderListResponse(loanLenders)); err != nil {
		utils.Render(w, r, utils.ErrRender(err))
		return
	}
}

func GetLoanLenders(w http.ResponseWriter, r *http.Request) {
	loanID := chi.URLParam(r, "loanId")

	intLoanID, err := strconv.ParseInt(loanID, 10, 64)
	if err != nil {
		utils.Render(w, r, utils.ErrBadRequest)
		return
	}

	loan := &models.Loan{ID: uint(intLoanID)}
	if err := loan.Get(); err != nil {
		utils.Render(w, r, utils.ErrBadRequestWithErr(err))
		return
	}

	loanLenders, err := models.GetLoanLendersByLoanID(uint(intLoanID))
	if err != nil {
		utils.Render(w, r, utils.ErrInternalServerWithErr(err))
		return
	}

	if err := render.RenderList(w, r, NewLoanLenderListResponse(loanLenders)); err != nil {
		utils.Render(w, r, utils.ErrRender(err))
		return
	}
}

func CreateLoanLender(w http.ResponseWriter, r *http.Request) {
	loanID := chi.URLParam(r, "loanId")

	intLoanID, err := strconv.ParseInt(loanID, 10, 64)
	if err != nil {
		utils.Render(w, r, utils.ErrBadRequest)
		return
	}

	loan := &models.Loan{ID: uint(intLoanID)}
	if err := loan.Get(); err != nil {
		utils.Render(w, r, utils.ErrBadRequestWithErr(err))
		return
	}

	req := &Request{}
	if err := render.Bind(r, req); err != nil {
		utils.Render(w, r, utils.ErrBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		utils.Render(w, r, utils.ErrBadRequestWithErr(err))
		return
	}

	// Make sure there is enough balance on the loan
	if req.LoanAmount > loan.GoalAmount-loan.AmountRaised {
		utils.Render(w, r, utils.ErrBadRequestWithErr(errors.New("loan amount exceeds the balance left on the loan")))
		return
	}

	// Verify lender exists
	lender := &models.Account{ID: req.LenderID}
	if err := lender.Get(); err != nil {
		utils.Render(w, r, utils.ErrBadRequestWithErr(err))
		return
	}

	loanLender := &models.LoanLender{
		LoanID:     loan.ID,
		LenderID:   lender.ID,
		LoanAmount: req.LoanAmount,
	}

	if err := loanLender.Create(); err != nil {
		utils.Render(w, r, utils.ErrInternalServerWithErr(err))
		return
	}

	// Update the loan amount raised
	loan.AmountRaised += req.LoanAmount
	if err := loan.Update(); err != nil {
		utils.Render(w, r, utils.ErrInternalServerWithErr(err))
		return
	}

	utils.Render(w, r, NewLoanLenderResponse(loanLender))
}
