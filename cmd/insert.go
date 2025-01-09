package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/c8763yee/mygo-backend/pkg/database"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run cmd/insert.go <*.(json|csv)>")
		os.Exit(1)
	}

	filename, err := filepath.Abs(os.Args[1])
	if err != nil {
		fmt.Printf("Error resolving file path: %v\n", err)
		os.Exit(1)
	}

	if err := database.ImportDataFromFile(database.DB, filename); err != nil {
		fmt.Printf("Error importing data: %v\n", err)
		os.Exit(1)
	}
}
