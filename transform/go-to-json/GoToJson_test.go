package go_to_json

import (
	"fmt"
	"testing"
)

func TestGoToJson(t *testing.T) {
	d := EventInstall{}
	got := GoToJson(d)
	fmt.Println(got)
}
