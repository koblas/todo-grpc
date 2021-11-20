import { InputGroup, Input, InputRightElement, IconButton, Icon, InputProps } from "@chakra-ui/react";
import React, { useState } from "react";
import { FiEyeOff, FiEye } from "react-icons/fi";

export const PasswordInput = React.forwardRef<HTMLInputElement, InputProps>((props, ref) => {
  const [passwordVisible, setPasswordVisible] = useState(false);

  function onPasswordVisible() {
    setPasswordVisible(!passwordVisible);
  }

  return (
    <InputGroup>
      <Input ref={ref} type={passwordVisible ? "text" : "password"} {...props} />
      <InputRightElement>
        <IconButton
          aria-label="view password"
          h="1.75rem"
          size="sm"
          onClick={() => onPasswordVisible()}
          icon={<Icon as={passwordVisible ? FiEyeOff : FiEye} />}
        />
      </InputRightElement>
    </InputGroup>
  );
});

export default PasswordInput;
