import React, { useState } from "react";
import { useHistory } from "react-router-dom";
import { Button, Checkbox, Flex, FormControl, FormLabel, Heading, Input, Link, Stack, Box } from "@chakra-ui/react";
import { useAuth } from "../hooks/auth";

export function AuthRegisterPage() {
  const [username, setUsername] = useState("");
  const [name, setName] = useState("");
  const [password, setPassword] = useState("");
  const auth = useAuth();
  const history = useHistory();

  function handleSubmit() {
    auth
      .register(username, password, name)
      .then(() => {
        history.replace("/todo");
        setUsername("");
        setPassword("");
      })
      .catch((err) => {
        console.log(err);
        // todo
      });
  }

  return (
    <Stack minH="100vh" direction={{ base: "column", md: "row" }}>
      <Flex p={8} flex={1} align="center" justify="center">
        <Stack spacing={4} w="full" maxW="md">
          <form
            onSubmit={(e) => {
              e.preventDefault();
              handleSubmit();
            }}
          >
            <Heading fontSize="2xl">Create an account</Heading>
            <FormControl id="name">
              <FormLabel>Your Full Name</FormLabel>
              <Input type="text" value={name} onChange={(e) => setName(e.target.value)} />
            </FormControl>
            <FormControl id="email">
              <FormLabel>Email address</FormLabel>
              <Input type="email" value={username} onChange={(e) => setUsername(e.target.value)} />
            </FormControl>
            <FormControl id="password">
              <FormLabel>Password</FormLabel>
              <Input type="password" value={password} onChange={(e) => setPassword(e.target.value)} />
            </FormControl>
            <Stack spacing={6}>
              <Stack direction={{ base: "column", sm: "row" }} align="start" justify="space-between">
                <Checkbox>Remember me</Checkbox>
                <Link color="blue.500">Forgot password?</Link>
              </Stack>
              <Button type="submit" colorScheme="blue" variant="solid">
                Sign in
              </Button>
            </Stack>
          </form>
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
