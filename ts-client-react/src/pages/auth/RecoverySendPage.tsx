import React, { useEffect, useState } from "react";
import { Link as RouterLink, useLocation } from "react-router-dom";
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
  FormErrorMessage,
  Alert,
  AlertIcon,
} from "@chakra-ui/react";
import { useForm } from "react-hook-form";
import { useAuth } from "../../hooks/auth";
import AuthWrapper from "./AuthWrapper";

type FormFields = {
  email: string;
};

export function AuthRecoverSendPage() {
  const location = useLocation();
  const {
    register,
    handleSubmit,
    formState: { errors },
    setError,
    setValue,
  } = useForm<FormFields>();
  const auth = useAuth();
  const [recoverSend, { loading }] = auth.mutations.useRecoverSend();
  const [completed, setCompleted] = useState(false);

  function onSubmit(data: { email: string }) {
    recoverSend(data, {
      onCompleted() {
        setCompleted(true);
      },
      onErrorField(serror: Record<string, string[]>) {
        const fields: (keyof FormFields)[] = ["email"];

        fields.forEach((key) => {
          const message = serror[key]?.[0];
          if (message) {
            setError(key, { message });
          }
        });
      },
    });
  }

  useEffect(() => {
    const params = new URLSearchParams(location.search);

    const email = params.get("email");
    if (email) {
      setValue("email", email);
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [location.search]);

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <AuthWrapper>
        <Flex p={8} flex={1} align="center" justify="center">
          <Stack spacing={8} w="full" maxW="md">
            <Heading fontSize="2xl">Forgot your password?</Heading>
            <Text>
              Entere the email address you signed up with, and we will send instructions on how to reset your password
            </Text>
            {completed ? (
              <Alert status="success">
                <AlertIcon />
                Instructions have been sent, please check your email.
              </Alert>
            ) : (
              <FormControl id="email" isInvalid={!!errors.email}>
                <FormLabel>Email address</FormLabel>
                <Input
                  autoFocus
                  type="email"
                  {...register("email", {
                    required: {
                      value: true,
                      message: "Email address is required",
                    },
                  })}
                />
                <FormErrorMessage>{errors.email?.message}</FormErrorMessage>
              </FormControl>
            )}
            <Stack spacing={6}>
              {!completed && (
                <Button isLoading={loading} disabled={completed} type="submit" colorScheme="blue" variant="solid">
                  Reset your password
                </Button>
              )}
              <Stack direction={{ base: "column", sm: "row" }} align="start" justify="space-between">
                <Text>
                  Remember your password?{" "}
                  <Link as={RouterLink} to="/auth/login" color="blue.500">
                    Sign-in
                  </Link>
                </Text>
              </Stack>
            </Stack>
          </Stack>
        </Flex>
      </AuthWrapper>
    </form>
  );
}
