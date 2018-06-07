package gsb

import "testing"

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func TestInsertInto(t *testing.T) {
	sql := Insert("t_users").
		Columns("username", "password").
		ToSQL()
	checkSQLMatches(sql, "INSERT INTO `t_users`(`username`, `password`) VALUES (?, ?);", t)
}

func TestInsertIntoValued(t *testing.T) {
	sql := Insert("t_users").
		Columns("username", "password").
		Values("yoojia", "123456").
		ToSQL()
	checkSQLMatches(sql, "INSERT INTO `t_users`(`username`, `password`) VALUES ('yoojia', '123456');", t)
}
