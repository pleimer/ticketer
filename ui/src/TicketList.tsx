
import React from 'react';
import {
  Box,
  Card,
  CardContent,
  Typography,
  Chip,
  List,
  ListItem,
  ListItemText,
} from '@mui/material';
import { Ticket, TicketPriority } from './clients/tickets/models';

interface TicketListProps {
  tickets: Ticket[];
}

const TicketList: React.FC<TicketListProps> = ({ tickets }) => {
  // Sort tickets by updated_at in descending order
  const sortedTickets = [...tickets].sort((a, b) => 
    new Date(b.updated_at).getTime() - new Date(a.updated_at).getTime()
  );

  return (
    <Box sx={{ maxWidth: 800, margin: 'auto', padding: 2 }}>
      <Typography variant="h4" gutterBottom>
        Tickets
      </Typography>
      <List>
        {sortedTickets.map((ticket) => (
          <ListItem key={ticket.id} disablePadding>
            <Card sx={{ width: '100%', mb: 2 }}>
              <CardContent>
                <Typography variant="h6" gutterBottom>
                  {ticket.title}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  #{ticket.id}
                </Typography>
                <Box sx={{ display: 'flex', justifyContent: 'space-between', mt: 2 }}>
                  <Chip label={`Status: ${ticket.status}`} color="primary" size="small" />
                  <Chip label={`Priority: ${ticket.priority}`} color={ticket.priority == TicketPriority.high ? "warning": "secondary"} size="small" />
                </Box>
                <List dense>
                  <ListItemText 
                    primary="Created by"
                    secondary={`${ticket.created_by} on ${new Date(ticket.created_at).toLocaleString()}`}
                  />
                  <ListItemText 
                    primary="Updated by"
                    secondary={`${ticket.updated_by} on ${new Date(ticket.updated_at).toLocaleString()}`}
                  />
                  {ticket.assignee && (
                    <ListItemText 
                      primary="Assignee"
                      secondary={ticket.assignee}
                    />
                  )}
                </List>
              </CardContent>
            </Card>
          </ListItem>
        ))}
      </List>
    </Box>
  );
};

export default TicketList;