package wtype

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseToTabrow(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want Tabrow
	}{
		{
			args: args{
				str: strings.TrimSpace(`
stat_date	string	数据日期，格式：YYYY-MM-DD
stat_hour	int	数据小时
charge	double	花费（元）
show	long	封面曝光数
photo_click	long	封面点击数
`),
			},
			want: Tabrow{
				[]string{
					"stat_date\tstring\t数据日期，格式：YYYY-MM-DD", "stat_date", "string", "数据日期，格式：YYYY-MM-DD",
				},
				[]string{
					"stat_hour\tint\t数据小时", "stat_hour", "int", "数据小时",
				},
				[]string{
					"charge\tdouble\t花费（元）", "charge", "double", "花费（元）",
				},
				[]string{
					"show\tlong\t封面曝光数", "show", "long", "封面曝光数",
				},
				[]string{
					"photo_click\tlong\t封面点击数", "photo_click", "long", "封面点击数",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseToTabrow(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseToTabRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
