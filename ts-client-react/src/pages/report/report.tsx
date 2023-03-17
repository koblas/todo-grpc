import React, { useState } from "react";
import {
  Heading,
  Box,
  Flex,
  FormControl,
  FormLabel,
  FormErrorMessage,
  Button,
  Checkbox,
  VStack,
  Radio,
  Input,
  Select,
  RadioGroup,
  IconButton,
  FormHelperText,
} from "@chakra-ui/react";
import { StarterKit } from "@tiptap/starter-kit";
import { useEditor } from "@tiptap/react";
import { DeleteIcon } from "@chakra-ui/icons";
import { useFieldArray, useForm } from "react-hook-form";
import { RichTextEditor } from "../../components/chipchap";
import { useGptApi } from "../../hooks/data/gptapi";
import { Dropzone } from "../../components/Dropzone";
import { FetchError } from "../../rpc/utils";

const INPUT_STYLE = {
  bg: "white",
  borderColor: "gray.400",
};

type FormFields = {
  // text: string;
  description: string;

  submitter: {
    kind: "anonymous" | "protected" | "shared" | null;
    name: string;
    email: string;
    phone: string;
    contact_preference: string;
    email_on_change: boolean;
  };
  individuals: { name: string }[];

  when: {
    date: string;
    time: string;
    approx: string;
  };
  relationship: {
    kind: "employee" | "former" | "other" | null;
    text: string;
  };
  location: {
    kind: "company" | "other" | null;
    name: string;
    address: string;
    city: string;
    region: string;
    postal: string;
    country: string;
  };
};

export function ReportContents() {
  const editor = useEditor({
    extensions: [
      StarterKit.configure({
        // history: false,
        // document: false,
      }),
    ],
    content: "",
  });
  const { useTextApi } = useGptApi();
  const [resultText, setResultText] = useState("");
  const [isSubmitting, setSubmitting] = useState(false);
  const [callTextApi] = useTextApi();
  const {
    register,
    handleSubmit,
    setError,
    control,
    watch,
    formState: { errors },
  } = useForm({
    defaultValues: {
      description: "",
      submitter: {
        name: "",
        email: "",
        phone: "",
        contact_preference: "email",
        email_on_change: false,
        kind: null,
      },
      individuals: [{ name: "" }],
      when: {
        date: "",
        time: "",
        approx: "",
      },
      relationship: {
        kind: null,
        text: "",
      },
      location: {
        kind: null,
        name: "",
        address: "",
        city: "",
        region: "",
        postal: "",
        country: "",
      },
    } as FormFields,
  });
  const {
    fields,
    append: appendIndividuals,
    remove: removeIndividuals,
  } = useFieldArray({
    control,
    name: "individuals",
  });

  const canContact = watch("submitter.kind");
  const showExtraContact = ["shared", "protected"].includes(canContact ?? "");
  const requireEmail =
    (showExtraContact && watch("submitter.contact_preference") === "email") || watch("submitter.email_on_change");
  const requireRelationshipText = watch("relationship.kind") === "other";

  function onSubmit(data: FormFields) {
    setSubmitting(true);
    // Manual validation because the hook form validation is "setup time"
    const text = editor?.getText();
    console.log("SUBMIT", { ...data, text });
    setSubmitting(false);
    // callTextApi(data, {
    //   onCompleted(result) {
    //     setResultText(result.text);
    //   },
    //   onError(error) {
    //     if (error instanceof FetchError) {
    //       const { code, msg, argument } = error.getInfo();

    //       if (code === "invalid_argument" && msg && argument) {
    //         setError(argument as keyof FormFields, { message: msg });
    //       }
    //     }
    //     setSubmitting(false);
    //   },
    //   onFinished() {
    //     setSubmitting(false);
    //   },
    // });
  }
  console.log("ERRORS", errors);

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <Box w="100%" p="8" bgColor="gray.100">
        <Box w="100%" bgColor="white" p="5">
          <Heading as="h3" size="xl" textColor="gray.800" textAlign="center" fontWeight="light" padding="5">
            Tell Us What Happened
          </Heading>

          <Box pt="5" alignItems="start" justifyContent="left" justifyItems="left">
            <Heading as="h3" size="md" textColor="gray.800" fontWeight="light" padding="5">
              Issue details and involved parties
            </Heading>
            <FormControl isInvalid={!!errors.description}>
              <FormLabel>Description</FormLabel>
              {/* <Flex> */}
              {/* <Textarea {...register("text")} {...INPUT_STYLE} /> */}
              {/* </Flex> */}
              <RichTextEditor editor={editor} menu="fixed" />
              <FormErrorMessage>{errors.description?.message}</FormErrorMessage>
            </FormControl>

            <FormControl isInvalid={!!errors.individuals}>
              <FormLabel>Who was involved</FormLabel>
              {fields.map((field, number) => (
                <Flex key={field.id}>
                  <Input
                    key={field.id}
                    {...register(`individuals.${number}.name`, { minLength: 2, maxLength: 200 })}
                    {...INPUT_STYLE}
                  />
                  {number !== 0 && (
                    <IconButton
                      aria-label="delete entry"
                      icon={<DeleteIcon />}
                      onClick={() => removeIndividuals(number)}
                    />
                  )}
                </Flex>
              ))}
              <FormErrorMessage>{errors.individuals?.message}</FormErrorMessage>
            </FormControl>
            <Button onClick={() => appendIndividuals({ name: "" })}>Add Individual</Button>
          </Box>

          <Flex pt="5" justify="left">
            <Box>{resultText}</Box>
          </Flex>
          <Box pt="5" alignItems="start" justifyContent="left" justifyItems="left">
            <Heading as="h3" size="md" textColor="gray.800" fontWeight="light" padding="5">
              Date and location
            </Heading>

            <FormControl>
              <RadioGroup>
                <Radio value="company" {...register("location.kind")}>
                  Company Location
                </Radio>
                <Radio value="other" {...register("location.kind")}>
                  Other Location
                </Radio>
              </RadioGroup>
            </FormControl>
            <li>Company location</li>
            <li>Other location</li>

            <li>Date with optional time</li>
            <li>approximate time (e.g. last week)</li>
          </Box>

          <Box pt="5" alignItems="start" justifyContent="left" justifyItems="left">
            <Heading as="h3" size="md" textColor="gray.800" fontWeight="light" padding="5">
              Photo or file uploads
            </Heading>
            <Dropzone />
          </Box>

          <Box pt="5" alignItems="start" justifyContent="left" justifyItems="left">
            <Heading as="h3" size="md" textColor="gray.800" fontWeight="light" padding="5">
              Tell Us About Yourself
            </Heading>

            <FormControl isInvalid={!!errors.submitter?.kind}>
              <RadioGroup>
                <VStack justifyContent="left" alignItems="start">
                  <FormErrorMessage>{errors.submitter?.kind?.message}</FormErrorMessage>
                  <Radio
                    value="shared"
                    {...register("submitter.kind", { required: "Please select one of these option" })}
                  >
                    Share your name and contact information
                  </Radio>
                  <Radio value="protected" {...register("submitter.kind")}>
                    Remain anonymous toward the organization
                  </Radio>
                  <Radio value="anonymous" {...register("submitter.kind")}>
                    Remain completly anonymous
                  </Radio>
                </VStack>
              </RadioGroup>
            </FormControl>

            <Box bg="gray.50" padding={5} hidden={!canContact}>
              <FormControl isInvalid={!!errors.submitter?.name} hidden={!showExtraContact} isRequired>
                <FormLabel>Name</FormLabel>
                <Flex>
                  <Input
                    type="text"
                    {...register("submitter.name", {
                      validate: {
                        required: (value) => (showExtraContact && !value ? "Name is required" : undefined),
                      },
                      minLength: 2,
                      maxLength: 200,
                    })}
                    {...INPUT_STYLE}
                  />
                </Flex>
                <FormErrorMessage>{errors.submitter?.name?.message}</FormErrorMessage>
              </FormControl>
              <FormControl
                isInvalid={!!errors.submitter?.email}
                hidden={!requireEmail && !showExtraContact}
                isRequired={requireEmail}
              >
                <FormLabel>Email</FormLabel>
                <Flex>
                  <Input
                    type="text"
                    {...register("submitter.email", {
                      validate: {
                        required: (value) => (requireEmail && !value ? "Email is required" : undefined),
                      },
                      maxLength: 200,
                    })}
                    {...INPUT_STYLE}
                  />
                </Flex>
                <FormErrorMessage>{errors.submitter?.email?.message}</FormErrorMessage>
              </FormControl>
              <FormControl
                isInvalid={!!errors.submitter?.phone}
                hidden={!showExtraContact}
                isRequired={showExtraContact}
              >
                <FormLabel>Phone number</FormLabel>
                <Flex>
                  <Input
                    type="text"
                    {...register("submitter.phone", {
                      validate: {
                        required: (value) => (showExtraContact && !value ? "Phone number is required" : undefined),
                      },
                      maxLength: 100,
                    })}
                    {...INPUT_STYLE}
                  />
                </Flex>
                <FormErrorMessage>{errors.submitter?.phone?.message}</FormErrorMessage>
              </FormControl>

              <FormControl isInvalid={!!errors.submitter?.contact_preference} hidden={!showExtraContact}>
                <FormLabel>Preferred contact method</FormLabel>
                <Select {...register("submitter.contact_preference")}>
                  <option value="email">Email</option>
                  <option value="phone">Phone</option>
                </Select>
              </FormControl>

              <FormControl>
                <Checkbox {...register("submitter.email_on_change")}>
                  I would like to receive updates via email
                </Checkbox>
              </FormControl>
            </Box>
          </Box>

          <Box>
            <Heading as="h3" size="md" textColor="gray.800" fontWeight="light" padding="5">
              Your Relationship to the Organization
            </Heading>
            <FormControl isInvalid={!!errors.relationship?.kind}>
              <RadioGroup>
                <VStack justifyContent="left" alignItems="start">
                  <FormErrorMessage>{errors.relationship?.kind?.message}</FormErrorMessage>
                  <Radio
                    value="employee"
                    {...register("relationship.kind", { required: "Please select one of these option" })}
                  >
                    I am a current employee
                  </Radio>
                  <Radio value="former" {...register("relationship.kind")}>
                    I am a former employee
                  </Radio>
                  <Radio value="other" {...register("relationship.kind")}>
                    Non employee (this includes contractor, student, supplier, partner, etc.)
                  </Radio>
                </VStack>
              </RadioGroup>
            </FormControl>

            <FormControl
              isRequired={requireRelationshipText}
              isInvalid={!!errors.relationship?.text}
              hidden={!requireRelationshipText}
            >
              <FormLabel>What is your relationship to this organization</FormLabel>
              <Input
                type="text"
                {...register("relationship.text", {
                  validate: {
                    required: (value) => (requireRelationshipText && !value ? "This field is required" : undefined),
                  },
                  maxLength: 100,
                })}
                {...INPUT_STYLE}
              />
              <FormErrorMessage>{errors.relationship?.text?.message}</FormErrorMessage>
              <FormHelperText>
                The answer you provide will be shared with the organization regardless of the anonymity preference you
                selected above
              </FormHelperText>
            </FormControl>
          </Box>

          <Flex pt="5" justify="right">
            <Button mt={4} colorScheme="teal" isLoading={isSubmitting} type="submit" onClick={handleSubmit(onSubmit)}>
              Send
            </Button>
          </Flex>
        </Box>
      </Box>
    </form>
  );
}
