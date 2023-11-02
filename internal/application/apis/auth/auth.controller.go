package auth

import (
	"encoding/json"
	"net/http"

	e "github.com/prcryx/raft-server/internal/common/err"
	utils "github.com/prcryx/raft-server/internal/common/utils"
	"github.com/prcryx/raft-server/internal/domain/types"
	"github.com/prcryx/raft-server/internal/domain/usecases"
)

type IAuthController interface {
	SendOtp(http.ResponseWriter, *http.Request)
	Login(http.ResponseWriter, *http.Request)
}

type AuthController struct {
	authUseCase *usecases.AuthUseCase
}

var _ IAuthController = (*AuthController)(nil)

func NewAuthController(authUseCase *usecases.AuthUseCase) *AuthController {
	return &AuthController{
		authUseCase: authUseCase,
	}
}

func (ac *AuthController) SendOtp(w http.ResponseWriter, request *http.Request) {

	//validating request
	otpReq := new(types.OtpReqBody)
	err := json.NewDecoder(request.Body).Decode(&otpReq)
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, e.InvalidBodyRequestException().Error())
		return
	}
	//checking the dto
	if ok := utils.ValidatePhone(otpReq); !ok {
		utils.ResponseWithError(w, http.StatusBadRequest, e.InvalidBodyRequestException().Error())
		return
	}

	verificationRes, err := ac.authUseCase.SendOtp(*otpReq)
	if err != nil {
		if appErr, ok := err.(e.AppError); ok {
			utils.ResponseWithError(w, appErr.GetCode(), appErr.Error())
			return
		}
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseWithJSONData(w, http.StatusOK, verificationRes)

}

func (ac *AuthController) Login(w http.ResponseWriter, request *http.Request) {
	//validating request
	otpVerificationReq := new(types.OtpVerificationReqBody)
	err := json.NewDecoder(request.Body).Decode(&otpVerificationReq)
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, e.InvalidBodyRequestException().Error())
		return
	}

	verificationRes, err := ac.authUseCase.Login(*otpVerificationReq)
	if err != nil {
		if appErr, ok := err.(e.AppError); ok {
			utils.ResponseWithError(w, appErr.GetCode(), appErr.Error())
			return
		}
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseWithJSONData(w, http.StatusOK, verificationRes)

}
