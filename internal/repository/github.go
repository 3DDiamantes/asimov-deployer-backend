package repository

type GithubRepository interface {
	DownloadAsset(repo string, tag string)
}

type githubRepository struct {
}

func NewGithubRepository() GithubRepository {
	return &githubRepository{}
}

func (r *githubRepository) DownloadAsset(repo string, tag string) {

}
