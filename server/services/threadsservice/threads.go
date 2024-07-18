package threadsservice

import (
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/pleimer/ticketer/server/integration/nylas"
	"go.uber.org/zap"
)

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=threads-models.cfg.yaml ../../../internal/api/threads.json
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=threads-service.cfg.yaml ../../../internal/api/threads.json

// ThreadsService is just a Nylas api passthrough
type ThreadsService struct {
	nylasCli *nylas.NylasClient
	*zap.Logger
}

func NewThreadsService(logger *zap.Logger, nylasCli *nylas.NylasClient) *ThreadsService {
	return &ThreadsService{
		Logger:   logger,
		nylasCli: nylasCli,
	}
}

// Get messages in Thread
// (GET threads/{threadId})
func (t *ThreadsService) GetThreadsThreadId(ctx echo.Context, threadId string) error {

	// TODO: paging
	r, err := t.nylasCli.ListThreadMessages(threadId)
	if err != nil {
		t.Error("connecting to email server", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, "connecting to email server")
	}

	reply := ThreadResponse{}
	err = copier.Copy(&reply, r)
	if err != nil {
		// dev error
		t.Fatal("mismatched structs")
	}

	return ctx.JSON(http.StatusOK, reply)
}
