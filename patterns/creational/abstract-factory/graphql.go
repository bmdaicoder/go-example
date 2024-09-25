package abstractfactory

import "fmt"

type GraphqlService interface {
	CreatePR() error
	MergePR() error
}

type GithubGraphqlService struct{}

func NewGithubGraphqlService() *GithubGraphqlService {
	return &GithubGraphqlService{}
}

func (s *GithubGraphqlService) CreatePR() error {
	fmt.Println("Request Github to CREATE a PR, using GraphQL")
	return nil
}

func (s *GithubGraphqlService) MergePR() error {
	fmt.Println("Request Github to MERGE a PR, using GraphQL")
	return nil
}

type GitlabGraphqlService struct{}

func NewGitlabGraphqlService() *GitlabGraphqlService {
	return &GitlabGraphqlService{}
}

func (s *GitlabGraphqlService) CreatePR() error {
	fmt.Println("Request Github to CREATE a PR, using GraphQL")
	return nil
}

func (s *GitlabGraphqlService) MergePR() error {
	fmt.Println("Request Github to MERGE a PR, using GraphQL")
	return nil
}
