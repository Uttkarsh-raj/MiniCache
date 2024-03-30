package handler

import (
	"fmt"
	"os"

	"github.com/Uttkarsh-raj/redis-go/database"
)

var (
	dirPath  = os.Getenv("USERPROFILE") + "/Cache/"
	fileName = "set.txt"
	filePath = dirPath + fileName
)

func StoreData(db *database.Database) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("File Does Not Exists")
		CreateAndStore(db)
	} else {
		fmt.Println("File Found")
		Store(db)
	}
}

func Store(db *database.Database) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println("Error opening the file")
	}
	defer file.Close()
	for i := range db.Db {
		data := []byte(i + " " + db.Db[i] + "\n")
		_, err = file.Write(data)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Println("Data written to file successfully!")
}

func CreateAndStore(db *database.Database) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
	}
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	Store(db)
	defer file.Close()
}
