"use client";

import { useEffect, useRef } from "react";
import styles from "./tilt.module.css";

export const TiltedDashboard = () => {
  const dashboardRef = useRef(null);

  useEffect(() => {
    const observerOptions = {
      root: null,
      rootMargin: "0px",
      threshold: 0.5,
    };

    const observer = new IntersectionObserver((entries) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          entry.target.classList.add(styles.hero_active as string);
        } else {
          entry.target.classList.remove(styles?.hero_active as string);
        }
      });
    }, observerOptions);

    if (dashboardRef.current) {
      observer.observe(dashboardRef.current);
    }

    return () => {
      if (dashboardRef.current) {
        observer.unobserve(dashboardRef.current);
      }
    };
  }, []);

  return (
    <section className="space-y-6">
      <div
        className={`container m-auto max-w-6xl relative ${styles.hero_container} px-4`}
      >
        <div
          ref={dashboardRef}
          className={`flex rounded-lg md:rounded-xl lg:rounded-3xl overflow-hidden will-change  opacity transition ease-in-out duration-1000  ${styles.hero}`}
        >
          <img
            src="/images/dashboard.jpg"
            alt="missing studio dashboard"
            loading="eager"
          />
        </div>
      </div>
    </section>
  );
};
