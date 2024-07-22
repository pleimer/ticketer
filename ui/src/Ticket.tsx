import { useParams } from "react-router-dom"
import ArrowDropDownIcon from '@mui/icons-material/ArrowDropDown';
import { useReadTicketSuspense, useUpdateTicket } from "./clients/tickets/tickets"
import SendIcon from '@mui/icons-material/Send';
import { TicketStatus, Ticket as TicketModel, TicketPriority} from "./clients/tickets/models"
import { Box, Card, CardContent, Chip, CircularProgress, Divider, Grid, IconButton, Menu, MenuItem, Popover, TextField, Typography } from "@mui/material"
import { ChangeEvent, Suspense, useState } from "react"
// import EdiText from 'react-editext'
import { Messages } from "./Messages";

export const Ticket = () => {
    const {id} =  useParams()

    // ticket information

    return (
          <Suspense fallback={<CircularProgress color="primary" />}>
            <TicketContent id={Number(id)}/>
          </Suspense>
    );
}

const TicketContent = ({id}: {id: number}) => {
  const {data: {data: ticket = {} as TicketModel} = {}, error, refetch} =  useReadTicketSuspense(Number(id))

  const {mutate: updateTicket} = useUpdateTicket()

  const [anchorStatusEl, setAnchorStatusEl] = useState<null | HTMLElement>(null);
  const statusOpen = Boolean(anchorStatusEl);

  const [anchorPriorityEl, setAnchorPriorityEl] = useState<null | HTMLElement>(null);
  const prioritoryOpen = Boolean(anchorPriorityEl);

  const handleClose = () => {
    setAnchorStatusEl(null);
  };
  
  const handleStatusChange = (newStatus: TicketStatus) => {
    updateTicket({id: id, data: {
      status: newStatus,
    }},{
      onSuccess: () => refetch(),
    })

    handleClose();
  };

  const handlePriorityChange = (newStatus: TicketPriority) => {
    updateTicket({id: id, data: {
      priority: newStatus,
    }},{
      onSuccess: () => refetch(),
    })

    handleClose();
  };

  return (
    error ? "error fetching ticket data" :
    // TODO: error boundaries
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
                onDelete={(event: React.MouseEvent<HTMLButtonElement>) => setAnchorStatusEl(event.currentTarget)}
                deleteIcon={<ArrowDropDownIcon />}
                sx={{ mb: 1 }}
              />
              <AttributeOptions 
                anchor={anchorStatusEl} 
                open={statusOpen} 
                options={TicketStatus} 
                onSelect={handleStatusChange} 
                onClose={() => setAnchorStatusEl(null)}
              />

              <Chip 
                label={`Priority: ${ticket.priority}`} 
                color={ticket.priority == TicketPriority.high ? "warning": "secondary"} 
                onDelete={(event: React.MouseEvent<HTMLButtonElement>) => setAnchorPriorityEl(event.currentTarget)}
                deleteIcon={<ArrowDropDownIcon />}
                sx={{ mb: 1}}
              />

              <AttributeOptions 
                anchor={anchorPriorityEl} 
                open={prioritoryOpen} 
                options={TicketPriority} 
                onSelect={handlePriorityChange} 
                onClose={() => setAnchorPriorityEl(null)}
              />

              <Assignee ticketId={ticket.id} assignee={ticket.assignee} refetchFn={refetch}/>
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
    
  </Box>
  )
}

type AssigneeProps = {
  ticketId: number
  refetchFn: () => void
  assignee?: string
}

const Assignee = ({assignee, ticketId, refetchFn}: AssigneeProps) => {

  const [anchor, setAnchor] = useState<null | HTMLElement>(null);
  const open = Boolean(anchor);

  const [value, setValue] = useState(assignee)

  const {mutate: updateTicket} = useUpdateTicket()

  const handleClose = () => {
    setAnchor(null)
    setValue("")
  }

  return (
    <div>
      <Chip
        label={`Assignee: ${assignee || "None"}`}
        onClick={(e) => setAnchor(e.currentTarget)}
      />
      <Popover 
        id="simpole-popover"
        open={open} 
        onClose={handleClose}
        anchorOrigin={{
          vertical: 'center',
          horizontal: 'right',
        }}

      >
        <TextField 
          id="outlined-basic" 
          label="Type name" 
          variant="outlined" 
          onChange={(event: ChangeEvent<HTMLInputElement>) => {
            setValue(event.target.value);
          }}
        />

        <IconButton color="secondary" onClick={() => {
          updateTicket({
            id: ticketId,
            data: {
              assignee: value,
            }
          },{
            onSuccess: () => {
              handleClose()
              refetchFn()
            }
          })
        }}>
          <SendIcon />
        </IconButton>

      </Popover>
    </div>
  )
}

type AttributeOptionsProps<T extends { [key: string]: string }> = {
  anchor: null | HTMLElement,
  open: boolean,
  options: T,
  onClose?: () => void,
  onSelect?: (option: T[keyof T]) => void,
}

const AttributeOptions = <T extends { [key: string]: string }>({
  anchor, 
  open, 
  options, 
  onClose, 
  onSelect
}: AttributeOptionsProps<T>) => {
  return (
    <Menu
      anchorEl={anchor}
      open={open}
      onClose={onClose}
    >
      {(Object.values(options) as T[keyof T][]).map((option) => (
        <MenuItem key={option} onClick={() => onSelect?.(option)}>
          {option}
        </MenuItem>
      ))}
    </Menu>
  )
}