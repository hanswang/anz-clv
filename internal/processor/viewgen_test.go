package processor_test

import (
	"bytes"
	"testing"

	"github.com/hanswang/clv/internal/processor"
	"github.com/hanswang/clv/internal/types"
)

func TestViewGenerate(t *testing.T) {
	type args struct {
		reports []*types.Report
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "View on empty report",
			args: args{},
			want: "",
		},
		{
			name: "View on single no breach report",
			args: args{
				[]*types.Report{
					{
						Name: "a",
						Entries: []string{"a"},
						DirectUsage: 5,
						Usage: 5,
						Allocation: 10,
					},
				},
			},
			want: "Entities: a:\n  No limit breaches\n",
		},
		{
			name: "View on single breach report",
			args: args{
				[]*types.Report{
					{
						Name: "a",
						Entries: []string{"a"},
						DirectUsage: 5,
						Usage: 16,
						Allocation: 10,
					},
				},
			},
			want: "Entities: a:\n  Limit breach at a (limit = 10, direct utilisation = 5, combined utilisation = 16).\n",
		},
		{
			name: "View on single warning report",
			args: args{
				[]*types.Report{
					{
						Name: "a",
						Entries: []string{"a"},
						DirectUsage: 5,
						Usage: 5,
						Allocation: 10,
						SubTotalLimit: 20,
					},
				},
			},
			want: "Entities: a:\n  Warning for limit at a (limit = 10, combined sub-entity limit = 20).\n  No limit breaches\n",
		},
		{
			name: "View on 2 reports",
			args: args{
				[]*types.Report{
					{
						Name: "a",
						Entries: []string{"a"},
						DirectUsage: 5,
						Usage: 5,
						Allocation: 10,
					},
					{
						Name: "b",
						Entries: []string{"b"},
						DirectUsage: 5,
						Usage: 5,
						Allocation: 6,
					},
				},
			},
			want: "Entities: a:\n  No limit breaches\n\nEntities: b:\n  No limit breaches\n",
		},
		{
			name: "View on 2 reports",
			args: args{
				[]*types.Report{
					{
						Name: "a",
						Entries: []string{"a", "b"},
						DirectUsage: 5,
						Usage: 14,
						Allocation: 10,
						SubTotalLimit: 3,
						SubReports: []*types.Report{
							{
								Name: "b",
								Entries: []string{"b"},
								DirectUsage: 9,
								Usage: 9,
								Allocation: 3,
							},
						},
					},
				},
			},
			want: "Entities: a/b:\n  Limit breach at a (limit = 10, direct utilisation = 5, combined utilisation = 14).\n  Limit breach at b (limit = 3, direct utilisation = 9, combined utilisation = 9).\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bak := processor.Out
			processor.Out = new(bytes.Buffer)
			defer func () { processor.Out = bak }()
			processor.ViewGenerate(tt.args.reports)
			if got := processor.Out.(*bytes.Buffer).String(); got != tt.want {
				t.Errorf("ViewGenerate() = %q, want %q", got, tt.want)
			}
		})
	}
}
