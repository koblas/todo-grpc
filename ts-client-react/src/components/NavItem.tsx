import React from "react";
import { Flex, Text, Icon, Link, Menu, MenuButton, As, Avatar } from "@chakra-ui/react";
// import NavHoverBox from "./NavHoverBox";

type Props = {
  // description?: string;
  active?: boolean;
  // navSize?: string;
  expanded: boolean;
  icon?: As;
  avatar?: string;
  onClick?: () => void;
};

export default function NavItem({
  avatar,
  icon,
  children,
  // description,
  active,
  expanded,
  onClick,
}: React.PropsWithChildren<Props>) {
  return (
    <Flex flexDir="column" w="100%" alignItems={expanded ? "flex-start" : "center"}>
      <Menu placement="right">
        <MenuButton
          w="100%"
          backgroundColor={active ? "#AEC8CA" : undefined}
          p={3}
          borderRadius={8}
          onClick={(e) => {
            e.preventDefault();
            onClick?.();
          }}
          _hover={{ textDecor: "none", backgroundColor: "#AEC8CA" }}
        >
          <Flex alignItems="left">
            {avatar && <Avatar src={avatar} size="sm" w="20px" h="20px" color={active ? "#82AAAD" : "gray.500"} />}
            {icon && <Icon as={icon} fontSize="xl" color={active ? "#82AAAD" : "gray.500"} />}
            <Text as="div" ml={5} align="left" display={expanded ? "flex" : "none"}>
              {children}
            </Text>
          </Flex>
        </MenuButton>
        {/* <MenuList py={0} border="none" w={200} h={200} ml={5}>
          <NavHoverBox title={children} avatar={avatar} icon={icon} description={description} />
        </MenuList> */}
      </Menu>
    </Flex>
  );
}
