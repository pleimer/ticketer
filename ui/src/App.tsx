import Container from "@mui/material/Container";
import Typography from "@mui/material/Typography";
import Box from "@mui/material/Box";
import Link from "@mui/material/Link";
import ProTip from "./ProTip";
import { useListTicket } from "./clients/tickets/tickets";
import { ListTicketStatus } from "./clients/tickets/models/listTicketStatus";

function Copyright() {
  return (
    <Typography variant="body2" color="text.secondary" align="center">
      {"Copyright © "}
      <Link color="inherit" href="https://mui.com/">
        Your Website
      </Link>{" "}
      {new Date().getFullYear()}.
    </Typography>
  );
  }

export default function App() {

  // TODO: api should be adjusted to allow batch querying
  const {data: notStartedTickets} = useListTicket({
    status: ListTicketStatus.not_started
  })

  const {data: inProgressTickets} = useListTicket({
    status: ListTicketStatus.in_progress,
  },{
  })

  console.log(notStartedTickets?.data)
  // console.log(inProgressTickets)
  
  return (
    <Container maxWidth="sm">
      {notStartedTickets?.data?.map((ticket) => {
        return <Box sx={{ my: 4 }}>
            <p>{ticket.id}</p>
            <Typography variant="h4" component="h1" sx={{ mb: 2 }}>
              {ticket.title}
            </Typography>
          </Box>
      })}
    </Container>
  );
}
