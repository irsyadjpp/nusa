import { useTranslation } from "react-i18next";

import { ListItemIcon, MenuItem } from "@mui/material";

import NiMessages from "@/icons/nexture/ni-messages";

export default function UserLanguageSwitch() {
  const {
    t,
    i18n: { language: locale },
  } = useTranslation();

  return (
    <MenuItem disabled>
      <ListItemIcon>
        <NiMessages size={20} />
      </ListItemIcon>
      <div className="w-full">{t("user-language")}</div>
      <div className="text-sm text-gray-600">{t(locale)}</div>
    </MenuItem>
  );
}
