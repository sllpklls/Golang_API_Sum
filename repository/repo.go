package repository

import (
	"context"

	"example.com/test/model"
	"example.com/test/model/response"
)

type Repo interface {
	Sum(context context.Context, input model.Input) (response.Output, error)
}
