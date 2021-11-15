import React, { useEffect, useState } from "react";
import { useHistory, Link as RouterLink, useParams } from "react-router-dom";
import { Text, Flex, Heading, Link, Stack, Box, Alert, AlertIcon, Spinner } from "@chakra-ui/react";
import { useForm } from "react-hook-form";
import { useAuth } from "../../hooks/auth";

export function AuthEmailConfirmPage() {
  const {
    register,
    handleSubmit,
    formState: { errors },
    setError,
  } = useForm();
  const { userId, token } = useParams<{ userId: string; token: string }>();
  const auth = useAuth();
  const [verified, setVerified] = useState(false);
  const [emailConfirm, { loading }] = auth.mutations.useEmailConfirm();

  useEffect(() => {
    console.log("VERIFY TOKEN", token);
    emailConfirm(
      { userId, token },
      {
        onCompleted() {
          setVerified(true);
        },
        onErrorField() {
          setVerified(false);
        },
      },
    );
  }, []);

  return (
    <Stack minH="100vh" direction={{ base: "column", md: "row" }}>
      <Flex p={8} flex={1} align="center" justify="center">
        <Stack spacing={8} w="full" maxW="md">
          <Heading fontSize="2xl">Confirm email address</Heading>
          <Text></Text>
          {loading && <Spinner />}
          {!loading && verified && <Alert type="success">Thank you</Alert>}
          {!loading && !verified && <Alert type="error">Email already confirmed</Alert>}
          <Stack spacing={6}>
            <Stack direction={{ base: "column", sm: "row" }} align="start" justify="space-between">
              <Link as={RouterLink} to="/auth/login" color="blue.500">
                Sign-in
              </Link>
              <Link as={RouterLink} to="/auth/recover/send" color="blue.500">
                Forgot password?
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
