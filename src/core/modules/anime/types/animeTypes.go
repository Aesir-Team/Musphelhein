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

type EpisodeInfo struct {
    Title string
    Link  string
}

type AnimeInfo struct {
    Title      string       `json:"title"`
    Year       string       `json:"year"`
    Episodes   string       `json:"episodes"`
    Audio      string       `json:"audio"`
    Genres     []string     `json:"genres"`
    Synopsis   string       `json:"synopsis"`
    ImageLink  string       `json:"image_link"`
    EpisodeList []EpisodeInfo `json:"episode_list"`
}