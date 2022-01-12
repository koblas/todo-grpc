import React, { useEffect } from "react";
import { Routes, Route } from "react-router-dom";
import { useToast } from "@chakra-ui/react";

import { AuthRecoverSendPage } from "./RecoverySendPage";
import LoginPage from "./LoginPage";
import OAuthPage from "./OAuthPage";
import { AuthLogoutPage } from "./LogoutPage";
import { AuthRegisterPage } from "./RegisterPage";
import { AuthRecoveryResetPage } from "./RecoveryResetPage";
import { AuthEmailConfirmPage } from "./EmailConfirmPage";
import { NotFoundPage } from "../NotFoundPage";
import { NetworkContextProvider, useNetworkContext } from "../../hooks/network";

function AuthPagesActions() {
  const toast = useToast();
  const network = useNetworkContext();

  useEffect(() => {
    function showToast() {
      toast({
        title: "Network error",
        status: "error",
        isClosable: true,
      });
    }

    network.setHandlers({
      onErrorNetwork: showToast,
    });
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <Routes>
      <Route path="/register" element={<AuthRegisterPage />} />
      <Route path="/login" element={<LoginPage />} />
      <Route path="/logout" element={<AuthLogoutPage />} />
      <Route path="/oauth/:provider" element={<OAuthPage />} />
      <Route path="/recover/send" element={<AuthRecoverSendPage />} />
      <Route path="/email/confirm/:userId/:token" element={<AuthEmailConfirmPage />} />
      <Route path="/recover/verify/:userId/:token" element={<AuthRecoveryResetPage />} />
      <Route path="/*" element={<NotFoundPage />} />
    </Routes>
  );
}

export function AuthPages() {
  return (
    <NetworkContextProvider>
      <AuthPagesActions />
    </NetworkContextProvider>
  );
}
