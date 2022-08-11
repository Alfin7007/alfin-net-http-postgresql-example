package presentation

import (
	"encoding/json"
	"fmt"
	"http/example/features/users"
	"http/example/features/users/presentation/request"
	"http/example/features/users/presentation/response"
	"http/example/helper"
	"http/example/middlewares"
	"io/ioutil"
	"net/http"
)

type UserHandler struct {
	userBussiness users.Bussiness
}

func NewUserHandler(userBussiness users.Bussiness) *UserHandler {
	return &UserHandler{
		userBussiness: userBussiness,
	}
}

func (uh UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		helper.FailedBadRequestWithMSG(fmt.Sprintf("need content-type 'application/json', but got '%s'", ct), w)
		return
	}

	userRequest := request.UserRequest{}
	bodyBytes, errReq := ioutil.ReadAll(r.Body)

	if errReq != nil {
		helper.FailedBadRequestWithMSG(errReq.Error(), w)
		return
	}
	defer r.Body.Close()

	errUnmarshal := json.Unmarshal(bodyBytes, &userRequest)
	if errUnmarshal != nil {
		helper.FailedBadRequestWithMSG(errUnmarshal.Error(), w)
		return
	}
	userCore := request.ToCore(userRequest)
	userID, token, errLogin := uh.userBussiness.Login(userCore)
	if errLogin != nil {
		helper.FailedNotFound(errLogin.Error(), w)
		return
	}
	helper.AuthOK(userID, token, w)
}

func (uh UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		helper.FailedBadRequestWithMSG(fmt.Sprintf("need content-type 'application/json', but got '%s'", ct), w)
		return
	}
	bodyBytes, errBody := ioutil.ReadAll(r.Body)

	if errBody != nil {
		helper.FailedBadRequestWithMSG(errBody.Error(), w)
		return
	}
	defer r.Body.Close()

	userRequest := request.UserRequest{}
	errUnmarshal := json.Unmarshal(bodyBytes, &userRequest)

	if errUnmarshal != nil {
		helper.FailedBadRequestWithMSG(errUnmarshal.Error(), w)
		return
	}

	userCore := request.ToCore(userRequest)
	err := uh.userBussiness.Register(userCore)
	if err != nil {
		helper.FailedBadRequestWithMSG(err.Error(), w)
		return
	}

	helper.SuccessCreateNoData(w)
}

func (uh UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	userID, _, errJWT := middlewares.JWTTokenCheck(r.Header.Get("Authorization"))

	if errJWT != nil {
		helper.FailedForbidden(w)
		return
	}
	result, err := uh.userBussiness.GetData(userID)
	if err != nil {
		helper.FailedBadRequestWithMSG(err.Error(), w)
		return
	}

	helper.SuccessGetData(response.FromCore(result), w)
}
