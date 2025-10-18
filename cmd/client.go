package main

import (
	"fmt"
	"github.com/gallyamow/go-torrent-client/pkg/bittorrent"
)

func main() {
	tf := bittorrent.TorrentFile[bittorrent.SingleFileInfo]{}
	fmt.Println("Hello, world!", tf)
}
