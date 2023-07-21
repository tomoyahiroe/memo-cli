package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	path := os.Getenv("MEMO_COMMAND_STORE_PATH")
	w := flag.String("w", "", "write a memo")
	r := flag.Bool("r", false, "read a memo")
	d := flag.Bool("d", false, "delete a memo")
	flag.Parse()
	fmt.Sprintln(*w, *r, *d)
	if *w != "" {
		writeMemo(*w, path)
	}

	if *r == true {
		readMemo(path)
	}
	if *d == true {
		clearMemo(path)
	}
}

func writeMemo(memo string, path string) {
	bar := "====================" + time.Now().Format("2006-01-02T15:04:05Z07:00") + "====================\n"
	session := bar + memo + "\n\n\n"

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	data := []byte(session)
	f.Write(data)
	fmt.Println(session, "\nwritten successfully")
}

func readMemo(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}

func clearMemo(path string) {
	os.Remove(path)
}

func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
