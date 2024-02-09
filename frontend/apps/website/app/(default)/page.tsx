import CTA from "~/components/cta";
import FeaturesAPIGateway from "~/components/features/gateway";
import FeaturesUniversalAPI from "~/components/features/universal-api";
import Hero from "~/components/hero";

export default function Page(): JSX.Element {
  return (
    <>
      <Hero />
      <FeaturesUniversalAPI />
      <FeaturesAPIGateway />
      <CTA />
    </>
  );
}
