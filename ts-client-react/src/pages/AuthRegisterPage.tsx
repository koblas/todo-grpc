import React, { useState } from "react";
import { useHistory, Link as RouterLink } from "react-router-dom";
import {
  Link,
  Text,
  Button,
  Flex,
  FormControl,
  FormLabel,
  Heading,
  Input,
  Stack,
  Box,
  FormErrorMessage,
} from "@chakra-ui/react";
import { useForm } from "react-hook-form";
import { useAuth } from "../hooks/auth";

export function AuthRegisterPage() {
  const {
    register,
    handleSubmit,
    formState: { errors },
    setError,
  } = useForm();
  const { mutations } = useAuth();
  const [authRegister] = mutations.useRegister();
  const history = useHistory();

  function onSubmit(data: { email: string; password: string; name: string }) {
    authRegister(data, {
      onCompleted() {
        history.replace("/todo");
      },
      onErrorField(serror: Record<string, string[]>) {
        console.log("FIELD ERRORS", serror);
        ["name", "email", "password"].forEach((key) => {
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
            <Heading fontSize="2xl">Create your account</Heading>
            <FormControl id="name" isInvalid={!!errors.name}>
              <FormLabel>Your Name</FormLabel>
              <Input
                type="text"
                {...register("name", {
                  required: {
                    value: true,
                    message: "Name is required",
                  },
                })}
              />
              <FormErrorMessage>{errors.name?.message}</FormErrorMessage>
            </FormControl>
            <FormControl id="email" isInvalid={!!errors.email}>
              <FormLabel>Email address</FormLabel>
              <Input
                type="email"
                {...register("email", {
                  required: {
                    value: true,
                    message: "Email is required",
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
            <Stack spacing={6}>
              <Button type="submit" colorScheme="blue" variant="solid">
                Register
              </Button>
            </Stack>

            <Text>
              Already have an account?{" "}
              <Link as={RouterLink} to="/auth/login" color="blue.500">
                Sign-in
              </Link>
            </Text>
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
