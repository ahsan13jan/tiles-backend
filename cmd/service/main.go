package main

import (
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Listing struct {
	Id 			int64	`json:"id"`
	ImageUrl    string 	`json:"imageUrl"`
	Description string 	`json:"description"`
	Category    string 	`json:"category"`
	Price       int 	`json:"price"`
}

var listings []Listing
func GetMainListings() []byte {
	b, _ :=  json.Marshal(listings)
	return b
}

func GetDetails(id int64) []byte {

	var listing Listing
	for _, l := range listings {
		if l.Id == id {
			listing = l
		}

	}

	b, _ :=  json.Marshal(listing)
	return b
}

func main() {
	populateData()

	r := mux.NewRouter()
	r.HandleFunc("/mainListings", handlerListings)
	r.HandleFunc("/getDetails", handlerDetails)

	log.Println("Server running")
	http.ListenAndServe(":8080", handlers.CORS()(r))


}

func handlerListings(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	log.Println(r)
	//enableCors(&w)
	w.Write(GetMainListings())
	w.Header().Set("Content-Type", "application/json")
}

func handlerDetails(w http.ResponseWriter, r *http.Request) {

	log.Println(r.Method)
	log.Println(r)
	//enableCors(&w)

	id, ok := r.URL.Query()["id"]

	if !ok || len(id[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	i, err := strconv.ParseInt(id[0], 10, 64)

	if err != nil {
		log.Println(err)
	}
	w.Write(GetDetails(i))
	w.Header().Set("Content-Type", "application/json")
}

func populateData()  {
	l1 := Listing {
		Id: 1,
		ImageUrl:    "https://hillceramic.se/storage/media/products/2020/06/marmor-klinker-emperador-morkgra-matt-60x60-cm-KLN5143.jpg",
		Description: "Gray exotic tiles in your bathroom",
		Category:    "Tiles",
		Price:       10000,
	}
	l2 := Listing {
		Id: 2,
		ImageUrl:    "https://i.pinimg.com/736x/43/c0/5f/43c05f1189c02595b15223e22c9cc530.jpg",
		Description: "Rio Grey EcoTec Matt Porcelain Floor Tile",
		Category:    "Luxury",
		Price:       20000,
	}

	l3 := Listing {
		Id: 3,
		ImageUrl:    "https://i.pinimg.com/736x/43/c0/5f/43c05f1189c02595b15223e22c9cc530.jpg",
		Description: "Swedish tiles",
		Category:    "Tiles",
		Price:       30000,
	}
	l4 := Listing {
		Id: 4,
		ImageUrl:    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTMcI93wHuFUXHWJKT29FRQYs7kIyp_fjQHkQ&usqp=CAU",
		Description: "Klinker Flodsten Mörkgrå Matt Rund",
		Category:    "Luxury",
		Price:       10000,
	}

	listings = append(listings, l1)
	listings = append(listings, l2)
	listings = append(listings, l3)
	listings = append(listings, l4)

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Authorization, X-Huzu-User, Content-Type, Accept")

}
