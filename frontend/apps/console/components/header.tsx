"use client";
import React from "react";
import { Toggle } from "~/components/theme";

import { Button } from "@missingstudio/ui/button";
import { BookAIcon, Shapes } from "lucide-react";
import Link from "next/link";

const links = [
  {
    name: "Playground",
    route: "/playground",
    icon: Shapes,
  },
  {
    name: "Docs",
    route: "https://docs.missing.studio",
    icon: BookAIcon,
  },
];

export default function Header() {
  return (
    <header className="relative flex h-20 p-4">
      <div className="inline-flex w-[200px] items-center font-bold tracking-wide">
        <Link href="/">AIstudio</Link>
      </div>
      <div className="absolute left-1/2 top-1/2 flex -translate-x-1/2 -translate-y-1/2 items-center gap-4">
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
