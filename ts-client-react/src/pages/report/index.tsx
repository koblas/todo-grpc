import React, { useState } from "react";
import {
  Heading,
  Box,
  Spinner,
  Flex,
  Textarea,
  FormControl,
  FormLabel,
  FormErrorMessage,
  Button,
} from "@chakra-ui/react";
import { ReportContents } from './report';

export function ReportPage() {
  return (
    <Box w="100%" p="8" bgColor="gray.100">
      <React.Suspense fallback={<Spinner />}>
        <ReportContents />
      </React.Suspense>
    </Box>
  );
}
