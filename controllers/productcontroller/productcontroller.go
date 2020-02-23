package productcontroller

import (
	"CartShop/models"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter,r *http.Request){
	var productModel models.ProductModel
	products,_:=productModel.FindAll()
	data:=map[string]interface{}{
		"products":products,
	}
	tmp,_:=template.ParseFiles("views/product/index.html")
	tmp.Execute(w,data)
}
