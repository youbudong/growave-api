package growaveapi

import (
	"reflect"
	"testing"
)

func TestRewardServiceOp_GetUserDiscounts(t *testing.T) {
	type args struct {
		email string
	}
	client := New(App{"", "", []string{}})
	tests := []struct {
		name    string
		s       *RewardServiceOp
		args    args
		want    *[]*Discount
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "", s: &RewardServiceOp{client}, args: args{email: ""}, want: &[]*Discount{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetUserDiscounts(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("RewardServiceOp.GetUserDiscounts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RewardServiceOp.GetUserDiscounts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRewardServiceOp_UserRedeemReward(t *testing.T) {
	type args struct {
		email  string
		ruleId int64
	}
	tests := []struct {
		name    string
		s       *RewardServiceOp
		args    args
		want    *Discount
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.UserRedeemReward(tt.args.email, tt.args.ruleId)
			if (err != nil) != tt.wantErr {
				t.Errorf("RewardServiceOp.UserRedeemReward() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RewardServiceOp.UserRedeemReward() = %v, want %v", got, tt.want)
			}
		})
	}
}
