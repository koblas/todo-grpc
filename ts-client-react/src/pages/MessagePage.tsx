import React, { useCallback, useEffect, useState } from "react";
import { Box, Spinner, Flex } from "@chakra-ui/react";
import { useMessages } from "../hooks/data/message";
import { Messages } from "../components/chat/Messages";
import { Header } from "../components/chat/Header";
import { Footer } from "../components/chat/Footer";
import { Divider } from "../components/chat/Divider";

// const baseMessage: Message[] = [
//   { id: "1", from: "computer", text: "Hi, My Name is HoneyChat" },
//   { id: "2", from: "me", text: "Hey there" },
//   { id: "3", from: "me", text: "Myself Ferin Patel" },
//   {
//     id: "4",
//     from: "computer",
//     text: "Nice to meet you. You can send me message and i'll reply you with same message.",
//   },
// ];

function Chat() {
  const [roomId, setRoomId] = useState("");
  const { messages, mutations } = useMessages();
  const [roomJoin] = mutations.useRoomJoin();
  const [listMessages] = mutations.useListMessages();
  const [addMessage] = mutations.useAddMessage();
  // const [messages, setMessages] = useState(baseMessage);
  const [inputMessage, setInputMessage] = useState("");

  console.log("ROOM ID", roomId, messages);

  useEffect(() => {
    const timer = setTimeout(() => {
      roomJoin(
        { name: "xyzzy" },
        {
          onCompleted({ room }) {
            console.log("COMPLETED ", room.id, roomId);
            setRoomId(room.id);
            listMessages({ room_id: room.id });
          },
        },
      );
    }, 10);
    return () => clearTimeout(timer);
  }, [roomId]);

  const handleSendMessage = useCallback(
    (msg: string) => {
      if (!msg.trim().length) {
        return;
      }
      addMessage({
        room_id: roomId,
        text: msg,
      });

      // addMessage("me", msg);
      setInputMessage("");

      // setTimeout(() => {
      //   addMessage("computer", msg);
      // }, 1000);
    },
    [addMessage, roomId],
  );

  return (
    <Flex w="100%" h="100vh" justify="center" align="center">
      <Flex w={["100%", "100%", "40%"]} h="90%" flexDir="column">
        <Header />
        <Divider />
        <Messages messages={messages.filter((m) => m.room_id === roomId)} />
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
