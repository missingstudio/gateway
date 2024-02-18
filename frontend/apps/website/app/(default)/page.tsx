import CTA from "components/cta";
import AIGateway from "components/features/ai-gateway";
import AIRouter from "components/features/ai-router";
import Hero from "components/hero";
import { TiltedDashboard } from "components/tilt";

export default function Page(): JSX.Element {
  return (
    <>
      <Hero />
      <TiltedDashboard />
      <AIRouter />
      <AIGateway />
      <CTA />
    </>
  );
}
