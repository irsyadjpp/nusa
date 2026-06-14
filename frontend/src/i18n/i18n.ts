import i18n from "i18next";
import { initReactI18next } from "react-i18next";

import { getClientLocale } from "@/i18n/locale";
import id from "@/i18n/messages/id.json";

i18n.use(initReactI18next).init({
  resources: {
    id: { translation: id },
  },
  lng: getClientLocale(),
  fallbackLng: "id",
  interpolation: { escapeValue: false },
});

export default i18n;
