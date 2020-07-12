package storage

import (
	"fmt"
	"testing"

	"github.com/ljb021027/tcc/util"
)

func Test(t *testing.T) {
	util.InitServiceConfig()
	util.InitDb()
	table := &BranchTable{
		BranchId: "111",
		Xid:      "111",
		Type:     "",
		Status:   StatusTry,
	}
	err := table.Insert()
	fmt.Println(err)
}

func TestQueryBranchRecord(t *testing.T) {
	util.InitServiceConfig()
	util.InitDb()
	record, err := QueryBranchRecord("1111", "111")
	fmt.Println(record)
	fmt.Println(err)
}
