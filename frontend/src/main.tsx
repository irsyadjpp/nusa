import "@/i18n/i18n";
import "@/style/global.css";
import "@fontsource/mulish/latin.css";
import "@fontsource/urbanist/latin.css";

import { StrictMode } from "react";
import { createRoot } from "react-dom/client";

import App from "@/App";
import { AuthProvider } from "@/features/auth";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <AuthProvider>
      <App />
    </AuthProvider>
  </StrictMode>,
);
