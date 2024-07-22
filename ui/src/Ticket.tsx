import { useParams } from "react-router-dom"
import { useGetThreadsThreadId } from "./clients/messages/messages"
import ArrowDropDownIcon from '@mui/icons-material/ArrowDropDown';
import { useReadTicket } from "./clients/tickets/tickets"
import { TicketStatus, Ticket as TicketModel, TicketPriority} from "./clients/tickets/models"
import { Message as MessageModel} from "./clients/messages/models"
import SendIcon from '@mui/icons-material/Send';
import { Box, Button, Card, CardContent, CardHeader, Chip, Divider, Grid, IconButton, List, ListItem, ListItemText, Menu, MenuItem, TextField, Typography } from "@mui/material"
import { useState } from "react"

export const Ticket = () => {
    const {id} =  useParams()

    // ticket information

    const {data: {data: ticket = {} as TicketModel} = {}} = useReadTicket(Number(id))

    // const threadResp = useGetThreadsThreadId(ticket?.data?.thread_id!, {
    //     query: {
    //         enabled: !!ticket?.data?.thread_id!!,
    //     }  
    // })


    // useGetThreadsThreadId()


    // const ticket = useTicket(id); // Custom hook to fetch ticket details
    // const comments = useComments(ticket?.thread_id) || [];
    const comments = [] as MessageModel[]
  
    const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
    const open = Boolean(anchorEl);

    const [newComment, setNewComment] = useState('');
  
    const handleClick = (event: React.MouseEvent<HTMLButtonElement>) => {
      setAnchorEl(event.currentTarget);
    };

    const handleCommentSubmit = () => {
        console.log('Submitting comment:', newComment);
        setNewComment('');
      };
  
    const handleClose = () => {
      setAnchorEl(null);
    };
  
    const handleStatusChange = (newStatus: TicketStatus) => {
      // Implement status change logic here
      console.log(`Changing status to: ${newStatus}`);
      handleClose();
    };
  
    if (!ticket) {
        // TODO: suspense
        return <Typography>Loading...</Typography>;
    }

    return (
        <Box sx={{ maxWidth: 1000, margin: 'auto', padding: 2 }}>
        <Card sx={{ mb: 4 }}>
          <CardContent>
            <Typography variant="h4" gutterBottom>
              {ticket.title}
            </Typography>
            <Typography variant="body2" color="text.secondary" gutterBottom>
              # {ticket.id}
            </Typography>
            <Grid container spacing={2}>
              <Grid item xs={12} sm={8}>
                <Typography variant="body1" paragraph>
                  {/* {ticket.description} */}
                  TODO: add ticket descriptions
                </Typography>
              </Grid>
              <Grid item xs={12} sm={4}>
                <Box sx={{ display: 'flex', flexDirection: 'column', alignItems: 'flex-end' }}>
                  <Chip
                    label={`Status: ${ticket.status}`}
                    color="primary"
                    // onClick={handleClick}
                    onDelete={handleClick}
                    deleteIcon={<ArrowDropDownIcon />}
                    sx={{ mb: 1 }}
                  />
                  <Menu
                    anchorEl={anchorEl}
                    open={open}
                    onClose={handleClose}
                  >
                    {Object.values(TicketStatus).map((status) => (
                      <MenuItem key={status} onClick={() => handleStatusChange(status)}>
                        {status}
                      </MenuItem>
                    ))}
                  </Menu>
                  <Chip label={`Priority: ${ticket.priority}`} color={ticket.priority == TicketPriority.high ? "warning": "secondary"} sx={{ mb: 1}} />
                  <Typography variant="body2" align="right">
                    Created by: {ticket.created_by?.split("@")[0]}
                  </Typography>
                </Box>
              </Grid>
            </Grid>
            <Divider sx={{ my: 2 }} />
            <Box sx={{ display: 'flex', justifyContent: 'space-between' }}>
              <Typography variant="body2" color="text.secondary">
                Last updated: {new Date(ticket.updated_at).toLocaleString()}
              </Typography>
              <Typography variant="body2" color="text.secondary">
                Updated by: {ticket.updated_by?.split("@")[0]}
              </Typography>
            </Box>
          </CardContent>
        </Card>
  
        <Typography variant="h5" gutterBottom>
          Comments
        </Typography>
        <Box sx={{ mb: 2 }}>
          {comments.sort((a, b) => (a.date || 0) - (b.date || 0)).map((message) => (
            <Card key={message.id} sx={{ mb: 2 }}>
              <CardContent>
                <Typography variant="subtitle1">
                  From: {message.from?.[0]?.name || 'Unknown'}
                </Typography>
                <Typography variant="caption" color="text.secondary">
                  {message.date ? new Date(message.date).toLocaleString() : 'No date'}
                </Typography>
                <Divider sx={{ my: 1 }} />
                <Typography variant="body1">
                  {message.body || 'No content'}
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
      </Box>



    );

}

