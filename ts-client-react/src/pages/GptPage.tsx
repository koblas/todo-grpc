import React, { useState } from "react";
import {
  Heading,
  Box,
  Spinner,
  Flex,
  Textarea,
  FormControl,
  FormLabel,
  FormErrorMessage,
  Button,
} from "@chakra-ui/react";
import { useForm } from "react-hook-form";
import { useGptApi } from "../hooks/data/gptapi";
import { FetchError } from "../rpc/utils";

const INPUT_STYLE = {
  bg: "white",
  borderColor: "gray.400",
};

type FormFields = {
  text: string;
};

export function GptBox() {
  const { useTextApi } = useGptApi();
  const [resultText, setResultText] = useState("");
  const [isSubmitting, setSubmitting] = useState(false);
  const [callTextApi] = useTextApi();
  const {
    register,
    handleSubmit,
    setError,
    formState: { errors },
  } = useForm({
    defaultValues: {
      text: "",
    },
  });

  function onSubmit(data: FormFields) {
    setSubmitting(true);
    callTextApi(data, {
      onCompleted(result) {
        setResultText(result.text);
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
      onFinished() {
        setSubmitting(false);
      },
    });
  }

  return (
    <Box w="100%" p="8" bgColor="gray.100">
      <Box w="100%" bgColor="white" p="5">
        <Heading as="h3" size="xl" textColor="gray.800" textAlign="center" fontWeight="light" padding="5">
          Gpt Testing
        </Heading>

        <FormControl isInvalid={!!errors.text}>
          <FormLabel>Prompt</FormLabel>
          <Flex>
            <Textarea {...register("text")} {...INPUT_STYLE} />
          </Flex>
          <FormErrorMessage>{errors.text?.message}</FormErrorMessage>
        </FormControl>

        <Flex pt="5" justify="left">
          <Box>{resultText}</Box>
        </Flex>

        <Flex pt="5" justify="right">
          <Button mt={4} colorScheme="teal" isLoading={isSubmitting} type="submit" onClick={handleSubmit(onSubmit)}>
            Send
          </Button>
        </Flex>
      </Box>
    </Box>
  );
}

export function GptPage() {
  return (
    <Box w="100%" p="8" bgColor="gray.100">
      <React.Suspense fallback={<Spinner />}>
        <GptBox />
      </React.Suspense>
    </Box>
  );
}
