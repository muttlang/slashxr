package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
)

func main() {
	fmt.Println("Running")
	currentUsername, err := user.Current()
	homeDir, err := os.ReadDir(string(currentUsername.HomeDir))
	os.Mkdir(fmt.Sprintf("%d\\Appdata\\Temp\\Local\\ElectronWebSocket\\", homeDir), 0777)
	logFile, err := os.Create(fmt.Sprintf("%d\\Appdata\\Temp\\Local\\ElectronWebSocket\\LOG", homeDir))
	for _, i := range homeDir {
		if !i.IsDir() {
			logFile.WriteString(i.Name())
			logFile.WriteString(fmt.Sprint(i))
			fmt.Sprintf("Writing %d to LOG", i)
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	if err != nil {
		log.Fatal(err)
	}
}
