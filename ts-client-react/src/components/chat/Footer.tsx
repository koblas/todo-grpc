import React from "react";
import { Flex, Input, Button } from "@chakra-ui/react";

type Props = {
  inputMessage: string;
  setInputMessage: (value: string) => void;
  handleSendMessage: (msg: string) => void;
};

export function Footer({ inputMessage, setInputMessage, handleSendMessage }: Props) {
  return (
    <Flex w="100%" mt="5">
      <Input
        placeholder="Type Something..."
        border="none"
        borderRadius="none"
        _focus={{
          border: "1px solid black",
        }}
        onKeyPress={(e) => {
          if (e.key === "Enter" && inputMessage) {
            handleSendMessage(inputMessage);
          }
        }}
        value={inputMessage}
        onChange={(e) => setInputMessage(e.target.value)}
      />
      <Button
        bg="black"
        color="white"
        borderRadius="none"
        _hover={{
          bg: "white",
          color: "black",
          border: "1px solid black",
        }}
        disabled={inputMessage.trim().length <= 0}
        onClick={() => (inputMessage ? handleSendMessage(inputMessage) : true)}
      >
        Send
      </Button>
    </Flex>
  );
}

export default Footer;
