package model

type userDomain struct {
	id       string
	email    string
	password string
	name     string
	age      int8
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetAge() int8 {
	return ud.age
}

func (ud *userDomain) GetId() string {
	return ud.id
}

func (ud *userDomain) SetId(id string) {
	ud.id = id
}

func (ud *userDomain) SetEmail(email string) {
	ud.email = email
}

func (ud *userDomain) SetPassword(password string) {
	ud.password = password
}

func (ud *userDomain) SetName(name string) {
	ud.name = name
}

func (ud *userDomain) SetAge(age int8) {
	ud.age = age
}
