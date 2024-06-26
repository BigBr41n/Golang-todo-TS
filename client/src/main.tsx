import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App.tsx";
import "./index.css";
import { MantineProvider } from "@mantine/core";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <MantineProvider theme={{}} defaultColorScheme="light">
    <React.StrictMode>
      <App />
    </React.StrictMode>
  </MantineProvider>
);
