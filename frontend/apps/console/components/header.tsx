"use client";
import React, { useMemo } from "react";
import { Toggle } from "~/components/theme";

import { Button } from "@missingstudio/ui/button";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@missingstudio/ui/popover";
import { Blocks, BookAIcon, Shapes } from "lucide-react";
import Link from "next/link";
import { usePathname } from "next/navigation";

const links = [
  {
    name: "Docs",
    route: "https://docs.missing.studio",
    icon: BookAIcon,
  },
];

const navigations = {
  playground: {
    name: "Playground",
    route: "/playground",
    icon: Shapes,
  },
  providers: {
    name: "Providers",
    route: "/providers",
    icon: Blocks,
  },
};

export default function Header() {
  const pathName = usePathname();
  const selectedPage = useMemo(() => {
    if (pathName.endsWith("/playground")) {
      return "playground";
    }

    if (pathName.endsWith("/providers")) {
      return "providers";
    }

    return "playground";
  }, [pathName]);

  const page = navigations[selectedPage];
  return (
    <header className="relative flex h-20 p-4">
      <div className="inline-flex w-[200px] items-center font-bold tracking-wide">
        <Link href="/">AIstudio</Link>
      </div>
      <div className="absolute left-1/2 top-1/2 flex -translate-x-1/2 -translate-y-1/2 items-center gap-4">
        <Popover>
          <PopoverTrigger asChild>
            <Button asChild key={"/playground"} variant="ghost">
              <div>
                {React.createElement(page.icon, { className: "mr-2 h-5 w-5" })}
                {page.name}
              </div>
            </Button>
          </PopoverTrigger>
          <PopoverContent className="flex flex-col gap-2 w-48">
            {Object.values(navigations).map((link) => (
              <Button
                asChild
                key={link.route}
                variant="ghost"
                className="justify-start"
              >
                <Link href={link.route}>
                  {React.createElement(link.icon, {
                    className: "mr-2 h-5 w-5",
                  })}
                  {link.name}
                </Link>
              </Button>
            ))}
          </PopoverContent>
        </Popover>

        {links.map((link) => (
          <Button asChild key={link.route} variant="ghost">
            <Link href={link.route}>
              {React.createElement(link.icon, { className: "mr-2 h-5 w-5" })}
              {link.name}
            </Link>
          </Button>
        ))}
        <Toggle />
      </div>
    </header>
  );
}
