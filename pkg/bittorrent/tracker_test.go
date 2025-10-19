package bittorrent

import (
	"testing"
)

func TestBuildTrackerURL(t *testing.T) {
	tf := TorrentFile[SingleFileInfo]{
		Announce: "https://example.com/announce",
	}

	var tests = []struct {
		name     string
		peerPort uint16
		peerID   [20]byte
		expected string
	}{
		{"first", 6432, [20]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, "https://example.com/announce?compact=1&downloaded=0&info_hash=%00%00%00%00%00%00%00%00%00%00%00%00%00%00%00%00%00%00%00%00%00%00%00%00%00%00%00%00%00%00%00%00&left=0&peer_id=%01%02%03%04%05%06%07%08%09%0A%0B%0C%0D%0E%0F%10%11%12%13%14&port=6432&uploaded=0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := buildTrackerUrl(tf, [20]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, tt.peerPort)

			if err != nil {
				t.Fatalf("buildTrackerUrl() error = %v", err)
			}

			if u != tt.expected {
				t.Errorf("buildTrackerUrl() = %q, want %q", u, tt.expected)
			}
		})
	}
}
