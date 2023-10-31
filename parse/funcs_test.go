package parse

import "testing"

func TestSnakeCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				str: "IGotInternAtGeeksForGeeks",
			},
			want: "i_got_intern_at_geeks_for_geeks",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SnakeCase(tt.args.str); got != tt.want {
				t.Errorf("SnakeCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceToString(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "String",
			args: args{v: "Hello, World!"},
			want: "Hello, World!",
		},
		{
			name: "Integer",
			args: args{v: 42},
			want: "42",
		},
		{
			name: "Float",
			args: args{v: 3.14},
			want: "3.14",
		},
		{
			name: "Float",
			args: args{v: 1.0},
			want: "1",
		},
		{
			name: "Boolean",
			args: args{v: true},
			want: "true",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterfaceToString(tt.args.v); got != tt.want {
				t.Errorf("InterfaceToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceToJsonString(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "String",
			args: args{v: "Hello, World!"},
			want: "\"Hello, World!\"",
		},
		{
			name: "Integer",
			args: args{v: 42},
			want: "42",
		},
		{
			name: "Float",
			args: args{v: 3.14},
			want: "3.14",
		},
		{
			name: "Float",
			args: args{v: 1.0},
			want: "1",
		},
		{
			name: "Boolean",
			args: args{v: true},
			want: "true",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterfaceToJsonString(tt.args.v); got != tt.want {
				t.Errorf("InterfaceToJsonString() = %v, want %v", got, tt.want)
			}
		})
	}
}
