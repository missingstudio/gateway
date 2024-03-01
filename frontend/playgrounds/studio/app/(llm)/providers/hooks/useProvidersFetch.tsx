import { useEffect, useState } from "react";
import { toast } from "sonner";

export interface Provider {
  name: string;
  title: string;
  description: string;
}

const BASE_URL = process.env.NEXT_PUBLIC_GATEWAY_URL ?? "http://localhost:3000";
export function useProvidersFetch() {
  const [providers, setProviders] = useState<Provider[]>([]);

  useEffect(() => {
    async function fetchModels() {
      try {
        const response = await fetch(`${BASE_URL}/api/v1/providers`);
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
