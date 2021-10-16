import React, { useState } from "react";
import {
  Box,
  Flex,
  Button,
  FormControl,
  FormLabel,
  Heading,
  Input,
  Link,
  Switch,
  Text,
  useColorModeValue,
} from "@chakra-ui/react";
import { useHistory } from "react-router-dom";
import { useAuth } from "../hooks/auth";
import { DoorImage } from "../components/DoorImage";

export function AuthLoginPage() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const auth = useAuth();
  const history = useHistory();

  function handleLogin() {
    auth
      .login(username, password)
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

  // Chakra color mode
  const titleColor = useColorModeValue("teal.300", "teal.200");
  const textColor = useColorModeValue("gray.400", "white");

  return (
    <Flex position="relative">
      <Flex h={{ sm: "initial", md: "100vh", lg: "100vh" }} w="100%" mx="auto" justifyContent="space-between">
        <Flex alignItems="center" justifyContent="start" style={{ userSelect: "none" }} w="50%">
          <Flex
            direction="column"
            background="transparent"
            p="48px"
            w="100%"
            mt={{ md: "150px", lg: "80px" }}
            pl="24"
            pr="24"
          >
            <Heading color={titleColor} fontSize="32px" mb="10px">
              Welcome Back
            </Heading>
            <Text mb="36px" ms="4px" color={textColor} fontWeight="bold" fontSize="14px">
              Enter your email and password to sign in
            </Text>
            <FormControl>
              <FormLabel ms="4px" fontSize="sm" fontWeight="normal">
                Email
              </FormLabel>
              <Input
                borderRadius="15px"
                mb="24px"
                fontSize="sm"
                type="text"
                placeholder="Your email adress"
                size="lg"
              />
              <FormLabel ms="4px" fontSize="sm" fontWeight="normal">
                Password
              </FormLabel>
              <Input
                borderRadius="15px"
                mb="36px"
                fontSize="sm"
                type="password"
                placeholder="Your password"
                size="lg"
              />
              {/* <FormControl display="flex" alignItems="center">
                <Switch id="remember-login" colorScheme="teal" me="10px" />
                <FormLabel htmlFor="remember-login" mb="0" ms="1" fontWeight="normal">
                  Remember me
                </FormLabel>
              </FormControl> */}
              <Button
                fontSize="10px"
                type="submit"
                bg="teal.300"
                w="100%"
                h="45"
                mb="20px"
                color="white"
                mt="20px"
                _hover={{
                  bg: "teal.200",
                }}
                _active={{
                  bg: "teal.400",
                }}
              >
                SIGN IN
              </Button>
            </FormControl>
            <Flex flexDirection="column" justifyContent="center" alignItems="center" maxW="100%" mt="0px">
              <Text color={textColor} fontWeight="medium">
                Don't have an account?
                <Link color={titleColor} as="span" ms="5px" fontWeight="bold">
                  Sign Up
                </Link>
              </Text>
            </Flex>
          </Flex>
        </Flex>
        <Flex
          flex="1 1 0%"
          h="100vh"
          // w="50%"
          display={{ base: "none", md: "flex" }}
          alignItems="center"
          justifyContent="center"
          right="0px"
          background="rgb(224,231,255,1)"
        >
          <DoorImage />
        </Flex>
      </Flex>
    </Flex>
  );

  //       <div className="hidden lg:flex items-center justify-center bg-indigo-100 flex-1 h-screen">
}
