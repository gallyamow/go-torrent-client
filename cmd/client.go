package main

import (
	"fmt"
	"github.com/gallyamow/go-torrent-client/pkg/bittorrent"
)

func main() {
	tf := bittorrent.TorrentFile[bittorrent.SingleFileInfo]{}
	fmt.Println("Hello, world!", tf)

	mf := bittorrent.MultipleFileInfo{}
	fmt.Println(mf)

	mf2 := bittorrent.MultipleFileInfo{Name: bittorrent.Ptr("string")}
	fmt.Println(mf2)
}
