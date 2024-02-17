import Balance from "react-wrap-balancer";

export default function CTA() {
  return (
    <section>
      <div className="max-w-6xl mx-auto px-4 sm:px-6">
        <div className="relative px-8 py-12 md:py-20 rounded-[3rem] overflow-hidden">
          <div
            className="absolute flex items-center justify-center top-0 -translate-y-1/2 left-1/2 -translate-x-1/2 pointer-events-none -z-10 w-1/3 aspect-square"
            aria-hidden="true"
          >
            <div className="absolute inset-0 translate-z-0 bg-violet-500 rounded-full blur-[120px] opacity-70" />
            <div className="absolute w-1/4 h-1/4 translate-z-0 bg-violet-400 rounded-full blur-[40px]" />
          </div>

          <div className="max-w-3xl mx-auto text-center">
            <div>
              <span className="relative inline-flex items-center font-bold uppercase tracking-[0.12rem] text-violet-400">
                playgrounds
              </span>
            </div>
            <h2 className="text-3xl md:text-6xl font-bold bg-clip-text !leading-tight text-transparent bg-gradient-to-r from-slate-200/60 via-violet-200 to-slate-200/60 mb-4">
              AI workflows
            </h2>
            <Balance className="text-sm sm:text-xl text-slate-400 tracking-wide mb-8 antialiased leading-snug">
              Use AI studio to create, experiment, deploy and scale LLM powered
              workflows using the latest AI models like GPT-4-turbo, DALL·E 3
              and Claude 2
            </Balance>
            <div>
              <a
                className="btn text-slate-900 bg-gradient-to-r from-white/80 via-white to-white/80 hover:bg-white transition duration-150 ease-in-out group"
                href="https://github.com/missingstudio/studio"
              >
                Get Started{" "}
              </a>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
