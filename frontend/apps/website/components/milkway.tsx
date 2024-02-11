"use client";

import { useEffect, useRef } from "react";

export const MilkeyWay = () => {
  return (
    <div className="absolute inset-0 z-[-1] w-full">
      {[...Array(20)].map((_, i) => (
        <Star key={i} />
      ))}
      {[...Array(3)].map((_, i) => (
        <AnimatedLine size="large" key={i} />
      ))}
    </div>
  );
};

export const Star = () => {
  const ref = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (!ref.current) return;

    // Generate a random number with high chances of being between 30 and 70,
    // and low chances of being between 0 and 100.
    function getPosition() {
      const rand = Math.random() * 100;

      if (rand < 30) {
        return Math.random() * 100;
      } else {
        return Math.random() * 40 + 30;
      }
    }

    const x = getPosition();
    const y = getPosition();
    const size = Math.random() * 2;
    const duration = Math.random() * 5 + 5;

    ref.current.style.top = `${y}%`;
    ref.current.style.left = `${x}%`;
    ref.current.style.width = `${1 + size}px`;
    ref.current.style.height = `${1 + size}px`;
    ref.current.style.animation = `star ${duration}s ease-in infinite`;
  }, [ref]);

  return <span ref={ref} className="absolute rounded-full bg-white" />;
};

type AnimatedLineProps = {
  size: "small" | "large";
};

export const AnimatedLine = ({ size }: AnimatedLineProps) => {
  const line = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (line.current) {
      const delay = (Math.random() * 3 + 3) * 1000;

      line.current.style.top = Math.random() * 50 + "%";
      line.current.style.left = Math.random() * 100 + "%";
      line.current.style.animation = `line-bottom-left ${delay}ms ease-in-out infinite`;

      const interval = setInterval(() => {
        if (line.current) {
          line.current.style.top = Math.random() * 100 + "%";
          line.current.style.left = Math.random() * 100 + "%";
        }
      }, delay);

      return () => {
        clearInterval(interval);
      };
    }
  }, [line, size]);

  if (size === "small") {
    return (
      <div
        ref={line}
        className="absolute hidden h-32 w-[1px] rotate-45 transform opacity-0 lg:block"
        style={{
          backgroundImage:
            "linear-gradient(to bottom, transparent, rgba(255, 255, 255, 0.2))",
        }}
      />
    );
  }

  return (
    <div
      ref={line}
      className="absolute hidden h-40 w-[2px] rotate-[-135deg] transform opacity-0 lg:block"
      style={{
        backgroundImage:
          "linear-gradient(to bottom, transparent, rgb(255, 255, 255))",
      }}
    />
  );
};
