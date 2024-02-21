"use client";

import { EyeClosedIcon, EyeOpenIcon } from "@radix-ui/react-icons";

import { Button } from "@missingstudio/ui/button";
import { Input } from "@missingstudio/ui/input";
import { Loader } from "lucide-react";
import Image from "next/image";
import Link from "next/link";
import { useParams } from "next/navigation";
import mergeDeepRight from "ramda/es/mergeDeepRight";
import pathOr from "ramda/es/pathOr";

import { Controller, SubmitHandler, useForm } from "react-hook-form";
import { toast } from "sonner";
import * as z from "zod";

import { Label } from "@missingstudio/ui/label";
import { HTMLInputTypeAttribute, useEffect, useState } from "react";
import ProvidersIcon from "~/public/providers-icon.svg";
import ProvidersImg from "~/public/providers-image.png";
import Star from "~/public/star.svg";
import { useProviderFetch } from "../hooks/useProviderFetch";

const initialValues = {
  config: {
    headers: {
      Authorization: "",
    },
  },
};
const formSchema = z.object({
  config: z.object({
    headers: z.object({
      Authorization: z.string({ required_error: "Authorization is required" }),
    }),
  }),
});
type FormValues = z.infer<typeof formSchema>;

const BASE_URL = process.env.NEXT_PUBLIC_GATEWAY_URL ?? "http://localhost:3000";
export default function SingleProvider() {
  const [authInputType, setAuthInputType] =
    useState<HTMLInputTypeAttribute>("password");
  const params = useParams<{ providerId: string }>();
  const { config, provider } = useProviderFetch(params.providerId);

  const defaultValue = mergeDeepRight(initialValues, provider ?? {});

  const {
    control,
    getValues,
    handleSubmit,
    formState: { errors, isValid, isLoading },
    reset,
  } = useForm({
    defaultValues: defaultValue,
    mode: "all",
  });

  useEffect(() => {
    reset(defaultValue);
  }, [provider]);

  const onSubmit: SubmitHandler<FormValues> = async () => {
    try {
      const data = getValues();
      const response = await fetch(
        `${BASE_URL}/v1/providers/${params.providerId}`,
        {
          method: "put",
          body: JSON.stringify(data),
          headers: {
            "Content-Type": "application/json",
          },
        }
      );
      await response.json();
      toast.info("Updated", {
        description: "Configuration has been updated",
      });
    } catch (error) {
      toast.error("Something went wrong", {
        description: "Not able to update configurations",
      });
    }
  };

  const headersProperties = pathOr({}, ["properties", "headers", "properties"])(
    config
  ) as Record<string, any>;

  const requiredHeadersProperties = pathOr(
    [],
    ["properties", "headers", "required"]
  )(config) as String[];

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
                      </div>
                    </article>
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
                      <div className="flex items-center justify-between space-x-2 py-3 border-t [border-image:linear-gradient(to_right,theme(colors.slate.700/.3),theme(colors.slate.700),theme(colors.slate.700/.3))1]" />
                      <form
                        onSubmit={handleSubmit(onSubmit)}
                        className="flex flex-col gap-2"
                      >
                        {Object.keys(headersProperties).map((hp: string) => {
                          const { title, description } = headersProperties[hp];
                          return (
                            <div key="hp" className="relative">
                              <Label>{title}</Label>
                              <Controller
                                key={hp}
                                render={({ field: { value, onChange } }) => (
                                  <>
                                    <Input
                                      className="col-span-2"
                                      type={authInputType}
                                      // @ts-ignore
                                      value={value}
                                      onChange={onChange}
                                      placeholder={description}
                                    />
                                    {authInputType == "password" ? (
                                      <EyeOpenIcon
                                        className="absolute bottom-3 right-2 bg-white p-[2px] w-4 cursor-pointer"
                                        onClick={() => setAuthInputType("text")}
                                      />
                                    ) : (
                                      <EyeClosedIcon
                                        className="absolute bottom-3 right-2 bg-white p-[2px] w-4 cursor-pointer"
                                        onClick={() =>
                                          setAuthInputType("password")
                                        }
                                      />
                                    )}
                                  </>
                                )}
                                // @ts-ignore
                                name={`config.headers.${hp}`}
                                control={control}
                                rules={{
                                  required:
                                    requiredHeadersProperties.includes(hp),
                                }}
                              />
                            </div>
                          );
                        })}

                        <Button
                          type="submit"
                          className="@lg:w-auto float-right w-full dark:bg-gray-100 dark:text-white dark:active:bg-gray-100"
                          disabled={!isValid}
                        >
                          {isLoading ? (
                            <Loader color="primary" />
                          ) : (
                            "update API key"
                          )}
                        </Button>
                      </form>
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
