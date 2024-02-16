import Image from "next/image";
import Link from "next/link";
import ProvidersImg from "~/public/logo.svg";
import Star from "~/public/star.svg";

export default function ProvidersList() {
  const items = [
    {
      img: ProvidersImg,
      name: "Openai",
      description:
        "Stellar makes it easy to build extensions by providing an authentication provider that handles the OAuth flow.",
      link: "/providers/openai",
      featured: false,
      category: "LLM",
    },
    {
      img: ProvidersImg,
      name: "Azure",
      description:
        "Stellar makes it easy to build extensions by providing an authentication provider that handles the OAuth flow.",
      link: "/providers/azure",
      featured: false,
      category: "LLM",
    },
    {
      img: ProvidersImg,
      name: "Anyscale",
      description:
        "Stellar makes it easy to build extensions by providing an authentication provider that handles the OAuth flow.",
      link: "/providers/anyscale",
      featured: false,
      category: "LLM",
    },
    {
      img: ProvidersImg,
      name: "Deepinfra",
      description:
        "Stellar makes it easy to build extensions by providing an authentication provider that handles the OAuth flow.",
      link: "/providers/deepinfra",
      featured: false,
      category: "LLM",
    },
    {
      img: ProvidersImg,
      name: "Togetherai",
      description:
        "Stellar makes it easy to build extensions by providing an authentication provider that handles the OAuth flow.",
      link: "/providers/togetherai",
      featured: false,
      category: "LLM",
    },
  ];

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
                {items.map(
                  (item, index) =>
                    item.category === "LLM" && (
                      <ProviderCard item={item} index={index} />
                    )
                )}
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}

type CardProps = {
  item: {
    category: string;
    img: string;
    name: string;
    featured: boolean;
    link: string;
    description: string;
  };
  index: number;
};

export function ProviderCard({ item, index }: CardProps) {
  return (
    <div
      key={index}
      className="bg-gradient-to-tr from-slate-200 to-slate-200/25 dark:from-slate-800 dark:to-slate-800/25 rounded-3xl border border-slate-200 dark:border-slate-800 hover:border-slate-700/60 transition-colors group relative"
    >
      <div className="flex flex-col p-5 h-full">
        <div className="flex items-center space-x-3 mb-3">
          <div className="relative">
            <Image src={item.img} width="40" height="40" alt={item.name} />
            {item.featured && (
              <Image
                className="absolute top-0 -right-1"
                src={Star}
                width={16}
                height={16}
                alt="Star"
                aria-hidden="true"
              />
            )}
          </div>
          <Link
            className="font-semibold bg-clip-text text-transparent bg-gradient-to-r from-slate-800/60 via-slate-800 to-slate-800/60 dark:from-slate-200/60 dark:via-slate-200 dark:to-slate-200/60 group-hover:before:absolute group-hover:before:inset-0"
            href={item.link}
          >
            {item.name}
          </Link>
        </div>
        <div className="grow">
          <div className="text-sm text-slate-400 dark:text-slate-600">
            {item.description}
          </div>
        </div>
      </div>
    </div>
  );
}
