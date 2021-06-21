package domain

type DeployBody struct {
	Owner string `json:"owner"`
	Repo  string `json:"repo"`
	Tag   string `json:"tag"`
	Scope string `json:"scope"`
}

func (body *DeployBody) IsValid() bool {
	return body.Owner != "" &&
		body.Repo != "" &&
		body.Tag != "" &&
		body.Scope != ""
}