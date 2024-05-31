package gov

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"net/http"
	"steller-api/pkg/database/models"
	"steller-api/pkg/utils"
	"strconv"
	"time"
)

type Request struct {
	Title       string `json:"title"`
	Description string `json:"description"`
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
	*models.GovernanceVote
}

func NewGovVoteResponse(gv *models.GovernanceVote) *Response {
	resp := &Response{GovernanceVote: gv}

	return resp
}

func (res *Response) Render(w http.ResponseWriter, r *http.Request) error {
	if res.GovernanceVote == nil {
		return errors.New("missing required governance_vote")
	}
	return nil
}

func NewGovVoteListResponse(gvs []models.GovernanceVote) []render.Renderer {
	list := []render.Renderer{}
	for idx := range gvs {
		list = append(list, NewGovVoteResponse(&gvs[idx]))
	}
	return list
}

func GetGovVotes(w http.ResponseWriter, r *http.Request) {
	loanId := chi.URLParam(r, "loanId")

	intLoanID, err := strconv.ParseInt(loanId, 10, 64)
	if err != nil {
		utils.Render(w, r, utils.ErrBadRequest)
		return
	}

	loan := &models.Loan{ID: uint(intLoanID)}
	if err := loan.Get(); err != nil {
		utils.Render(w, r, utils.ErrBadRequestWithErr(err))
		return
	}

	gvs, err := models.GetGVsByLoanID(uint(intLoanID))
	if err != nil {
		utils.Render(w, r, utils.ErrInternalServerWithErr(err))
		return
	}

	if err := render.RenderList(w, r, NewGovVoteListResponse(gvs)); err != nil {
		utils.Render(w, r, utils.ErrRender(err))
		return
	}
}

func CreateGovVote(w http.ResponseWriter, r *http.Request) {
	loanId := chi.URLParam(r, "loanId")

	intLoanID, err := strconv.ParseInt(loanId, 10, 64)
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

	// Check if funding vote, if funding vote 75% of goal must've been reached
	if req.Title == "" {
		if loan.AmountRaised < loan.GoalAmount*0.75 {
			utils.Render(w, r, utils.ErrBadRequestWithErr(errors.New("funding goal not reached")))
			return
		}
	}

	gv := &models.GovernanceVote{
		LoanID:       loan.ID,
		AmountRaised: loan.AmountRaised,
		VotingActive: true,
		VotePassed:   false,
		UUID:         uuid.NewString(),
		Title:        req.Title,
		Description:  req.Description,
		EndAt:        time.Now().AddDate(0, 0, 7),
	}

	if err := gv.Create(); err != nil {
		utils.Render(w, r, utils.ErrInternalServerWithErr(err))
		return
	}

	utils.Render(w, r, NewGovVoteResponse(gv))
}
