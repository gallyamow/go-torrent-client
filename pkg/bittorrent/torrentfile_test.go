package bittorrent

import (
	"testing"
)

func TestOpenTorrentFile(t *testing.T) {
	_, err := OpenTorrentFile[SingleFileInfo]("../../kubuntu-7.10-alternate-amd64.iso.torrent")
	if err != nil {
		t.Fatalf("OpenTorrentFile() error = %v", err)
	}

	//fmt.Printf("%#v\n", tf)
}
