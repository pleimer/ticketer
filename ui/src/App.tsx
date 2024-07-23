import { useNavigate} from "react-router-dom";
import TicketList from "./TicketList";
import { useListTicket} from "./clients/tickets/tickets";
import { Status } from "./clients/tickets/models/status";



export default function App() {

  // TODO: api should be adjusted to allow batch querying
  const {data: tickets} = useListTicket({
    'status[]': [Status.not_started, Status.in_progress],
  })

  const nav = useNavigate()

  return (
    <TicketList tickets={tickets?.data || []} onCardClick={(t) => nav(`/ticket/${t.id}`)} />
  );
}
