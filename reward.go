package growaveapi

import (
	"errors"
)

const (
	rewardPath    = "/reward"
	redeemPath    = "/redeem"
	discountsPath = "/discounts"
)

type discountResult struct {
	Result
	Data Discount
}

type RewardService interface {
	UserRedeemReward(string, int64) (*Discount, error)
	GetUserDiscounts(string) ([]*Discount, error)
}

type RewardServiceOp struct {
	client *Client
}

type Discount struct {
	Title        string `json:"title,omitempty"`
	DiscountCode string `json:"discount_code,omitempty"`
	Type         string `json:"type,omitempty"`
}

func (s *RewardServiceOp) UserRedeemReward(email string, ruleId int64) (*Discount, error) {
	var discountResult *discountResult
	var errResult *Result
	_, err := s.client.Client.R().SetResult(&discountResult).SetBody(map[string]interface{}{"email": email, "rule_id": ruleId}).SetError(&errResult).Post(rewardPath + redeemPath)
	if err != nil {
		return nil, err
	}
	if errResult != nil {
		return nil, errors.New(errResult.Message)
	}
	return &discountResult.Data, nil
}

func (s *RewardServiceOp) GetUserDiscounts(email string) ([]*Discount, error) {
	var discounts []*Discount
	var errResult *Result
	_, err := s.client.Client.R().SetResult(&Result{Data: &discounts}).SetQueryParam("email", email).SetError(&errResult).Get(rewardPath + discountsPath)
	if err != nil {
		return nil, err
	}
	if errResult != nil {
		return nil, errors.New(errResult.Message)
	}
	return discounts, nil
}
