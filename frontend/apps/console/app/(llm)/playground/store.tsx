import { create } from "zustand";

export interface Log {
  provider: string;
  model: string;
  latency: string;
  prompt_tokens: number;
  completion_tokens: number;
  total_tokens: number;
}

interface Store {
  status: string;
  logs: Log[];
  input: string;
  output: string;
  provider: string;
  model: string;
  parameters: Record<string, unknown>;
  setStatus: (status: string) => void;
  setLogs: (logs: Log[]) => void;
  setInput: (input: string) => void;
  setOutput: (output: string, isChunk?: boolean) => void;
  setModel: (model: string) => void;
  setProvider: (provider: string) => void;
  setParameter: (parameter: string, value: any) => void;
  getParameter: (
    parameter: string
  ) => string | number | readonly string[] | undefined;
}

export const useStore = create<Store>()((set, get) => ({
  status: "idle",
  logs: [],
  input: "",
  output: "",
  provider: "Open AI",
  model: "gpt-3.5-turbo",
  parameters: {},
  setStatus: (status) => set({ status: status }),
  setLogs: (logs: Log[]) => set({ logs: logs }),
  setInput: (input) => set({ input: input }),
  setOutput: (output, isChunk = false) =>
    set((state) => ({
      output: isChunk ? state.output + output : output,
    })),
  setModel: (model) => set({ model: model }),
  setProvider: (provider) => set({ provider: provider }),
  setParameter: (parameter, value) => {
    set((state) => ({
      parameters: {
        ...state.parameters,
        [parameter]: Number(value),
      },
    }));
  },
  getParameter: (parameter: string) =>
    get().parameters[parameter] as
      | string
      | number
      | readonly string[]
      | undefined,
}));
