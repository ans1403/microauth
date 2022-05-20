package service_test

import (
	"golang-devcontainer/src/service"
	"testing"
)

func TestCalculateServiceImpl_Service(t *testing.T) {
	type args struct {
		loopNumber int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "1を渡した時に1が返却される",
			args: args{loopNumber: 1},
			want: 1,
		},
		{
			name: "10を渡した時に55が返却される",
			args: args{loopNumber: 10},
			want: 55,
		},
		{
			name: "100を渡した時に5050が返却される",
			args: args{loopNumber: 100},
			want: 5050,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &service.CalculateServiceImpl{}
			if got := c.Service(tt.args.loopNumber); got != tt.want {
				t.Errorf("CalculateServiceImpl.Service() = %v, want %v", got, tt.want)
			}
		})
	}
}
