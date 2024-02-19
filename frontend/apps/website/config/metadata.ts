import type { Metadata } from "next";
import { metaConfig } from "./meta";

const defaultMetadata: Metadata = {
  title: {
    template: `%s | ${metaConfig.applicationName}`,
    default: metaConfig.applicationName,
  },
  description: metaConfig.description,
  applicationName: metaConfig.applicationName,
  keywords: [
    "missing studio",
    "llmops",
    "ai studio",
    "ai router",
    "gateway",
    "ai gateway",
    "ai workflow",
    "universal api",
    "unified api",
    "local inference",
    "generative ai",
  ],
  icons: {
    icon: "/favicon/favicon.ico",
    shortcut: "/favicon-32x32.png",
    apple: "/apple-touch-icon.png",
  },
  creator: metaConfig.creator,
  twitter: metaConfig.twitter,
  openGraph: metaConfig.openGraph,
  alternates: { canonical: metaConfig.url },
  metadataBase: new URL(metaConfig.url),
};

export default defaultMetadata;
