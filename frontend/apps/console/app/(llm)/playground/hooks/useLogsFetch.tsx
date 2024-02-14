import { useEffect } from "react";
import { useStore } from "~/app/(llm)/playground/store";

export function useLogsFetch() {
  const { status, setLogs } = useStore();

  useEffect(() => {
    async function fetchLogs() {}
    fetchLogs();
  }, [status, setLogs]);
}
