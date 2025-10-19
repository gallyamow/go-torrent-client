package bittorrent

import (
	"fmt"
	"net/url"
)

// Tracker is an HTTP/HTTPS service which responds to HTTP GET requests.
// The response includes a peer list that helps the client participate in the torrent.
// The base URL consists of the "announce URL" as defined in the metainfo (.torrent) file.
// The parameters are then added to this URL, using standard CGI methods (i.e. a '?' after the announce URL,
// followed by 'param=value' sequences separated by '&').
type Tracker struct {
}

func buildTrackerUrl[T SingleFileInfo](tf TorrentFile[T]) (string, error) {
	parsedUrl, err := url.Parse(tf.Announce)
	if err != nil {
		return "", err
	}
	fmt.Println(parsedUrl)

	return "", nil
}
