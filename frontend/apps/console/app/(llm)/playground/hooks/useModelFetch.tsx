import { useEffect, useState } from "react";
import { toast } from "sonner";
import models from "./data/models.json";

interface ModelType {
  name: string;
  models: string[];
}

export function useModelFetch() {
  const [providers, setProviders] = useState<ModelType[]>([]);

  useEffect(() => {
    async function fetchModels() {
      try {
        var data = (models as Record<string, ModelType>) || {};
        const fetchedProviders: ModelType[] = Object.keys(data).map((key) => ({
          name: (data[key] as ModelType).name,
          models: (data[key] as ModelType).models.map(
            (modelName: string, index: number) => modelName
          ),
        }));
        setProviders(fetchedProviders);
      } catch (e) {
        toast.error("AIstudio is not running");
      }
    }

    fetchModels();
  }, []);

  return { providers };
}
