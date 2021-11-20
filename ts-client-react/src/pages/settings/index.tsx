import React from "react";
import { Switch, useRouteMatch, Route } from "react-router-dom";
import { Box } from "@chakra-ui/react";

import { ProfileSettings } from "./ProfileSettings";
import { SecuritySettings } from "./SecuritySettings";
import { NotFoundPage } from "../NotFoundPage";

export function SettingsPage() {
  const { path } = useRouteMatch();

  return (
    <Box w="100%" p="8" bgColor="gray.100">
      <Switch>
        <Route path={`${path}/`} exact>
          <ProfileSettings />
        </Route>
        <Route path={`${path}/security`}>
          <SecuritySettings />
        </Route>
        <Route path={`${path}/*`}>
          <NotFoundPage />
        </Route>
      </Switch>
    </Box>
  );
}
