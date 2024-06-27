package transaction

import (
	"context"
)

type Transaction interface {
	DoInTx(context.Context, any, func(context.Context) (any, error)) (any, error)
}
