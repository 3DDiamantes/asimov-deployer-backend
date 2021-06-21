package defines

const (
	GithubURLBase             = "https://api.github.com"
	GithubPathParamRepository = "repository"
	GithubPathParamOwner      = "owner"
	GithubPathParamAssetID    = "asset_id"
	GithubPathParamTag        = "tag"

	GithubURLGetReleaseAsset = GithubURLBase + "/repos/{" + GithubPathParamOwner + "}/{" + GithubPathParamRepository + "}/releases/assets/{" + GithubPathParamAssetID + "}"
	GithubURLGetReleaseByTag = GithubURLBase + "/repos/{" + GithubPathParamOwner + "}/{" + GithubPathParamRepository + "}/releases/tags/{" + GithubPathParamTag + "}"

	GithubHeaderAccept = "application/vnd.github.v3+json"
)
