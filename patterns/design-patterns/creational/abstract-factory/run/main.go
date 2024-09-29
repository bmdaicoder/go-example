package main

import abstractfactory "go-example/patterns/design-patterns/creational/abstract-factory"

func main() {
	github, err := abstractfactory.GetGitServiceFactory("github")
	if err != nil {
		return
	}

	rest := github.CreateRestService()
	if err := rest.MergePR(); err != nil {
		return
	}

	gitlab, err := abstractfactory.GetGitServiceFactory("gitlab")
	if err != nil {
		return
	}

	graphql := gitlab.CreateGraphqlService()
	if err := graphql.CreatePR(); err != nil {
		return
	}
}
