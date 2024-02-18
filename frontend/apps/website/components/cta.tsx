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
            <div className="absolute inset-0 translate-z-0 bg-neutral-500 rounded-full blur-[120px] opacity-10" />
          </div>

          <div className="max-w-3xl mx-auto text-center">
            <div>
              <span className="relative inline-flex items-center font-bold uppercase tracking-[0.12rem] text-violet-400">
                playgrounds
              </span>
            </div>
            <h2 className="text-4xl md:text-6xl font-bold text-center bg-clip-text text-transparent bg-gradient-to-b from-neutral-50 to-neutral-400 bg-opacity-50 mb-4 !leading-tight">
              LLM Playground
            </h2>
            <Balance className="text-base text-neutral-300 mb-8 tracking-wide">
              If you are experimenting with LLMs, getting started with them can
              be cumbersome. We provides playground studio to play with and to
              make decision to use in a production-ready.
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
