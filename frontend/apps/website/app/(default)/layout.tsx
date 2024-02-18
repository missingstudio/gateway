"use client";

import AOS from "aos";
import "aos/dist/aos.css";
import Footer from "components/footer";
import Header from "components/header";
import { Spotlight } from "components/spotlight";
import { useEffect } from "react";

export default function DefaultLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  useEffect(() => {
    AOS.init({
      once: true,
      disable: "phone",
      duration: 1000,
      easing: "ease-out-cubic",
    });
  });

  return (
    <>
      <Header />
      <Spotlight
        className="-top-40 left-0 lg:left-[25rem] lg:-top-[4rem]"
        fill="white"
      />
      <main className="grow">{children}</main>
      <Footer />
    </>
  );
}
