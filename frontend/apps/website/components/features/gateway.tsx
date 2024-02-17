import Highlighter, { HighlighterItem } from "components/highlighter";
import Particles from "components/particles";
import Balance from "react-wrap-balancer";

export default function Features() {
  return (
    <section className="relative">
      <div className="absolute left-1/2 -translate-x-1/2 top-0 -z-10 w-80 h-80 -mt-24 -ml-32">
        <Particles
          className="absolute inset-0 -z-10"
          quantity={6}
          staticity={30}
        />
      </div>

      <div className="max-w-6xl mx-auto px-4 sm:px-6 ">
        <div className="pt-16 md:pt-32 md:pb-32">
          <div className="max-w-3xl mx-auto text-center pb-12 md:pb-20">
            <h2 className="text-3xl md:text-6xl font-bold bg-clip-text !leading-tight text-transparent bg-gradient-to-r from-slate-200/60 via-slate-200 to-slate-200/60 mb-4">
              AI Gateway
            </h2>
            <Balance className="text-sm sm:text-xl text-slate-400 tracking-wide mb-8 antialiased leading-snug">
              Missing studio AI gateway allow you gain visibility, control and
              insights about API Application uses. It allow you control how your
              application scales with features such as caching, rate limiting,
              as well as request retries, model fallback, and more
            </Balance>
          </div>

          <div className="relative pb-12 md:pb-20">
            <Highlighter className="grid md:grid-cols-12 gap-6 group">
              <div className="md:col-span-12" data-aos="fade-down">
                <HighlighterItem>
                  <div className="relative h-full bg-slate-900 rounded-[inherit] z-20 overflow-hidden">
                    <div className="flex flex-col md:flex-row md:items-center md:justify-between">
                      <div
                        className="absolute right-0 top-0 blur-2xl"
                        aria-hidden="true"
                      >
                        <svg
                          xmlns="http://www.w3.org/2000/svg"
                          width="342"
                          height="393"
                        >
                          <defs>
                            <linearGradient
                              id="bs-a"
                              x1="19.609%"
                              x2="50%"
                              y1="14.544%"
                              y2="100%"
                            >
                              <stop offset="0%" stopColor="#6366F1" />
                              <stop
                                offset="100%"
                                stopColor="#6366F1"
                                stopOpacity="0"
                              />
                            </linearGradient>
                          </defs>
                          <path
                            fill="url(#bs-a)"
                            fillRule="evenodd"
                            d="m104 .827 461 369-284 58z"
                            transform="translate(0 -112.827)"
                            opacity=".7"
                          />
                        </svg>
                      </div>

                      <div className="md:max-w-[480px] shrink-0 order-1 md:order-none p-6 pt-0 md:p-8 md:pr-0">
                        <div className="mb-5">
                          <div>
                            <h3 className="h3 bg-clip-text text-transparent bg-gradient-to-r from-slate-200/60 via-slate-200 to-slate-200/60">
                              Analytics for LLMs
                            </h3>
                            <p className="text-sm sm:text-lg text-slate-400 mb-8 tracking-wide">
                              Understand your costs per users and models for any
                              provider — plus, caching and rate limiting out of
                              the box.
                            </p>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </HighlighterItem>
              </div>
            </Highlighter>
          </div>

          <div className="grid md:grid-cols-3 gap-8 md:gap-12s">
            <div>
              <div className="flex items-center space-x-2 mb-1">
                <h4 className="h4 bg-clip-text text-transparent bg-gradient-to-r from-slate-200/80 via-slate-200 to-slate-200/80">
                  Usage monitoring
                </h4>
              </div>
              <p className="text-sm sm:text-lg text-slate-400 mb-8 tracking-wide">
                Track and analyze usage, requests, tokens, cache, and more.
              </p>
            </div>

            <div>
              <div className="flex items-center space-x-2 mb-1">
                <svg
                  className="shrink-0 fill-slate-300"
                  xmlns="http://www.w3.org/2000/svg"
                  width="16"
                  height="16"
                >
                  <path d="M11 0c1.3 0 2.6.5 3.5 1.5 1 .9 1.5 2.2 1.5 3.5 0 1.3-.5 2.6-1.4 3.5l-1.2 1.2c-.2.2-.5.3-.7.3-.2 0-.5-.1-.7-.3-.4-.4-.4-1 0-1.4l1.1-1.2c.6-.5.9-1.3.9-2.1s-.3-1.6-.9-2.2C12 1.7 10 1.7 8.9 2.8L7.7 4c-.4.4-1 .4-1.4 0-.4-.4-.4-1 0-1.4l1.2-1.1C8.4.5 9.7 0 11 0ZM8.3 12c.4-.4 1-.5 1.4-.1.4.4.4 1 0 1.4l-1.2 1.2C7.6 15.5 6.3 16 5 16c-1.3 0-2.6-.5-3.5-1.5C.5 13.6 0 12.3 0 11c0-1.3.5-2.6 1.5-3.5l1.1-1.2c.4-.4 1-.4 1.4 0 .4.4.4 1 0 1.4L2.9 8.9c-.6.5-.9 1.3-.9 2.1s.3 1.6.9 2.2c1.1 1.1 3.1 1.1 4.2 0L8.3 12Zm1.1-6.8c.4-.4 1-.4 1.4 0 .4.4.4 1 0 1.4l-4.2 4.2c-.2.2-.5.3-.7.3-.2 0-.5-.1-.7-.3-.4-.4-.4-1 0-1.4l4.2-4.2Z" />
                </svg>
                <h4 className="h4 bg-clip-text text-transparent bg-gradient-to-r from-slate-200/80 via-slate-200 to-slate-200/80">
                  Cost analytics
                </h4>
              </div>
              <p className="text-sm sm:text-lg text-slate-400 mb-8 tracking-wide">
                Understand how much users, requests and models are costing you.
              </p>
            </div>
            <div>
              <div className="flex items-center space-x-2 mb-1">
                <svg
                  className="shrink-0 fill-slate-300"
                  xmlns="http://www.w3.org/2000/svg"
                  width="16"
                  height="16"
                >
                  <path d="M14 0a2 2 0 0 1 2 2v4a1 1 0 0 1-2 0V2H2v12h4a1 1 0 0 1 0 2H2a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2h12Zm-1.957 10.629 3.664 3.664a1 1 0 0 1-1.414 1.414l-3.664-3.664-.644 2.578a.5.5 0 0 1-.476.379H9.5a.5.5 0 0 1-.48-.362l-2-7a.5.5 0 0 1 .618-.618l7 2a.5.5 0 0 1-.017.965l-2.578.644Z" />
                </svg>
                <h4 className="h4 bg-clip-text text-transparent bg-gradient-to-r from-slate-200/80 via-slate-200 to-slate-200/80">
                  LLM optimization
                </h4>
              </div>
              <p className="text-sm sm:text-lg text-slate-400 mb-8 tracking-wide">
                Easily add remote cache, rate limits, and auto retries.
              </p>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
