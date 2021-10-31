package aggregator_test

import (
	"reflect"
	"testing"

	"github.com/hanswang/clv/internal/aggregator"
	"github.com/hanswang/clv/internal/types"
)

func TestAggregator_GenerateReport(t *testing.T) {
	type args struct {
		entities *map[string]types.Entity
	}
	anchor := "a"
	tests := []struct {
		name string
		args args
		want []*types.Report
	}{
		{
			name: "Aggregate on empty map",
			args: args{
				entities: &map[string]types.Entity{},
			},
			want: []*types.Report{},
		},
		{
			name: "Aggregate on 1 entity",
			args: args{
				entities: &map[string]types.Entity{
					"a": {
						Name:     "a",
						Limit:    10,
						Utilised: 5,
					},
				},
			},
			want: []*types.Report{
				{
					Name:        "a",
					Entries:     []string{"a"},
					DirectUsage: 5,
					Usage:       5,
					Allocation:  10,
				},
			},
		},
		{
			name: "Aggregate on 2 entries with 2 entities",
			args: args{
				entities: &map[string]types.Entity{
					"a": {
						Name:     "a",
						Limit:    10,
						Utilised: 5,
					},
					"b": {
						Name:     "b",
						Limit:    2,
						Utilised: 4,
					},
				},
			},
			want: []*types.Report{
				{
					Name:        "a",
					Entries:     []string{"a"},
					DirectUsage: 5,
					Usage:       5,
					Allocation:  10,
				},
				{
					Name:        "b",
					Entries:     []string{"b"},
					DirectUsage: 4,
					Usage:       4,
					Allocation:  2,
				},
			},
		},
		{
			name: "Aggregate on 1 entry with 2 entities",
			args: args{
				entities: &map[string]types.Entity{
					"a": {
						Name:     "a",
						Limit:    10,
						Utilised: 5,
					},
					"b": {
						Name:     "b",
						Parent:   &anchor,
						Limit:    3,
						Utilised: 9,
					},
				},
			},
			want: []*types.Report{
				{
					Name:          "a",
					Entries:       []string{"a", "b"},
					DirectUsage:   5,
					Usage:         14,
					Allocation:    10,
					SubTotalLimit: 3,
					SubReports: []*types.Report{
						{
							Name:        "b",
							Entries:     []string{"b"},
							DirectUsage: 9,
							Usage:       9,
							Allocation:  3,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &aggregator.Aggregator{}
			if got := a.GenerateReport(tt.args.entities); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Aggregator.GenerateReport() = %v, want %v", got, tt.want)
			}
		})
	}
}
