package votes

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
	LenderID uint `json:"lender_id" validate:"required"`
	Approves bool `json:"approves"`
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
	*models.Vote
}

func NewVoteResponse(vote *models.Vote) *Response {
	resp := &Response{Vote: vote}

	return resp
}

func (res *Response) Render(w http.ResponseWriter, r *http.Request) error {
	if res.Vote == nil {
		return errors.New("missing required vote")
	}
	return nil
}

func NewVoteListResponse(votes []models.Vote) []render.Renderer {
	list := []render.Renderer{}
	for idx := range votes {
		list = append(list, NewVoteResponse(&votes[idx]))
	}
	return list
}

func GetVotes(w http.ResponseWriter, r *http.Request) {
	loanId := chi.URLParam(r, "loanId")
	govId := chi.URLParam(r, "govId")

	intLoanID, err := strconv.ParseInt(loanId, 10, 64)
	if err != nil {
		utils.Render(w, r, utils.ErrBadRequest)
		return
	}

	intGovID, err := strconv.ParseInt(govId, 10, 64)
	if err != nil {
		utils.Render(w, r, utils.ErrBadRequest)
		return
	}

	gv := &models.GovernanceVote{ID: uint(intGovID)}
	if err := gv.Get(); err != nil {
		utils.Render(w, r, utils.ErrBadRequestWithErr(err))
		return
	}

	if gv.LoanID != uint(intLoanID) {
		utils.Render(w, r, utils.ErrBadRequest)
		return
	}

	votes, err := models.GetVotesByGovernanceVoteID(uint(intGovID))
	if err != nil {
		utils.Render(w, r, utils.ErrInternalServerWithErr(err))
		return
	}

	if err := render.RenderList(w, r, NewVoteListResponse(votes)); err != nil {
		utils.Render(w, r, utils.ErrRender(err))
		return
	}
}

func Vote(w http.ResponseWriter, r *http.Request) {
	loanId := chi.URLParam(r, "loanId")
	govId := chi.URLParam(r, "govId")

	intLoanID, err := strconv.ParseInt(loanId, 10, 64)
	if err != nil {
		utils.Render(w, r, utils.ErrBadRequest)
		return
	}

	intGovID, err := strconv.ParseInt(govId, 10, 64)
	if err != nil {
		utils.Render(w, r, utils.ErrBadRequest)
		return
	}

	gv := &models.GovernanceVote{ID: uint(intGovID)}
	if err := gv.Get(); err != nil {
		utils.Render(w, r, utils.ErrBadRequestWithErr(err))
		return
	}

	if gv.LoanID != uint(intLoanID) {
		utils.Render(w, r, utils.ErrBadRequest)
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

	vote := &models.Vote{
		GovernanceVoteID: gv.ID,
		LenderID:         req.LenderID,
		Approves:         req.Approves,
	}

	if err := vote.Create(); err != nil {
		utils.Render(w, r, utils.ErrInternalServerWithErr(err))
		return
	}

	if err := render.Render(w, r, NewVoteResponse(vote)); err != nil {
		utils.Render(w, r, utils.ErrRender(err))
		return
	}
}
