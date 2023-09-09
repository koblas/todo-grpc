// Inspired from
//   https://dribbble.com/shots/15186840-Setting-page-example

import React from "react";
import { Route, Routes, useLocation, useNavigate } from "react-router-dom";
import { As, Box, Flex, Heading, Icon, Spinner, Text } from "@chakra-ui/react";

import { ProfileSettings } from "./ProfileSettings";
import { SecuritySettings } from "./SecuritySettings";
import { NotificationSettings } from "./NotificationSettings";
import { NotFoundPage } from "../NotFoundPage";
import { BellIcon, GearIcon, KeyIcon } from "../../components/icons";

function SettingItem({
  active,
  icon,
  title,
  description,
  onClick,
}: React.PropsWithChildren<{
  icon: As;
  title: string;
  description: string;
  active?: boolean;
  onClick?: () => void;
}>) {
  return (
    <Box
      p="2"
      borderBottom="1px"
      borderColor="gray.200"
      bgColor={active ? "gray.100" : "white"}
      onClick={(e) => {
        e.preventDefault();
        onClick?.();
      }}
    >
      <Flex>
        <Icon as={icon} mr={2} mt={1} />
        <Box>
          <Text fontWeight="bold">{title}</Text>
          <Text fontSize="sm" fontWeight="light">
            {description}
          </Text>
        </Box>
      </Flex>
    </Box>
  );
}

export function SettingsPage() {
  const { pathname } = useLocation();
  const navigate = useNavigate();

  return (
    <Flex w="100%">
      <Box w="35%" bgColor="white" paddingTop="4" borderLeft="1px" borderRight="1px" borderColor="gray.200">
        <Box p="2" borderBottom="1px" borderColor="gray.200">
          <Heading size="md" fontWeight="bold">
            Settings
          </Heading>
        </Box>
        <SettingItem
          active={pathname === "/settings/"}
          icon={GearIcon}
          title="Account"
          description="Some descriptive text"
          onClick={() => {
            navigate("/settings/");
          }}
        />
        <SettingItem
          active={pathname === "/settings/security"}
          icon={KeyIcon}
          title="Security"
          description="Some descriptive text"
          onClick={() => {
            navigate("/settings/security");
          }}
        />
        <SettingItem
          active={pathname === "/settings/notifications"}
          icon={BellIcon}
          title="Notifications"
          description="Some descriptive text"
          onClick={() => {
            navigate("/settings/notifications");
          }}
        />
      </Box>
      <Box bgColor="gray.100" w="100%">
        <React.Suspense fallback={<Spinner />}>
          <Routes>
            <Route path="/" element={<ProfileSettings />} />
            <Route path="/security" element={<SecuritySettings />} />
            <Route path="/notifications" element={<NotificationSettings />} />
            <Route path="*" element={<NotFoundPage />} />
          </Routes>
        </React.Suspense>
      </Box>
    </Flex>
  );
}
