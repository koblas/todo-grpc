import React, { useCallback } from "react";
import { Heading, Box, Spinner, Flex } from "@chakra-ui/react";
import { useDropzone } from "react-dropzone";
import { useUploadFile } from "../hooks/data/file";

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

function Dropzone() {
  const { useUploadSend, useUploadUrl } = useUploadFile();

  const [uploaderUrl] = useUploadUrl();
  const [uploaderFile] = useUploadSend();
  const onDrop = useCallback((acceptedFiles: File[]) => {
    if (acceptedFiles.length !== 0) {
      uploaderUrl(
        { type: "profile_image" },
        {
          onCompleted(data) {
            console.log("UPLOAD URL SUCCESS", data.url);
            uploaderFile(
              {
                url: data.url,
                file: acceptedFiles[0],
              },
              {
                onCompleted(data2) {
                  console.log("UPLOAD SUCCESS", data2.id);
                },
                onError(err) {
                  console.log("UPLOAD FAILED", err);
                },
              },
            );
          },
          onError(err) {
            console.log("UPLOAD FAILED", err);
          },
        },
      );
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

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

  // return (
  //   <section className="container">
  //     <div {...getRootProps({ className: "dropzone" })}>
  //       <input {...getInputProps()} />
  //       <p>Drag 'n' drop some files here, or click to select files</p>
  //     </div>
  //     <aside>
  //       <h4>Files</h4>
  //       <ul>{files}</ul>
  //     </aside>
  //   </section>
  // );
}

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
