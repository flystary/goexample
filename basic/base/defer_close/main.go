package main

import (
	"io"
	"log"
	"os"
)

func solution() {
	f, err := os.Open("/gopher.txt")
	if err != nil {
		return
	}
	defer f.Close()
}

func solution01() error {
	f, err := os.Create("./gopher.txt")
	if err != nil {
		return err
	}
	if _, err = io.WriteString(f, "hello gopher"); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}

func solution02() (err error) {
	f, err := os.Create("./gopher.txt")
	if err != nil {
		return
	}

	defer func() {
		closeErr := f.Close()
		if err == nil {
			err = closeErr
		}
	}()

	_, err = io.WriteString(f, "hello gopher")
	return
}

func solution03() error {
	f, err := os.Create("./gopher.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := io.WriteString(f, "hello gophrt"); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

func solution04() error {
	f, err := os.Create("./gopher.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = io.WriteString(f, "hello World"); err != nil {
		return err
	}
	return f.Sync()
}

func main() {
	err := solution01()
	log.Fatalf("%v", err)

}
