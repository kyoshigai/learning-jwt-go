package main

import "github.com/k-yoshigai/learning-jwt-go/database"

func main() {
	connectionsString := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	database.Connect(connectionsString)
	database.Migrate()
}
