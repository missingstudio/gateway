import { useEffect, useState } from "react";
import { toast } from "sonner";

export interface Provider {
  name: string;
  title: string;
  description: string;
}

export function useProviderFetch(pname: string) {
  const [provider, setProvider] = useState<Provider>();
  const [config, setConfig] = useState();

  useEffect(() => {
    async function fetchProvider() {
      try {
        const response = await fetch(
          `http://localhost:3000/v1/providers/${pname}`,
          {
            headers: {
              "x-ms-provider": "openai",
            },
          }
        );
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
          `http://localhost:3000/v1/providers/${pname}/config`,
          {
            headers: {
              "x-ms-provider": "openai",
            },
          }
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
