export const Roles = ["user", "assistant", "system"] as const;
export type Role = (typeof Roles)[number];

export type PromptTemplate = {
  role: Role;
  message: string;
};

export type Prompt = {
  metadata: { [key: string]: unknown } | undefined;

  messages: PromptTemplate[];
};
