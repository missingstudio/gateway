"use client";

import { useState } from "react";

import { Transition } from "@headlessui/react";
import Particles from "~/components/particles";

export default function Features() {
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
          <div>
            <div className="max-w-xl mx-auto md:max-w-none flex flex-col md:flex-row md:space-x-8 lg:space-x-16 xl:space-x-20 space-y-8 space-y-reverse md:space-y-0">
              <div
                className="md:w-7/12 lg:w-1/2 order-1 md:order-none max-md:text-center"
                data-aos="fade-down"
              >
                <div>
                  <div className="inline-flex font-medium bg-clip-text text-transparent bg-gradient-to-r from-violet-500 to-violet-200 pb-3 uppercase">
                    Universal API
                  </div>
                </div>
                <h3 className="h3 bg-clip-text text-transparent bg-gradient-to-r from-slate-200/60 via-slate-200 to-slate-200/60 pb-3">
                  Universal API for providers
                </h3>
                <p className="text-lg text-slate-400 mb-8">
                  Familiar with spending hundreds of hours with multiple
                  providers to get things right. Try our universal API for
                  seamless integration with 100+ providers like OpenAI,
                  Anthropic, Cohere, and more.
                </p>
              </div>

              <div
                className="md:w-5/12 lg:w-1/2"
                data-aos="fade-up"
                data-aos-delay="100"
              >
                <div className="relative py-24 -mt-12">
                  <Particles
                    className="absolute inset-0 -z-10"
                    quantity={8}
                    staticity={30}
                  />

                  <div className="flex items-center justify-center">
                    <div className="relative w-48 h-48 flex justify-center items-center">
                      <svg
                        className="absolute inset-0 left-1/2 top-1/2 transform -translate-x-1/2 -translate-y-1/2 will-change-transform pointer-events-none blur-md"
                        width="480"
                        height="480"
                        viewBox="0 0 480 480"
                        xmlns="http://www.w3.org/2000/svg"
                      >
                        <defs>
                          <linearGradient
                            id="pulse-a"
                            x1="50%"
                            x2="50%"
                            y1="100%"
                            y2="0%"
                          >
                            <stop offset="0%" stopColor="#A855F7" />
                            <stop offset="76.382%" stopColor="#FAF5FF" />
                            <stop offset="100%" stopColor="#6366F1" />
                          </linearGradient>
                        </defs>
                        <g fillRule="evenodd">
                          <path
                            className="pulse"
                            fill="url(#pulse-a)"
                            fillRule="evenodd"
                            d="M240,0 C372.5484,0 480,107.4516 480,240 C480,372.5484 372.5484,480 240,480 C107.4516,480 0,372.5484 0,240 C0,107.4516 107.4516,0 240,0 Z M240,88.8 C156.4944,88.8 88.8,156.4944 88.8,240 C88.8,323.5056 156.4944,391.2 240,391.2 C323.5056,391.2 391.2,323.5056 391.2,240 C391.2,156.4944 323.5056,88.8 240,88.8 Z"
                          />
                          <path
                            className="pulse pulse-1"
                            fill="url(#pulse-a)"
                            fillRule="evenodd"
                            d="M240,0 C372.5484,0 480,107.4516 480,240 C480,372.5484 372.5484,480 240,480 C107.4516,480 0,372.5484 0,240 C0,107.4516 107.4516,0 240,0 Z M240,88.8 C156.4944,88.8 88.8,156.4944 88.8,240 C88.8,323.5056 156.4944,391.2 240,391.2 C323.5056,391.2 391.2,323.5056 391.2,240 C391.2,156.4944 323.5056,88.8 240,88.8 Z"
                          />
                          <path
                            className="pulse pulse-2"
                            fill="url(#pulse-a)"
                            fillRule="evenodd"
                            d="M240,0 C372.5484,0 480,107.4516 480,240 C480,372.5484 372.5484,480 240,480 C107.4516,480 0,372.5484 0,240 C0,107.4516 107.4516,0 240,0 Z M240,88.8 C156.4944,88.8 88.8,156.4944 88.8,240 C88.8,323.5056 156.4944,391.2 240,391.2 C323.5056,391.2 391.2,323.5056 391.2,240 C391.2,156.4944 323.5056,88.8 240,88.8 Z"
                          />
                        </g>
                      </svg>

                      <div className="absolute inset-0 left-1/2 top-1/2 transform -translate-x-1/2 -translate-y-1/2 pointer-events-none w-[500px] h-[500px] rounded-full overflow-hidden [mask-image:_radial-gradient(black,_transparent_60%)]">
                        <div className="h-[200%] animate-endless">
                          <div className="absolute inset-0 [background:_repeating-linear-gradient(transparent,_transparent_48px,_theme(colors.white)_48px,_theme(colors.white)_49px)] blur-[2px] opacity-20" />
                          <div className="absolute inset-0 [background:_repeating-linear-gradient(transparent,_transparent_48px,_theme(colors.violet.500)_48px,_theme(colors.violet.500)_49px)]" />
                          <div className="absolute inset-0 [background:_repeating-linear-gradient(90deg,transparent,_transparent_48px,_theme(colors.white)_48px,_theme(colors.white)_49px)] blur-[2px] opacity-20" />
                          <div className="absolute inset-0 [background:_repeating-linear-gradient(90deg,transparent,_transparent_48px,_theme(colors.violet.500)_48px,_theme(colors.violet.500)_49px)]" />
                        </div>
                      </div>

                      <Transition
                        show
                        className="absolute"
                        enter="transition ease-[cubic-bezier(0.68,-0.3,0.32,1)] duration-700 order-first"
                        enterFrom="opacity-0 -rotate-[60deg]"
                        enterTo="opacity-100 rotate-0"
                        leave="transition ease-[cubic-bezier(0.68,-0.3,0.32,1)] duration-700"
                        leaveFrom="opacity-100 rotate-0"
                        leaveTo="opacity-0 rotate-[60deg]"
                      >
                        <div className="relative flex items-center justify-center w-16 h-16 border border-transparent rounded-2xl shadow-2xl -rotate-[14deg] [background:linear-gradient(theme(colors.slate.900),_theme(colors.slate.900))_padding-box,_conic-gradient(theme(colors.slate.400),_theme(colors.slate.700)_25%,_theme(colors.slate.700)_75%,_theme(colors.slate.400)_100%)_border-box] before:absolute before:inset-0 before:bg-slate-800/30 before:rounded-2xl">
                          <svg
                            className="relative fill-slate-200"
                            xmlns="http://www.w3.org/2000/svg"
                            width="23"
                            height="25"
                          >
                            <path
                              fillRule="nonzero"
                              d="M10.55 15.91H.442L14.153.826 12.856 9.91h10.107L9.253 24.991l1.297-9.082Zm.702-8.919L4.963 13.91h7.893l-.703 4.918 6.289-6.918H10.55l.702-4.918Z"
                            />
                          </svg>
                        </div>
                      </Transition>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
