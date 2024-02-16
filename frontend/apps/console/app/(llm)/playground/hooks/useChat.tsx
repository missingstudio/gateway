import { useCallback } from "react";
import { toast } from "sonner";
import { useStore } from "~/app/(llm)/playground/store";

export const useChat = (): (() => Promise<void>) => {
  const { input, provider, model, parameters, setOutput, setStatus } =
    useStore();

  const submitChat = useCallback(async () => {
    setStatus("waiting");
    setOutput("");

    async function chatCompletions() {
      try {
        const response = await fetch(
          "http://localhost:3000/v1/chat/completions",
          {
            method: "post",
            headers: {
              "x-ms-provider": "openai",
              "Content-Type": "application/json",
            },
            body: JSON.stringify({
              messages: [
                {
                  role: "user",
                  content: input,
                },
              ],
              model: model,
              temperature: parameters?.temperature,
              max_tokens: parameters?.max_tokens,
              top_p: parameters?.top_p,
              frequency_penalty: parameters?.frequency_penalty,
              presence_penalty: parameters?.presence_penalty,
            }),
          }
        );
        const data = await response.json();
        if (data?.choices.length) {
          setOutput(data.choices[0]?.message?.content);
        } else {
          setOutput("");
        }
      } catch (e: any) {
        toast.error(e.message);
        setStatus("error");
      } finally {
        setStatus("done");
      }
    }

    chatCompletions();
  }, [input, model, parameters, provider, setStatus, setOutput]);

  return submitChat;
};
