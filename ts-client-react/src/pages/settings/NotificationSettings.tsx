// https://dribbble.com/shots/15186840-Setting-page-example

import React, { useState } from "react";
import { Text, Heading, FormControl, FormLabel, Button, Flex, Box, useToast, FormErrorMessage } from "@chakra-ui/react";
import { useForm } from "react-hook-form";
import { useUser } from "../../hooks/data/user";
import { FetchError } from "../../rpc/utils";

const INPUT_STYLE = {
  bg: "white",
  borderColor: "gray.400",
};

type FormFields = {
  oldPassword: string;
  newPassword: string;
  passwordAgain: string;
};

export function NotificationSettings() {
  const { mutations } = useUser();
  const toast = useToast();

  const [isSubmitting, setSubmitting] = useState(false);
  const {
    register,
    handleSubmit,
    setError,
    formState: { errors },
  } = useForm<FormFields>({
    defaultValues: {},
  });

  function onSubmit(data: FormFields) {
    setSubmitting(true);
    mutations.updateUser.mutate(data, {
      onSettled() {
        setSubmitting(false);
      },
      onError(error) {
        if (error instanceof FetchError) {
          const { code, msg, argument } = error.getInfo();

          if (code === "invalid_argument" && msg && argument) {
            setError(argument as keyof FormFields, { message: msg });
          }
        }
        setSubmitting(false);
      },
      onSuccess() {
        toast({
          position: "top",
          title: "Profile Updated",
          status: "success",
          isClosable: true,
        });
      },
    });
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <Box w="100%" p="5">
        <Flex justify="center">
          <Box w="80%">
            <Heading fontSize="3xl">Notifications</Heading>
            <Box pt="5">
              <Heading fontSize="lg">Profile</Heading>
              <Text>When we will contact you</Text>
            </Box>

            <Flex pt="5" justify="right">
              <Button mt={4} colorScheme="teal" isLoading={isSubmitting} type="submit" onClick={handleSubmit(onSubmit)}>
                Save
              </Button>
            </Flex>
          </Box>
        </Flex>
      </Box>
    </form>
  );
}
