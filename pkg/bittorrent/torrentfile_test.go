package bittorrent

import (
	"fmt"
	"testing"
)

func TestOpenTorrentFile(t *testing.T) {
	tf, err := OpenTorrentFile[SingleFileInfo]("../../kubuntu-7.10-alternate-amd64.iso.torrent")
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%#v\n", tf)
}
