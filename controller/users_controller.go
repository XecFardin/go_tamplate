package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/XecFardin/go_tamplate/services"
	"github.com/XecFardin/go_tamplate/utils"
)

func GetUsers(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("id"), 10, 64)
	if err != nil {
		userError, _ := json.Marshal(&utils.ApplicationError{
			Message:    "id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		})
		resp.Write(userError)
		return
	}
	user, Apierr := services.GetUser(userId)
	if Apierr != nil {
		ApiError, _ := json.Marshal(Apierr)
		resp.WriteHeader(Apierr.StatusCode)
		resp.Write([]byte(ApiError))
		return
	}
	jsonValue, _ := json.Marshal(user)
	resp.WriteHeader(http.StatusOK)
	resp.Write(jsonValue)

}
