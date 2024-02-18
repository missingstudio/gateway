"use client";
import { BentoGrid, BentoGridItem } from "components/bento-grid";
import Particles from "components/particles";
import { cn } from "lib/utils";
import { useState } from "react";
import Balance from "react-wrap-balancer";
import {
  SkeletonFour,
  SkeletonOne,
  SkeletonThree,
  SkeletonTwo,
  TextRevealCardPreview,
} from "./skeletons";

const items = [
  {
    title: "Load balancing",
    description: (
      <span className="text-sm">
        Efficiently distribute incoming requests among multiple models or models
      </span>
    ),
    header: <SkeletonOne />,
    className: "md:col-span-1",
  },
  {
    title: "Exponential Retries",
    description: (
      <span className="text-sm">Atomatic Retries with exponential backoff</span>
    ),
    header: <SkeletonTwo />,
    className: "md:col-span-1",
  },
  {
    title: "Automatic Fallback",
    description: (
      <span className="text-sm">
        Improve reliability by automatically switching to alternative models
      </span>
    ),
    header: <SkeletonThree />,
    className: "md:col-span-1",
  },

  {
    header: <TextRevealCardPreview />,
    className: "md:col-span-2",
  },

  {
    title: "Semantic Caching",
    description: (
      <span className="text-sm">
        Reduce costs and latency by efficient caching.
      </span>
    ),
    header: <SkeletonFour />,
    className: "md:col-span-1",
  },
];

export default function AIRouter() {
  const [tab, setTab] = useState<number>(1);

  return (
    <section>
      <div className="relative max-w-6xl mx-auto px-4 sm:px-8">
        <div className="pt-16 pb-12 md:pt-52 md:pb-20 overflow-hidden relative">
          <div
            className="absolute flex items-center justify-center top-0 -translate-y-1/2 left-1/2 -translate-x-1/2 pointer-events-none -z-10 w-1/3 aspect-square"
            aria-hidden="true"
          >
            <div className="absolute inset-0 translate-z-0 bg-slate-300 rounded-full blur-[120px] opacity-70" />
            <div className="absolute w-1/4 h-1/4 translate-z-0 bg-slate-200 rounded-full blur-[40px]" />
          </div>

          <div className="max-w-xl mx-auto md:max-w-none flex flex-col md:flex-row md:space-x-8 lg:space-x-16 xl:space-x-20 space-y-8 space-y-reverse md:space-y-0 pb-24">
            <div
              className="md:w-7/12 lg:w-1/2 order-1 md:order-none max-md:text-center"
              data-aos="fade-down"
            >
              <div className="text-sm inline-flex font-bold bg-clip-text text-transparent bg-gradient-to-r from-violet-500 to-violet-200 pb-3 uppercase">
                Universal API
              </div>
              <h1 className="text-4xl md:text-5xl font-bold bg-clip-text text-transparent bg-gradient-to-b from-neutral-50 to-neutral-400 bg-opacity-50 pb-3">
                AI Router
              </h1>
              <Balance className="text-base text-neutral-300 mb-8 tracking-wide">
                Improve reliability and eliminate need to learn and manage
                multiple APIs for different LLM provider. Use a single API to
                access various LLMs using AI Router.
                <div className="mb-4" />
                Try our unified API for seamless integration with 100+ providers
                like OpenAI, Anthropic, Cohere, and more.
              </Balance>
            </div>

            <div
              className="hidden md:block md:w-5/12 lg:w-1/2"
              data-aos="fade-up"
              data-aos-delay="100"
            >
              <div className="relative py-24 ">
                <Particles
                  className="absolute inset-0 -z-10"
                  quantity={8}
                  staticity={30}
                />
              </div>
            </div>
          </div>

          <BentoGrid className="max-w-4xl mx-auto md:auto-rows-[20rem]">
            {items.map((item, i) => (
              <BentoGridItem
                key={i}
                title={item.title}
                description={item.description}
                header={item.header}
                className={cn("[&>p:text-lg]", item.className)}
              />
            ))}
          </BentoGrid>
        </div>
      </div>
    </section>
  );
}
