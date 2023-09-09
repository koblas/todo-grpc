import React, { useState } from "react";
import { useForm } from "react-hook-form";
import {
  Flex,
  Avatar,
  AvatarBadge,
  Text,
  Box,
  IconButton,
  PopoverTrigger,
  Popover,
  PopoverContent,
  FocusLock,
  PopoverArrow,
  PopoverCloseButton,
  Stack,
  ButtonGroup,
  Button,
  FormControl,
  FormLabel,
  Input,
  useDisclosure,
  FormErrorMessage,
  useToast,
  useMergeRefs,
} from "@chakra-ui/react";
import { useTeamMutations } from "../../hooks/data/team";
import { PlusIcon } from "../icons";

type FormFields = {
  email: string;
};

export function InviteUserPopup({
  firstFieldRef,
  onCancel,
}: {
  firstFieldRef: React.Ref<HTMLElement>;
  onCancel: () => void;
}) {
  const [isSubmitting, setSubmitting] = useState(false);
  const toast = useToast();
  const mutations = useTeamMutations();
  const [teamInviteMember] = mutations.useInviteUser();
  const {
    register,
    handleSubmit,
    setError,
    formState: { errors },
  } = useForm<FormFields>({
    defaultValues: {
      email: "",
    },
  });
  // const [] = mutations.use

  function onSubmit(data: FormFields) {
    setSubmitting(true);
    // TODO
    teamInviteMember(
      { team_id: "xyzzy", email: data.email },
      {
        onCompleted() {
          toast({
            position: "top",
            title: "User Invited",
            status: "success",
            isClosable: true,
          });
          onCancel();
        },
        onError() {
          setError("email", { message: "Problem sending invite" });
        },
        onFinished() {
          setSubmitting(false);
        },
      },
    );
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <Stack spacing={4}>
        <FormControl isInvalid={!!errors.email}>
          <FormLabel>Email Address</FormLabel>
          <Input
            type="text"
            autoComplete="off"
            placeholder="user@example.com"
            {...register("email", {
              required: {
                value: true,
                message: "Email is required",
              },
              pattern: {
                value: /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}$/i,
                message: "Entered value does not match email format",
              },
            })}
            ref={useMergeRefs(firstFieldRef, register("email").ref)}
          />
          <FormErrorMessage>{errors.email?.message}</FormErrorMessage>
        </FormControl>

        <ButtonGroup display="flex" justifyContent="flex-end">
          <Button isDisabled={isSubmitting} variant="outline" onClick={onCancel}>
            Cancel
          </Button>
          <Button isDisabled={isSubmitting} colorScheme="teal" type="submit" onClick={handleSubmit(onSubmit)}>
            Save
          </Button>
        </ButtonGroup>
      </Stack>
    </form>
  );
}

export function Header() {
  const { onOpen, onClose, isOpen } = useDisclosure();
  const firstFieldRef = React.useRef(null);

  return (
    <Flex w="100%" justifyContent="space-between">
      <Box>
        <Avatar size="lg" name="Dan Abrahmov" src="https://bit.ly/dan-abramov">
          <AvatarBadge boxSize="1.25em" bg="green.500" />
        </Avatar>
        <Flex flexDirection="column" mx="5" justify="center">
          <Text fontSize="lg" fontWeight="bold">
            Ferin Patel
          </Text>
          <Text color="green.500">Online</Text>
        </Flex>
      </Box>
      <Box>
        <Popover
          isOpen={isOpen}
          initialFocusRef={firstFieldRef}
          onOpen={onOpen}
          onClose={onClose}
          placement="right"
          closeOnBlur={false}
        >
          <PopoverTrigger>
            <IconButton colorScheme="blue" size="md" aria-label="Invite User" icon={<PlusIcon />} onClick={onOpen} />
          </PopoverTrigger>
          <PopoverContent p={5}>
            <FocusLock persistentFocus={false}>
              <PopoverArrow />
              <PopoverCloseButton />
              <InviteUserPopup firstFieldRef={firstFieldRef} onCancel={onClose} />
            </FocusLock>
          </PopoverContent>
        </Popover>
      </Box>
    </Flex>
  );
}
