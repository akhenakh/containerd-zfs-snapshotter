package main

import (
	"log"

	zfs "github.com/bicomsystems/go-libzfs"
)

func main() {
	d, err := zfs.DatasetOpen("zp00/containerd/98551")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(d)
}
