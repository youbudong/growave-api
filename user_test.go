package growaveapi

import (
	"reflect"
	"testing"
)

func TestUserServiceOp_Search(t *testing.T) {
	type args struct {
		field string
		value string
	}
	client := New(App{"", "", []string{}})
	tests := []struct {
		name    string
		s       *UserServiceOp
		args    args
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "", s: &UserServiceOp{client}, args: args{field: "email", value: ""}, want: &User{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Search(tt.args.field, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserServiceOp.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserServiceOp.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
