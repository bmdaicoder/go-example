package abstractfactory

import "fmt"

type RestService interface {
	CreatePR() error
	MergePR() error
}

type GithubRestService struct{}

func NewGithubRestService() *GithubRestService {
	return &GithubRestService{}
}

func (s *GithubRestService) CreatePR() error {
	fmt.Println("Request Github to CREATE a PR, using REST")
	return nil
}

func (s *GithubRestService) MergePR() error {
	fmt.Println("Request Github to MERGE a PR, using REST")
	return nil
}

type GitlabRestService struct{}

func NewGitlabRestService() *GitlabRestService {
	return &GitlabRestService{}
}

func (s *GitlabRestService) CreatePR() error {
	fmt.Println("Request Github to CREATE a PR, using REST")
	return nil
}

func (s *GitlabRestService) MergePR() error {
	fmt.Println("Request Github to MERGE a PR, using REST")
	return nil
}
