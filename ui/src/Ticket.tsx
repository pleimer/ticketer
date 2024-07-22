import { useParams } from "react-router-dom"
import ArrowDropDownIcon from '@mui/icons-material/ArrowDropDown';
import { useReadTicket, useReadTicketSuspense } from "./clients/tickets/tickets"
import { TicketStatus, Ticket as TicketModel, TicketPriority} from "./clients/tickets/models"
import { Message as MessageModel} from "./clients/messages/models"
import SendIcon from '@mui/icons-material/Send';
import { Box, Card, CardContent, Chip, CircularProgress, Divider, Grid, IconButton, Menu, MenuItem, TextField, Typography } from "@mui/material"
import { Suspense, useState } from "react"
import { useListThreadMessages } from "./clients/messages/messages";
import { Messages } from "./Messages";

export const Ticket = () => {
    const {id} =  useParams()

    // ticket information

    return (
        <Box sx={{ maxWidth: 1000, margin: 'auto', padding: 2 }}>
          <Suspense fallback={<CircularProgress color="primary" />}>
            <TicketContent id={Number(id)}/>
          </Suspense>
        </Box>
    );
}

const TicketContent = ({id}: {id: number}) => {
  const {data: {data: ticket = {} as TicketModel} = {}, error} =  useReadTicketSuspense(Number(id))

  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
  const open = Boolean(anchorEl);

  const handleClick = (event: React.MouseEvent<HTMLButtonElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };
  
  const handleStatusChange = (newStatus: TicketStatus) => {
    handleClose();
  };

  const [newComment, setNewComment] = useState('');

  const handleCommentSubmit = () => {
    setNewComment('');
  };

  return (
    error ? "error fetching ticket data" :
    // TODO: error boundaries
    <>
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
              <Chip 
                label={`Priority: ${ticket.priority}`} 
                color={ticket.priority == TicketPriority.high ? "warning": "secondary"} 
                onDelete={handleClick}
                deleteIcon={<ArrowDropDownIcon />}
                sx={{ mb: 1}}
              />
              <Typography variant="body2" align="right">
                Assignee: {ticket.assignee || "None"}
              </Typography>
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
      Messages
    </Typography>

    <Suspense fallback={<CircularProgress color="secondary" />} >
      <Messages threadID={ticket.thread_id}/>
    </Suspense>
    
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

