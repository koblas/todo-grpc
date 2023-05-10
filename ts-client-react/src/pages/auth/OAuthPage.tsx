import React, { useEffect, useState } from "react";
import { Alert, Text, Box, Flex, Heading, Link, Stack, Spinner } from "@chakra-ui/react";
import { Link as RouterLink, useNavigate, useLocation, useParams } from "react-router-dom";
import { useAuth } from "../../hooks/auth";
import AuthWrapper from "./AuthWrapper";

export default function OAuthPage() {
  const auth = useAuth();
  const navigate = useNavigate();
  const [success, setSuccess] = useState(true);
  const [loading, setLoading] = useState(true);
  const { provider } = useParams<{ provider: string }>();
  const [oauthLogin, { loading: loadingLogin }] = auth.mutations.useOauthLogin();
  const [oauthRedirect, { loading: loadingRedirect }] = auth.mutations.useOauthRedirect();

  const redirectUrl = `${window.location.origin}/auth/oauth/${provider}`;
  const { search: queryString } = useLocation();
  const search = new URLSearchParams(queryString);
  const code = search.get("code");
  const state = search.get("state") ?? "";

  if (!provider) {
    throw new Error("Missing provider");
  }

  useEffect(() => {
    // FIXME -- `next` parameter triggers a google error with redirects, so not implemented
    // const query = new URLSearchParams(search);
    // const next = query.get("next") ?? "";

    // const nextQuery = next ? `?next=${encodeURIComponent(next)}` : "";
    // eslint-disable-next-line @typescript-eslint/naming-convention
    // const redirect_url = `${window.location.origin}/auth/oauth/${provider}${nextQuery}`;
    // No code, redice to the OAuth provider

    // This is a total hack, since useEffect in React 18 is called twice we need to
    // handle the create/destroy cycle and only run for the one that sticks around.
    // There probably is a better way, but ugh.
    const timer = setTimeout(() => {
      if (!code) {
        oauthRedirect(
          { provider, redirect_url: redirectUrl },
          {
            onCompleted({ url }) {
              setTimeout(() => {
                // window.location.href = url;
                window.location.assign(url);
              }, 0);
            },
            onError() {
              setLoading(false);
            },
          },
        );

        return;
      }

      oauthLogin(
        { provider, code, redirect_url: redirectUrl, state },
        {
          onCompleted(data) {
            if (data.created) {
              // TODO -- If created we should flag this as a new user
              // For now we just have a lame message
              setSuccess(true);
            } else {
              navigate(search.get("next") ?? "/", { replace: true });
            }
          },
          onFinished() {
            setLoading(false);
          },
        },
      );
    }, 2);

    return () => clearTimeout(timer);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [provider, code, state]);

  // When redirected from Google/GitHub/etc.. Show a spinner while we figure out
  //  if everything is valid and the account exists or needs to be created
  if (code && loading) {
    return (
      <Box width="100%" height="100vh" alignContent="center" alignItems="center">
        <Flex p={8} flex={1} align="center" justify="center">
          <Spinner />;
        </Flex>
      </Box>
    );
  }

  return (
    <AuthWrapper loading={loadingLogin || loadingRedirect || loading}>
      <Flex p={8} flex={1} align="center" justify="center">
        <Stack spacing={8} w="full" maxW="md">
          {success ? (
            <>
              <Heading>Welcome!</Heading>
              <Text>Your account is now created, to get started</Text>
              <Link as={RouterLink} to="/" color="blue.500">
                Visit your homepage
              </Link>
            </>
          ) : (
            <>
              <Heading>An error has occured</Heading>
              <Alert>Error creating account</Alert>
              <Text>
                Please return to the{" "}
                <Link as={RouterLink} to="/auth/login" color="blue.500">
                  Sign-in
                </Link>{" "}
                page.
              </Text>
            </>
          )}
        </Stack>
      </Flex>
    </AuthWrapper>
  );
}
