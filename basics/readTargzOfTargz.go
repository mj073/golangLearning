package main

import (
	"os"
	"fmt"
	"compress/gzip"
	"archive/tar"
)

func main(){
	name := "targzOfTargz.tar.gz"
	f, err := os.Open(name)
	if err != nil {
		fmt.Println("failed to open file:",name,"ERROR:",err)
		os.Exit(1)
	}

	r,err := gzip.NewReader(f)
	if err != nil {
		fmt.Println("failed to create gzip reader ERROR:",err)
		os.Exit(1)
	}
	t := tar.NewReader(r)
	if t != nil {
		t.Next()
	}

}