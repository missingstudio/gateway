"use client";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { Provider, useProvidersFetch } from "./hooks/useProvidersFetch";

export default function ProvidersList() {
  const { providers } = useProvidersFetch();

  return (
    <section>
      <div className="max-w-6xl mx-auto px-4 sm:px-6">
        <div className="pb-12 md:pb-20">
          <div>
            <div className="mt-12 md:mt-16">
              <h3
                id="engineering"
                className="scroll-mt-8 text-2xl font-bold inline-flex bg-clip-text text-transparent bg-gradient-to-r from-slate-800/60 via-slate-800 to-slate-800/60 dark:from-slate-200/60 dark:via-slate-200 dark:to-slate-200/60 pb-8"
              >
                <div className="flex justify-between items-center py-6 border-b [border-image:linear-gradient(to_right,transparent,theme(colors.slate.800),transparent)1] space-x-8 overflow-x-scroll no-scrollbar">
                  Providers
                </div>
              </h3>

              <div className="grid gap-6 sm:grid-cols-2 lg:grid-cols-3">
                {providers.map((provider, index) => (
                  <ProviderCard
                    key={provider.name}
                    provider={provider}
                    index={index}
                  />
                ))}
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}

type CardProps = {
  provider: Provider;
  index: number;
};

export function ProviderCard({ provider, index }: CardProps) {
  const pathname = usePathname();

  return (
    <div
      key={index}
      className="bg-gradient-to-tr from-slate-200 to-slate-200/25 dark:from-slate-800 dark:to-slate-800/25 rounded-3xl border border-slate-200 dark:border-slate-800 hover:border-slate-700/60 transition-colors group relative"
    >
      <div className="flex flex-col p-5 h-full">
        <div className="flex items-center space-x-3 mb-3">
          <Link
            className="font-semibold bg-clip-text text-transparent bg-gradient-to-r from-slate-800/60 via-slate-800 to-slate-800/60 dark:from-slate-200/60 dark:via-slate-200 dark:to-slate-200/60 group-hover:before:absolute group-hover:before:inset-0"
            href={`${pathname}/${provider.name}`}
          >
            {provider.title}
          </Link>
        </div>
        <div className="grow">
          <div className="text-sm text-slate-400 dark:text-slate-600">
            {provider.description}
          </div>
        </div>
      </div>
    </div>
  );
}
