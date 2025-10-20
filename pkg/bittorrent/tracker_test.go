package bittorrent

import "testing"

func TestDecodeDictTrackerPeer(t *testing.T) {
	peer := decodeDictTrackerPeer(map[string]any{
		"peer id": "peerA",
		"ip":      "192.168.1.10",
		"port":    int64(6881),
	})

	if peer.PeerID != "peerA" {
		t.Errorf("expected %v peer, got %v", "peerA", peer.PeerID)
	}
	if peer.IP != "192.168.1.10" {
		t.Errorf("expected %v peer, got %v", "192.168.1.10", peer.IP)
	}
	if peer.Port != 6881 {
		t.Errorf("expected %v peer, got %v", 6881, peer.Port)
	}
}

func TestDecodeBinaryTrackerPeers(t *testing.T) {
	peers := decodeBinaryTrackerPeers([]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c})

	if len(peers) != 2 {
		t.Errorf("expected %v peers, got %v", 3, len(peers))
	}

	peer := peers[0]
	if peer.IP != "1.2.3.4" {
		t.Errorf("expected %v peers, got %v", "1.2.3.4", peer.IP)
	}
	if peer.Port != 1286 { // 0x5 0x6 = 5*256+6 {
		t.Errorf("expected %v peers, got %v", 1286, peer.Port)
	}

	peer = peers[1]
	if peer.IP != "7.8.9.10" {
		t.Errorf("expected %v peers, got %v", "7.8.9.10", peer.IP)
	}
	if peer.Port != 2828 { // 0xB 0x1C = 11*256+12
		t.Errorf("expected %v peers, got %v", 2828, peer.Port)
	}
}
