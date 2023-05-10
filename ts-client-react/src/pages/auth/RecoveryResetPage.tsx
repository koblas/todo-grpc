import React, { useEffect, useState } from "react";
import { Link as RouterLink, useLocation, useNavigate, useParams } from "react-router-dom";
import {
  Text,
  Button,
  Flex,
  FormControl,
  FormLabel,
  Heading,
  Input,
  Link,
  Stack,
  Box,
  FormErrorMessage,
  Alert,
  AlertIcon,
  Spinner,
} from "@chakra-ui/react";
import { useForm } from "react-hook-form";
import { useAuth } from "../../hooks/auth";
import { assert } from "../../util/assert";

type FormFields = {
  password: string;
};

function Verify({ loading, success }: { loading: boolean; success: boolean }): JSX.Element | null {
  if (loading) {
    return <Spinner />;
  }
  if (success) {
    return null;
  }

  return <Alert>It appears that this link has already been used</Alert>;
}

export function AuthRecoveryResetPage() {
  const {
    register,
    handleSubmit,
    formState: { errors },
    setError,
  } = useForm<FormFields>();
  const { user_id: userId, token } = useParams<{ user_id: string; token: string }>();
  const auth = useAuth();
  const { search } = useLocation();
  const navigate = useNavigate();
  const [verified, setVerified] = useState(false);
  const [recoverVerify, { loading: loadingVerify }] = auth.mutations.useRecoveryVerify();
  const [recoverUpdate, { loading: loadingUpdate }] = auth.mutations.useRecoveryUpdate();
  const [completed, setCompleted] = useState(false);

  assert(userId && token);

  useEffect(() => {
    recoverVerify(
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

  function onSubmit({ password }: { password: string }) {
    assert(userId && token);

    recoverUpdate(
      { user_id: userId, token, password },
      {
        onCompleted() {
          setCompleted(true);
          const query = new URLSearchParams(search);
          navigate(query.get("next") ?? "/", { replace: true });
        },
        onErrorField(serror: Record<string, string[]>) {
          const fields: (keyof FormFields)[] = ["password"];
          fields.forEach((key) => {
            const message = serror[key]?.[0];
            if (message) {
              setError(key, { message });
            }
          });
        },
      },
    );
  }

  let body;
  if (!verified) {
    body = <Verify loading={loadingVerify} success={verified} />;
  } else if (!completed) {
    body = (
      <FormControl id="password" isInvalid={!!errors.password}>
        <FormLabel>Enter your new password</FormLabel>
        <Input
          type="password"
          {...register("password", {
            required: {
              value: true,
              message: "Please provide a password",
            },
            minLength: {
              value: 8,
              message: "Passwords must have 8 characters",
            },
          })}
        />
        <FormErrorMessage>{errors.password?.message}</FormErrorMessage>
      </FormControl>
    );
  } else {
    body = (
      <Alert status="success">
        <AlertIcon />
        Instructions have been sent, please check your email.
      </Alert>
    );
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <Stack minH="100vh" direction={{ base: "column", md: "row" }}>
        <Flex p={8} flex={1} align="center" justify="center">
          <Stack spacing={8} w="full" maxW="md">
            <Heading fontSize="2xl">Reset your password?</Heading>
            <Text>
              RECOVERY RESET! Entere the email address you signed up with, and we will send instructions on how to reset
              your password
            </Text>
            {body}
            <Stack spacing={6}>
              {!completed && verified && (
                <Button isLoading={loadingUpdate} disabled={completed} type="submit" colorScheme="blue" variant="solid">
                  Reset your password
                </Button>
              )}
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
    </form>
  );
}
