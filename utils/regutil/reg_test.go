package regutil

import (
	"testing"
)

func TestShowRegRes(t *testing.T) {
	type args struct {
		expr string
		s    string
	}
	tests := []struct {
		name    string
		args    args
		want    RegRes
		wantErr bool
	}{
		{
			args: args{
				expr: `-\s+\[[-a-zA-Z0-9_ ]+]\(https://github.com/([-a-zA-Z0-9_ ]+)/([-a-zA-Z0-9_ ]+)\)`,
				s:    "- [Uniqush-Push](https://github.com/uniqush/uniqush-push) ",
			},
		},
		{
			args: args{
				expr: "p([a-z]+)e",
				s:    "Golang regular expressions example",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ShowRegRes(tt.args.expr, tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShowRegRes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("ShowRegRes() got = %v, want %v", got, tt.want)
			//}
			t.Logf("got = %s", got.String())
		})
	}
}
