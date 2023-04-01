import React from "react";
import { Flex } from "@chakra-ui/react";
import { Routes, Route, Outlet } from "react-router-dom";

import { AuthPages } from "./pages/auth";
import { SettingsPage } from "./pages/settings";
import { ReportPage } from "./pages/report";
import { TodoPage } from "./pages/TodoPage";
import { GptPage } from "./pages/GptPage";
import { MessagePage } from "./pages/MessagePage";
import { UploadPage } from "./pages/UploadPage";
import { HomePage } from "./pages/HomePage";
import { Sidebar } from "./components/Sidebar";
import ProtectedRoute from "./components/ProtectedRoute";

function SiteLayout(): JSX.Element {
  return (
    <ProtectedRoute>
      <Flex w="100%">
        <Sidebar />
        <Outlet />
      </Flex>
    </ProtectedRoute>
  );
}

function Site() {
  return (
    <Routes>
      <Route path="/" element={<SiteLayout />}>
        <Route index element={<HomePage />} />
        <Route path="settings/*" element={<SettingsPage />} />
        <Route path="todo/*" element={<TodoPage />} />
        <Route path="gpt/*" element={<GptPage />} />
        <Route path="message/*" element={<MessagePage />} />
        <Route path="upload/*" element={<UploadPage />} />
        <Route path="*" element={<HomePage />} />
      </Route>
    </Routes>
  );
}

export function AppRoutes(): React.ReactElement {
  return (
    <Routes>
      <Route path="/auth/*" element={<AuthPages />} />
      <Route path="/report" element={<ReportPage />} />
      <Route path="*" element={<Site />} />
    </Routes>
  );
}
