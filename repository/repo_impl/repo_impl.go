package repo_impl

import (
	"context"

	"example.com/test/model"
	"example.com/test/model/response"
	"example.com/test/repository"
)

type RepoImpl struct {
}

func NewRepo() repository.Repo {
	return &RepoImpl{}
}

func (r *RepoImpl) Sum(context context.Context, input model.Input) (response.Output, error) {
	result := input.Number1 + input.Number2

	output := response.Output{}

	output.Sum = result
	return output, nil

}
