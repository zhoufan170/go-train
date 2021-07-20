package main

import (
	"error/mysql"
	"fmt"

	"github.com/pkg/errors"
)

// 1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

// 我认为sql.ErrNoRows 不需要抛给上层，这应该被当作一个业务上的正常结果返回给调用方
// 至于业务上是否正确，应该留给调用方来判断，下层应该直接把这个error处理了

func main() {
	db, err := mysql.Open()
	if err != nil {
		fmt.Printf("open database fail, root error: %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("trace: \n%+v\n", err)
		return
	}

	var sqlString = "select * from test_table where 1=1"
	tableResultSlice, err := mysql.Query(sqlString, db)
	if err != nil {
		fmt.Printf("query data fail, root error: %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("trace: \n%+v\n", err)
		return
	}

	if len(tableResultSlice) == 0 {
		// do something to handle 0 row
	}

	fmt.Printf("%v", tableResultSlice)

	defer db.Close()
}
