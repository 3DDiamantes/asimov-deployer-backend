package domain

type Asset struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	ID   uint64 `json:"id"`
}

type GithubGetReleaseByTagResponse struct {
	ID     int64   `json:"id"`
	Assets []Asset `json:"assets"`
}
