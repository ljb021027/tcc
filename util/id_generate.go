package util

import (
	"strconv"
	"sync/atomic"
	"time"
)

var seqIndex uint64
var branchIdIndex uint64

func GenerateSeq() string {
	nano := time.Now().UnixNano()
	addUint64 := atomic.AddUint64(&seqIndex, 1)
	return strconv.FormatInt(nano, 10) + strconv.FormatUint(addUint64, 10)
}

func GenerateBranchId() string {
	nano := time.Now().UnixNano()
	addUint64 := atomic.AddUint64(&branchIdIndex, 1)
	return strconv.FormatInt(nano, 10) + strconv.FormatUint(addUint64, 10)
}
