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
const [isFineTuning, setIsFineTuning] = useState<boolean>(false);

  useEffect(() => {
    const fetchEndpoint = isFineTuning ? `${BASE_URL}/api/v1/finetune/models` : `${BASE_URL}/api/v1/models`;
    async function fetchModels() {
      try {
        const response = await fetch(fetchEndpoint);
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
