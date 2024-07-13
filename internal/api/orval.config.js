// TODO: move to root
module.exports = {

    tickets: {
      input: './tickets.yaml',
      output: {
        target: '../../ui/src/clients/tickets.ts',
        schemas: '../../ui/src/models',
        client: 'react-query',
      }
    },
 
  };