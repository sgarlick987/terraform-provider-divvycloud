package divvycloud

import (
	"github.com/sgarlick987/godivvycloud/client"
	"github.com/sgarlick987/godivvycloud/client/users"
	"github.com/sgarlick987/godivvycloud/models"
	"log"
)

type LoginConfig struct {
	Username string
	Password string
	Client   *client.DivvyCloudV2
}

// Wrap the generated divvycloud client with a token that is returned from the login call at client setup
// TODO: figure out if we can use the token auth support in go-swagger. the client swagger.json will need to support and it'll need to work with the login process
type ClientTokenWrapper struct {
	Token string
	*client.DivvyCloudV2
}

// I've no idea if this is proper doing the login here to get the session token
// divvycloud requires a login with a username/password
// and then taking the session id returned and using it as a token in X-Auth-Token header
func (c *LoginConfig) WrappedClient() (interface{}, error) {
	log.Print("[DEBUG] creating login params")
	params := users.NewPublicUserLoginPostParams().WithBody(&models.LoginRequest{
		Password: &c.Password,
		Username: &c.Username,
	})

	log.Print("[DEBUG] calling user login")
	resp, err := c.Client.Users.PublicUserLoginPost(params)

	if err != nil {
		return nil, err
	}

	log.Print("[DEBUG] retrieving SessionID from user login")
	token := resp.Payload.SessionID

	return &ClientTokenWrapper{
		Token:        *token,
		DivvyCloudV2: c.Client,
	}, nil
}
