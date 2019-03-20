package main

import (
	"archive/zip"
	"fmt"
	"os"
)

func main() {
	r, err := zip.OpenReader("zipOfZip.zip")
	if err != nil {
		fmt.Println("failed to open zip file..ERROR:",err)
		os.Exit(1)
	}
	defer r.Close()
	for _, f := range r.File{
		fmt.Println("name:",f.Name)
		rr, err := zip.OpenReader(f.Name)
		if err != nil {
			fmt.Println("failed to open zipped file..ERROR:",err)
			os.Exit(1)
		}

		if rr != nil {
			for _, ff := range rr.File {
				fmt.Println("rr name:",ff.Name)
				rr.Close()
			}
		}
	}
}



