package feed

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
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	CreatedBy   uint   `json:"created_by" validate:"required"`
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
	*models.Feed
}

func NewFeedResponse(feed *models.Feed) *Response {
	resp := &Response{Feed: feed}

	return resp
}

func (res *Response) Render(w http.ResponseWriter, r *http.Request) error {
	if res.Feed == nil {
		return errors.New("missing required feed")
	}
	return nil
}

func NewFeedListResponse(feeds []models.Feed) []render.Renderer {
	list := []render.Renderer{}
	for idx := range feeds {
		list = append(list, NewFeedResponse(&feeds[idx]))
	}
	return list
}

func GetFeed(w http.ResponseWriter, r *http.Request) {
	feedType := chi.URLParam(r, "feedType")
	if feedType != "loan" && feedType != "governance" && feedType != "account" {
		utils.Render(w, r, utils.ErrBadRequest)
		return
	}

	id := chi.URLParam(r, "id")

	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.Render(w, r, utils.ErrBadRequest)
		return
	}

	feed, err := models.GetFeed(uint(intID), feedType)
	if err != nil {
		utils.Render(w, r, utils.ErrInternalServerWithErr(err))
		return
	}

	if err := render.RenderList(w, r, NewFeedListResponse(feed)); err != nil {
		utils.Render(w, r, utils.ErrRender(err))
		return
	}
}

func CreateFeed(w http.ResponseWriter, r *http.Request) {
	feedType := chi.URLParam(r, "feedType")
	if feedType != "loan" && feedType != "governance" && feedType != "account" {
		utils.Render(w, r, utils.ErrBadRequest)
		return
	}

	id := chi.URLParam(r, "id")

	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
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

	f := &models.Feed{
		Title:       req.Title,
		Description: req.Description,
		CreatedBy:   req.CreatedBy,
		FeedType:    feedType,
	}

	if feedType == "loan" {
		f.LoanID = uint(intID)
	}
	if feedType == "governance" {
		f.GovernanceID = uint(intID)
	}
	if feedType == "account" {
		f.AccountID = uint(intID)
	}

	if err := f.Create(); err != nil {
		utils.Render(w, r, utils.ErrInternalServerWithErr(err))
		return
	}

	utils.Render(w, r, NewFeedResponse(f))
}
