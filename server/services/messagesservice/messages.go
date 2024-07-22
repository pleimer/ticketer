package messagesservice

import (
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/pleimer/ticketer/server/integration/nylas"
	"go.uber.org/zap"
)

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=messages-models.cfg.yaml ../../../internal/api/messages.json
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=messages-service.cfg.yaml ../../../internal/api/messages.json

// MessagesService is just a Nylas api passthrough
type MessagesService struct {
	nylasCli *nylas.NylasClient
	*zap.Logger
}

func NewMessagesService(logger *zap.Logger, nylasCli *nylas.NylasClient) *MessagesService {
	return &MessagesService{
		Logger:   logger,
		nylasCli: nylasCli,
	}
}

// Get messages in Thread
// (GET threads/{threadId})
func (t *MessagesService) ListThreadMessages(ctx echo.Context, threadId string) error {

	// TODO: paging
	r, err := t.nylasCli.ListThreadMessages(threadId)
	if err != nil {
		t.Error("connecting to email server", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, "connecting to email server")
	}

	// filter out messages not in 'inbox' e.g. in 'sent'
	folders, err := t.nylasCli.GetFolders()
	if err != nil {
		t.Error("folders from mail server", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, "connecting to email server")
	}

	folderIDSet := map[string]string{}
	for _, f := range folders.Data {
		folderIDSet[f.ID] = f.Name
	}

	reply := []Message{}

	for _, msg := range r.Data {
		// only want messages in "Inbox"
		if name, ok := folderIDSet[msg.Folders[0]]; ok && name == "Inbox" { // only 1 with outlook
			var rMsg Message
			err = copier.Copy(&rMsg, msg)
			if err != nil {
				// dev error
				t.Fatal("mismatched structs")
			}
			reply = append(reply, rMsg)
		}
	}

	return ctx.JSON(http.StatusOK, reply)
}

// reply to thread
// (POST /threads/reply)
func (t *MessagesService) ReplyToThread(ctx echo.Context) error {

	var replyToThread ReplyToThreadJSONBody

	if err := ctx.Bind(&replyToThread); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	// TODO: make this transactional

	// fetch all participants in the thread so that we can "replyAll"

	thread, err := t.nylasCli.GetThread(*replyToThread.ThreadId)
	if err != nil {
		t.Error("reading thread from mail server", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, "mail server")
	}

	r, err := t.nylasCli.SendMessage(&nylas.SendMessageRequest{
		To:               thread.Data.Participants,
		ReplyToMessageID: thread.Data.LatestDraftOrMessage.ID, // reply to latest message to keep thread going
		Body:             *replyToThread.Body,
		TrackingOptions: &nylas.TrackingOptions{
			ThreadReplies: true,
		},
	})
	if err != nil {
		t.Error("submitting thread reply", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, "submitting reply")
	}

	return ctx.JSON(http.StatusOK, r)
}
