import Container from "@mui/material/Container";
import Typography from "@mui/material/Typography";
import Box from "@mui/material/Box";
import { useListTicket } from "./clients/tickets/tickets";
import { ListTicketStatus } from "./clients/tickets/models/listTicketStatus";
import { Button, ButtonGroup } from "@mui/material";
import { useNavigate, useParams } from "react-router-dom";
import TicketList from "./TicketList";



export default function App() {

  // TODO: api should be adjusted to allow batch querying
  const {data: tickets} = useListTicket({
    'status[]': [ListTicketStatus.not_started, ListTicketStatus.in_progress],
  })

  const nav = useNavigate()

  return (
    <TicketList tickets={tickets?.data || []} onCardClick={(t) => nav(`/ticket/${t.id}`)} />
  );
}
