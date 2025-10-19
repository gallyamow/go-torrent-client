package bittorrent

import "testing"

func TestParsePeers(t *testing.T) {
	var tests = []struct {
		name     string
		input    []byte
		expected []string
	}{
		{"one", []byte{8, 8, 8, 8, 0xFF, 0xFF}, []string{"8.8.8.8:65535"}},
		{"two", []byte{8, 8, 4, 4, 0x00, 0x50, 7, 7, 7, 7, 0x1F, 0x90}, []string{"8.8.4.4:80", "7.7.7.7:8080"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			peers, err := ParsePeers(tt.input)
			if err != nil {
				t.Fatalf("ParsePeers() error = %v", err)
			}
			if len(peers) != len(tt.expected) {
				t.Errorf("expected %d peer, got %d", len(tt.expected), len(peers))
			}

			for i, _ := range tt.expected {
				if peers[i].String() != tt.expected[i] {
					t.Errorf("expected %s, got %s", tt.expected[i], peers[i])
				}
			}
		})
	}
}

func TestPeerString(t *testing.T) {
	peer := Peer{IP: []byte{8, 8, 8, 8}, Port: 65535}
	if peer.String() != "8.8.8.8:65535" {
		t.Errorf("expected 8.8.8.8:65535, got %s", peer.String())
	}
}
