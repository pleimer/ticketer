package nylas

// Response model structs
type ThreadResponse struct {
	RequestID  string `json:"request_id"`
	Data       Thread `json:"data"`
	NextCursor string `json:"next_cursor"`
}

type Thread struct {
	GrantID                   string        `json:"grant_id"`
	ID                        string        `json:"id"`
	Object                    string        `json:"object"`
	HasAttachments            bool          `json:"has_attachments"`
	HasDrafts                 bool          `json:"has_drafts"`
	EarliestMessageDate       int64         `json:"earliest_message_date"`
	LatestMessageReceivedDate int64         `json:"latest_message_received_date"`
	LatestMessageSentDate     int64         `json:"latest_message_sent_date"`
	Participants              []Participant `json:"participants"`
	Snippet                   string        `json:"snippet"`
	Starred                   bool          `json:"starred"`
	Subject                   string        `json:"subject"`
	Unread                    bool          `json:"unread"`
	MessageIDs                []string      `json:"message_ids"`
	DraftIDs                  []string      `json:"draft_ids"`
	Folders                   []string      `json:"folders"`
	LatestDraftOrMessage      Message       `json:"latest_draft_or_message"`
}

type Participant struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Message struct {
	Body        string        `json:"body"`
	CC          []Participant `json:"cc"`
	Date        int64         `json:"date"`
	Attachments []Attachment  `json:"attachments"`
	Folders     []string      `json:"folders"`
	From        []Participant `json:"from"`
	GrantID     string        `json:"grant_id"`
	ID          string        `json:"id"`
	Object      string        `json:"object"`
	ReplyTo     []Participant `json:"reply_to"`
	Snippet     string        `json:"snippet"`
	Starred     bool          `json:"starred"`
	Subject     string        `json:"subject"`
	ThreadID    string        `json:"thread_id"`
	To          []Participant `json:"to"`
	Unread      bool          `json:"unread"`
}

type Attachment struct {
	Content     string `json:"content,omitempty"`
	ContentType string `json:"content_type"`
	ID          string `json:"id"`
	Size        int    `json:"size"`
	Filename    string `json:"filename,omitempty"`
}
