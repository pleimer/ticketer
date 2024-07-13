/**
 * Generated by orval v6.31.0 🍺
 * Do not edit manually.
 * Ticket Management API
 * API for managing support tickets
 * OpenAPI spec version: 1.0.0
 */
import axios from 'axios'
import type {
  AxiosRequestConfig,
  AxiosResponse
} from 'axios'
export interface Comment {
  content?: string;
  created_at?: string;
  id?: number;
  ticket_id?: number;
}

export type TicketStatus = typeof TicketStatus[keyof typeof TicketStatus];


// eslint-disable-next-line @typescript-eslint/no-redeclare
export const TicketStatus = {
  open: 'open',
  in_progress: 'in_progress',
  resolved: 'resolved',
  closed: 'closed',
} as const;

export type TicketPriority = typeof TicketPriority[keyof typeof TicketPriority];


// eslint-disable-next-line @typescript-eslint/no-redeclare
export const TicketPriority = {
  low: 'low',
  medium: 'medium',
  high: 'high',
} as const;

export interface Ticket {
  created_at?: string;
  description?: string;
  id?: number;
  priority?: TicketPriority;
  status?: TicketStatus;
  title?: string;
  updated_at?: string;
}

export interface CommentCreate {
  content: string;
}

export type TicketStatusUpdateStatus = typeof TicketStatusUpdateStatus[keyof typeof TicketStatusUpdateStatus];


// eslint-disable-next-line @typescript-eslint/no-redeclare
export const TicketStatusUpdateStatus = {
  open: 'open',
  in_progress: 'in_progress',
  resolved: 'resolved',
  closed: 'closed',
} as const;

export interface TicketStatusUpdate {
  status: TicketStatusUpdateStatus;
}

export type TicketCreatePriority = typeof TicketCreatePriority[keyof typeof TicketCreatePriority];


// eslint-disable-next-line @typescript-eslint/no-redeclare
export const TicketCreatePriority = {
  low: 'low',
  medium: 'medium',
  high: 'high',
} as const;

export interface TicketCreate {
  description: string;
  priority?: TicketCreatePriority;
  title: string;
}





  /**
 * @summary Create a new ticket
 */
export const postTickets = <TData = AxiosResponse<Ticket>>(
    ticketCreate: TicketCreate, options?: AxiosRequestConfig
 ): Promise<TData> => {
    return axios.post(
      `/tickets`,
      ticketCreate,options
    );
  }

/**
 * @summary Update the status of a ticket
 */
export const putTicketsTicketIdStatus = <TData = AxiosResponse<Ticket>>(
    ticketId: number,
    ticketStatusUpdate: TicketStatusUpdate, options?: AxiosRequestConfig
 ): Promise<TData> => {
    return axios.put(
      `/tickets/${ticketId}/status`,
      ticketStatusUpdate,options
    );
  }

/**
 * @summary Add a comment to a ticket
 */
export const postTicketsTicketIdComments = <TData = AxiosResponse<Comment>>(
    ticketId: number,
    commentCreate: CommentCreate, options?: AxiosRequestConfig
 ): Promise<TData> => {
    return axios.post(
      `/tickets/${ticketId}/comments`,
      commentCreate,options
    );
  }

export type PostTicketsResult = AxiosResponse<Ticket>
export type PutTicketsTicketIdStatusResult = AxiosResponse<Ticket>
export type PostTicketsTicketIdCommentsResult = AxiosResponse<Comment>
