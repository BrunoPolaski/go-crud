package controller

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	model "github.com/BrunoPolaski/go-crud/src/model/user"
	"github.com/BrunoPolaski/go-crud/src/tests/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserController_FindUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	service := mocks.NewMockUserDomainService(ctrl)
	controller := NewUserController(service)
	t.Run("shall_return_error_when_email_is_incorrect", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := mocks.GetTestGinContext(recorder)

		param := gin.Params{
			{
				Key:   "userEmail",
				Value: "TEST_ERROR",
			},
		}

		mocks.MakeRequest(context, param, url.Values{}, "GET", nil, nil)
		controller.FindUserByEmailController(context)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("shall_return_error_when_service_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := mocks.GetTestGinContext(recorder)

		param := gin.Params{
			{
				Key:   "userEmail",
				Value: "test@test.com",
			},
		}

		service.EXPECT().FindUserByEmailService("test@test.com").Return(
			nil,
			rest_err.NewInternalServerError("Error finding user by email"),
		)

		mocks.MakeRequest(context, param, url.Values{}, "GET", nil, nil)
		controller.FindUserByEmailController(context)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("shall_return_success_when_email_is_found", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := mocks.GetTestGinContext(recorder)

		param := gin.Params{
			{
				Key:   "userEmail",
				Value: "test@test.com",
			},
		}

		service.EXPECT().FindUserByEmailService("test@test.com").Return(
			model.NewUserDomain(
				"email",
				"password",
				"name",
				19,
			),
			nil,
		)

		mocks.MakeRequest(context, param, url.Values{}, "GET", nil, nil)
		controller.FindUserByEmailController(context)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestUserController_FindUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	service := mocks.NewMockUserDomainService(ctrl)
	controller := NewUserController(service)

	t.Run("shall_return_error_when_id_is_incorrect", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := mocks.GetTestGinContext(recorder)

		param := gin.Params{
			{
				Key:   "userId",
				Value: "TEST_ERROR",
			},
		}

		mocks.MakeRequest(context, param, url.Values{}, "GET", nil, nil)
		controller.FindUserByIdController(context)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("shall_return_error_when_service_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := mocks.GetTestGinContext(recorder)

		id := primitive.NewObjectID().Hex()

		param := gin.Params{
			{
				Key:   "userId",
				Value: id,
			},
		}

		service.EXPECT().FindUserByIDService(id).Return(
			nil,
			rest_err.NewInternalServerError("Error finding user by email"),
		)

		mocks.MakeRequest(context, param, url.Values{}, "GET", nil, nil)
		controller.FindUserByIdController(context)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("shall_return_success_when_id_is_found", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := mocks.GetTestGinContext(recorder)

		id := primitive.NewObjectID().Hex()

		param := gin.Params{
			{
				Key:   "userId",
				Value: id,
			},
		}

		service.EXPECT().FindUserByIDService(id).Return(
			model.NewUserDomain(
				"email",
				"password",
				"name",
				19,
			),
			nil,
		)

		mocks.MakeRequest(context, param, url.Values{}, "GET", nil, nil)
		controller.FindUserByIdController(context)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
