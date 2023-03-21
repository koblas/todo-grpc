import * as Sentry from "@sentry/react";
import React, { useCallback, useRef } from "react";
import { Flex, useToast } from "@chakra-ui/react";
import { useDropzone } from "react-dropzone";
import { useFileListener, useUploadFile } from "../hooks/data/file";

const focusedStyle = {
  borderColor: "#2196f3",
};

const acceptStyle = {
  borderColor: "#00e676",
};

const rejectStyle = {
  borderColor: "#ff1744",
};

const dragStyle = {
  borderColor: "green",
};

export function Dropzone() {
  const { useUploadSend, useUploadUrl } = useUploadFile();
  const toast = useToast();
  const fileIds = useRef([] as string[]);

  const [uploaderUrl] = useUploadUrl();
  const [uploaderFile] = useUploadSend();
  const onDrop = useCallback(
    (acceptedFiles: File[]) => {
      if (acceptedFiles.length !== 0) {
        const file = acceptedFiles[0];

        uploaderUrl(
          { type: "profile_image", contentType: file.type },
          {
            onCompleted(data) {
              fileIds.current.push(data.id);
              uploaderFile(
                {
                  url: data.url,
                  file,
                },
                {
                  onError(err) {
                    Sentry.captureException(err);
                    toast({
                      position: "top",
                      title: "Upload failed",
                      status: "error",
                      isClosable: true,
                    });
                  },
                },
              );
            },
            onError(err) {
              Sentry.captureException(err);
              toast({
                position: "top",
                title: "Upload failed",
                status: "error",
                isClosable: true,
              });
            },
          },
        );
      }
    },
    // eslint-disable-next-line react-hooks/exhaustive-deps
    [fileIds],
  );
  const action = (id: string, error?: string) => {
    if (fileIds.current.includes(id)) {
      if (error) {
        toast({
          position: "top",
          title: "Upload failed",
          status: "error",
          isClosable: true,
        });
      } else {
        toast({
          position: "top",
          title: "File uploaded",
          status: "success",
          isClosable: true,
        });
      }
    }
  };
  useFileListener(action);

  const { acceptedFiles, getRootProps, getInputProps, isDragActive, isFocused, isDragAccept, isDragReject } =
    useDropzone({ onDrop });
  // const { acceptedFiles, getRootProps, getInputProps } = useDropzone();

  const files = acceptedFiles.map((file) => (
    <li key={file.name}>
      {file.name} - {file.size} bytes
    </li>
  ));

  return (
    <>
      <Flex
        w="100%"
        borderRadius="2px"
        border="2px"
        p="8"
        borderColor="#eeeeee"
        borderStyle="dashed"
        color="#bdbdbd"
        bgColor="#fafafa"
        {...getRootProps({
          ...(isFocused ? focusedStyle : {}),
          ...(isDragAccept ? acceptStyle : {}),
          ...(isDragReject ? rejectStyle : {}),
          ...(isDragActive ? dragStyle : {}),
        })}
      >
        <input {...getInputProps()} />
        <p>Drag 'n' drop some files here, or click to select files</p>
      </Flex>
      <span>{files.length !== 0 ? files[0] : null}</span>
    </>
  );
}
