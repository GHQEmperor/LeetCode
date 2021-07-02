package _833_maxIceCream

import "testing"

func Test_maxIceCream(t *testing.T) {
	type args struct {
		costs []int
		coins int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{
			name: "costs = [1,3,2,4,1], coins = 7",
			args: args{
				costs: []int{1, 3, 2, 4, 1},
				coins: 7,
			},
			wantResult: 4,
		},
		{
			name: "costs = [10,6,8,7,7,8], coins = 5",
			args: args{
				costs: []int{10, 6, 8, 7, 7, 8},
				coins: 5,
			},
			wantResult: 0,
		},
		{
			name: "costs = [1,6,3,1,2,5], coins = 20",
			args: args{
				costs: []int{1, 6, 3, 1, 2, 5},
				coins: 20,
			},
			wantResult: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := maxIceCream(tt.args.costs, tt.args.coins); gotResult != tt.wantResult {
				t.Errorf("maxIceCream() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
