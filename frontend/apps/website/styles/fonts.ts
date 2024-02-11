import { Lora, Open_Sans, Ubuntu, Work_Sans } from "next/font/google";
import localFont from "next/font/local";

export const opensans = Open_Sans({
  variable: "--font-sans",
  subsets: ["latin"],
  weight: ["300", "400", "500", "700"],
});

export const inter = Ubuntu({
  variable: "--font-sans",
  subsets: ["latin"],
  weight: ["300", "400", "500", "700"],
});

export const calsans = localFont({
  src: "../assets/fonts/CalSans-SemiBold.woff2",
  variable: "--font-heading",
});

export const lora = Lora({
  variable: "--font-heading",
  subsets: ["latin"],
  weight: "600",
  display: "swap",
});

export const work = Work_Sans({
  variable: "--font-heading",
  subsets: ["latin"],
  weight: "600",
  display: "swap",
});
