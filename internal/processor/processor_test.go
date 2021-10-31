package processor_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hanswang/clv/internal/mocks"
	"github.com/hanswang/clv/internal/processor"
)

func TestNewProcessor(t *testing.T) {
	type args struct {
		parser     *mocks.MockParserManager
		aggregator *mocks.MockAggregatorManager
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tests := []struct {
		name string
		args args
		want *processor.Processor
	}{
		{
			name: "create new processor",
			args: args{
				mocks.NewMockParserManager(ctrl),
				mocks.NewMockAggregatorManager(ctrl),
			},
			want: &processor.Processor{
				mocks.NewMockParserManager(ctrl),
				mocks.NewMockAggregatorManager(ctrl),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processor.NewProcessor(tt.args.parser, tt.args.aggregator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProcessor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcessor_Run(t *testing.T) {
	type fields struct {
		parser     *mocks.MockParserManager
		aggregator *mocks.MockAggregatorManager
	}
	type args struct {
		body []byte
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		args    args
		wantErr bool
	}{
		{
			name: "Parse only header",
			args: args{
				body: []byte("dummyline"),
			},
			wantErr: true,
		},
		{
			name: "Parse two lines with parser error",
			prepare: func(f *fields) {
				f.parser.EXPECT().ParseCSV([]string{"dummyline2"}).Return(nil, fmt.Errorf("dummyErr"))
			},
			args: args{
				body: []byte("dummyline1\ndummyline2"),
			},
			wantErr: true,
		},
		{
			name: "Parse two lines with parser success",
			prepare: func(f *fields) {
				f.parser.EXPECT().ParseCSV([]string{"dummyline2"}).Return(nil, nil)
				f.aggregator.EXPECT().GenerateReport(gomock.Any()).Return(nil)
			},
			args: args{
				body: []byte("dummyline1\ndummyline2"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				parser: mocks.NewMockParserManager(ctrl),
				aggregator: mocks.NewMockAggregatorManager(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			p := processor.NewProcessor(f.parser, f.aggregator)
			if err := p.Run(tt.args.body); (err != nil) != tt.wantErr {
				t.Errorf("Processor.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
