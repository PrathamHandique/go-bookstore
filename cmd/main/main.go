package main
import (
	"github.com/PrathamHandique/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	"net/http"
	"log"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
func main() {
	r:=mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:8080",r))
}
