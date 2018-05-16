package sqlx

import (
	"fmt"
	"testing"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func checkSQLMatches(sql string, shouldBe string, t *testing.T) {
	fmt.Println(sql)
	if sql != shouldBe {
		t.Error("Output sql not match")
	}
}

func newWhereTest(conditions SQLStatement) *WhereBuilder {
	return newWhere(nil, conditions)
}
