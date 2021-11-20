import React, { useEffect } from "react";
import { useHistory, Link as RouterLink, useLocation } from "react-router-dom";
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
  InputGroup,
  Icon,
} from "@chakra-ui/react";
import { useForm } from "react-hook-form";
import { FaGithub } from "react-icons/fa";
import { FcGoogle } from "react-icons/fc";
import { useAuth } from "../../hooks/auth";
import AuthWrapper from "./AuthWrapper";
import { PasswordInput } from "../../components/PasswordInput";

export default function AuthLoginPage() {
  const {
    register,
    handleSubmit,
    formState: { errors },
    setError,
    getValues,
  } = useForm();
  const auth = useAuth();
  const [login, { loading }] = auth.mutations.useLogin();
  const history = useHistory();
  const { search } = useLocation();

  useEffect(() => {
    if (auth.isAuthenticated) {
      const query = new URLSearchParams(search);

      history.replace(query.get("next") ?? "/");
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [auth.isAuthenticated]);

  function onSubmit(data: { email: string; password: string }) {
    login(data, {
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

  function onOauthButton(provider: string) {
    history.push(`/auth/oauth/${provider}`);
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <AuthWrapper>
        <Flex p={8} flex={1} align="center" justify="center">
          <Stack spacing={8} w="full" maxW="md">
            <Heading fontSize="2xl">Welcome back to ProjectX</Heading>
            <Text>Sign in to your account</Text>
            <Stack direction={{ base: "row" }}>
              <Button variant="outline" onClick={() => onOauthButton("google")} className="">
                <Icon as={FcGoogle} size="8" mr="2" /> Sign in with Google
              </Button>
              <Button variant="outline" onClick={() => onOauthButton("github")} className="">
                <Icon as={FaGithub} size="8" mr="2" /> Sign in with Github
              </Button>
            </Stack>
            <FormControl id="email" isInvalid={!!errors.email}>
              <FormLabel>Email address</FormLabel>
              <InputGroup>
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
              </InputGroup>
              <FormErrorMessage>{errors.email?.message}</FormErrorMessage>
            </FormControl>
            <FormControl id="password" isInvalid={!!errors.password}>
              <FormLabel>Password</FormLabel>
              <PasswordInput
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
              <Button isLoading={loading} type="submit" colorScheme="blue" variant="solid">
                Sign in
              </Button>
              <Stack direction={{ base: "column", sm: "row" }} align="start" justify="space-between">
                <Text>
                  Don't have an account?{" "}
                  <Link as={RouterLink} to="/auth/register" color="blue.500">
                    Register
                  </Link>
                </Text>
                <Link
                  onClick={(e) => {
                    e.preventDefault();
                    const { email } = getValues();

                    history.push({
                      pathname: "/auth/recover/send",
                      ...(email ? { search: `?email=${encodeURIComponent(email)}` } : {}),
                    });
                  }}
                  to="/auth/recover/send"
                  color="blue.500"
                >
                  Forgot password?
                </Link>
              </Stack>
            </Stack>
          </Stack>
        </Flex>
      </AuthWrapper>
    </form>
  );
}
