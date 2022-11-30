package factories

import (
	"fmt"

	"github.com/bxcodec/faker/v4"
)

type FakeUser struct {
	Email    string `faker:"email"`
	Nickname string `faker:"username"`
	Password string `faker:"password"`
}

func GenUser() FakeUser {
	user := FakeUser{}
	err := faker.FakeData(&user)
	if err != nil {
		fmt.Println(err)
	}
	return user
}
