import { Box, Card, CardContent, Divider, IconButton, TextField, Typography } from "@mui/material"
import { Message } from "./clients/messages/models"
import { getListThreadMessagesQueryKey, getListThreadMessagesSuspenseQueryOptions, useListThreadMessagesSuspense, useReplyToThread } from "./clients/messages/messages"
import SendIcon from '@mui/icons-material/Send';
import { useState } from "react";
import { useQueryClient } from "@tanstack/react-query";
import { AxiosResponse } from "axios";

type MessagesProps = {
    threadID: string
}

export const Messages = ({threadID}: MessagesProps) => {
    const {data: {data: messages = [] as Message[]} = {}, error, refetch} = useListThreadMessagesSuspense(threadID || "",)

    const [newComment, setNewComment] = useState('');

    const {mutate: replyToThread} = useReplyToThread()

    const queryClient = useQueryClient();

    const [sending, setSending] = useState(false)

    const handleCommentSubmit = () => {
      setSending(true)

      replyToThread({
        data: {
          thread_id: threadID,
          body: newComment,
        }
      },{
        onSuccess: (r: any) => {
          setNewComment('');
          // const qKey = getListThreadMessagesQueryKey(threadID)

          // // directly update the cache so we don't hit the mail server too often
          // queryClient.setQueryData(qKey, (prevMessages: AxiosResponse<Message[], any> | undefined) => {
          //   return prevMessages?.data ? [...prevMessages.data, r.data.data] : [r.data.data]
          // })
          refetch()
          setSending(false)
        },
        onError: (e) => {
          setNewComment(`There was an error sending the message: ${e.message}`)
          setSending(false)
        }
      })

    };
    return (
        error ? 
            // TODO: error boundaries
            "Error fetching messages" : 
            <>
            <Box sx={{ mb: 2 }}>
              {messages.sort((a, b) => (a.date || 0) - (b.date || 0)).map((message) => (
                <Card key={message.id} sx={{ mb: 2 }}>
                  <CardContent>
                    <Typography variant="subtitle1">
                      {message.from?.[0]?.name || 'Unknown'}
                    </Typography>
                    <Typography variant="caption" color="text.secondary">
                      {message.date ? new Date(message.date).toLocaleString() : 'No date'}
                    </Typography>
                    <Divider sx={{ my: 1 }} />
                    <Typography 
                      variant="body1"
                      sx={{ whiteSpace: 'pre-wrap' }}
                    >
                      {message.snippet || 'No content'}
                    </Typography>
                  </CardContent>
                </Card>
              ))}
            </Box>
            <Box sx={{ display: 'flex', alignItems: 'flex-start' }}>
              <TextField
                disabled={sending}
                fullWidth
                multiline
                rows={3}
                variant="outlined"
                placeholder="Write a comment..."
                value={newComment}
                onChange={(e) => setNewComment(e.target.value)}
                sx={{ mr: 1 }}
              />
              <IconButton disabled={sending} color="primary" onClick={handleCommentSubmit}>
                <SendIcon />
              </IconButton>
            </Box>
            </>
    )
}