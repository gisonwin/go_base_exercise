package main

import(
	_ "github.com/Go-SQL-Driver/MySQL"
	"database/sql"
	"fmt"
)
/**
*  @Author <a href="mailto:gisonwin@qq.com">GisonWin</a>
*  @Date  2019/8/19 17:53
 */
func main() {
	db,err:= sql.Open("mysql","microuser:RtaQbTrwaZC@tcp(49.4.28.140:14397)/micro?charset=utf8")
	checkErr(err)
	//insert data
	stmt, err := db.Prepare("insert userinfo set username=?,departname=?,created=?")
	checkErr(err)

	result, err := stmt.Exec("gisonwin", "技术开发部", "2018-08-13")
	checkErr(err)
	id, err := result.LastInsertId()
	checkErr(err)
	fmt.Println(id)
 //update data
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	result, err = stmt.Exec("admin", id)
	checkErr(err)
	rowsAffected, err := result.RowsAffected()
	checkErr(err)
	fmt.Println(rowsAffected)
	//query data
	rows, err := db.Query("select * from userinfo")
	checkErr(err)
	for rows.Next(){
		var uid int
		var username string
		var departname string
		var created string
		err := rows.Scan(&uid, &username, &departname, &created)
		checkErr(err)
		fmt.Println(uid,username,departname,created)
	}
	//delete data
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)
	affect, err := stmt.Exec(id)
	checkErr(err)
	fmt.Println(affect.RowsAffected())

	db.Close()

}

func checkErr(e error) {
	if e!= nil {
		panic(e)
	}
}

