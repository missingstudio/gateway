"use client";
import Image from "next/image";
import Client01 from "public/images/client-01.svg";
import Marquee from "react-fast-marquee";

import Particles from "./particles";
export default function Clients() {
  return (
    <section>
      <div className="relative max-w-6xl mx-auto px-4 sm:px-6">
        <div className="absolute inset-0 max-w-6xl mx-auto px-4 sm:px-6">
          <Particles className="absolute inset-0 -z-10" quantity={5} />
        </div>

        <div className="py-12 md:py-16">
          <Marquee
            direction={"left"}
            speed={30}
            gradient={true}
            gradientColor="#0d1424"
            gradientWidth="100px"
            autoFill
          >
            <Image
              src={Client01}
              alt="Client 01"
              width={110}
              height={21}
              className="mx-8"
            />
          </Marquee>
        </div>
      </div>
    </section>
  );
}
