"use client";

import { Button } from "@missingstudio/ui/button";
import { Text } from "@missingstudio/ui/text";
import { motion } from "framer-motion";
import Image from "next/image";
import Link from "next/link";
import logo from "~/public/mstudio_light.svg";

export const Header = () => {
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
      className="container z-40 bg-background py-8"
    >
      <div className="container flex h-8 items-center ">
        <div className="flex gap-6 md:gap-10 w-full justify-between">
          <Link href="/" className="items-center space-x-2 md:flex">
            <Image priority src={logo} height={24} alt="missing studio logo" />
          </Link>
          <Button variant="alternative" className="!p-0">
            <a
              target="_blank"
              href="https://github.com/missingstudio/studio"
              className="px-5 py-2"
            >
              <Text size={2} bold>
                Try AI Studio
              </Text>
            </a>
          </Button>
        </div>
      </div>
    </motion.header>
  );
};
