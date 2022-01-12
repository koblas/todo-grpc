import React from "react";
import { Route, Routes } from "react-router-dom";
import { Box } from "@chakra-ui/react";

import { ProfileSettings } from "./ProfileSettings";
import { SecuritySettings } from "./SecuritySettings";
import { NotFoundPage } from "../NotFoundPage";

export function SettingsPage() {
  console.log("SETTING APGE");
  return (
    <Box w="100%" p="8" bgColor="gray.100">
      <Routes>
        <Route path="/" element={<ProfileSettings />} />
        <Route path="/security" element={<SecuritySettings />} />
        <Route path="*" element={<NotFoundPage />} />
      </Routes>
    </Box>
  );
}
