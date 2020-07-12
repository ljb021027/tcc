package storage

import (
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/ljb021027/tcc/util"
)

var (
	AlreadyExistsError = errors.New("alaready exists record ")
)

const (
	_ = iota
	StatusTry
	StatusCommit
	StatusCancel
)

type BranchTable struct {
	BranchId string
	Xid      string
	Type     string
	Status   int
}

func (*BranchTable) TableName() string {
	return "branch_table"
}

func SaveBranchRecord(branchId, xid string, s int) error {
	table := &BranchTable{
		BranchId: branchId,
		Xid:      xid,
		Type:     "",
		Status:   s,
	}
	err := table.Insert()
	me, ok := err.(*mysql.MySQLError)
	if !ok {
		return err
	}
	if me.Number == 1062 {
		return AlreadyExistsError
	}
	if err != nil {
		return err
	}

	return nil
}

func QueryBranchRecord(branchId, xid string) (int, error) {
	branchTable := BranchTable{
		BranchId: branchId,
		Xid:      xid,
	}
	query, err := branchTable.Query()
	if err != nil {
		return -1, err
	}
	if query != nil && len(query) == 1 {
		return query[0].Status, nil
	} else {
		return 0, nil
	}
}

func (bt *BranchTable) Insert() error {
	err := util.GloalDb.Save(bt).Error
	return err
}

func (bt *BranchTable) Query() ([]*BranchTable, error) {
	results := make([]*BranchTable, 0)
	err := util.GloalDb.Find(&results, bt).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}
