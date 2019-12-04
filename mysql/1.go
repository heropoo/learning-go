package main

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "fmt"

type User struct {
	ID int64 `db:"id"`
	//Username sql.NullString `db:"username"`  //如果在mysql的users表中name没有设置为NOT NULL,所以name可能为null,在查询过程中会返回nil，如果是string类型则无法接收nil,但sql.NullString则可以接收nil值
	Username  string `db:"username"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

func main() {
	DB, err := sql.Open("mysql", "root:root123456@tcp(localhost:3306)/test")
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(DB)
	var id int64
	id = 1
	queryOne(DB, id)
	queryMulti(DB)
	insertData(DB)
	updateData(DB)
	deleteData(DB)
}

//查询单行
func queryOne(DB *sql.DB, id int64) {
	user := new(User)
	row := DB.QueryRow("select * from user where id=?", id)
	if err := row.Scan(&user.ID, &user.Username, &user.CreatedAt, &user.UpdatedAt); err != nil {
		fmt.Printf("scan failed, err:%v", err)
		return
	}

	fmt.Println(*user)
}

//查询多行
func queryMulti(DB *sql.DB) {
	user := new(User)
	rows, err := DB.Query("select * from user where id > ?", 1)
	defer func() {
		if rows != nil {
			rows.Close() //可以关闭掉未scan连接一直占用
		}
	}()

	if err != nil {
		fmt.Printf("Query failed,err:%v", err)
		return
	}

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.CreatedAt, &user.UpdatedAt) //不scan会导致连接不释放
		if err != nil {
			fmt.Printf("Scan failed,err:%v", err)
			return
		}
		fmt.Println(*user)
	}
}

//插入数据
func insertData(DB *sql.DB){
    result,err := DB.Exec("insert INTO user(username,created_at,updated_at) values(?,?,?)","小明","2019-12-04 17:23:01", "2019-12-04 17:23:01")
    if err != nil{
        fmt.Printf("Insert failed,err:%v",err)
        return
    }
    lastInsertID,err := result.LastInsertId()
    if err != nil {
        fmt.Printf("Get lastInsertID failed,err:%v",err)
        return
    }
    fmt.Println("LastInsertID:",lastInsertID)
    rowsaffected,err := result.RowsAffected()
    if err != nil {
        fmt.Printf("Get RowsAffected failed,err:%v",err)
        return
    }
    fmt.Println("RowsAffected:",rowsaffected)
}

//更新数据
func updateData(DB *sql.DB){
    result,err := DB.Exec("UPDATE user set updated_at=? where id=?","2019-12-04 17:25:01", 1)
    if err != nil{
        fmt.Printf("Insert failed,err:%v",err)
        return
    }
    rowsaffected,err := result.RowsAffected()
    if err != nil {
        fmt.Printf("Get RowsAffected failed,err:%v",err)
        return
    }
    fmt.Println("RowsAffected:",rowsaffected)
}

//删除数据
func deleteData(DB *sql.DB){
    result,err := DB.Exec("delete from user where id=?",3)
    if err != nil{
        fmt.Printf("Insert failed,err:%v",err)
        return
    }
    rowsaffected,err := result.RowsAffected()
    if err != nil {
        fmt.Printf("Get RowsAffected failed,err:%v",err)
        return
    }
    fmt.Println("RowsAffected:",rowsaffected)
}