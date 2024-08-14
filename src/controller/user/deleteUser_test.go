package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"github.com/BrunoPolaski/go-crud/src/tests/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserController_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	service := mocks.NewMockUserDomainService(ctrl)
	controller := NewUserController(service)

	t.Run("shall_return_error_when_service_fails", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := mocks.GetTestGinContext(recorder)

		id := primitive.NewObjectID().Hex()

		param := gin.Params{
			{
				Key:   "userId",
				Value: id,
			},
		}

		mocks.MakeRequest(context, param, nil, "DELETE", nil, nil)

		service.EXPECT().DeleteUserService(id).Return(
			rest_err.NewInternalServerError("Error trying to delete user"),
		)

		controller.DeleteUserController(context)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("shall_return_error_when_id_is_invalid", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := mocks.GetTestGinContext(recorder)

		param := gin.Params{
			{
				Key:   "userId",
				Value: "invalid_id",
			},
		}

		mocks.MakeRequest(context, param, nil, "DELETE", nil, nil)

		controller.DeleteUserController(context)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("shall_return_success_when_id_is_valid", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := mocks.GetTestGinContext(recorder)

		id := primitive.NewObjectID().Hex()

		param := gin.Params{
			{
				Key:   "userId",
				Value: id,
			},
		}

		mocks.MakeRequest(context, param, nil, "DELETE", nil, nil)
		service.EXPECT().DeleteUserService(id).Return(nil)
		controller.DeleteUserController(context)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
