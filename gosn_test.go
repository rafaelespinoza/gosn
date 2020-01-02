package gosn

import (
	"fmt"
	"os"
	"time"
)

const (
	// SN_SERVER is an env var key.
	SN_SERVER = "SN_SERVER"
)

var (
	// _Conf organizes various shared free variables in one namespace.
	_Conf = struct {
		APIServer    string
		GoodEmail    string
		GoodPassword string
	}{
		APIServer:    os.Getenv(SN_SERVER),
		GoodEmail:    makeStubEmail(""),
		GoodPassword: "testpassword1234",
	}

	_SignInInput SignInInput
)

func init() {
	user, err := registerNewUser("shared")
	if err != nil {
		panic(err)
	}
	_SignInInput = SignInInput{
		APIServer: user.APIServer,
		Email:     user.Email,
		Password:  user.Password,
	}
}

func makeStubEmail(tag string) string {
	return fmt.Sprintf(
		"testuser-%s-%s@example.com",
		time.Now().Format("20060102150405"),
		tag,
	)
}

type User struct{ APIServer, Email, Password string }

func registerNewUser(tag string) (user *User, err error) {
	rInput := RegisterInput{
		APIServer: _Conf.APIServer,
		Email:     makeStubEmail(tag),
		Password:  "secret",
	}
	if _, err = rInput.Register(); err != nil {
		return
	}
	user = &User{
		APIServer: rInput.APIServer,
		Email:     rInput.Email,
		Password:  rInput.Password,
	}
	return
}
