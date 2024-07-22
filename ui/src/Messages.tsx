import { Box, Card, CardContent, Divider, Typography } from "@mui/material"
import { Message } from "./clients/messages/models"
import { useListThreadMessagesSuspense } from "./clients/messages/messages"

type MessagesProps = {
    threadID?: string
}

export const Messages = ({threadID}: MessagesProps) => {
    const {data: {data: messages = [] as Message[]} = {}, error} = useListThreadMessagesSuspense(threadID || "",)
    return (
        error ? 
            // TODO: error boundaries
            "Error fetching messages" : 
            <Box sx={{ mb: 2 }}>
              {messages.sort((a, b) => (a.date || 0) - (b.date || 0)).map((message) => (
                <Card key={message.id} sx={{ mb: 2 }}>
                  <CardContent>
                    <Typography variant="subtitle1">
                      From: {message.from?.[0]?.name || 'Unknown'}
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
    )
}