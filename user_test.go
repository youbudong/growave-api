package growaveapi

import (
	"reflect"
	"testing"
)

func TestUserServiceOp_Search(t *testing.T) {
	type args struct {
		options interface{}
	}
	tests := []struct {
		name    string
		s       *UserServiceOp
		args    args
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Search(tt.args.options)
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
