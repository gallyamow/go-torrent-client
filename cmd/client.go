package main

import (
	"flag"
	"fmt"
	"github.com/gallyamow/go-torrent-client/pkg/bittorrent"
	"log"
)

func main() {
	cmdArgs, err := readCmdArgs()
	if err != nil {
		log.Fatalf("failed to parse command %v", err)
	}

	tf, err := bittorrent.OpenTorrentFile[bittorrent.SingleFileInfo](cmdArgs.torrentFile)
	if err != nil {
		log.Fatalf("failed to open file %v", err)
	}

	url, _ := bittorrent.BuildTrackerUrl(tf, [20]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, 6881, 0, 0)
	fmt.Println("url", url)
}

type cmdArgs struct {
	torrentFile string
	verbose     bool
}

func readCmdArgs() (cmdArgs, error) {
	torrentFile := flag.String("path", "./kubuntu-7.10-alternate-amd64.iso.torrent", "Torrent file")
	verbose := flag.Bool("verbose", false, "Verbose mode")
	flag.Parse()

	return cmdArgs{
		torrentFile: *torrentFile,
		verbose:     *verbose,
	}, nil
}
