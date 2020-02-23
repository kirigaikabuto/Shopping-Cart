package main

import (
	"CartShop/controllers/cartcontroller"
	"CartShop/controllers/productcontroller"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	routes := mux.NewRouter()
	routes.HandleFunc("/", productcontroller.Index)
	routes.HandleFunc("/product", productcontroller.Index)
	routes.HandleFunc("/product/index", productcontroller.Index)

	routes.HandleFunc("/cart", cartcontroller.Index)
	routes.HandleFunc("/cart/index", cartcontroller.Index)
	routes.HandleFunc("/cart/buy", cartcontroller.Buy)
	routes.HandleFunc("/cart/remove", cartcontroller.Remove)
	fs := http.FileServer(http.Dir("static"))
	routes.Handle("/static/", http.StripPrefix("\\static\\", fs))
	fmt.Println("Server is started")
	http.ListenAndServe(":8080", routes)
}
