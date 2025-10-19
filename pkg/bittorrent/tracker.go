package bittorrent

import (
	"fmt"
	"net/url"
	"strconv"
)

// Tracker is an HTTP/HTTPS service which responds to HTTP GET requests.
// The response includes a peer list that helps the client participate in the torrent.
type Tracker struct {
}

// buildTrackerUrl builds torrent file URL.
// The base URL consists of the "announce URL" as defined in the metainfo (.torrent) file.
// The parameters are then added to this URL, using standard CGI methods (i.e. a '?' after the announce URL,
// followed by 'param=value' sequences separated by '&').
func buildTrackerUrl[T FileInfo](tf TorrentFile[T], peerID [20]byte, peerPort uint16) (string, error) {
	pu, err := url.Parse(tf.Announce)
	if err != nil {
		return "", err
	}
	fmt.Println(pu)

	qs := url.Values{}
	hash := tf.Info.Hash()
	qs.Add("info_hash", string(hash[:]))
	qs.Add("peer_id", string(peerID[:]))
	qs.Add("port", strconv.Itoa(int(peerPort)))
	qs.Add("uploaded", "0")
	qs.Add("downloaded", "0")

	qs.Add("left", strconv.FormatInt(tf.Info.Size(), 10))
	// compact: Setting this to 1 indicates that the client accepts a compact response.
	// The peers list is replaced by a peers string with 6 bytes per peer.
	qs.Add("compact", "1")
	// остальные опции - Optional

	pu.RawQuery = qs.Encode()

	return pu.String(), nil
}
