import { useCallback } from "react";
import { useStore } from "~/app/(llm)/playground/store";

export const useChat = (): (() => Promise<void>) => {
  const { input, provider, model, parameters, setOutput, setStatus } =
    useStore();

  const submitChat = useCallback(async () => {
    setStatus("waiting");
    setOutput("");
  }, [input, model, parameters, provider, setStatus, setOutput]);

  return submitChat;
};
