package bittorrent

import (
	"encoding/binary"
	"errors"
	"net"
	"strconv"
)

type Peer struct {
	IP   net.IP
	Port uint16 // port is positive
}

func (p Peer) String() string {
	return net.JoinHostPort(p.IP.String(), strconv.Itoa(int(p.Port)))
}

type Peers []Peer

// ParsePeers parses a list of peers from a byte slice.
// Peers: (binary model) Instead of using the dictionary model described above, the peers value may be a string
// consisting of multiples of 6 bytes. First 4 bytes are the IP address and last 2 bytes are the port number.
// All in network (big endian) notation.
func ParsePeers(bytes []byte) (Peers, error) {
	const peerSize = 6
	peersCount := len(bytes) / peerSize

	if len(bytes)%6 != 0 {
		return nil, errors.New("invalid peers length")
	}

	peers := make(Peers, peersCount)
	for i := 0; i < peersCount; i++ {
		offset := i * peerSize
		peers[i] = Peer{
			IP:   net.IP(bytes[offset : offset+4]),                      // 4 bytes are the IP address
			Port: binary.BigEndian.Uint16(bytes[offset+4 : offset+4+2]), // last 2 bytes are the port number
		}
	}

	return peers, nil
}
