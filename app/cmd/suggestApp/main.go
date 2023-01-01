package main
 
import (
	"log"
	"net/http"
	sql "github.com/Babarin0512/SuggestMenuApp/tree/main/app/pkg/backend/sql" //パッケージ名が長いので別名で参照する
)
 
func main() {
	sql.ShowShokuzaiID()



	/*port := "8080"
	http.Handle("/", http.FileServer(http.Dir("./docs/")))
	log.Printf("Listen on port: %s", port)
	http.ListenAndServe(":"+port, nil)*/
}