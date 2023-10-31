package pageutil

import (
	"reflect"
	"testing"
)

func TestPageLoop(t *testing.T) {
	originData := []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
	}

	getData := func(page, size int) (values []string, total int) {
		start := (page - 1) * size
		if start >= len(originData) {
			return nil, len(originData)
		}
		end := start + size
		if end > len(originData) {
			end = len(originData)
		}
		return originData[start:end], len(originData)
	}

	type args struct {
		page int
		size int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantRes []string
	}{
		{
			name: "多次分页获取",
			args: args{
				page: 1,
				size: 3,
			},
			wantRes: originData,
		},
		{
			name: "一次分页获取，恰好取完",
			args: args{
				page: 1,
				size: 10,
			},
			wantRes: originData,
		},
		{
			name: "一次分页获取，大于分页数据",
			args: args{
				page: 1,
				size: 20,
			},
			wantRes: originData,
		},
		{
			name: "大于分页数据，取不到数据",
			args: args{
				page: 2,
				size: 10,
			},
			wantRes: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := make([]string, 0)
			f := func(page, size int) (total int, err error) {
				values, total := getData(page, size)
				t.Logf("execute f func, page:%d, size:%d, total:%d, len(values):%d", page, size, total, len(values))
				res = append(res, values...)
				return total, nil
			}

			if err := PageLoop(f, tt.args.page, tt.args.size); (err != nil) != tt.wantErr {
				t.Errorf("PageLoop() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(res, tt.wantRes) {
				t.Errorf("res = %v, wantRes= %v", res, tt.wantRes)
			}
		})
	}
}
