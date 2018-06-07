package gsb

import "testing"

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func TestUpdateShouldPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("TestUpdateShouldPanic should have panicked!")
		}
	}()
	Update("db.t_user").Columns("username").ToSQL()
}

func TestUpdate(t *testing.T) {
	sql := Update("db.t_user").
		Columns("username").
		ColumnAndValue("age", 18).
		YesYesYesForceUpdate().
		Compile()
	checkSQLMatches(sql, "UPDATE `db.t_user` SET `username`=?, `age`=18", t)
}

func TestUpdateBuilder_Where(t *testing.T) {
	sql := Update("db.t_user").
		Columns("username").
		Where(GreaterEqualThen("age").
			Or().LessThenTo("height", 50)).
		ToSQL()
	checkSQLMatches(sql, "UPDATE `db.t_user` SET `username`=? WHERE `age` >= ? OR `height` < 50;", t)
}
