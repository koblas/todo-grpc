import React, { useCallback, useEffect, useState } from "react";
import { Box, Spinner, Flex } from "@chakra-ui/react";
import { useMessages } from "../hooks/data/message";
import { Messages, Message } from "../components/chat/Messages";
import { Header } from "../components/chat/Header";
import { Footer } from "../components/chat/Footer";
import { Divider } from "../components/chat/Divider";

const baseMessage: Message[] = [
  { id: "1", from: "computer", text: "Hi, My Name is HoneyChat" },
  { id: "2", from: "me", text: "Hey there" },
  { id: "3", from: "me", text: "Myself Ferin Patel" },
  {
    id: "4",
    from: "computer",
    text: "Nice to meet you. You can send me message and i'll reply you with same message.",
  },
];

function Chat() {
  const roomId = "xyzzy";
  const { messages, mutations } = useMessages();
  const [listMessages] = mutations.useListMessages();
  const [addMessage] = mutations.useAddMessage();
  // const [messages, setMessages] = useState(baseMessage);
  const [inputMessage, setInputMessage] = useState("");

  useEffect(() => {
    const timer = setTimeout(() => {
      listMessages({ roomId });
    }, 1);
    return () => clearTimeout(timer);
  }, []);

  const handleSendMessage = useCallback(
    (msg: string) => {
      if (!msg.trim().length) {
        return;
      }
      addMessage({
        roomId,
        text: msg,
      });

      // addMessage("me", msg);
      setInputMessage("");

      // setTimeout(() => {
      //   addMessage("computer", msg);
      // }, 1000);
    },
    [addMessage],
  );

  return (
    <Flex w="100%" h="100vh" justify="center" align="center">
      <Flex w={["100%", "100%", "40%"]} h="90%" flexDir="column">
        <Header />
        <Divider />
        <Messages messages={messages} />
        <Divider />
        <Footer inputMessage={inputMessage} setInputMessage={setInputMessage} handleSendMessage={handleSendMessage} />
      </Flex>
    </Flex>
  );
}

export function MessagePage() {
  return (
    <Box w="100%" p="8" bgColor="gray.100">
      <React.Suspense fallback={<Spinner />}>
        <Chat />
      </React.Suspense>
    </Box>
  );
}