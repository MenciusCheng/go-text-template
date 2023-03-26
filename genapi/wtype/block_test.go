package wtype

import "testing"

func Test_BlockDefReg_Match(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Empty string",
			args: args{b: []byte("")},
			want: false,
		},
		{
			name: "Valid string with one key-value pair",
			args: args{b: []byte("@=id:23")},
			want: true,
		},
		{
			name: "Valid string with multiple key-value pairs",
			args: args{b: []byte("@=id:23,name:pp,type:yaml")},
			want: true,
		},
		{
			name: "Valid string with missing value for last key",
			args: args{b: []byte("@=id:23,name:")},
			want: true,
		},
		{
			name: "Invalid string without @= prefix",
			args: args{b: []byte("id:23,name:pp,type:yaml")},
			want: false,
		},
		{
			name: "Invalid string with extra whitespace",
			args: args{b: []byte("@= id:23 , name:pp, type:yaml ")},
			want: false,
		},
		{
			name: "Invalid string with extra newline characters",
			args: args{b: []byte("@=id:23,\nname:pp,\ntype:yaml\n")},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BlockDefReg.Match(tt.args.b); got != tt.want {
				t.Errorf("IsBlockDefLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
