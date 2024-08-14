package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"github.com/BrunoPolaski/go-crud/src/controller/model/request"
	model "github.com/BrunoPolaski/go-crud/src/model/user"
	"github.com/BrunoPolaski/go-crud/src/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserController_CreateUserController(t *testing.T) {
	ctrl := gomock.NewController(t)
	service := mocks.NewMockUserDomainService(ctrl)
	controller := NewUserController(service)

	t.Run("shall_return_error_when_body_is_wrong", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := mocks.GetTestGinContext(recorder)

		body := request.UserRequest{
			Email:    "email",
			Password: "pwd",
			Age:      0,
			Name:     "john doe",
		}

		json, _ := json.Marshal(body)
		stringReader := io.NopCloser(strings.NewReader(string(json)))

		mocks.MakeRequest(context, nil, nil, "POST", stringReader, nil)
		controller.CreateUserController(context)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("shall_return_error_when_email_is_found", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := mocks.GetTestGinContext(recorder)

		body := request.UserRequest{
			Email:    "test@test.com",
			Password: "Pwd123#",
			Age:      19,
			Name:     "john doe",
		}

		json, _ := json.Marshal(body)
		stringReader := io.NopCloser(strings.NewReader(string(json)))

		service.EXPECT().CreateUserService(gomock.Any()).Return(
			nil,
			rest_err.NewBadRequestError("User already exists"),
		)

		mocks.MakeRequest(context, nil, nil, "POST", stringReader, nil)
		controller.CreateUserController(context)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("shall_return_success_when_user_is_created", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := mocks.GetTestGinContext(recorder)

		body := request.UserRequest{
			Email:    "test@test.com",
			Password: "Pwd123#",
			Age:      19,
			Name:     "john doe",
		}

		json, _ := json.Marshal(body)
		stringReader := io.NopCloser(strings.NewReader(string(json)))

		service.EXPECT().CreateUserService(gomock.Any()).Return(
			model.NewUserDomain(
				body.Email,
				body.Password,
				body.Name,
				body.Age,
			),
			nil,
		)

		mocks.MakeRequest(context, nil, nil, "POST", stringReader, nil)
		controller.CreateUserController(context)

		assert.Equal(t, http.StatusCreated, recorder.Code)
	})
}
