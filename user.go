package growaveapi

import (
	"errors"
)

const (
	userPath   = "/users"
	searchPath = "/search"
)

type UserResult struct {
	Result
	Data User `json:"data,omitempty"`
}

type UserService interface {
	Search(field string, value string) (*User, error)
}

type UserServiceOp struct {
	client *Client
}

type User struct {
	UserId         int64  `json:"user_id,omitempty"`
	CustomerId     string `json:"customer_id,omitempty"`
	Email          string `json:"email,omitempty"`
	FirstName      string `json:"first_name,omitempty"`
	LastName       string `json:"last_name,omitempty"`
	PhotoUrl       string `json:"photo_url,omitempty"`
	About          string `json:"about,omitempty"`
	Birthdate      string `json:"birthdate,omitempty"`
	ProfileAddress string `json:"profile_address,omitempty"`
	Username       string `json:"username,omitempty"`
	Privacy        string `json:"privacy,omitempty"`
	Points         string `json:"points,omitempty"`
	ReferLink      string `json:"refer_link,omitempty"`
}

func (s *UserServiceOp) Search(field string, value string) (*User, error) {
	var userResult *UserResult
	var errResult *Result
	_, err := s.client.Client.R().SetResult(userResult).SetError(errResult).SetQueryParams(map[string]string{"field": field, "value": value}).Get(userPath + searchPath)
	if err != nil {
		return nil, err
	}
	if errResult != nil {
		return nil, errors.New(errResult.Message)
	}
	return &userResult.Data, nil
}
