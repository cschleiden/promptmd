import { prepare, Vars } from "./template";
import { Message, Prompt } from "./types";

export type preparedPromptFunc = (args: {
  [variableName: string]: unknown;
}) => Message[];

export function preparePrompt(p: Prompt): preparedPromptFunc {
  const preparedMessages = p.messages.map((m) => prepare(m.message));

  return (vars: Vars) => {
    return p.messages.map((m, i) => {
      return {
        role: m.role,
        message: preparedMessages[i](vars),
      };
    });
  };
}
