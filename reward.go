package growaveapi

type RewardService interface {
	Get(int64, interface{}) (*Reward, error)
	Put(Reward) (*Reward, error)
	List(interface{}) ([]Reward, error)
	Search(string) (*Reward, error)
}

type Reward struct {
	RewardId int64 `json:"user_id,omitempty"`
	// CustomerId     string `json:"customer_id,omitempty"`
	// Email          string `json:"email,omitempty"`
	// FirstName      string `json:"first_name,omitempty"`
	// LastName       string `json:"last_name,omitempty"`
	// PhotoUrl       string `json:"photo_url,omitempty"`
	// About          string `json:"about,omitempty"`
	// Birthdate      string `json:"birthdate,omitempty"`
	// ProfileAddress string `json:"profile_address,omitempty"`
	// Rewardname     string `json:"username,omitempty"`
	// Privacy        string `json:"privacy,omitempty"`
	// Points         string `json:"points,omitempty"`
	// ReferLink      string `json:"refer_link,omitempty"`
}
