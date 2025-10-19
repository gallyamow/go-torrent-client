package bittorrent

import (
	"github.com/gallyamow/go-bencoder"
	"os"
)

type TorrentFile[T FileInfo] struct {
	Announce     string  // The announce URL of the tracker (string)
	CreationDate int64   // (optional) the creation time of the torrent, in standard UNIX epoch format (integer, seconds since 1-Jan-1970 00:00:00 UTC)
	Comment      *string // (optional) free-form textual comments of the author (string)
	CreatedBy    *string // (optional) name and version of the program used to create the .torrent (string)
	Encoding     *string // (optional) the string encoding format used to generate the pieces part of the info dictionary in the .torrent metafile (string)
	Info         T
}

type FileInfo interface {
	SingleFileInfo | MultipleFileInfo // union-type реализует один из
	Size() int64
	Hash() [32]byte
	Parse(decoded map[string]any)
}

func OpenTorrentFile[T FileInfo](path string) (*TorrentFile[T], error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	tmp, err := bencoder.Decode(file)
	if err != nil {
		return nil, err
	}

	decoded := tmp.(map[string]any)

	tf := TorrentFile[T]{}

	if val, ok := decoded["creation date"].(int64); ok {
		tf.CreationDate = val
	}

	if val, ok := decoded["announce"].(string); ok {
		tf.Announce = val
	}

	if val, ok := decoded["comment"].(string); ok {
		tf.Comment = &val
	}

	if val, ok := decoded["created by"].(string); ok {
		tf.CreatedBy = &val
	}

	if val, ok := decoded["encoding"].(string); ok {
		tf.Encoding = &val
	}

	var info T
	info.Parse(decoded["info"].(map[string]any))
	tf.Info = info

	return &tf, nil
}
