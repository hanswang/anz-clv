package parser_test

import (
	"reflect"
	"testing"

	"github.com/hanswang/clv/internal/parser"
	"github.com/hanswang/clv/internal/types"
)

func TestParser_ParseCSV(t *testing.T) {
	type args struct {
		rows []string
	}
	anchor := "a"
	tests := []struct {
		name    string
		args    args
		want    *map[string]types.Entity
		wantErr bool
	}{
		{
			name: "Parse on incorrect col number",
			args: args{
				rows: []string{"1,2,3"},
			},
			want: nil,
			wantErr: true,
		},
		{
			name: "Parse on incorrect limit type",
			args: args{
				rows: []string{"1,2,a,4"},
			},
			want: nil,
			wantErr: true,
		},
		{
			name: "Parse on incorrect utilised type",
			args: args{
				rows: []string{"1,2,3,a"},
			},
			want: nil,
			wantErr: true,
		},
		{
			name: "Parse on unspecified parent value",
			args: args{
				rows: []string{"1,2,3,4"},
			},
			want: nil,
			wantErr: true,
		},
		{
			name: "Parse on success entities",
			args: args{
				rows: []string{"a,,3,4","b,a,3,4"},
			},
			want: &map[string]types.Entity{
				"a": {
					Name: "a",
					Limit: 3,
					Utilised: 4,
				},
				"b": {
					Name: "b",
					Parent: &anchor,
					Limit: 3,
					Utilised: 4,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser.Parser{}
			got, err := p.ParseCSV(tt.args.rows)
			
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.ParseCSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.ParseCSV() = %v, want %v", got, tt.want)
			}
		})
	}
}
