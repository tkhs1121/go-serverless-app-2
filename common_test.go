package main

import "testing"

func TestParseRequest(t *testing.T) {
	tests := []struct{
		name string
		payload Payload
		want string
	}{
		{
			name: "Parse Request URL",
			payload: Payload{`{"url": "https://www.amazon.co.jp"}`},
			want: "https://www.amazon.co.jp",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := parseRequest(tt.payload); got != tt.want {
				if err != nil {
					t.Errorf("%v", err)
				}
				t.Errorf("parseRequest = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestCheckAmazonURL(t *testing.T) {
	tests := []struct{
		name string
		url  string
		want bool
	}{
		{
			name: "Correct URL",
			url:  "https://www.amazon.co.jp",
			want: true,
		},
		{
			name: "Invalid URL",
			url:  "http://konozama.co.jp",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := checkAmazonURL(tt.url); got != tt.want {
				if err != nil {
					t.Errorf("%v", err)
				}
				t.Errorf("checkAmazonURL = %v, want %v", got, tt.want)
			}
		})
	}

}
