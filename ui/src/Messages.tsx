import { Box, Card, CardContent, Divider, IconButton, TextField, Typography } from "@mui/material"
import { Message } from "./clients/messages/models"
import { useListThreadMessagesSuspense, useReplyToThread } from "./clients/messages/messages"
import SendIcon from '@mui/icons-material/Send';
import { useState } from "react";
import { ControlPointSharp } from "@mui/icons-material";

type MessagesProps = {
    threadID?: string
}

export const Messages = ({threadID}: MessagesProps) => {
    const {data: {data: messages = [] as Message[]} = {}, error, refetch} = useListThreadMessagesSuspense(threadID || "",)

    const [newComment, setNewComment] = useState('');

    const {mutate: replyToThread} = useReplyToThread()

    const handleCommentSubmit = () => {

      replyToThread({
        data: {
          thread_id: threadID,
          body: newComment,
        }
      },{
        onSuccess: (r) => {
          setNewComment('');
          console.log(r.data)
          
          // refetch()
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
                fullWidth
                multiline
                rows={3}
                variant="outlined"
                placeholder="Write a comment..."
                value={newComment}
                onChange={(e) => setNewComment(e.target.value)}
                sx={{ mr: 1 }}
              />
              <IconButton color="primary" onClick={handleCommentSubmit}>
                <SendIcon />
              </IconButton>
            </Box>
            </>
    )
}