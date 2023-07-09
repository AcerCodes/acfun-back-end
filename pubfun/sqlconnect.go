package sqlconnect

import (
	"database/sql"

	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

type u struct {
	id                 int
	user_name          string
	user_register_time string
}

func Init() {
	fmt.Println("imp-init() come here.")
}

func Query() {
	rows, err := db.Query("SELECT id,user_name,user_register_time FROM tbl_user")
	if err != nil {
		panic(err)
	}
	//延迟关闭rows
	defer rows.Close()
	// for rows.Next() {
	// 	cols, _ := rows.Columns()
	// 	for i := range cols {
	// 		fmt.Print(cols[i])
	// 		fmt.Print("\t")
	// 	}
	// }

	for rows.Next() {
		user := u{}
		err := rows.Scan(&user.id, &user.user_name, &user.user_register_time)
		if err != nil {
			panic(err)
		}
		fmt.Printf("id = %v, name = %v, time = %v\n", user.id, user.user_name, user.user_register_time)
	}
}

func DbOpen() {
	var err error
	//参数根据自己的数据库进行修改
	connStr := "host=8.142.172.39 port=5432 user=studentdb password=studentdb dbname=studentdb sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connect to database!")
}
