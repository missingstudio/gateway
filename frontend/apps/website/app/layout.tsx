import { GoogleAnalytics } from "@next/third-parties/google";
import defaultMetadata from "config/metadata";
import type { Metadata } from "next";
import { outfit } from "styles/fonts";
import "styles/globals.css";
import { Providers } from "./providers";

export const metadata: Metadata = defaultMetadata;
export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}): JSX.Element {
  return (
    <html lang="en" suppressHydrationWarning>
      <body
        className={`${outfit.className} antialiased bg-[#101010] text-neutral-100 tracking-tight`}
      >
        <div className="flex flex-col min-h-screen overflow-hidden">
          <Providers>{children}</Providers>
        </div>
      </body>
      <GoogleAnalytics gaId="G-Y922M67ET7" />
    </html>
  );
}
