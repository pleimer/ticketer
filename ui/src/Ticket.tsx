import { useParams } from "react-router-dom"
import { useGetThreadsThreadId } from "./clients/messages/messages"
import { useReadTicket } from "./clients/tickets/tickets"

export const Ticket = () => {
    const {id} =  useParams()

    // ticket information

    const {data: ticket} = useReadTicket(Number(id))

    const threadResp = useGetThreadsThreadId(ticket?.data?.thread_id!, {
        query: {
            enabled: !!ticket?.data?.thread_id!!,
        }  
    })

    // useGetThreadsThreadId()

    console.log(threadResp.data)

    return <div>
        <div>{`Ticket ${ticket?.data?.title}`}</div>
        <div>{`Ticket ${ticket?.data?.id}`}</div>
        <div>{`Ticket ${ticket?.data?.assignee}`}</div>
        <div>{`Ticket ${ticket?.data?.opened_by}`}</div>
        <div>{`Ticket ${ticket?.data?.priority}`}</div>
        <div>{`ThreadID ${ticket?.data?.thread_id}`}</div>

        <h1>Thread Info</h1>
        {/* <div>{threadResp?.data?.data?.data}</div> */}
    </div>
}