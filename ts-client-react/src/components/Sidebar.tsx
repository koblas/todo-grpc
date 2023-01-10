import React, { useState } from "react";
import { Flex, IconButton, Divider, Box, Button, Icon, Spinner } from "@chakra-ui/react";
import {
  BsGear,
  BsCalendar,
  BsPerson,
  BsHouseDoor,
  BsBoxArrowRight,
  BsCurrencyDollar,
  BsBriefcase,
  BsList,
  BsChevronLeft,
} from "react-icons/bs";
import { useNavigate, useLocation } from "react-router";
import NavItem from "./NavItem";
import { useUser } from "../hooks/data/user";

function Username() {
  const { user } = useUser();

  return <span>{user ? user.name : "Unknown"}</span>;
}

// https://github.com/bjcarlson42/chakra-left-responsive-navbar

export default function Sidebar() {
  const { pathname } = useLocation();
  const navigate = useNavigate();
  const [isExpanded, setExpanded] = useState(true);
  const [showExpando, setShowExpando] = useState(false);
  const { user } = useUser();

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
            mt={4}
            _hover={{ background: "none" }}
            icon={<BsList />}
            onClick={() => {
              setExpanded(!isExpanded);
            }}
          />
          <NavItem
            expanded={isExpanded}
            icon={BsHouseDoor}
            active={pathname === "/todo"}
            onClick={() => {
              navigate("/todo");
            }}
          >
            Dashboard
          </NavItem>
          <NavItem
            expanded={isExpanded}
            icon={BsCalendar}
            active={pathname === "/upload"}
            onClick={() => {
              navigate("/upload");
            }}
          >
            Upload
          </NavItem>
          <NavItem expanded={isExpanded} icon={BsPerson}>
            Clients
          </NavItem>
          <NavItem expanded={isExpanded} icon={BsCurrencyDollar}>
            Stocks
          </NavItem>
          <NavItem expanded={isExpanded} icon={BsBriefcase}>
            Reports
          </NavItem>
        </Flex>

        <Flex flexDir="column" w="100%" alignItems={isExpanded ? "flex-start" : "center"}>
          <Divider display={isExpanded ? "flex" : "none"} />
          <NavItem expanded={isExpanded} avatar={user?.avatar_url ?? "avatar-1.jpg"}>
            <React.Suspense fallback={<Spinner />}>
              <Username />
            </React.Suspense>
          </NavItem>
          <NavItem
            expanded={isExpanded}
            icon={BsGear}
            active={pathname.startsWith("/settings")}
            onClick={() => {
              navigate("/settings/");
            }}
          >
            Settings
          </NavItem>
          <NavItem
            expanded={isExpanded}
            icon={BsBoxArrowRight}
            onClick={() => {
              navigate("/auth/logout");
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
        />
        <Button h="100%" width="24px" p="0" border="0" bgColor="transparent" cursor="ew-resize" minWidth="0">
          <Box cursor="ew-resize" h="100%" w="2px" transition="background-color 200ms" />
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
              <Icon as={BsChevronLeft} strokeWidth="1" color="rgb(107, 119, 140)" w={6} h={6} />
            </Box>
            <Box position="absolute" left="-8px" right="-12px" bottom="-8px" top="-8px" />
          </Button>
        </Box>
      </Box>
    </Flex>
  );
}
