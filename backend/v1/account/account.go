package account

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
	//PublicKey string `json:"public_key"`
	//AccountID uint   `json:"id"`
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
	*models.Account
}

func NewAccountResponse(acc *models.Account) *Response {
	resp := &Response{Account: acc}

	return resp
}

func (res *Response) Render(w http.ResponseWriter, r *http.Request) error {
	if res.Account == nil {
		return errors.New("missing required account")
	}
	return nil
}

func NewAccountListResponse(accounts []models.Account) []render.Renderer {
	list := []render.Renderer{}
	for idx := range accounts {
		list = append(list, NewAccountResponse(&accounts[idx]))
	}
	return list
}

func GetAccount(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		utils.Render(w, r, utils.ErrBadRequest)
		return
	}

	acc := models.Account{
		PublicKey: id,
	}
	if err := acc.Get(); err == nil {
		utils.Render(w, r, NewAccountResponse(&acc))
		return
	}

	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.Render(w, r, utils.ErrBadRequest)
		return
	}

	acc = models.Account{
		ID: uint(intId),
	}
	if err := acc.Get(); err == nil {
		utils.Render(w, r, NewAccountResponse(&acc))
		return
	}

	utils.Render(w, r, utils.ErrNotFound)
}
