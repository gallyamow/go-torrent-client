package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/gallyamow/go-torrent-client/pkg/bittorrent"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cmdArgs, err := readCmdArgs()
	if err != nil {
		log.Fatalf("failed to parse command %v", err)
	}

	tf, err := bittorrent.OpenTorrentFile[bittorrent.SingleFileInfo](cmdArgs.torrentFile)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	fmt.Printf("torrent file: %s\n", tf)

	tracker := bittorrent.NewTracker()

	resp, err := tracker.RequestAnnounce(ctx, tf)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}

	fmt.Printf("announce response: %s\n", resp)
	if resp.FailureReason != nil {
		log.Fatalf("failed to request announce: %q", *resp.FailureReason)
	}
}

type cmdArgs struct {
	torrentFile string
	verbose     bool
}

func readCmdArgs() (cmdArgs, error) {
	torrentFile := flag.String("path", "./example.torrent", "Torrent file")
	verbose := flag.Bool("verbose", false, "Verbose mode")
	flag.Parse()

	return cmdArgs{
		torrentFile: *torrentFile,
		verbose:     *verbose,
	}, nil
}
