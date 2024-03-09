import { useEffect } from "react";
import { toast } from "sonner";
import { useStore } from "~/app/(llm)/playground/store";

const BASE_URL = process.env.NEXT_PUBLIC_GATEWAY_URL ?? "http://localhost:3000";
export function useLogsFetch() {
  const { status, setLogs } = useStore();

  useEffect(() => {
    async function fetchLogs() {
      try {
        const response = await fetch(`${BASE_URL}/api/v1/tracking/logs`);
        const { logs = [] } = await response.json();
        setLogs(logs);
      } catch (e) {
        toast.error("Tracking API is not available");
      }
    }

    fetchLogs();
  }, [status, setLogs]);
}
