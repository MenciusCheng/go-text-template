package demo

import "testing"

func TestPrintInventory(t *testing.T) {
	type args struct {
		inventory Inventory
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				inventory: Inventory{
					Material: "wool",
					Count:    17,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PrintInventory(tt.args.inventory); (err != nil) != tt.wantErr {
				t.Errorf("PrintInventory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTextAndSpaces(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			args: args{
				text: "{{23 -}} < {{- 45}}",
			},
			want: "23<45",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TextAndSpaces(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("TextAndSpaces() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TextAndSpaces() got = %v, want %v", got, tt.want)
			}
		})
	}
}
