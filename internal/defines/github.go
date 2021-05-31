package defines

const (
	GithubToken = "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"

	GithubURLBase             = "https://api.github.com"
	GithubPathParamRepository = "repository"
	GithubPathParamOwner      = "owner"
	GithubPathParamAssetID    = "asset_id"
	GithubPathParamTag        = "tag"

	GithubURLGetReleaseAsset = "/repos/{" + GithubPathParamOwner + "}/{" + GithubPathParamRepository + "}/releases/assets/{" + GithubPathParamAssetID + "}"
	GithubURLGetReleaseByTag = "/repos/{" + GithubPathParamOwner + "}/{" + GithubPathParamRepository + "}/releases/tags/{" + GithubPathParamTag + "}"
)
