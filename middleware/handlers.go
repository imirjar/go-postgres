package middleware

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/gorilla/mux"
)

type response struct {
	ID		int64 	`json:"id, omitempty"`
	Message string 	`json:"message, omitempty"`
}

func CreateConnection() *sql.DB {
	err := gotdotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Succesfully connected to postgres")
	return db
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Fatal("Unable to decode the request body. %v", err)
	}

	insertID := insertStock(stock)

	res := response{
		ID: insertID,
		Message: "Stock created succesfully",

	}

	json.NewEncoder(w).Encode(res)

}

func GetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	stock, err := getStock(int64(id))

	if err != nil {
		log.Fatalf("Unable to get stock. %v", err)
	}
}

func GetAllStock(w http.ResponseWriter, r *http.Request) {

}

func UpdateStock(w http.ResponseWriter, r *http.Request) {

}

func DeleteStock(w http.ResponseWriter, r *http.Request) {

}