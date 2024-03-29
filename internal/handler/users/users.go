package users

import (
	"net/http"

	"github.com/eifzed/joona/internal/entity/users"
	"github.com/eifzed/joona/lib/common/commonerr"
)

func (h *UsersHandler) RegisterNewAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	registerData := users.UserRegistration{}
	err := bindingBind(r, &registerData)
	if err != nil {
		err = commonerr.ErrorBadRequest("invalid_params", "invalid registration params")
		commonwriterRespondError(ctx, w, err)
		return
	}
	if err := registerData.ValidateInput(); err != nil {
		commonwriterRespondError(ctx, w, err)
		return
	}

	auth, err := h.UsersUC.RegisterNewUser(ctx, registerData)
	if err != nil {
		commonwriterRespondError(ctx, w, err)
		return
	}
	commonwriterRespondOKWithData(ctx, w, auth)
}

func (h *UsersHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loginData := users.UserLogin{}
	err := bindingBind(r, &loginData)
	if err != nil {
		err = commonerr.ErrorBadRequest("invalid_params", "invalid login params")
		commonwriterRespondError(ctx, w, err)
		return
	}

	auth, err := h.UsersUC.LoginUser(ctx, loginData)
	if err != nil {
		commonwriterRespondError(ctx, w, err)
		return
	}
	commonwriterRespondOKWithData(ctx, w, auth)
}
