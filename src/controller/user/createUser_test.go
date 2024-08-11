package controller

// func TestUserController_CreateUserController(t *testing.T) {
// ctrl := gomock.NewController(t)
// service := mocks.NewMockUserDomainService(ctrl)
// controller := NewUserController(service)

// t.Run("shall_return_error_when_email_is_found", func(t *testing.T) {
// 	recorder := httptest.NewRecorder()
// 	context := mocks.GetTestGinContext(recorder)

// 	body := request.UserRequest{
// 		Email:    "email",
// 		Password: "pwd",
// 		Age:      0,
// 		Name:     "bruneco",
// 	}

// 	json, _ := json.Marshal(body)
// 	stringReader := io.NopCloser(strings.NewReader(string(json)))

// 	service.EXPECT().CreateUserService(gomock.Any()).Return(nil, rest_err.NewBadRequestError("Error creating user"))

// 	mocks.MakeRequest(context, nil, nil, "POST", stringReader, nil)
// 	controller.CreateUserController(context)

// 	assert.Equal(t, http.StatusBadRequest, recorder.Code)
// })
// }
