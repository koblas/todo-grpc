// https://dribbble.com/shots/15186840-Setting-page-example

import React, { useState } from "react";
import {
  Text,
  Heading,
  FormControl,
  FormLabel,
  Input,
  Button,
  Flex,
  Box,
  Grid,
  Textarea,
  useToast,
  FormErrorMessage,
} from "@chakra-ui/react";
import { useForm } from "react-hook-form";
import { useUser } from "../../hooks/data/user";
import { FetchError } from "../../rpc/utils";

const INPUT_STYLE = {
  bg: "white",
  borderColor: "gray.400",
};

type FormFields = {
  name: string;
  about: string;
  email: string;
  country: string;
  language: string;
  phone: string;
};

export function ProfileSettings() {
  const { user, mutations } = useUser();
  const toast = useToast();
  const [updateUser] = mutations.useUpdateUser();

  const [isSubmitting, setSubmitting] = useState(false);
  const {
    register,
    handleSubmit,
    setError,
    formState: { errors },
  } = useForm<FormFields>({
    defaultValues: {
      name: user?.name ?? "",
      email: user?.email ?? "",
    },
  });

  function onSubmit(data: FormFields) {
    setSubmitting(true);
    updateUser(data, {
      onFinished() {
        setSubmitting(false);
      },
      onErrorField(serror: Record<string, string[]>) {
        const fields: (keyof FormFields)[] = ["email", "name", "email", "about", "country", "language", "phone"];
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
            <Heading fontSize="3xl">Account</Heading>
            <Box pt="5">
              <Heading fontSize="lg">Profile</Heading>
              <Text>This information will be displayed publically so be careful what you share</Text>
            </Box>

            <Box pt="5">
              <FormControl isInvalid={!!errors.name}>
                <FormLabel>Name</FormLabel>
                <Flex>
                  <Input
                    type="text"
                    {...register("name", {
                      required: {
                        value: true,
                        message: "Name is required",
                      },
                      minLength: {
                        value: 2,
                        message: "Name must be at least 2 characters",
                      },
                    })}
                    {...INPUT_STYLE}
                  />
                </Flex>
                <FormErrorMessage>{errors.name?.message}</FormErrorMessage>
              </FormControl>
            </Box>

            <Box pt="5">
              <FormControl isInvalid={!!errors.about}>
                <FormLabel>About</FormLabel>
                <Flex>
                  <Textarea {...register("about")} {...INPUT_STYLE} />
                </Flex>
                <FormErrorMessage>{errors.about?.message}</FormErrorMessage>
              </FormControl>
            </Box>

            <Box pt="5">
              <Heading fontSize="lg">Personal Information</Heading>
              <Text>This information will be displayed publically so be careful what you share</Text>
            </Box>

            <Grid pt="5" templateColumns="repeat(auto-fit, minmax(300px, 1fr))" gap="6">
              <FormControl isInvalid={!!errors.email}>
                <FormLabel>Email address</FormLabel>
                <Flex>
                  <Input
                    type="email"
                    {...register("email", {
                      required: {
                        value: true,
                        message: "Email is required",
                      },
                    })}
                    {...INPUT_STYLE}
                  />
                </Flex>
                <FormErrorMessage>{errors.email?.message}</FormErrorMessage>
              </FormControl>

              <FormControl isInvalid={!!errors.phone}>
                <FormLabel>Phone number</FormLabel>
                <Flex>
                  <Input type="text" {...register("phone")} {...INPUT_STYLE} />
                </Flex>
                <FormErrorMessage>{errors.phone?.message}</FormErrorMessage>
              </FormControl>

              <FormControl isInvalid={!!errors.country}>
                <FormLabel>Country</FormLabel>
                <Flex>
                  <Input type="text" {...register("country")} {...INPUT_STYLE} />
                </Flex>
                <FormErrorMessage>{errors.country?.message}</FormErrorMessage>
              </FormControl>

              <FormControl isInvalid={!!errors.language}>
                <FormLabel>Language</FormLabel>
                <Flex>
                  <Input type="text" {...register("language")} {...INPUT_STYLE} />
                </Flex>
                <FormErrorMessage>{errors.language?.message}</FormErrorMessage>
              </FormControl>
            </Grid>

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
