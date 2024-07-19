import Container from "@mui/material/Container";
import Typography from "@mui/material/Typography";
import Box from "@mui/material/Box";
import { useListTicket } from "./clients/tickets/tickets";
import { ListTicketStatus } from "./clients/tickets/models/listTicketStatus";
import { Button, ButtonGroup } from "@mui/material";

export default function App() {

  // TODO: api should be adjusted to allow batch querying
  const {data: notStartedTickets} = useListTicket({
    status: ListTicketStatus.not_started
  })

  const {data: inProgressTickets} = useListTicket({
    status: ListTicketStatus.in_progress,
  },{
  })

  return (
    <Container maxWidth="sm">
      {notStartedTickets?.data?.map((ticket) => {
        return <Box key={ticket.id} sx={{ my: 4 }}>
          <ButtonGroup>
            <Button onClick={() => console.log("clicked")}>
              <p>{`# ${ticket.id}`}</p>
              <Typography variant="h4" component="h1" sx={{ mb: 2 }}>
                {ticket.title}
              </Typography>
              <h5>{`${ticket.status}`}</h5>
            </Button>
          </ButtonGroup>
          </Box>
      })}
    </Container>
  );
}
