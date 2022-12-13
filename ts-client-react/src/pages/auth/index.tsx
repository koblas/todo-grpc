import React from "react";
import { Routes, Route } from "react-router-dom";

import { AuthRecoverSendPage } from "./RecoverySendPage";
import LoginPage from "./LoginPage";
import OAuthPage from "./OAuthPage";
import { AuthLogoutPage } from "./LogoutPage";
import { AuthRegisterPage } from "./RegisterPage";
import { AuthRecoveryResetPage } from "./RecoveryResetPage";
import { AuthEmailConfirmPage } from "./EmailConfirmPage";
import { NotFoundPage } from "../NotFoundPage";

function AuthPagesActions() {
  return (
    <Routes>
      <Route path="/register" element={<AuthRegisterPage />} />
      <Route path="/login" element={<LoginPage />} />
      <Route path="/logout" element={<AuthLogoutPage />} />
      <Route path="/oauth/:provider" element={<OAuthPage />} />
      <Route path="/recover/send" element={<AuthRecoverSendPage />} />
      <Route path="/recover/verify/:userId/:token" element={<AuthRecoveryResetPage />} />
      <Route path="/email/confirm/:userId/:token" element={<AuthEmailConfirmPage />} />
      <Route path="/*" element={<NotFoundPage />} />
    </Routes>
  );
}

export function AuthPages() {
  return <AuthPagesActions />;
}
