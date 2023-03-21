import React from "react";
import { Heading, Box, Spinner } from "@chakra-ui/react";
import { Dropzone } from "../components/Dropzone";

export function UploadDetail() {
  return (
    <Box w="100%" p="8" bgColor="gray.100">
      <Box w="100%" bgColor="white" p="5">
        <Heading as="h3" size="xl" textColor="gray.800" textAlign="center" fontWeight="light" padding="5">
          Upload File
        </Heading>
        <Box>
          <Dropzone />
        </Box>
      </Box>
    </Box>
  );
}

export function UploadPage() {
  return (
    <Box w="100%" p="8" bgColor="gray.100">
      <React.Suspense fallback={<Spinner />}>
        <UploadDetail />
      </React.Suspense>
    </Box>
  );
}
