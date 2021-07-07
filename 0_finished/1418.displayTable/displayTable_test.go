package displaytable

import (
	"reflect"
	"testing"
)

func Test_displayTable(t *testing.T) {
	type args struct {
		orders [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "1",
			args: args{
				orders: [][]string{
					{"David", "3", "Ceviche"},
					{"Corina", "10", "Beef Burrito"},
					{"David", "3", "Fried Chicken"},
					{"Carla", "5", "Water"},
					{"Carla", "5", "Ceviche"},
					{"Rous", "3", "Ceviche"},
				},
			},
			want: [][]string{
				{"Table", "Beef Burrito", "Ceviche", "Fried Chicken", "Water"},
				{"3", "0", "2", "1", "0"},
				{"5", "0", "1", "0", "1"},
				{"10", "1", "0", "0", "0"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := displayTable(tt.args.orders); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("displayTable() = %v, want %v", got, tt.want)
			}
		})
	}
}
