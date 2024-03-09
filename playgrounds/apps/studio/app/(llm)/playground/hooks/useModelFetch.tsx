import { useEffect, useState } from "react";
import { toast } from "sonner";

interface ModelType {
  name: string;
  models: Model[];
}

interface Model {
  name: string;
  value: string;
}

const BASE_URL = process.env.NEXT_PUBLIC_GATEWAY_URL ?? "http://localhost:3000";
export function useModelFetch() {
  const [providers, setProviders] = useState<ModelType[]>([]);

  useEffect(() => {
    async function fetchModels() {
      try {
        const response = await fetch(`${BASE_URL}/api/v1/models`);
        const { models } = await response.json();
        const fetchedProviders: ModelType[] = Object.keys(models).map(
          (key) => ({
            name: models[key].name,
            models: models[key].models,
          })
        );
        setProviders(fetchedProviders);
      } catch (e) {
        toast.error("AI Studio is not running");
      }
    }

    fetchModels();
  }, []);

  return { providers };
}
