import * as React from "react";
import * as ReactDOM from "react-dom/client";
import { ThemeProvider } from "@emotion/react";
import { CssBaseline } from "@mui/material";
import theme from "./theme";
import App from "./App";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import axios from 'axios';

import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import { Ticket } from "./Ticket";

// Set default base URL
axios.defaults.baseURL = 'http://localhost:8080';

const queryClient = new QueryClient();

const router = createBrowserRouter([
  {
    path: "/",
    element: (<App />),
  },
  {
    path: "ticket/:id",
    element: <Ticket />
  },
]);

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <QueryClientProvider client={queryClient}>
        <RouterProvider router={router} />
      </QueryClientProvider>
    </ThemeProvider>
  </React.StrictMode>
);
