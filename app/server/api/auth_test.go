package api

import (
	"net/http"
	"testing"

	"github.com/google/uuid"

	"github.com/ismdeep/go-rest-api-demo/app/server/store"
	"github.com/ismdeep/go-rest-api-demo/internal/request"
)

func TestSignUp(t *testing.T) {
	username := uuid.New().String()

	defer func() {
		_ = store.User.Delete(username)
	}()

	d := request.SignUp{
		Username: username,
		Password: uuid.New().String(),
	}
	testRequest(t, http.MethodPost, "/api/v1/auth/sign-up", d, 200, 0)
}
