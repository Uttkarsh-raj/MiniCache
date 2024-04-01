package handler

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Uttkarsh-raj/minicache/database"
	"github.com/Uttkarsh-raj/minicache/model"
)

var (
	dirPath      = os.Getenv("USERPROFILE") + "\\Cache\\"
	fileNameSet  = "set.txt"
	fileNameList = "list.txt"
	filePathSet  = dirPath + fileNameSet
	filePathList = dirPath + fileNameList
)

// Initialize
func Initialize(db *database.Database, newList *model.CommandsList) {
	if _, err := os.Stat(filePathSet); os.IsNotExist(err) {
		fmt.Println("Starting a new DB.")
	} else {
		InitilizeDB(db)
		InitializeList(newList)
	}
}

func InitilizeDB(db *database.Database) {
	file, err := os.Open(filePathSet)
	if err != nil {
		fmt.Println("Error opening the file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		var set []string
		set = append(set, "SET")
		parts = append(set, parts...)
		SetRequestHandler(parts, db)
	}
}

func InitializeList(newList *model.CommandsList) {
	file, err := os.Open(filePathList)
	if err != nil {
		fmt.Println("Error opening the file")
	}
	defer file.Close()
	commandList := make(map[string]*model.List)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		list := &model.List{}
		key := parts[0]
		list.List = parts[1:]
		commandList[key] = list
	}
	newList.Store = commandList
}

// Store in set.txt

func StoreData(db *database.Database) {
	if _, err := os.Stat(filePathSet); os.IsNotExist(err) {
		fmt.Println("File Does Not Exists")
		CreateAndStore(db)
	} else {
		fmt.Println("File Found")
		Store(db)
	}
}

func Store(db *database.Database) {
	file, err := os.OpenFile(filePathSet, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
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
	file, err := os.Create(filePathSet)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	Store(db)
	defer file.Close()
}

// Store in list.txt

func StoreDataList(newList *model.CommandsList) {
	if _, err := os.Stat(filePathList); os.IsNotExist(err) {
		fmt.Println("File Does Not Exists")
		CreateAndStoreList(newList)
	} else {
		fmt.Println("File Found")
		StoreList(newList)
	}
}
func CreateAndStoreList(newList *model.CommandsList) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
	}
	file, err := os.Create(filePathList)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	StoreList(newList)
	defer file.Close()
}
func StoreList(newList *model.CommandsList) {
	file, err := os.OpenFile(filePathList, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println("Error opening the file")
	}
	defer file.Close()
	for i := range newList.Store {
		var ans string
		for j := 0; j < len(newList.Store[i].List); j++ {
			ans += newList.Store[i].List[j]
			ans += " "
		}
		data := []byte(i + " " + ans + "\n")
		_, err = file.Write(data)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Println("Data written to file successfully!")
}
