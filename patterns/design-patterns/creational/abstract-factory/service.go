package abstractfactory

import "fmt"

type GitServiceFactory interface {
	CreateRestService() RestService
	CreateGraphqlService() GraphqlService
}

type GithubServiceFactory struct{}

func NewGithubServiceFactory() *GithubServiceFactory {
	return &GithubServiceFactory{}
}

func (f *GithubServiceFactory) CreateRestService() RestService {
	return NewGithubRestService()
}
func (f *GithubServiceFactory) CreateGraphqlService() GraphqlService {
	return NewGithubGraphqlService()
}

type GitlabServiceFactory struct{}

func NewGitlabServiceFactory() *GitlabServiceFactory {
	return &GitlabServiceFactory{}
}

func (f *GitlabServiceFactory) CreateRestService() RestService {
	return NewGitlabRestService()
}
func (f *GitlabServiceFactory) CreateGraphqlService() GraphqlService {
	return NewGitlabGraphqlService()
}

func GetGitServiceFactory(platform string) (GitServiceFactory, error) {
	switch platform {
	case "github":
		return NewGithubServiceFactory(), nil
	case "gitlab":
		return NewGitlabServiceFactory(), nil
	default:
		return nil, fmt.Errorf("platform is wrong")
	}
}
