package main

import "fmt"
import "os"

var path = "/home/reza/go/src/problem3/folder1/tes1.txt"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}
func createFile() {
	// deteksi apakah file sudah ada
	var _, err = os.Stat(path)
	// buat file baru jika belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}
	fmt.Println(path, "NEW")
}
func main() {
	createFile()
}
