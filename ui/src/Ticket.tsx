import { useParams } from "react-router-dom"
import { useGetThreadsThreadId } from "./clients/messages/messages"
import { useReadTicket } from "./clients/tickets/tickets"

export const Ticket = () => {
    const {id} =  useParams()

    // ticket information

    const {data: ticket} = useReadTicket(Number(id))

    const threadResp = useGetThreadsThreadId(ticket?.data?.threadID!, {
        query: {
            enabled: !!ticket?.data?.threadID!!,
        }  
    })

    // useGetThreadsThreadId()


    return <div>
        <div>{`Ticket ${ticket?.data?.title}`}</div>
        <div>{`Ticket ${ticket?.data?.id}`}</div>
        <div>{`Ticket ${ticket?.data?.assignee}`}</div>
        <div>{`Ticket ${ticket?.data?.opened_by}`}</div>
        <div>{`Ticket ${ticket?.data?.priority}`}</div>
        <div>{`ThreadID ${ticket?.data?.threadID}`}</div>
        <h1>Thread Info</h1>
        <div>{threadResp?.data?.data?.data?.message_ids}</div>
    </div>
}