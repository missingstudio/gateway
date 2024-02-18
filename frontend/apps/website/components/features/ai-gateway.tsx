"use client";
import { HoverEffect } from "components/card-hover-effect";
import Particles from "components/particles";
import Balance from "react-wrap-balancer";
export const projects = [
  {
    title: "Observability",
    description: (
      <div>
        <div className="text-sm mb-2">
          <b>Logging</b>: Keep track of all requests for monitoring and
          debugging.
        </div>
        <span className="text-sm">
          <b>Requests Tracing</b>: Understand the journey of each request for
          optimization.
        </span>
      </div>
    ),
  },
  {
    title: "Usage monitoring",
    description: "Track and analyze usage, requests, tokens, cache, and more.",
  },
  {
    title: "Performance monitoring",
    description:
      "Understand how much users, requests and models are costing you.",
  },
  {
    title: "LLM optimization",
    description: "Easily add remote cache, rate limits, and auto retries.",
  },
  {
    title: "API key Managment",
    description:
      "Keep your primary credentials away. Can be easily revoked or renewed for better access control",
  },
];

export default function AIGateway() {
  return (
    <section className="relative">
      <div className="absolute left-1/2 -translate-x-1/2 top-0 -z-10 w-80 h-80 -mt-24 -ml-32">
        <Particles
          className="absolute inset-0 -z-10"
          quantity={6}
          staticity={30}
        />
      </div>

      <div
        className="absolute top-0 -translate-y-1/4 left-1/2 -translate-x-1/2 blur-2xl opacity-50 pointer-events-none -z-10"
        aria-hidden="true"
      >
        <svg xmlns="http://www.w3.org/2000/svg" width="434" height="427">
          <defs>
            <linearGradient
              id="bs3-a"
              x1="19.609%"
              x2="50%"
              y1="14.544%"
              y2="100%"
            >
              <stop offset="0%" stopColor="#6366F1" />
              <stop offset="100%" stopColor="#6366F1" stopOpacity="0" />
            </linearGradient>
          </defs>
          <path
            fill="url(#bs3-a)"
            fillRule="evenodd"
            d="m410 0 461 369-284 58z"
            transform="matrix(1 0 0 -1 -410 427)"
          />
        </svg>
      </div>
      <div className="max-w-6xl mx-auto px-4 sm:px-6 ">
        <div className="pt-16 md:pt-32 md:pb-32">
          <div className="max-w-3xl mx-auto text-center pb-12 md:pb-20">
            <h2 className="text-4xl md:text-4xl font-bold bg-clip-text text-transparent bg-gradient-to-b from-neutral-50 to-neutral-400 bg-opacity-50 pb-3 mb-4">
              AI Gateway
            </h2>
            <Balance className="text-base text-neutral-300 mb-8 tracking-wide">
              Missing studio AI gateway allow you gain visibility, control and
              insights about API Application uses. Track all your LLM requests
              transparently and conveys the insights needed to make data-driven
              decisions.
            </Balance>
          </div>

          <HoverEffect items={projects} />
        </div>
      </div>
    </section>
  );
}
