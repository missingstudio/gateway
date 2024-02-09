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
              <div className="inline-flex font-medium bg-clip-text text-transparent bg-gradient-to-r from-violet-500 to-violet-200 pb-3">
                playgorund
              </div>
            </div>
            <h2 className="h2 bg-clip-text text-transparent bg-gradient-to-r from-slate-200/60 via-slate-200 to-slate-200/60 pb-4">
              AI workflows
            </h2>
            <p className="text-lg text-slate-400 mb-8">
              Use LLM studio to create, experiment, deploy and scale LLM powered
              workflows using the latest AI models like GPT-4-turbo, DALL·E 3
              and Claude 2
            </p>
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
