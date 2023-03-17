import * as React from "react";
import { IconButton, IconButtonProps, Tooltip } from "@chakra-ui/react";

interface ToolbarButtonProps extends Omit<IconButtonProps, "aria-label"> {
  label: string;
}

export function ToolbarButton({ label, ...rest }: ToolbarButtonProps) {
  return (
    <Tooltip label={label}>
      <IconButton variant="ghost" colorScheme="gray" aria-label={label} {...rest} />
    </Tooltip>
  );
}
