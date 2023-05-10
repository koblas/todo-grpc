import React, { useEffect, useState } from "react";
import { Link as RouterLink, useParams } from "react-router-dom";
import { Text, Flex, Heading, Link, Stack, Box, Alert, Spinner } from "@chakra-ui/react";
import { useAuth } from "../../hooks/auth";
import { assert } from "../../util/assert";

export function AuthEmailConfirmPage() {
  const { userId, token } = useParams<{ userId: string; token: string }>();
  const auth = useAuth();
  const [verified, setVerified] = useState(false);
  const [emailConfirm, { loading }] = auth.mutations.useEmailConfirm();

  assert(userId && token);

  useEffect(() => {
    emailConfirm(
      { user_id: userId, token },
      {
        onCompleted() {
          setVerified(true);
        },
        onErrorField() {
          setVerified(false);
        },
      },
    );
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [userId, token]);

  return (
    <Stack minH="100vh" direction={{ base: "column", md: "row" }}>
      <Flex p={8} flex={1} align="center" justify="center">
        <Stack spacing={8} w="full" maxW="md">
          <Heading fontSize="2xl">Confirm email address</Heading>
          <Text>TODO: Insert some nice text to keep you interested</Text>
          {loading && <Spinner />}
          {!loading && verified && <Alert status="success">Thank you</Alert>}
          {!loading && !verified && <Alert status="error">This confirmation link has already been used</Alert>}
          <Stack spacing={6}>
            <Stack direction={{ base: "column", sm: "row" }} align="start" justify="space-between">
              <Link as={RouterLink} to="/auth/login" color="blue.500">
                Sign-in
              </Link>
            </Stack>
          </Stack>
        </Stack>
      </Flex>
      <Flex flex={1}>
        <Box
          width="100%"
          bgImage="url(https://source.unsplash.com/random)"
          bgRepeat="no-repeat"
          bgPosition="center"
          // bgColor={(t) => (t.palette.mode === "light" ? t.palette.grey[50] : t.palette.grey[900])},
          bgSize="cover"
        />
      </Flex>
    </Stack>
  );
}
