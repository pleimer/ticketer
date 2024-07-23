import { Box, Card, CardContent, Divider, IconButton, TextField, Typography } from "@mui/material"
import { Message } from "./clients/messages/models"
import {useListThreadMessagesSuspense, useReplyToThread } from "./clients/messages/messages"
import SendIcon from '@mui/icons-material/Send';
import { AxiosResponse } from "axios";
import { useState } from "react";

type MessagesProps = {
    threadID: string
}

export const Messages = ({threadID}: MessagesProps) => {
    // const {data: {data: messages = [] as Message[]} = {}, error} = useListThreadMessagesSuspense(threadID || "",)
    const {data: {data: messages = []} = {}, error} = useListThreadMessagesSuspense(threadID || "")


    // bit of a hack, but this ensures a new message is seen on screen right after submitting it
    // Typically, the approach would be to update the react query cache or invalidate it,
    // but there was an issue attempting this. This also causes some extra re-renders so not ideal
    const [messagesSynced, setMessagesSynced] = useState<Array<Message>>(messages)

    const [newComment, setNewComment] = useState('');

    const {mutate: replyToThread} = useReplyToThread()

    // const queryClient = useQueryClient();

    const [sending, setSending] = useState(false)

    const handleCommentSubmit = () => {
      setSending(true)

      replyToThread({
        data: {
          thread_id: threadID,
          body: newComment,
        }
      },{
        onSuccess: (r: AxiosResponse<Message, any>) => {
          r.data.date = Date.now()
          r.data.snippet = newComment

          setMessagesSynced([...messagesSynced, r.data])
          setSending(false)
          setNewComment('');

          // const qKey = getListThreadMessagesQueryKey(threadID)
          // queryClient.invalidateQueries(qKey)

          // // directly update the cache so we don't hit the mail server too often
          // queryClient.setQueryData(qKey, (prevMessages: AxiosResponse<Message[], any> | undefined) => {
          //   console.log(prevMessages?.data)
          //   return prevMessages?.data ? [...prevMessages.data, r.data] : [r.data]
          // })
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
              {messagesSynced.sort((a, b) => (a.date || 0) - (b.date || 0)).map((message) => (
                <Card key={message.id || 0} sx={{ mb: 2 }}>
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
                placeholder="Send a message..."
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