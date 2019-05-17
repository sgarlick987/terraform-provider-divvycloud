package divvycloud

import (
	"errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/sgarlick987/godivvycloud/client"
	"github.com/sgarlick987/godivvycloud/client/users"
	"github.com/sgarlick987/godivvycloud/models"
	"testing"
)

func TestClient(t *testing.T) {
	token := "123456"
	username := "username"
	password := "password"

	transport := &FakeDivvycloudClientTransport{
		ParamsMap: map[string]runtime.ClientRequestWriter{},
		CalledMap: map[string]bool{
			"PublicUserLoginPost": false,
		},
		ErrorMap: map[string]error{
			"PublicUserLoginPost": nil,
		},
		ResponseMap: map[string]interface{}{
			"PublicUserLoginPost": &users.PublicUserLoginPostOK{
				Payload: &models.OriginalAuth{
					SessionID: &token,
				},
			},
		},
	}

	c := &LoginConfig{
		Username: username,
		Password: password,
		Client:   client.New(transport, strfmt.Default),
	}

	wrappedClient, err := c.WrappedClient()

	if err != nil {
		t.Fatalf("wrapped client returned err: %s", err)
	}

	if !transport.CalledMap["PublicUserLoginPost"] {
		t.Error("user login should be called")
	}

	response := transport.ParamsMap["PublicUserLoginPost"].(*users.PublicUserLoginPostParams)

	if *response.Body.Username != username {
		t.Error("should be called with configured username")
	}

	if *response.Body.Password != password {
		t.Error("should be called with configured password")
	}

	tokenClient := wrappedClient.(*ClientTokenWrapper)

	if tokenClient.Token != token {
		t.Error("client token shoud be one returned from login call")
	}
}
func TestClientLoginError(t *testing.T) {
	transport := &FakeDivvycloudClientTransport{
		ParamsMap: map[string]runtime.ClientRequestWriter{},
		CalledMap: map[string]bool{},
		ErrorMap: map[string]error{
			"PublicUserLoginPost": errors.New("login failed"),
		},
		ResponseMap: map[string]interface{}{},
	}

	c := &LoginConfig{
		Username: "",
		Password: "",
		Client:   client.New(transport, strfmt.Default),
	}

	_, err := c.WrappedClient()

	if err == nil {
		t.Error("an error should be returned when login fails")
	}
}
