//go:generate mockgen -package=mocks -destination=../../../mocks/handler_mock.go -source=handler.go

package port

import "context"

type JobHandlerClient interface {
	Run(ctx context.Context) error
}
