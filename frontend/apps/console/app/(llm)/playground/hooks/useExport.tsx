import { useCallback } from "react";
import { useStore } from "~/app/(llm)/playground/store";

export const useExport = (): ((selected?: any[]) => void) => {
  const { logs } = useStore();
  const exportLogs = useCallback((selected?: any[]): void => {}, [logs]);
  return exportLogs;
};
