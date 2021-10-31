import React, { useState } from "react";
import { useHistory, Link as RouterLink } from "react-router-dom";
import {
  Text,
  Button,
  Checkbox,
  Flex,
  FormControl,
  FormLabel,
  Heading,
  Input,
  Link,
  Stack,
  Box,
  FormErrorMessage,
} from "@chakra-ui/react";
import { useForm } from "react-hook-form";
import { useAuth } from "../hooks/auth";

export function AuthLoginPage() {
  const {
    register,
    handleSubmit,
    formState: { errors },
    setError,
  } = useForm();
  const auth = useAuth();
  const [login] = auth.mutations.useLogin();
  const history = useHistory();

  function onSubmit(data: { email: string; password: string }) {
    login(data, {
      onCompleted() {
        history.replace("/todo");
      },
      onErrorField(serror: Record<string, string[]>) {
        ["email", "password"].forEach((key) => {
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
            <Heading fontSize="2xl">Welcome back to ProjectX</Heading>
            <Text>Sign in to your account</Text>
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
            <FormControl id="password" isInvalid={!!errors.password}>
              <FormLabel>Password</FormLabel>
              <Input
                type="password"
                {...register("password", {
                  required: {
                    value: true,
                    message: "Password is required",
                  },
                })}
              />
              <FormErrorMessage>{errors.password?.message}</FormErrorMessage>
            </FormControl>
            <Stack spacing={6}>
              <Button type="submit" colorScheme="blue" variant="solid">
                Sign in
              </Button>
              <Stack direction={{ base: "column", sm: "row" }} align="start" justify="space-between">
                <Text>
                  Don't have an account?{" "}
                  <Link as={RouterLink} to="/auth/register" color="blue.500">
                    Register
                  </Link>
                </Text>
                <Link as={RouterLink} to="/auth/recover" color="blue.500">
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
