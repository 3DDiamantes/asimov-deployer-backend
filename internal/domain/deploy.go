package domain

type DeployBody struct {
	Repo  string `json:"repo"`
	Tag   string `json:"tag"`
	Scope string `json:"scope"`
}
