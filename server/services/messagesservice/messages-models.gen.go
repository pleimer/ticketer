// Package messagesservice provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package messagesservice

// Attachment defines model for Attachment.
type Attachment struct {
	Content     *string `json:"content,omitempty"`
	ContentType *string `json:"content_type,omitempty"`
	Filename    *string `json:"filename,omitempty"`
	Id          *string `json:"id,omitempty"`
	Size        *int    `json:"size,omitempty"`
}

// Message defines model for Message.
type Message struct {
	Attachments *[]Attachment  `json:"attachments,omitempty"`
	Body        *string        `json:"body,omitempty"`
	Cc          *[]Participant `json:"cc,omitempty"`
	Date        *int           `json:"date,omitempty"`
	Folders     *[]string      `json:"folders,omitempty"`
	From        *[]Participant `json:"from,omitempty"`
	Id          *string        `json:"id,omitempty"`
	Object      *string        `json:"object,omitempty"`
	ReplyTo     *[]Participant `json:"reply_to,omitempty"`
	Snippet     *string        `json:"snippet,omitempty"`
	Starred     *bool          `json:"starred,omitempty"`
	Subject     *string        `json:"subject,omitempty"`
	ThreadId    *string        `json:"thread_id,omitempty"`
	To          *[]Participant `json:"to,omitempty"`
	Unread      *bool          `json:"unread,omitempty"`
}

// Participant defines model for Participant.
type Participant struct {
	Email *string `json:"email,omitempty"`
	Name  *string `json:"name,omitempty"`
}

// Thread defines model for Thread.
type Thread struct {
	DraftIds                  *[]string      `json:"draft_ids,omitempty"`
	EarliestMessageDate       *int           `json:"earliest_message_date,omitempty"`
	Folders                   *[]string      `json:"folders,omitempty"`
	HasAttachments            *bool          `json:"has_attachments,omitempty"`
	HasDrafts                 *bool          `json:"has_drafts,omitempty"`
	Id                        *string        `json:"id,omitempty"`
	LatestDraftOrMessage      *Message       `json:"latest_draft_or_message,omitempty"`
	LatestMessageReceivedDate *int           `json:"latest_message_received_date,omitempty"`
	LatestMessageSentDate     *int           `json:"latest_message_sent_date,omitempty"`
	MessageIds                *[]string      `json:"message_ids,omitempty"`
	Object                    *string        `json:"object,omitempty"`
	Participants              *[]Participant `json:"participants,omitempty"`
	Snippet                   *string        `json:"snippet,omitempty"`
	Starred                   *bool          `json:"starred,omitempty"`
	Subject                   *string        `json:"subject,omitempty"`
	Unread                    *bool          `json:"unread,omitempty"`
}

// ThreadResponse defines model for ThreadResponse.
type ThreadResponse struct {
	Data       *Thread `json:"data,omitempty"`
	NextCursor *string `json:"next_cursor,omitempty"`
	RequestId  *string `json:"request_id,omitempty"`
}