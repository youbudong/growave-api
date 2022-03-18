package growaveapi

const (
	userPath   = "/users"
	searchPath = "/search"
)

type UserService interface {
	Get(int64, interface{}) (*User, error)
	Put(User) (*User, error)
	List(interface{}) ([]User, error)
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

func (s *UserServiceOp) Search(options interface{}) (*User, error) {
	var user User
	_, err := s.client.Client.R().SetResult(&user).Get(userPath + searchPath)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
