import Image from "next/image";
import Link from "next/link";
import ProvidersIcon from "~/public/providers-icon.svg";
import ProvidersImg from "~/public/providers-image.png";
import Star from "~/public/star.svg";

export default function SingleProvider() {
  return (
    <section className="relative">
      <div className="max-w-6xl mx-auto px-4 sm:px-6">
        <div className="pt-32 md:pt-40">
          <div className="md:flex md:justify-between">
            <div className="md:grow pb-12 md:pb-20">
              <div className="max-w-[720px]">
                <div className="flex flex-col lg:flex-row space-y-6 lg:space-y-0 lg:space-x-16">
                  <div className="shrink-0">
                    <div className="sticky top-6">
                      <Link
                        className="flex items-center justify-center w-9 h-9 group border border-transparent rounded-full [background:linear-gradient(theme(colors.slate.900),_theme(colors.slate.900))_padding-box,_conic-gradient(theme(colors.slate.400),_theme(colors.slate.700)_25%,_theme(colors.slate.700)_75%,_theme(colors.slate.400)_100%)_border-box] relative before:absolute before:inset-0 before:bg-slate-800/30 before:rounded-full before:pointer-events-none"
                        href="/providers"
                      >
                        <span className="sr-only">Go back</span>
                        <svg
                          className="w-4 h-4 fill-purple-500"
                          viewBox="0 0 16 16"
                          xmlns="http://www.w3.org/2000/svg"
                        >
                          <path d="M6.7 14.7l1.4-1.4L3.8 9H16V7H3.8l4.3-4.3-1.4-1.4L0 8z" />
                        </svg>
                      </Link>
                    </div>
                  </div>

                  <div>
                    <article className="pb-12 mb-12 border-b [border-image:linear-gradient(to_right,transparent,theme(colors.slate.800),transparent)1]">
                      <figure className="bg-slate-300/20 dark:bg-slate-700/20 border border-slate-700/10 dark:border-slate-300/10 p-4 rounded-3xl mb-8">
                        <Image
                          className="w-full rounded-2xl"
                          src={ProvidersImg}
                          width={586}
                          height={316}
                          alt="Provider image"
                        />
                      </figure>

                      <h1 className="sr-only">GitHub</h1>

                      <div className="prose max-w-none text-slate-600 dark:text-slate-400 prose-headings:text-slate-50 prose-h2:text-xl prose-h2:mt-8 prose-h2:mb-4 prose-p:leading-relaxed prose-a:text-purple-500 prose-a:no-underline hover:prose-a:underline prose-strong:text-slate-50 prose-strong:font-medium prose-blockquote:pl-5 prose-blockquote:xl:-ml-5 prose-blockquote:border-l-2 prose-blockquote:border-purple-500 prose-blockquote:font-medium prose-blockquote:text-slate-300 prose-blockquote:italic">
                        <h2>Overview</h2>
                        <p>
                          This powerful GitHub provider keeps your work in sync
                          in both applications. It links features to Pull
                          Requests so that details update automatically from In
                          Progress to Done as the PR moves from drafted to
                          merged - there is no need to update the issue in
                          Stellar at all.
                        </p>
                        <p>
                          Move even faster by using a keyboard shortcut that
                          creates the issue's git branch name, assigns the issue
                          and moves the issue to In Progress in one step. This
                          provider supports <a href="#0">GitHub Enterprise</a>.
                        </p>
                        <h2 className=" text-black">How it works</h2>
                        <p>
                          Configure the GitHub provider{" "}
                          <strong className="text-slate-900 dark:text-slate-50 font-medium">
                            in a couple of minutes
                          </strong>{" "}
                          by authenticating and selecting which repos to connect
                          to Stellar. Once you do, anyone in the workspace can
                          start linking PRs.
                        </p>
                        <p>
                          Link PRs to issues in a few ways. Use the issue ID in
                          the branch name or PR title before creating the PR. To
                          link multiple PRs, or link issues after creating the
                          PR, use magic words along with the issue ID in the PR
                          description.
                        </p>
                      </div>
                    </article>

                    <aside className="pl-6 border-l-2 border-purple-500">
                      <p className="inline-flex font-medium italic text-lg bg-clip-text text-transparent bg-gradient-to-r from-slate-800/60 via-slate-800 to-slate-800/60 dark:from-slate-200/60 dark:via-slate-200 dark:to-slate-200/60 pb-4">
                        “ This provider is so perfect it tickles my brain. It
                        ensures that your work is in sync across your entire
                        team. ”
                      </p>
                      <footer className="flex items-center space-x-4">
                        <Image
                          className="shrink-0 rounded-full"
                          src={ProvidersImg}
                          width={32}
                          height={32}
                          alt="Author"
                        />
                        <div className="text-sm font-medium text-slate-700 dark:text-slate-300">
                          Mike Hunt{" "}
                          <span className="text-slate-300 dark:text-slate-700">
                            -
                          </span>{" "}
                          <a
                            className="text-purple-500 hover:underline"
                            href="#0"
                          >
                            Thunderbolt
                          </a>
                        </div>
                      </footer>
                    </aside>
                  </div>
                </div>
              </div>
            </div>

            <aside className="md:w-64 lg:w-80 md:shrink-0 md:pt-[3.75rem] lg:pt-0 pb-12 md:pb-20">
              <div className="sticky top-6 md:pl-6 lg:pl-10">
                <div className="space-y-6">
                  <div className="bg-gradient-to-tr from-slate-200 to-slate-200/25 dark:from-slate-800 dark:to-slate-800/25 rounded-3xl border border-slate-200 dark:border-slate-800">
                    <div className="px-5 py-6">
                      <div className="text-center mb-5">
                        <div className="mb-4">
                          <div className="relative inline-flex">
                            <Image
                              src={ProvidersIcon}
                              width={80}
                              height={80}
                              alt="Icon 08"
                            />
                            <Image
                              className="absolute top-0 -right-1"
                              src={Star}
                              width={24}
                              height={24}
                              alt="Star"
                              aria-hidden="true"
                            />
                          </div>
                        </div>
                      </div>
                      <ul className="text-sm">
                        <li className="flex items-center justify-between space-x-2 py-3 border-t [border-image:linear-gradient(to_right,theme(colors.slate.700/.3),theme(colors.slate.700),theme(colors.slate.700/.3))1]">
                          <span className="text-slate-600 dark:text-slate-400">
                            Company
                          </span>
                          <span className="text-slate-700 dark:text-slate-300 font-medium">
                            Github
                          </span>
                        </li>
                        <li className="flex items-center justify-between space-x-2 py-3 border-t [border-image:linear-gradient(to_right,theme(colors.slate.700/.3),theme(colors.slate.700),theme(colors.slate.700/.3))1]">
                          <span className="text-slate-600 dark:text-slate-400">
                            Website
                          </span>
                          <a
                            className="text-purple-500 font-medium flex items-center space-x-1"
                            href="#0"
                          >
                            <span>github.com</span>
                            <svg
                              className="fill-current"
                              xmlns="http://www.w3.org/2000/svg"
                              width="9"
                              height="9"
                            >
                              <path d="m1.285 8.514-.909-.915 5.513-5.523H1.663l.01-1.258h6.389v6.394H6.794l.01-4.226z" />
                            </svg>
                          </a>
                        </li>
                        <li className="flex items-center justify-between space-x-2 py-3 border-t [border-image:linear-gradient(to_right,theme(colors.slate.700/.3),theme(colors.slate.700),theme(colors.slate.700/.3))1]">
                          <span className="text-slate-600 dark:text-slate-400">
                            Industry
                          </span>
                          <span className="text-slate-700 dark:text-slate-300 font-medium">
                            Technology
                          </span>
                        </li>
                        <li className="flex items-center justify-between space-x-2 py-3 border-t [border-image:linear-gradient(to_right,theme(colors.slate.700/.3),theme(colors.slate.700),theme(colors.slate.700/.3))1]">
                          <span className="text-slate-600 dark:text-slate-400">
                            Installs
                          </span>
                          <span className="text-slate-700 dark:text-slate-300 font-medium">
                            100K+
                          </span>
                        </li>
                      </ul>
                    </div>
                  </div>
                </div>
              </div>
            </aside>
          </div>
        </div>
      </div>
    </section>
  );
}
