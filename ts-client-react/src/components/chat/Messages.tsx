import React, { useEffect, useRef } from "react";
import { Avatar, Flex, Text } from "@chakra-ui/react";
import { useUser } from "../../hooks/data/user";

export type Message = {
  id: string;
  sender: string;
  text: string;
};

function AlwaysScrollToBottom() {
  const elementRef = useRef<HTMLDivElement | null>(null);

  useEffect(() => elementRef.current?.scrollIntoView());

  return <div ref={elementRef} />;
}

export function Messages({ messages }: { messages: Message[] }) {
  const { user } = useUser();

  return (
    <Flex w="100%" h="80%" overflowY="scroll" flexDirection="column" p="3">
      {messages.map((item) => {
        if (item.sender === user?.id) {
          return (
            <Flex key={item.id} w="100%" justify="flex-end">
              <Flex bg="black" color="white" minW="100px" maxW="350px" my="1" p="3">
                <Text>{item.text}</Text>
              </Flex>
            </Flex>
          );
        }

        return (
          <Flex key={item.id} w="100%">
            <Avatar
              name="Computer"
              src="https://avataaars.io/?avatarStyle=Transparent&topType=LongHairStraight&accessoriesType=Blank&hairColor=BrownDark&facialHairType=Blank&clotheType=BlazerShirt&eyeType=Default&eyebrowType=Default&mouthType=Default&skinColor=Light"
              bg="blue.300"
            />
            <Flex bg="gray.100" color="black" minW="100px" maxW="350px" my="1" p="3">
              <Text>{item.text}</Text>
            </Flex>
          </Flex>
        );
      })}
      <AlwaysScrollToBottom />
    </Flex>
  );
}
