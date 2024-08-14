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
	"github.com/BrunoPolaski/go-crud/src/tests/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserController_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	service := mocks.NewMockUserDomainService(ctrl)
	controller := NewUserController(service)

	t.Run("shall_return_error_when_service_fails", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := mocks.GetTestGinContext(recorder)

		params := gin.Params{
			{
				Key: "userId", Value: primitive.NewObjectID().Hex(),
			},
		}

		body := request.UserRequest{
			Email:    "test@test.com",
			Password: "Pwd123#",
			Age:      19,
			Name:     "john doe",
		}

		json, _ := json.Marshal(body)
		stringReader := io.NopCloser(strings.NewReader(string(json)))

		mocks.MakeRequest(context, params, nil, "PUT", stringReader, nil)

		service.EXPECT().UpdateUserService(gomock.Any(), gomock.Any()).Return(
			rest_err.NewInternalServerError("Error trying to update user"),
		)

		controller.UpdateUserController(context)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("shall_return_error_when_id_is_invalid", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := mocks.GetTestGinContext(recorder)

		params := gin.Params{
			{
				Key: "userId", Value: "123",
			},
		}

		mocks.MakeRequest(context, params, nil, "PUT", nil, nil)

		controller.UpdateUserController(context)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("shall_return_error_when_body_is_invalid", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := mocks.GetTestGinContext(recorder)

		params := gin.Params{
			{
				Key: "userId", Value: primitive.NewObjectID().Hex(),
			},
		}

		body := request.UserUpdateRequest{
			Name: "john doe",
			Age:  -1,
		}

		json, _ := json.Marshal(body)
		stringReader := io.NopCloser(strings.NewReader(string(json)))

		mocks.MakeRequest(context, params, nil, "PUT", stringReader, nil)

		controller.UpdateUserController(context)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("shall_return_success_when_user_is_updated", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := mocks.GetTestGinContext(recorder)

		params := gin.Params{
			{
				Key: "userId", Value: primitive.NewObjectID().Hex(),
			},
		}

		body := request.UserUpdateRequest{
			Name: "john doe",
			Age:  19,
		}

		json, _ := json.Marshal(body)
		stringReader := io.NopCloser(strings.NewReader(string(json)))

		mocks.MakeRequest(context, params, nil, "PUT", stringReader, nil)

		service.EXPECT().UpdateUserService(gomock.Any(), gomock.Any()).Return(nil)

		controller.UpdateUserController(context)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
