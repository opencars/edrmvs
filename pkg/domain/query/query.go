package query

import (
	"github.com/opencars/schema"

	"github.com/opencars/edrmvs/pkg/domain/model"
)

var (
	source = schema.Source{
		Service: "edrmvs",
		Version: "1.0",
	}
)

type Query interface {
	Prepare()
	Validate() error
}

func Process(q Query) error {
	q.Prepare()

	return model.Validate(q, "request")
}
