import React, { useEffect } from "react";
import { Switch, useRouteMatch, Route } from "react-router-dom";
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
  const { path } = useRouteMatch();
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
    <>
      <Switch>
        <Route path={`${path}/register`}>
          <AuthRegisterPage />
        </Route>
        <Route path={`${path}/login`}>
          <LoginPage />
        </Route>
        <Route path={`${path}/logout`}>
          <AuthLogoutPage />
        </Route>
        <Route path={`${path}/oauth/:provider`}>
          <OAuthPage />
        </Route>
        <Route path={`${path}/recover/send`}>
          <AuthRecoverSendPage />
        </Route>
        <Route path={`${path}/email/confirm/:userId/:token`}>
          <AuthEmailConfirmPage />
        </Route>
        <Route path={`${path}/recover/verify/:userId/:token`}>
          <AuthRecoveryResetPage />
        </Route>
        <Route path={`${path}/*`}>
          <NotFoundPage />
        </Route>
      </Switch>
    </>
  );
}

export function AuthPages() {
  return (
    <NetworkContextProvider>
      <AuthPagesActions />
    </NetworkContextProvider>
  );
}
