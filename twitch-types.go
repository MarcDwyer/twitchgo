package twitchgo

import "time"

type TResponse struct {
	Total   int        `json:"_total"`
	Streams []TStreams `json:"streams"`
}
type SingleResponse struct {
	Stream *TStreams `json:"stream"`
}
type TStreams struct {
	ID         int64     `json:"_id"`
	AverageFps float64   `json:"average_fps"`
	Channel    TChannel  `json:"channel"`
	CreatedAt  time.Time `json:"created_at"`
	Delay      int       `json:"delay"`
	Game       string    `json:"game"`
	IsPlaylist bool      `json:"is_playlist"`
	Preview    struct {
		Large    string `json:"large"`
		Medium   string `json:"medium"`
		Small    string `json:"small"`
		Template string `json:"template"`
	} `json:"preview"`
	VideoHeight int `json:"video_height"`
	Viewers     int `json:"viewers"`
}
type TChannel struct {
	ID                           int         `json:"_id"`
	BroadcasterLanguage          string      `json:"broadcaster_language"`
	CreatedAt                    time.Time   `json:"created_at"`
	DisplayName                  string      `json:"display_name"`
	Followers                    int         `json:"followers"`
	Game                         string      `json:"game"`
	Language                     string      `json:"language"`
	Logo                         string      `json:"logo"`
	Mature                       bool        `json:"mature"`
	Name                         string      `json:"name"`
	Partner                      bool        `json:"partner"`
	ProfileBanner                string      `json:"profile_banner"`
	ProfileBannerBackgroundColor interface{} `json:"profile_banner_background_color"`
	Status                       string      `json:"status"`
	UpdatedAt                    time.Time   `json:"updated_at"`
	URL                          string      `json:"url"`
	VideoBanner                  string      `json:"video_banner"`
	Views                        int         `json:"views"`
}
