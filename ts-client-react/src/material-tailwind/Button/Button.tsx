import React, { forwardRef } from "react";
import Ripple from "material-ripple-effects";
import classNames from "classnames";
import {
  Colors,
  filledBgActiveColors,
  filledBgColors,
  filledBgFocusColors,
  filledBgHoverColors,
  filledShadowColors,
  filledShadowHoverColors,
  outlineBgActiveColors,
  outlineBgHoverColors,
  outlineBorderColors,
  outlineBorderHoverColors,
  colors,
  outlineTextHoverColors,
} from "../shared";

/* eslint-disable react/require-default-props,react/jsx-props-no-spreading */

interface BtnProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  children: React.ReactNode;
  color?: Colors;
  buttonType?: "outline" | "link" | "filled";
  size?: "sm" | "md" | "lg";
  iconOnly?: boolean;
  rounded?: boolean;
  block?: boolean;
  ripple?: string;
}

export const Button = forwardRef<HTMLButtonElement, BtnProps>(
  (
    {
      children,
      type = "button",
      color = "lightBlue",
      buttonType = "filled",
      size = "md",
      rounded = false,
      iconOnly = false,
      block = false,
      ripple,
      className,
      ...rest
    }: BtnProps,
    ref,
  ) => {
    const rippleEffect = new Ripple();

    const sharedClasses = [
      block && "w-full",
      "flex",
      "items-center",
      "justify-center",
      "gap-1",
      "font-bold",
      "outline-none",
      "uppercase",
      "tracking-wider",
      "focus:outline-none",
      "focus:shadow-none",
      "transition-all",
      "duration-300",
      rounded ? "rounded-full" : "rounded-lg",
    ];

    const buttonFilled = [
      "text-white",
      filledBgColors[color],
      filledBgHoverColors[color],
      filledBgFocusColors[color],
      filledBgActiveColors[color],
      filledShadowColors[color],
      filledShadowHoverColors[color],
    ];

    const buttonOutline = [
      "bg-transparent",
      "border",
      "border-solid",
      "shadow-none",
      colors[color],
      outlineBorderColors[color],
      outlineBgHoverColors[color],
      outlineBorderHoverColors[color],
      outlineTextHoverColors[color],
      outlineBgHoverColors[color],
      outlineBgActiveColors[color],
    ];

    const buttonLink = [
      `bg-transparent`,
      colors[color],
      outlineBgHoverColors[color],
      outlineTextHoverColors[color],
      outlineBgHoverColors[color],
      outlineBgActiveColors[color],
    ];

    const buttonSM = [iconOnly ? "w-8 h-8 p-0 grid place-items-center" : "py-1.5 px-4", "text-xs", "leading-normal"];
    const buttonRegular = [
      iconOnly ? "w-10 h-10 p-0 grid place-items-center" : "py-2.5 px-6",
      "text-xs",
      "leading-normal",
    ];
    const buttonLG = [iconOnly ? "w-12 h-12 p-0 grid place-items-center" : "py-3 px-7", "text-sm", "leading-relaxed"];

    const classValue = classNames(
      sharedClasses,
      size === "sm" && buttonSM,
      size === "md" && buttonRegular,
      size === "lg" && buttonLG,
      buttonType === "outline" && buttonOutline,
      buttonType === "link" && buttonLink,
      buttonType === "filled" && buttonFilled,
      className,
    );

    return (
      <button
        {...rest}
        className={classValue}
        type={type === "button" ? "button" : "submit"}
        ref={ref}
        onMouseUp={(e) => {
          if (ripple === "dark") rippleEffect.create(e, "dark");
          if (ripple === "light") rippleEffect.create(e, "light");
        }}
      >
        {children}
      </button>
    );
  },
);
