import React, { useState } from "react";
import { Flex, IconButton, Divider, Avatar, Heading, Box, Button, Icon } from "@chakra-ui/react";
import {
  FiMenu,
  FiHome,
  FiCalendar,
  FiUser,
  FiDollarSign,
  FiBriefcase,
  FiSettings,
  FiChevronLeft,
  FiLogOut,
} from "react-icons/fi";
import { IoPawOutline } from "react-icons/io5";
import { useHistory, useLocation } from "react-router";
import NavItem from "./NavItem";

// https://github.com/bjcarlson42/chakra-left-responsive-navbar

export default function Sidebar() {
  const { pathname } = useLocation();
  const history = useHistory();
  const [isExpanded, setExpanded] = useState(true);
  const [showExpando, setShowExpando] = useState(false);

  return (
    <Flex pos="sticky" h="100vh">
      <Flex
        // w={isExpanded ? "180px" : "50px"}
        flexDir="column"
        justifyContent="space-between"
      >
        <Flex bgColor="white" flexDir="column" w="100%" alignItems={isExpanded ? "flex-start" : "center"} as="nav">
          <IconButton
            background="none"
            aria-label="Toggle sidebar"
            mt={5}
            _hover={{ background: "none" }}
            icon={<FiMenu />}
            onClick={() => {
              setExpanded(!isExpanded);
            }}
          />
          <NavItem
            expanded={isExpanded}
            icon={FiHome}
            active={pathname === "/todo"}
            onClick={() => {
              history.push("/todo");
            }}
          >
            Dashboard
          </NavItem>
          <NavItem
            expanded={isExpanded}
            icon={FiCalendar}
            active={pathname === "/calendar"}
            onClick={() => {
              history.push("/calendar");
            }}
          >
            Calendar
          </NavItem>
          <NavItem expanded={isExpanded} icon={FiUser}>
            Clients
          </NavItem>
          <NavItem expanded={isExpanded} icon={IoPawOutline}>
            Animals
          </NavItem>
          <NavItem expanded={isExpanded} icon={FiDollarSign}>
            Stocks
          </NavItem>
          <NavItem expanded={isExpanded} icon={FiBriefcase}>
            Reports
          </NavItem>
        </Flex>

        <Flex flexDir="column" w="100%" alignItems={isExpanded ? "flex-start" : "center"}>
          <Divider display={isExpanded ? "flex" : "none"} />
          <NavItem expanded={isExpanded} avatar="avatar-1.jpg">
            <Heading as="h3" size="sm" style={{ textOverflow: "truncate" }}>
              Sylwia Weller
            </Heading>
          </NavItem>
          <NavItem
            expanded={isExpanded}
            icon={FiSettings}
            active={pathname.startsWith("/settings")}
            onClick={() => {
              history.push("/settings");
            }}
          >
            Settings
          </NavItem>
          <NavItem
            expanded={isExpanded}
            icon={FiLogOut}
            onClick={() => {
              history.push("/auth/logout");
            }}
          >
            Logout
          </NavItem>
        </Flex>
      </Flex>
      {/* Sidebar handle */}
      <Box
        pos="absolute"
        left="100%"
        top="0"
        bottom="0"
        onMouseOut={() => {
          setShowExpando(false);
        }}
        onMouseOver={() => {
          setShowExpando(true);
        }}
      >
        <Box
          pos="absolute"
          left="-1px"
          top="0"
          bottom="0"
          opacity="0.5"
          width="3px"
          bgColor="linear-gradient(to left, rgba(0,0,0,0.2) 0px, rgba(0,0,0,0.2) 1px, rgba(0,0,0,0.1) 1px, rgba(0,0,0,0) 100% )"
          transitionDuration="0.22s"
          transitionProperty="left,opacity,width"
          transitionTimingFunction="cubic-bezier(0.2,0,0,1)"
        ></Box>
        <Button h="100%" width="24px" p="0" b="0" bgColor="transparent" cursor="ew-resize" minWidth="0">
          <Box cursor="ew-resize" h="100%" w="2px" transition="background-color 200ms"></Box>
        </Button>
        <Box>
          <Button
            bgColor="white"
            transform={isExpanded ? "" : "rotate(180deg)"}
            opacity={showExpando ? 100 : 0}
            position="absolute"
            top="32px"
            left="0"
            display="block"
            p="0"
            outline="0"
            lineHeight="0"
            border="0"
            color="#6B778C"
            borderRadius="50%"
            boxShadow="0 0 0 1px rgba(9,30,66,0.08),0 2px 4px 1px rgba(9,30,66,0.08)"
            cursor="pointer"
            height="24px"
            width="24px"
            minWidth="0px"
            transformOrigin="7px"
            transition="background-color 100ms linear, color 100ms linear, opacity 350ms cubic-bezier(0.2,0,0,1)"
            alignItems="flex-start"
            onClick={() => {
              setExpanded(!isExpanded);
            }}
          >
            <Box as="span" display="inline-block" flexShrink={0} lineHeight="1">
              <Icon as={FiChevronLeft} strokeWidth="1" color="rgb(107, 119, 140)" w={6} h={6} />
            </Box>
            <Box position="absolute" left="-8px" right="-12px" bottom="-8px" top="-8px"></Box>
          </Button>
        </Box>
      </Box>
    </Flex>
  );
}
