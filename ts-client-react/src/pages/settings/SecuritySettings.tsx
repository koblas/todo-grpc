// https://dribbble.com/shots/15186840-Setting-page-example

import React, { useState } from "react";
import { Text, Heading, FormControl, FormLabel, Button, Flex, Box, useToast, FormErrorMessage } from "@chakra-ui/react";
import { useForm } from "react-hook-form";
import { useUser, useUserMutations } from "../../hooks/data/user";
import { PasswordInput } from "../../components/PasswordInput";

const INPUT_STYLE = {
  bg: "white",
  borderColor: "gray.400",
};

type FormFields = {
  password: string;
  passwordNew: string;
  passwordAgain: string;
};

export function SecuritySettings() {
  const { useUpdateUser } = useUserMutations();
  const toast = useToast();
  const [updateUser] = useUpdateUser();

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
    if (data.passwordAgain !== data.passwordNew) {
      setError("passwordNew", { message: "Password doesn't match retryped" });
      return;
    }
    setSubmitting(true);
    updateUser(data, {
      onFinished() {
        setSubmitting(false);
      },
      onErrorField(serror: Record<string, string[]>) {
        const fields: (keyof FormFields)[] = ["password", "passwordAgain", "passwordNew"];
        fields.forEach((key: keyof FormFields) => {
          const message = serror[key]?.[0];
          if (message) {
            setError(key, { message });
          }
        });
      },
      onCompleted() {
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
            <Heading fontSize="3xl">Security</Heading>
            <Box pt="5">
              <Heading fontSize="lg">Profile</Heading>
              <Text>Please help keep your account secure</Text>
            </Box>

            <Box pt="5">
              <FormControl isInvalid={!!errors.password}>
                <FormLabel>Old Password</FormLabel>
                <Flex>
                  <PasswordInput
                    {...register("password", {
                      required: {
                        value: true,
                        message: "Old password is required",
                      },
                    })}
                    {...INPUT_STYLE}
                  />
                </Flex>
                <FormErrorMessage>{errors.password?.message}</FormErrorMessage>
              </FormControl>
            </Box>

            <Box pt="5">
              <FormControl isInvalid={!!errors.passwordNew}>
                <FormLabel>New Password</FormLabel>
                <Flex>
                  <PasswordInput
                    {...register("passwordNew", {
                      required: {
                        value: true,
                        message: "Password is required",
                      },
                      minLength: {
                        value: 8,
                        message: "Password must have 8 characters",
                      },
                    })}
                    {...INPUT_STYLE}
                  />
                </Flex>
                <FormErrorMessage>{errors.passwordNew?.message}</FormErrorMessage>
              </FormControl>
            </Box>

            <Box pt="5">
              <FormControl isInvalid={!!errors.passwordAgain}>
                <FormLabel>Repeat New Password</FormLabel>
                <Flex>
                  <PasswordInput
                    {...register("passwordAgain", {
                      required: {
                        value: true,
                        message: "Retyping password is required",
                      },
                      minLength: {
                        value: 8,
                        message: "Password must have 8 characters",
                      },
                    })}
                    {...INPUT_STYLE}
                  />
                </Flex>
                <FormErrorMessage>{errors.passwordAgain?.message}</FormErrorMessage>
              </FormControl>
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
