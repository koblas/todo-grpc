/* eslint-disable react/require-default-props */
import classNames from "classnames";
import React, { forwardRef } from "react";
import { colors, Colors } from "../shared";

interface Props extends React.HTMLAttributes<HTMLHeadingElement> {
  size?: "4xl" | "3xl" | "2xl" | "xl" | "lg" | "md" | "sm" | "xs";
  as?: "h1" | "h2" | "h3" | "h4" | "h5" | "h6";
  color?: Colors;
  className?: string;
}

export const Heading = forwardRef<HTMLHeadingElement, Props>(
  ({ children, size = "4xl", as = "h1", color = "gray", className, ...rest }: Props, ref) => {
    const names = classNames(
      colors[color],
      "font-bold leading-normal mt-0 mb-2",
      {
        "text-6xl": size === "4xl",
        "text-5xl": size === "3xl",
        "text-4xl": size === "2xl",
        "text-3xl": size === "xl",
        "text-2xl": size === "lg",
        "text-xl": size === "md",
        "text-lg": size === "sm",
        "text-md": size === "xs",
      },
      className,
    );

    return React.createElement(as, { className: names, ...rest, ref }, children);
  },
);

Heading.displayName = "Heading";
