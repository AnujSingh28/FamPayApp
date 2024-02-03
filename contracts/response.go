package contracts

import "time"

type GetVideoResponse struct {
	Title string `json:"title"`
	Description string `json:"description"`
	VideoLink string `json:"video_link"`
	PublishedAt time.Time `json:"published_at"`
	Thumbnail string `json:"thumbnail"`
}
