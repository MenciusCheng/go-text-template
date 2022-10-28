package gen

import "testing"

func TestGenByDirConfig(t *testing.T) {
	type args struct {
		dirPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{args: args{dirPath: "../config/crud_dao"}},
		//{args: args{dirPath: "../config/router"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GenByDirConfig(tt.args.dirPath); (err != nil) != tt.wantErr {
				t.Errorf("GenByDirConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
