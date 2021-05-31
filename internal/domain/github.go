package domain

type Asset struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type GithubGetReleaseByTagResponse struct {
	ID     int64   `json:"id"`
	Assets []Asset `json:"assets"`
}
