package mysql

import (
	"testing"
)

func TestProdure(t *testing.T) {

	engine := GetXormEngine()

	result, err := engine.Exec("call t_message_get_increment_index(?,?,?,?,@p_out);", 32, 1, "a meesge for go test", 1)
	if err != nil {
		t.Error(err.Error())
	}
	result = result
}
