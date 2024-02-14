import type { Metadata } from "next";
import Header from "~/components/header";
import { ThemeProvider } from "~/components/theme-provider";
import { Toaster } from "~/components/toaster";
import { spaceGrotesk } from "~/styles/fonts";
import "~/styles/globals.css";

export const metadata: Metadata = {
  title: "AIstudio",
  description: "A Robust Open Source AI studio",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en" suppressHydrationWarning>
      <body className={spaceGrotesk.className} suppressHydrationWarning>
        <ThemeProvider attribute="class" defaultTheme="dark" enableSystem>
          <Header />
          <Toaster />
          {children}
        </ThemeProvider>
      </body>
    </html>
  );
}
