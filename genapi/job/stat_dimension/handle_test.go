package stat_dimension

import "testing"

func Test_checkList(t *testing.T) {
	if got := checkList(); got != true {
		t.Errorf("checkList() = %v, want %v", got, true)
	}
}

func TestGenRespDetailField(t *testing.T) {
	GenRespDetailField()
}

func TestGenReqDetailField(t *testing.T) {
	GenReqDetailField()
}

func TestAddRemarkToExpByField(t *testing.T) {
	AddRemarkToExpByField()
}
