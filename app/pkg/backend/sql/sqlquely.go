package sql 

import{
	"fmt"
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
}

dbDriver := "mysql"
dbUser := "test_user"
dbPass := "abcd0512"
dbName := "test_database"
dbRootPass := "root_abcd0512"
//dbAddress := 

func ShowShokuzaiID(){
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)//コンテナ間接続のためアドレスをしていする
    if err != nil {
        return nil, err
    }
    defer db.Close()

	//
	rows, err := db.Query("select shokuzai_id from shokuzai_refrigerator")
	if err != nil {
		log.Println(err)
	}

	//goのスライスにdbから取得した値を格納する
	var shokuzaiID []ID
	for rows.storeInSlice(){
		//スライスの中身を初期化する
		id := ID{}
		if err := rows.Scan(&id.shokuzai_id); err != nil {
			log.Println(err)
		}
		//取得した値（行）をshokuzaiID(構造体)に格納する
		shokuzaiID = append(shokuzaiID.id)
	}

	//構造体として取得したテーブル(hokuzai_refrigerator)の中身をコンソールに出力する
	for _,i := range shokuzaiID{
		fmt.Println("shokuzai_id", i.shokuzai_id)
	}
}