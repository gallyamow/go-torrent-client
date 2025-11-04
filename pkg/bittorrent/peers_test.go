package bittorrent

import "testing"

func TestDecodeDictTrackerPeer(t *testing.T) {
	peer := decodeDictPeer(map[string]any{
		"peer id": "peerA",
		"ip":      "192.168.1.10",
		"port":    int64(6881),
	})

	if peer.PeerID != "peerA" {
		t.Errorf("peer.PeerID = %q, want %q", peer.PeerID, "peerA")
	}
	if peer.IP.String() != "192.168.1.10" {
		t.Errorf("peer.IP = %q, want %q", peer.IP.String(), "192.168.1.10")
	}
	if peer.Port != 6881 {
		t.Errorf("peer.Port = %q, want %q", peer.Port, 6881)
	}
}

func TestDecodeBinaryTrackerPeers(t *testing.T) {
	peers := decodeBinaryPeer([]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c})

	if len(peers) != 2 {
		t.Errorf("len() = %q, want %q", len(peers), 2)
	}

	peer := peers[0]
	if peer.IP.String() != "1.2.3.4" {
		t.Errorf("peer.IP = %q, want %q", peer.IP.String(), "1.2.3.4")
	}
	if peer.Port != 1286 { // 0x5 0x6 = 5*256+6 {
		t.Errorf("peer.Port = %q, want %q", peer.Port, 1286)
	}

	peer = peers[1]
	if peer.IP.String() != "7.8.9.10" {
		t.Errorf("peer.IP = %q, want %q", peer.IP.String(), "7.8.9.10")
	}
	if peer.Port != 2828 { // 0xB 0x1C = 11*256+12
		t.Errorf("peer.Port = %q, want %q", peer.Port, 1286)
	}
}
