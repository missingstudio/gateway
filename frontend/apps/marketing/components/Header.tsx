"use client";

import { motion } from "framer-motion";
import Image from "next/image";
import Link from "next/link";
import logo from "~/public/next.svg";

interface HeaderProps {}

export const Header = ({}: HeaderProps) => {
  return (
    <motion.header
      initial={{
        opacity: 0,
        transform: "translateY(-50px)",
      }}
      animate={{
        opacity: 1,
        transform: "translateY(0px)",
      }}
      className="z-40 bg-background py-8"
    >
      <div className="container flex h-8 items-center justify-between">
        <div className="flex gap-6 md:gap-10">
          <Link href="/" className="items-center space-x-2 md:flex">
            <Image priority src={logo} height={24} alt="missing studio logo" />
          </Link>
        </div>
      </div>
    </motion.header>
  );
};
