import { useEffect, useState } from "react";
import { toast } from "sonner";

export interface Provider {
  name: string;
  title: string;
  description: string;
}

const BASE_URL = process.env.NEXT_PUBLIC_GATEWAY_URL ?? "http://localhost:3000";
export function useProviderFetch(pname: string) {
  const [provider, setProvider] = useState<Provider | undefined>();
  const [config, setConfig] = useState();

  useEffect(() => {
    async function fetchProvider() {
      try {
        const response = await fetch(`${BASE_URL}/api/v1/providers/${pname}`);
        const { provider } = await response.json();

        setProvider(provider);
      } catch (e) {
        toast.error("AI Studio is not running");
      }
    }

    fetchProvider();
  }, []);
  useEffect(() => {
    async function fetchProviderConfig() {
      try {
        const response = await fetch(
          `${BASE_URL}/api/v1/providers/${pname}/config`
        );
        const { config } = await response.json();

        setConfig(config);
      } catch (e) {
        toast.error("AI Studio is not running");
      }
    }

    fetchProviderConfig();
  }, []);

  return { provider, config };
}
