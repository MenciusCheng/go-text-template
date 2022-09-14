package parse

import (
	"fmt"
	"testing"
)

func TestStringToMap(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				text: "{\"rows\":[{\"1\":1,\"2\":\"b\"},{\"1\":1,\"2\":\"b\"},{\"1\":1,\"2\":\"b\"}]}",
			},
		},
		{
			args: args{
				text: "{\"rows\":[[1,\"b\"],[1,\"b\"]]}",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToMap(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Printf("got = %+v\n", got)
		})
	}
}

func TestPrintJsonRows(t *testing.T) {
	type args struct {
		jsonString string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				jsonString: "{\"rows\":[[1,\"b\"],[1,\"b\"]]}",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PrintJsonRows(tt.args.jsonString); (err != nil) != tt.wantErr {
				t.Errorf("PrintJsonRows() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
