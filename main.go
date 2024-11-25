package main

import (
    "fmt"
    "salary_mangemet_site/database"
)

func main() {
    fmt.Println("Connecting to the database...")
    database.Connect()
}
