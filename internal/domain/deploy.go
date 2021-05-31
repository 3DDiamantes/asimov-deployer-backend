package domain

type DeployBody struct {
	Owner string `json:"owner"`
	Repo  string `json:"repo"`
	Tag   string `json:"tag"`
	Scope string `json:"scope"`
}
