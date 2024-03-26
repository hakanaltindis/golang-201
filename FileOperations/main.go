package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	// 01 - Create a new file
	newFile, err := os.Create("myfile.txt")
	if err != nil {
		log.Fatal(err)
	}
	newFile.Close()

	// 02 - Get File Stats
	fileInfo, err := os.Stat("myfile.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("File Size:", fileInfo.Size())
	fmt.Println("File Permissions:", fileInfo.Mode())
	fmt.Println("File Last Modified:", fileInfo.ModTime())
	fmt.Println("Is Directory:", fileInfo.IsDir())

	// 03 - Rename file
	originalPath := "myfile.txt"
	newPath := "./moved/test.txt"
	linkErr := os.Rename(originalPath, newPath)
	if linkErr != nil {
		log.Fatal(linkErr)
	}

	// 04 - Check file existance
	info, err := os.Stat("myfile.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exists")
		}
	}
	log.Println("File does exists. File Information: ")
	log.Println(info)

	// 05 - Open and Close File
	file, err := os.Open("myfile.txt")
	if err != nil {
		log.Fatal(err)
	}
	// I think, I need to check err and than I can write defer statment. Otherwise file will be null while is calling file.Close()
	defer file.Close()

	// 05_02 - Open and Close File
	openedfile, err := os.OpenFile("myfile.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	openedfile.Close()

	// 06 - Check read & write permissions
	readFile, err := os.OpenFile("myfile.txt", os.O_WRONLY, os.ModeAppend)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Hata: Yazma izni reddedildi.")
		}
		log.Fatal(err)
	}
	readFile.Close()

	// 07 - Copy file
	originalFile, err := os.Open("myfile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()

	copyFile, err := os.Create("myfile-copy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer copyFile.Close()

	bytesWritten, err := io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesWritten)

	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}

	// 08 - Write to file
	writerFile, err := os.OpenFile("myfile.txt", os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer writerFile.Close()

	byteSlice := []byte("Bu dosyaya yazdık\n")
	len, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", len)
	// sample: get a value from user using console application and write it to file

	// 09 - Fast Write to file
	err = os.WriteFile("myfile.txt", []byte("Hello, World!"), 0666)
	if err != nil {
		log.Fatal(err)
	}

	// 10 - Work with temporary file and folder
	tempDirPath, err := os.MkdirTemp("", "tempFolder")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Temporary directory is created:", tempDirPath)

	tempFile, err := os.CreateTemp(tempDirPath, "tempfile.txt")
	fmt.Println("Temporary file is created:", tempFile)
	if err != nil {
		log.Fatal(err)
	}

	err = tempFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Deleting...
	err = os.Remove(tempFile.Name())
	if err != nil {
		log.Fatal(err)
	} else {
		log.Fatalf("%s is deleted.", tempFile.Name())
	}

	err = os.Remove(tempDirPath)
	if err != nil {
		log.Fatal(err)
	}

}
