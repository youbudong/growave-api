package growaveapi

import (
	"errors"
)

const (
	rewardPath     = "/reward"
	earnPath       = "/earn"
	redeemPath     = "/redeem"
	discountsPath  = "/discounts"
	activitiesPath = "/activities"
)

type discountResult struct {
	Result
	Data Discount
}

type RewardService interface {
	UserRedeemReward(email string, ruleId int64) (*Discount, error)
	RedeemEarn(email string, ruleType string, points float64) (*EarnedPoints, error)
	GetUserDiscounts(email string) ([]*Discount, error)
	GetUserActivities(email string) ([]*UserActivitie, error)
}

type RewardServiceOp struct {
	client *Client
}

type EarnedPoints struct {
	EarnedPoints int64 `json:"earned_points"`
}

type Discount struct {
	Title        string `json:"title,omitempty"`
	DiscountCode string `json:"discount_code,omitempty"`
	Type         string `json:"type,omitempty"`
}

type UserActivitie struct {
	Id            int64       `json:"id"`
	Type          string      `json:"type,omitempty"`
	RewardingType string      `json:"rewarding_type,omitempty"`
	Spend         int64       `json:"spend,omitempty"`
	GiftCardCode  interface{} `json:"gift_card_code,omitempty"`
	SpendingRule  interface{} `json:"spending_rule"`
}

// 兑换奖励
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

// 发放积分
func (s *RewardServiceOp) RedeemEarn(email string, ruleType string, points float64) (*EarnedPoints, error) {
	var pointsR *EarnedPoints
	var errResult *Result
	_, err := s.client.Client.R().SetResult(&Result{Data: &pointsR}).SetBody(map[string]interface{}{"email": email, "rule_type": ruleType, "points": points}).SetError(&errResult).Post(rewardPath + earnPath)
	if err != nil {
		return nil, err
	}
	if errResult != nil {
		return nil, errors.New(errResult.Message)
	}
	return pointsR, nil
}

// 获取用户折扣列表
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

// 获取用户活动列表
func (s *RewardServiceOp) GetUserActivities(email string) ([]*UserActivitie, error) {
	var userActivities []*UserActivitie
	var errResult *Result
	_, err := s.client.Client.R().SetResult(&Result{Data: &userActivities}).SetQueryParam("email", email).SetError(&errResult).Get(rewardPath + activitiesPath)
	if err != nil {
		return nil, err
	}
	if errResult != nil {
		return nil, errors.New(errResult.Message)
	}
	return userActivities, nil
}
