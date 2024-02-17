import { MilkeyWay } from "../milkway";

export default function Hero() {
  return (
    <section>
      <div className="relative max-w-6xl mx-auto px-4 sm:px-6">
        <MilkeyWay />
        <div className="pt-32 pb-16 md:pt-52 md:pb-32">
          <div className="max-w-3xl mx-auto text-center">
            <div className="mb-4" data-aos="fade-down">
              <span className="relative inline-flex items-center font-bold uppercase tracking-[0.12rem] text-violet-400">
                AI Studio for developer
              </span>
            </div>
            <h1
              className="text-3xl md:text-6xl font-bold bg-clip-text !leading-tight text-transparent bg-gradient-to-r from-slate-200/60 via-violet-200 to-slate-200/60 mb-4"
              data-aos="fade-down"
            >
              Open-source
              <br />
              AI studio for AI Apps
            </h1>
            <p
              className="text-sm sm:text-xl text-slate-400 tracking-wide mb-8 antialiased leading-snug"
              data-aos="fade-down"
              data-aos-delay="200"
            >
              An open-source AI studio for rapid development and robust
              deployment of generative AI applications that seamlessly
              integrates with top providers like OpenAI, Anthropic etc
            </p>
            <div
              className="max-w-xs mx-auto sm:max-w-none sm:inline-flex sm:justify-center space-y-4 sm:space-y-0 sm:space-x-4"
              data-aos="fade-down"
              data-aos-delay="400"
            >
              <div>
                <a
                  className="btn text-slate-200 hover:text-white bg-slate-900 bg-opacity-25 hover:bg-opacity-30 w-full transition duration-150 ease-in-out"
                  href="https://docs.missing.studio"
                  target="_blank"
                >
                  <svg
                    className="shrink-0 fill-slate-300 mr-3"
                    xmlns="http://www.w3.org/2000/svg"
                    width="16"
                    height="16"
                  >
                    <path d="m1.999 0 1 2-1 2 2-1 2 1-1-2 1-2-2 1zM11.999 0l1 2-1 2 2-1 2 1-1-2 1-2-2 1zM11.999 10l1 2-1 2 2-1 2 1-1-2 1-2-2 1zM6.292 7.586l2.646-2.647L11.06 7.06 8.413 9.707zM0 13.878l5.586-5.586 2.122 2.121L2.12 16z" />
                  </svg>
                  <span>Read the docs</span>
                </a>
              </div>
              <div>
                <a
                  className="btn text-slate-900 bg-gradient-to-r from-white/80 via-white to-white/80 hover:bg-white w-full transition duration-150 ease-in-out group"
                  href="https://github.com/missingstudio/studio"
                  target="_blank"
                >
                  Checkout AI Studio
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
