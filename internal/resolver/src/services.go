package resolver

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	ytdl "github.com/kkdai/youtube/v2"
)

type Resolver interface {
	Resolve(identifier string) (io.Reader, error)
}

type idResolver struct{}

func NewIDResolver() Resolver {
	return &idResolver{}
}

func (r idResolver) Resolve(identifier string) (io.Reader, error) {
	var videoID string
	var err error

	if strings.HasPrefix(identifier, "http") {
		// Handle YouTube URL
		videoID = extractVideoID(identifier)
		if videoID == "" {
			return nil, fmt.Errorf("invalid YouTube link")
		}
	} else if strings.HasPrefix(identifier, "ytsearch:") {
		// Handle search query with ytsearch: prefix
		query := strings.TrimPrefix(identifier, "ytsearch:")
		videoID, err = searchVideo(query)
		if err != nil {
			return nil, fmt.Errorf("search failed: %w", err)
		}
	} else {
		// Treat as bare video ID
		videoID = identifier
	}

	client := &ytdl.Client{}
	video, err := client.GetVideo(videoID)
	if err != nil {
		return nil, fmt.Errorf("failed to get video: %w", err)
	}

	formats := video.Formats
	var bestFormat *ytdl.Format

	for i, f := range formats {
		if strings.Contains(f.MimeType, "audio") && strings.Contains(f.MimeType, "opus") {
			if bestFormat == nil || f.Bitrate > bestFormat.Bitrate {
				bestFormat = &formats[i]
			}
		}
	}

	if bestFormat == nil {
		return nil, fmt.Errorf("no Opus audio format found")
	}

	stream, _, err := client.GetStream(video, bestFormat)
	if err != nil {
		return nil, fmt.Errorf("failed to get stream: %w", err)
	}

	return stream, nil
}

func extractVideoID(url string) string {
	if strings.Contains(url, "youtube.com/watch") {
		if idx := strings.Index(url, "v="); idx != -1 {
			id := url[idx+2:]
			if ampIdx := strings.Index(id, "&"); ampIdx != -1 {
				id = id[:ampIdx]
			}
			return id
		}
	}

	if strings.Contains(url, "youtu.be/") {
		idx := strings.LastIndex(url, "/")
		if idx != -1 {
			return url[idx+1:]
		}
	}

	return ""
}

type YouTubeSearchResult struct {
	Items []struct {
		ID struct {
			VideoID string `json:"videoId"`
		} `json:"id"`
	} `json:"items"`
}

func searchVideo(query string) (string, error) {
	searchURL := fmt.Sprintf("https://www.youtube.com/results?search_query=%s", url.QueryEscape(query))

	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to search: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	bodyStr := string(body)
	idx := strings.Index(bodyStr, `"videoId":"`)
	if idx == -1 {
		return "", fmt.Errorf("no video results found for query: %s", query)
	}

	startIdx := idx + len(`"videoId":"`)
	endIdx := strings.Index(bodyStr[startIdx:], `"`)
	if endIdx == -1 {
		return "", fmt.Errorf("invalid video ID format in search results")
	}

	videoID := bodyStr[startIdx : startIdx+endIdx]
	if videoID == "" {
		return "", fmt.Errorf("empty video ID from search results")
	}

	return videoID, nil
}
