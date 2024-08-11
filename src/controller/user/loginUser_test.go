package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	model "github.com/BrunoPolaski/go-crud/src/model/user"
	"github.com/BrunoPolaski/go-crud/src/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserController_LoginUserController(t *testing.T) {
	ctrl := gomock.NewController(t)
	service := mocks.NewMockUserDomainService(ctrl)
	controller := NewUserController(service)

	t.Run("shall_return_error_when_base64_is_incorrect", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := mocks.GetTestGinContext(recorder)

		mocks.MakeRequest(context, nil, nil, "POST", nil, nil)

		controller.LoginUserController(context)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("shall_return_error_when_service_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := mocks.GetTestGinContext(recorder)

		mocks.MakeRequest(context, nil, nil, "POST", nil, &mocks.BasicAuth{
			Username: "user",
			Password: "password",
		})

		login := model.NewUserLoginDomain(
			"user",
			"password",
		)

		service.EXPECT().LoginUserService(login).Return(
			nil,
			"",
			rest_err.NewBadRequestError("Error logging in user"),
		)

		controller.LoginUserController(context)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("shall_return_success_when_login_is_fine", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := mocks.GetTestGinContext(recorder)

		mocks.MakeRequest(context, nil, nil, "POST", nil, &mocks.BasicAuth{
			Username: "user",
			Password: "password",
		})

		login := model.NewUserLoginDomain(
			"user",
			"password",
		)

		service.EXPECT().LoginUserService(login).Return(
			model.NewUserDomain(
				"user",
				"password",
				"bruno",
				19,
			),
			"token",
			nil,
		)

		controller.LoginUserController(context)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
