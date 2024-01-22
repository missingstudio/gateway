"use client";

import { AnimatePresence, motion } from "framer-motion";
import { Maven_Pro } from "next/font/google";
import styles from "./hero.module.css";

const mavenpro = Maven_Pro({ subsets: ["latin"], variable: "--font-maven" });

export default function Page() {
  return (
    <AnimatePresence>
      <section className="space-y-6 pt-6 md:pt-10 lg:pt-32">
        <div className="container flex max-w-[64rem] flex-col items-center gap-6 text-center">
          <div className="flex flex-col gap-2 items-center">
            <motion.h1
              className={`font-extrabold font-heading text-3xl sm:text-4xl md:text-5xl lg:text-6xl ${mavenpro.className} ${styles.heading}`}
              initial={{ opacity: 0, top: 20 }}
              animate={{ opacity: 1, top: 0 }}
              transition={{ duration: 0.8, ease: "easeOut" }}
            >
              Supercharging startups with <br /> next-gen AI applications
            </motion.h1>
          </div>
          <motion.p
            className="max-w-[42rem] leading-normal text-muted-foreground sm:text-lg sm:leading-8"
            initial={{ opacity: 0, top: 20 }}
            animate={{ opacity: 1, top: 0 }}
            transition={{ duration: 0.6, delay: 0.2, ease: "easeOut" }}
          >
            Missing studio is a robust open source platform to build a
            developer-first way to create AI application - all in one platform.
          </motion.p>
        </div>
      </section>
    </AnimatePresence>
  );
}
