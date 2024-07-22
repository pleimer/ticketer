import { defineConfig} from 'orval';

// TODO: move to root
export default defineConfig({
    tickets: {
      input: '../internal/api/tickets.json',
      output: {
        target: './src/clients/tickets/tickets.ts',
        schemas: './src/clients/tickets/models',
        baseUrl: '/api/v1/tickets',
        client: 'react-query',
        override: {
          query: {
            useQuery: true,
            useSuspenseQuery: true,
          },
          components: {
            responses: {
              suffix: ""
            }
          }
        },
      },
    },
    messages: {
      input: '../internal/api/messages.json',
      output: {
        target: './src/clients/messages/messages.ts',
        schemas: './src/clients/messages/models',
        baseUrl: '/api/v1/messages',
        client: 'react-query',
        override: {
          query: {
            useQuery: true,
            useSuspenseQuery: true,
          },
          components: {
            responses: {
              suffix: ""
            }
          }
        },
      }
    },
  })
