import React, { useEffect, useState } from "react";
import { Alert, Text, Box, Flex, Heading, Link, Stack, Spinner } from "@chakra-ui/react";
import { Link as RouterLink, useNavigate, useLocation, useParams } from "react-router-dom";
import { useAuth } from "../../hooks/auth";
import AuthWrapper from "./AuthWrapper";

export default function OAuthPage() {
  const auth = useAuth();
  const location = useLocation();
  const navigate = useNavigate();
  const [loading, setLoading] = useState(true);
  const [success, setSuccess] = useState(false);
  // const [lError, setLError] = useState<string>("");
  const { provider } = useParams<{ provider: string }>();
  const [oauthLogin, { loading: loadingLogin }] = auth.mutations.useOauthLogin();
  const [oauthRedirect, { loading: loadingRedirect }] = auth.mutations.useOauthRedirect();

  const redirectUrl = `${window.location.origin}/auth/oauth/${provider}`;
  const search = new URLSearchParams(location.search);
  const code = search.get("code");
  const state = search.get("state") ?? "";

  if (!provider) {
    throw new Error("Missing provider");
  }

  useEffect(() => {
    // No code, redice to the OAuth provider
    if (!code) {
      const returnUrl = `${window.location.origin}/auth/oauth/${provider}`;

      oauthRedirect(
        { provider, returnUrl },
        {
          onCompleted({ url }) {
            window.location.href = url;
          },
          onError() {
            setLoading(false);
          },
        },
      );

      return;
    }

    oauthLogin(
      { provider, code, redirectUrl, state },
      {
        onCompleted(data) {
          setSuccess(true);

          if (!data?.created) {
            // TODO - should be "next"
            navigate("/", { replace: true });
          }
        },
        onFinished() {
          setLoading(false);
        },
      },
    );
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [provider, code, redirectUrl]);

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
