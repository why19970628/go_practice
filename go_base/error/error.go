package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
)

func CopyFile(src, dst string) error {
	r, err := os.Open(src)
	if err != nil {
		return errors.Wrap(err, "open source")
	}
	defer r.Close()

	w, err := os.Create(dst)
	if err != nil {
		return errors.Wrap(err, "create dest")
	}

	nn, err := io.Copy(w, r)
	if err != nil {
		w.Close()
		os.Remove(dst)
		return errors.Wrap(err, "copy body")
	}

	if err := w.Close(); err != nil {
		os.Remove(dst)
		return errors.Wrapf(err, "close dest, nn=%v", nn)
	}

	return nil
}

func LoadSystem() error {
	src, dst := "src.txt", "dst.txt"
	if err := CopyFile(src, dst); err != nil {
		fmt.Println("LoadSystem err:", err)
		return errors.WithMessage(err, fmt.Sprintf("load src=%v, dst=%v", src, dst))
	}

	//if err := CopyFile(src, dst); err != nil {
	//	//fmt.Println("LoadSystem err:", err)
	//	return errors.Wrapf(err, fmt.Sprintf("load src=%v, dst=%v", src, dst))
	//}

	// Do other jobs.

	return nil
}

func main() {
	if err := LoadSystem(); err != nil {
		fmt.Printf("main err: %+v\n", err)
	}
}
