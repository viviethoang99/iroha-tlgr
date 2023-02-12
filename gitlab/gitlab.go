package gitlab

import (
	"github.com/xanzy/go-gitlab"
	"iroha-tlgr/utils"
	"log"
)

func GetInfoUserCreateMergeRequest(config utils.Config, shouldGetData bool) (*gitlab.MergeRequest, error) {
	if !shouldGetData {
		return nil, nil
	}
	git, err := gitlab.NewClient(
		config.GitlabConfig.AccessToken,
		gitlab.WithBaseURL(config.GitlabConfig.BaseUrl),
	)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	opts := gitlab.ListProjectMergeRequestsOptions{
		Scope: gitlab.String("all"),
		State: gitlab.String("merged"),
	}
	mergeRequests, _, err := git.MergeRequests.ListProjectMergeRequests(
		config.GitlabConfig.IdProject,
		&opts,
	)
	if err != nil {
		log.Fatal("error", err)
		return nil, err
	}
	return mergeRequests[0], nil
}
