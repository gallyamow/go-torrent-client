package bittorrent

import (
	"encoding/binary"
	"fmt"
	"net"
	"strings"
)

type Peer struct {
	// peer id: peer's self-selected ID, as described above for the tracker RequestAnnounce (string)
	// @see NoPeerID
	PeerID string
	// ip: peer's IP address either IPv6 (hexed) or IPv4 (dotted quad) or DNS name (string)
	IP net.IP
	// port: peer's port number (integer)
	Port uint16
}

func (p Peer) String() string {
	return fmt.Sprintf("Peer{ PeerID: %s, IP: %x, Port: %d }", p.PeerID, p.IP, p.Port)
}

type Peers []Peer

func (ps Peers) String() string {
	var sb strings.Builder
	for _, p := range ps {
		sb.WriteString(p.String())
		sb.WriteString(", ")
	}
	return sb.String()
}

func decodeDictPeer(decoded map[string]any) Peer {
	var p Peer

	if v, ok := decoded["peer id"].(string); ok {
		p.PeerID = v
	}
	if v, ok := decoded["ip"].(string); ok {
		p.IP = net.ParseIP(v)
	}
	if v, ok := decoded["port"].(int64); ok {
		p.Port = uint16(v)
	}

	return p
}

func decodeBinaryPeer(data []byte) []Peer {
	const peerLen = 6
	count := len(data) / peerLen
	peers := make([]Peer, 0, count)

	for i := 0; i < count; i++ {
		offset := i * peerLen
		ip := net.IP(data[offset : offset+4])
		port := binary.BigEndian.Uint16(data[offset+4 : offset+6])
		peers = append(peers, Peer{IP: ip, Port: port})
	}

	return peers
}
