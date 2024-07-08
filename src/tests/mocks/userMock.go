package mocks

import model "github.com/BrunoPolaski/go-crud/src/model/user"

var UserMock model.UserDomainInterface = model.NewUserDomain(
	"test@test.com",
	"test",
	"bruno",
	19,
)
