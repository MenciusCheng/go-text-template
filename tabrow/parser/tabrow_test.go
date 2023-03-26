package parser

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseToTabRow(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want TabRow
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
			want: TabRow{
				Rows: []Row{
					{
						Origin: "stat_date\tstring\t数据日期，格式：YYYY-MM-DD",
						Cells:  []string{"stat_date", "string", "数据日期，格式：YYYY-MM-DD"},
					},
					{
						Origin: "stat_hour\tint\t数据小时",
						Cells:  []string{"stat_hour", "int", "数据小时"},
					},
					{
						Origin: "charge\tdouble\t花费（元）",
						Cells:  []string{"charge", "double", "花费（元）"},
					},
					{
						Origin: "show\tlong\t封面曝光数",
						Cells:  []string{"show", "long", "封面曝光数"},
					},
					{
						Origin: "photo_click\tlong\t封面点击数",
						Cells:  []string{"photo_click", "long", "封面点击数"},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseToTabRow(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				//gotJson, _ := json.Marshal(got)
				t.Errorf("ParseToTabRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
