package bittorrent

import (
	"github.com/gallyamow/go-bencoder"
	"os"
	"reflect"
)

type TorrentFile[T SingleFileInfo | MultipleFileInfo] struct {
	Announce     string  // The announce URL of the tracker (string)
	CreationDate int64   // (optional) the creation time of the torrent, in standard UNIX epoch format (integer, seconds since 1-Jan-1970 00:00:00 UTC)
	Comment      *string // (optional) free-form textual comments of the author (string)
	CreatedBy    *string // (optional) name and version of the program used to create the .torrent (string)
	Encoding     *string // (optional) the string encoding format used to generate the pieces part of the info dictionary in the .torrent metafile (string)
	Info         *T
}

func OpenTorrentFile[T SingleFileInfo | MultipleFileInfo](path string) (*TorrentFile[T], error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	tmp, err := bencoder.Decode(file)
	if err != nil {
		return nil, err
	}

	torrentFile, err := parseTorrentFile[T](tmp.(map[string]any))
	if err != nil {
		return nil, err
	}

	return &torrentFile, nil
}

func parseTorrentFile[T SingleFileInfo | MultipleFileInfo](decoded map[string]any) (TorrentFile[T], error) {
	torrentFile := TorrentFile[T]{}

	if val, ok := decoded["creation date"].(int64); ok {
		torrentFile.CreationDate = val
	}

	if val, ok := decoded["announce"].(string); ok {
		torrentFile.Announce = val
	}

	if val, ok := decoded["comment"].(string); ok {
		torrentFile.Comment = &val
	}

	if val, ok := decoded["created by"].(string); ok {
		torrentFile.CreatedBy = &val
	}

	if val, ok := decoded["encoding"].(string); ok {
		torrentFile.Encoding = &val
	}

	var t T
	switch reflect.TypeOf(t) {
	case reflect.TypeOf(SingleFileInfo{}):
		info := parseSingleFileInfo(decoded["info"].(map[string]any))
		torrentFile.Info = any(&info).(*T)
	case reflect.TypeOf(MultipleFileInfo{}):
		panic("not implemented")
	}

	return torrentFile, nil
}
