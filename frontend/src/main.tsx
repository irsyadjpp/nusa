import "@/i18n/i18n";
import "@/style/global.css";
import "@fontsource/mulish/latin.css";
import "@fontsource/urbanist/latin.css";

import { StrictMode } from "react";
import { createRoot } from "react-dom/client";

import App from "@/App";
import { AuthProvider } from "@/features/auth";
import { QueryClientProvider } from "@/shared/query-client";

// Handle Vite preload errors - refresh page on dynamic import failures
window.addEventListener('vite:preloadError', () => {
  console.warn('Vite preload error detected, refreshing page...');
  window.location.reload();
});

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <QueryClientProvider>
      <AuthProvider>
        <App />
      </AuthProvider>
    </QueryClientProvider>
  </StrictMode>,
);
