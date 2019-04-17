package divvycloud

import (
	"github.com/sgarlick987/godivvycloud/client"
	"github.com/sgarlick987/godivvycloud/client/users"
	"github.com/sgarlick987/godivvycloud/models"
	"log"
	"net/url"
)

type Config struct {
	Address  string
	Username string
	Password string
}

type ClientTokenWrapper struct {
	Token string
	*client.DivvyCloudV2
}

func (c *Config) Client() (interface{}, error) {
	address, err := url.Parse(c.Address)

	if err != nil {
		return nil, err
	}

	transport := &client.TransportConfig{
		BasePath: address.Path,
		Host:     address.Host,
		Schemes:  []string{address.Scheme},
	}
	log.Print("[DEBUG] setting up divvycloud http client")
	divvycloud := client.NewHTTPClientWithConfig(nil, transport)

	log.Print("[DEBUG] creating login params")
	params := users.NewPublicUserLoginPostParams().WithBody(&models.LoginRequest{
		Password: &c.Password,
		Username: &c.Username,
	})

	log.Print("[DEBUG] calling user login")
	resp, err := divvycloud.Users.PublicUserLoginPost(params)

	if err != nil {
		log.Fatal(err)
	}

	log.Print("[DEBUG] retreiving SessionID from user login")
	token := resp.Payload.SessionID

	return &ClientTokenWrapper{
		Token:        *token,
		DivvyCloudV2: divvycloud,
	}, nil
}
