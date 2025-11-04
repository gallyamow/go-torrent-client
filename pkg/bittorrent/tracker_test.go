package bittorrent

import (
	"testing"
)

func TestBuildRequestUrl(t *testing.T) {
	var tests = []struct {
		name string
		req  TrackerRequest
		url  string
	}{
		{
			"compact",
			TrackerRequest{
				InfoHash: [20]byte{'a', 'b', 'c', 'd', 'e', 'f', 'h', 'g', 'i', 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
				PeerID:   [20]byte{'G', 'T', 'C', 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
				Compact:  true,
			},
			"https://test/announce?compact=1&downloaded=0&event=&info_hash=abcdefhgi%00%01%02%03%04%05%06%07%08%09%00&left=0&no_peer_id=0&peer_id=GTC%04%05%06%07%08%09%0A%0B%0C%0D%0E%0F%10%11%12%13%14&port=0&uploaded=0",
		},
		{
			"not compact as default",
			TrackerRequest{
				InfoHash: [20]byte{'a', 'b', 'c', 'd', 'e', 'f', 'h', 'g', 'i', 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
				PeerID:   [20]byte{'G', 'T', 'C', 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			},
			"https://test/announce?compact=0&downloaded=0&event=&info_hash=abcdefhgi%00%01%02%03%04%05%06%07%08%09%00&left=0&no_peer_id=0&peer_id=GTC%04%05%06%07%08%09%0A%0B%0C%0D%0E%0F%10%11%12%13%14&port=0&uploaded=0",
		},
		{
			"downloaded and uploaded",
			TrackerRequest{
				InfoHash:   [20]byte{'a', 'b', 'c', 'd', 'e', 'f', 'h', 'g', 'i', 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
				PeerID:     [20]byte{'G', 'T', 'C', 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
				Downloaded: 10,
				Uploaded:   7,
			},
			"https://test/announce?compact=0&downloaded=10&event=&info_hash=abcdefhgi%00%01%02%03%04%05%06%07%08%09%00&left=0&no_peer_id=0&peer_id=GTC%04%05%06%07%08%09%0A%0B%0C%0D%0E%0F%10%11%12%13%14&port=0&uploaded=7",
		},
	}

	for _, tt := range tests {
		u, err := buildRequestUrl("https://test/announce", tt.req)
		if err != nil {
			t.Fatalf("buildRequestUrl() error = %v", err)
		}

		if u != tt.url {
			t.Errorf("url = %q, want %q", u, tt.url)
		}
	}
}
