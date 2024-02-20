import { useEffect, useState } from "react";
import { toast } from "sonner";

export interface Provider {
  name: string;
  title: string;
  description: string;
}

export function useProvidersFetch() {
  const [providers, setProviders] = useState<Provider[]>([]);

  useEffect(() => {
    async function fetchModels() {
      try {
        const response = await fetch("http://localhost:3000/v1/providers", {
          headers: {
            "x-ms-provider": "openai",
          },
        });
        const { providers } = await response.json();

        setProviders(providers);
      } catch (e) {
        toast.error("AI Studio is not running");
      }
    }

    fetchModels();
  }, []);

  return { providers };
}
