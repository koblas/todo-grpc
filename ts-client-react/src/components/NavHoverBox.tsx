import React, { ReactElement } from "react";
import { Flex, Heading, Text, Icon, As, Avatar } from "@chakra-ui/react";

type Props = {
  title?: React.ReactNode;
  icon?: As;
  avatar?: string;
  description?: string;
};

export default function NavHoverBox({ title, avatar, icon, description }: Props) {
  return (
    <>
      <Flex
        pos="absolute"
        mt="calc(100px - 7.5px)"
        ml="-10px"
        width={0}
        height={0}
        borderTop="10px solid transparent"
        borderBottom="10px solid transparent"
        borderRight="10px solid #82AAAD"
      />
      <Flex
        h={200}
        // w={200}
        w="100%"
        flexDir="column"
        alignItems="center"
        justify="center"
        backgroundColor="#82AAAD"
        borderRadius="10px"
        color="#fff"
        textAlign="center"
      >
        {icon && <Icon as={icon} fontSize="3xl" mb={4} />}
        {avatar && <Avatar src={avatar} mb={4} />}
        <Heading size="md" fontWeight="normal">
          {title ?? ""}
        </Heading>
        <Text>{description}</Text>
      </Flex>
    </>
  );
}
