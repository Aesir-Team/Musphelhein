package types

type Anime struct {
	Title     string `json:"title"`
	Link      string `json:"link"`
	ImageLink string `json:"image_link"` 
}
type AnimeReleases struct {
	Title         string `json:"title"`
	Link          string `json:"link"`
	ImageLink     string `json:"image_link"`
	EpisodeTitle  string `json:"episode_title"`
	EpisodeNumber string `json:"episode_number"`
	Quality       string `json:"quality"`
}