import { useEffect } from "react";
import { toast } from "sonner";
import { useStore } from "~/app/(llm)/playground/store";

export function useLogsFetch() {
  const { status, setLogs } = useStore();

  useEffect(() => {
    async function fetchLogs() {
      try {
        const response = await fetch("http://localhost:3000/v1/tracking/logs", {
          headers: {
            "x-ms-provider": "openai",
          },
        });
        const { logs = [] } = await response.json();
        setLogs(logs);
      } catch (e) {
        toast.error("Tracking API is not available");
      }
    }

    fetchLogs();
  }, [status, setLogs]);
}
