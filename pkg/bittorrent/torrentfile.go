package bittorrent

import (
	"errors"
	"fmt"
	"github.com/gallyamow/go-bencoder"
	"os"
)

type FileInfo interface {
	SingleFileInfo | MultipleFileInfo // union-type реализует один из
	Size() int64
	Hash() [20]byte
	Parse(decoded map[string]any)
	String() string
}

type TorrentFile[T FileInfo] struct {
	Announce     string  // The announce URL of the tracker (string)
	CreationDate int64   // (optional) the creation time of the torrent, in standard UNIX epoch format (integer, seconds since 1-Jan-1970 00:00:00 UTC)
	Comment      *string // (optional) free-form textual comments of the author (string)
	CreatedBy    *string // (optional) name and version of the program used to create the .torrent (string)
	Encoding     *string // (optional) the string encoding format used to generate the pieces part of the info dictionary in the .torrent metafile (string)
	Info         T
}

func (tf *TorrentFile[T]) String() string {
	return fmt.Sprintf("TorrentFile{ Announce: %s, CreationDate: %d, Comment: %s, CreatedBy: %s, Encoding: %s, Info: %s }",
		tf.Announce, tf.CreationDate, StringifyPtr(tf.Comment), StringifyPtr(tf.CreatedBy), StringifyPtr(tf.Encoding), tf.Info.String())
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

	if tmp == nil {
		return nil, errors.New("failed to decode torrent file")
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

	if val, ok := decoded["info"].(map[string]any); ok {
		var info T
		info.Parse(val)
		tf.Info = info
	}

	return &tf, nil
}
