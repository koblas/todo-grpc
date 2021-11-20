import React, { PropsWithChildren } from "react";
import { Flex, Stack, Box, Spinner } from "@chakra-ui/react";

export default function AuthWrapper({ children, loading }: PropsWithChildren<{ loading?: boolean }>) {
  return (
    <Stack minH="100vh" direction={{ base: "column", md: "row" }}>
      <Flex p={8} flex={1} align="center" justify="center">
        {loading ? <Spinner /> : children}
      </Flex>
      <Flex flex={1}>
        <Box
          width="100%"
          bgImage="url(https://source.unsplash.com/random)"
          bgRepeat="no-repeat"
          bgPosition="center"
          // bgColor={(t) => (t.palette.mode === "light" ? t.palette.grey[50] : t.palette.grey[900])},
          bgSize="cover"
        />
      </Flex>
    </Stack>
  );
}
