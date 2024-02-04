"use client";

import { Button } from "@missingstudio/ui/button";
import { Text } from "@missingstudio/ui/text";
import { AnimatePresence, motion } from "framer-motion";
import { Maven_Pro } from "next/font/google";
import styles from "./hero.module.css";

const mavenpro = Maven_Pro({ subsets: ["latin"], variable: "--font-maven" });

export default function Page() {
  return (
    <AnimatePresence>
      <section className="space-y-6 pt-6 md:pt-10 lg:pt-32">
        <div className="container flex max-w-[64rem] flex-col items-center gap-8 text-center">
          <div className="flex flex-col gap-2 items-center">
            <motion.h1
              className={`font-extrabold font-heading text-3xl sm:text-4xl md:text-5xl lg:text-6xl ${mavenpro.className} ${styles.heading} ${styles.headergradient}`}
              initial={{ opacity: 0, top: 20 }}
              animate={{ opacity: 1, top: 0 }}
              transition={{ duration: 0.8, ease: "easeOut" }}
            >
              A Open Source <br />
              For Enterprise AI studio
            </motion.h1>
            <motion.p
              className="max-w-[42rem] leading-normal text-muted-foreground sm:text-lg sm:leading-8"
              initial={{ opacity: 0, top: 20 }}
              animate={{ opacity: 1, top: 0 }}
              transition={{ duration: 0.6, delay: 0.2, ease: "easeOut" }}
            >
              A Robust Open Source AI studio with Universal API for inferencing
              100+ LLMs(OpenAI, Azure, Anthropic, HuggingFace, Replicate, Stable
              Diffusion).
            </motion.p>
          </div>

          <motion.div
            className="flex justify-center"
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            transition={{ duration: 1.2, delay: 0.2, ease: "easeOut" }}
          >
            <Button variant="alternative" className="!p-0">
              <a
                target="_blank"
                href="https://github.com/missingstudio/studio"
                className="px-5 py-2"
              >
                <Text size={2} bold>
                  We're Open Source
                </Text>
              </a>
            </Button>
          </motion.div>
        </div>
      </section>
    </AnimatePresence>
  );
}
