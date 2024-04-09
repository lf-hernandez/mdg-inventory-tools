import React from "react";
import ReactDOM from "react-dom/client";
import * as amplitude from "@amplitude/analytics-browser";

import App from "./App";
import "./index.css";

amplitude.init(import.meta.env.VITE_AMPLITUDE_API_KEY);

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
);
