/* eslint-disable react/jsx-props-no-spreading */
/* eslint-disable react/require-default-props */
import React from "react";
import classNames from "classnames";
import { borderColors, Colors, mtInputColors, mtInputOutlineColors } from "../shared";

export interface InputProps extends Omit<React.InputHTMLAttributes<HTMLInputElement>, "size"> {
  //   children: React.ReactNode;
  placeholder: string;
  color?: Colors;
  size?: "sm" | "md" | "lg";
  outline?: boolean;
  error?: string;
  success?: string;
  className?: string;
}

export function Input({
  placeholder,
  color = "lightBlue",
  size = "md",
  outline,
  className,
  error,
  success,
  ...rest
}: InputProps) {
  let asColor = color;
  if (error) {
    asColor = "red";
  } else if (success) {
    asColor = "green";
  }

  const mtInputBorderColor = mtInputColors[asColor];
  const mtInputOutlineColor = mtInputOutlineColors[asColor];
  const mtInputOutlineFocusColor = borderColors[asColor];

  const labelStyle = {
    [borderColors.red]: !!error,
    [borderColors.green]: !!success,
    "border-grey-300": !success && !error,
  };

  const containerNames = classNames("w-full", "relative", {
    "h-9": size === "sm",
    "h-11": size === "md",
    "h-12": size === "lg",
  });

  const inputNames = classNames(
    "w-full",
    "h-full",
    "text-gray-800",
    "leading-normal",
    "shadow-none",
    "outline-none",
    "focus:outline-none",
    "focus:ring-0",
    "focus:text-gray-800",
    outline ? "px-3" : "px-0",
    size === "sm" && {
      "pt-1.5 pb-0.5": outline,
      "text-sm": true,
    },
    size === "md" && {
      "pt-2.5 pb-1.5": outline,
    },
    size === "lg" && {
      "pt-3.5 pb-2.5": outline,
    },
    {
      "mt-input-outline-error": !!error,
      "mt-input-outline-success": !!success,
    },
    "bg-transparent",
    outline
      ? [
          mtInputOutlineColor,
          "mt-input-outline border border-1 border-gray-300 rounded-lg focus:border-2",
          `focus:${mtInputOutlineFocusColor}`,
        ]
      : [mtInputBorderColor, "mt-input border-none"],
    labelStyle,
    className,
  );

  const labelNames = classNames(
    "text-gray-400",
    "absolute",
    "left-0",
    "w-full",
    "h-full",
    labelStyle,
    "pointer-events-none",
    {
      "border border-t-0 border-l-0 border-r-0 border-b-1": !outline,
      "-top-1.5": outline,
      "-top-0.5": !outline,
      "text-sm": outline && size === "sm",
      "flex leading-10 transition-all duration-300": outline,
    },
  );

  return (
    <div className={containerNames}>
      <input {...rest} placeholder=" " className={inputNames} />
      <label className={labelNames}>
        {outline ? (
          placeholder
        ) : (
          <span className={classNames(size === "sm" && "text-sm", "absolute top-1/4 transition-all duration-300")}>
            {placeholder}
          </span>
        )}
      </label>
      {error && <span className="block mt-1 text-xs text-red-500">{error}</span>}
      {success && <span className="block mt-1 text-xs text-green-500">{success}</span>}
    </div>
  );
}
