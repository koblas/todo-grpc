import React, { useState } from "react";
import { useHistory, Link as RouterLink } from "react-router-dom";
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
} from "@chakra-ui/react";
import { useForm } from "react-hook-form";
import { useAuth } from "../../hooks/auth";

export function AuthRecoverSendPage() {
  const {
    register,
    handleSubmit,
    formState: { errors },
    setError,
  } = useForm();
  const auth = useAuth();
  const [recoverSend, { loading }] = auth.mutations.useRecoverSend();
  const [completed, setCompleted] = useState(false);

  function onSubmit(data: { email: string }) {
    recoverSend(data, {
      onCompleted() {
        setCompleted(true);
      },
      onErrorField(serror: Record<string, string[]>) {
        ["email"].forEach((key) => {
          const message = serror[key]?.[0];
          if (message) {
            setError(key, { message });
          }
        });
      },
    });
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <Stack minH="100vh" direction={{ base: "column", md: "row" }}>
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
