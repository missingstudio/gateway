"use client";

import AOS from "aos";
import "aos/dist/aos.css";
import { useEffect } from "react";
import Footer from "~/components/footer";
import Header from "~/components/header";

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
      <main className="grow">{children}</main>
      <Footer />
    </>
  );
}
