import React, { useCallback, useEffect, useState } from "react";
import {
  Box,
  Spinner,
  Flex,
  Text,
  Menu,
  MenuButton,
  MenuList,
  MenuItem,
  IconButton,
  HStack,
  VStack,
} from "@chakra-ui/react";
import { useMessageList, useMessageMutations, useRoomsList } from "../hooks/data/message";
import { Messages } from "../components/chat/Messages";
import { Header } from "../components/chat/Header";
import { Footer } from "../components/chat/Footer";
import { Divider } from "../components/chat/Divider";
import { DeleteIcon, EditIcon, KebabIcon } from "../components/icons";

function ChatMessages({ roomId }: { roomId: string }) {
  const messages = useMessageList(roomId);
  const mutations = useMessageMutations();
  const [addMessage] = mutations.useAddMessage();
  const [inputMessage, setInputMessage] = useState("");

  const handleSendMessage = useCallback(
    (msg: string) => {
      const value = msg.trim();
      if (!value.length) {
        return;
      }
      addMessage({
        room_id: roomId,
        text: value,
      });

      setInputMessage("");
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

function Chat({ roomId }: { roomId: string | null }) {
  if (!roomId) {
    return <Spinner />;
  }

  return (
    <React.Suspense fallback={<Spinner />}>
      <ChatMessages roomId={roomId} />;
    </React.Suspense>
  );
}

function Rooms({ roomId, onRoomChange }: { roomId: string | null; onRoomChange: (id: string) => void }) {
  const rooms = useRoomsList();
  const mutations = useMessageMutations();
  const [roomJoin] = mutations.useRoomJoin();

  useEffect(() => {
    if (roomId === null && rooms.length === 0) {
      // Create a default room
      roomJoin({ name: "Homebase" });
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [roomId, rooms]);

  if (roomId === null && rooms.length !== 0) {
    // Select the first room if no rooms are selected
    onRoomChange(rooms[0].id);
  }

  return (
    <Box w="20%">
      <Text>Rooms</Text>
      <VStack align="normal">
        {rooms.map((room) => (
          <HStack justify="space-between">
            <Text
              display="inline-block"
              fontWeight={roomId === room.id ? "bold" : ""}
              onClick={() => {
                onRoomChange(room.id);
              }}
            >
              {room.name || "MISSING"}
            </Text>

            <Menu>
              <MenuButton as={IconButton} size="xs" aria-label="Options" icon={<KebabIcon />} />
              <MenuList>
                <MenuItem icon={<EditIcon />}>Rename</MenuItem>
                <MenuItem icon={<DeleteIcon />}>Delete</MenuItem>
              </MenuList>
            </Menu>
          </HStack>
        ))}
      </VStack>
    </Box>
  );
}

export function MessagePage() {
  const [roomId, setRoomId] = useState<string | null>(null);

  return (
    <Box w="100%" p="8" bgColor="gray.100">
      <React.Suspense fallback={<Spinner />}>
        <Flex>
          <Rooms
            roomId={roomId}
            onRoomChange={(id) => {
              setRoomId(id);
            }}
          />
          <Chat roomId={roomId} />
        </Flex>
      </React.Suspense>
    </Box>
  );
}
