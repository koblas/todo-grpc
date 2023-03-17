import * as React from "react";
import { Box, Flex, StyleConfig, useStyleConfig } from "@chakra-ui/react";
import { EditorContent, Editor } from "@tiptap/react";
import { Toolbar } from "./Toolbar";

export type Menu = null | "fixed"; // | "bubble" | "floating";

export interface EditorProps {
  editor: Editor | null;
  menu?: Menu;
}

const styleConfig: StyleConfig = {
  baseStyle: ({ theme, colorMode }) => ({
    borderColor: colorMode === "dark" ? theme.colors.whiteAlpha[900] : theme.colors.gray[300],
    borderWidth: theme.borders["2px"],
    rounded: theme.radii.base,
  }),
};

export function RichTextEditor({ editor, menu = "fixed" }: EditorProps) {
  const { borderColor, borderWidth } = useStyleConfig("RichTextEditor", {
    styleConfig,
  });

  if (!editor) {
    return null;
  }

  return (
    <Flex border={`${borderWidth} ${borderColor}`} direction="column" maxH="md" borderRadius="md">
      {menu === "fixed" ? <Toolbar editor={editor} /> : null}
      <Box
        borderTop={`${borderWidth} ${borderColor}`}
        p=".25rem 0.5rem"
        flex="1 1 auto"
        overflowX="hidden"
        overflowY="auto"
      >
        <EditorContent width="100%" height="100%" editor={editor} />
      </Box>
    </Flex>
  );
}
