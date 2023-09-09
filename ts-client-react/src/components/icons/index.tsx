import React from "react";
import { IconProps, Icon } from "@chakra-ui/react";
import { FaGithub } from "react-icons/fa";
import { FcGoogle } from "react-icons/fc";
import { BsGear, BsBell, BsKey, BsPlus, BsThreeDotsVertical } from "react-icons/bs";
import {
  FiEyeOff,
  FiEye,
  FiHome,
  FiTrendingUp,
  FiCompass,
  FiStar,
  FiSettings,
  FiChevronsDown,
  FiMenu,
  FiEdit,
  FiDelete,
} from "react-icons/fi";

export function GoogleIcon(props: IconProps) {
  return <Icon as={FcGoogle} {...props} />;
}

export function GithubIcon(props: IconProps) {
  return <Icon as={FaGithub} {...props} />;
}

export function GearIcon(props: IconProps) {
  return <Icon as={BsGear} {...props} />;
}

export function BellIcon(props: IconProps) {
  return <Icon as={BsBell} {...props} />;
}

export function KeyIcon(props: IconProps) {
  return <Icon as={BsKey} {...props} />;
}

export function PlusIcon(props: IconProps) {
  return <Icon as={BsPlus} {...props} />;
}

export function EyeIcon(props: IconProps) {
  return <Icon as={FiEye} {...props} />;
}

export function EyeOffIcon(props: IconProps) {
  return <Icon as={FiEyeOff} {...props} />;
}

export function HomeIcon(props: IconProps) {
  return <Icon as={FiHome} {...props} />;
}

export function TrendingIcon(props: IconProps) {
  return <Icon as={FiTrendingUp} {...props} />;
}

export function CompassIcon(props: IconProps) {
  return <Icon as={FiCompass} {...props} />;
}

export function StarIcon(props: IconProps) {
  return <Icon as={FiStar} {...props} />;
}

export function SettingsIcon(props: IconProps) {
  return <Icon as={FiSettings} {...props} />;
}

export function ChevronDownIcon(props: IconProps) {
  return <Icon as={FiChevronsDown} {...props} />;
}

export function MenuIcon(props: IconProps) {
  return <Icon as={FiMenu} {...props} />;
}

export function EditIcon(props: IconProps) {
  return <Icon as={FiEdit} {...props} />;
}

export function DeleteIcon(props: IconProps) {
  return <Icon as={FiDelete} {...props} />;
}

export function KebabIcon(props: IconProps) {
  return <Icon as={BsThreeDotsVertical} {...props} />;
}
