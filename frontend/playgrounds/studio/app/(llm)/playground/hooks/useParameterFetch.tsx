import { useEffect, useState } from "react";
import { toast } from "sonner";
import { useStore } from "~/app/(llm)/playground/store";
import parametersData from "./data/parameters.json";

export interface ParameterType {
  id: string;
  name: string;
  type: string;
  default: number;
  min?: number;
  max?: number;
  step?: number;
}

export function useParameterFetch() {
  const [parameters, setParameters] = useState<ParameterType[]>([]);
  const { provider } = useStore();

  useEffect(() => {
    async function fetchParameters() {
      try {
        var data = (parametersData as ParameterType[]) || [];
        setParameters(data);
      } catch (e) {
        toast.error("AIstudio Engine is not running");
      }
    }

    if (provider) fetchParameters();
  }, [provider]);

  return { parameters };
}
