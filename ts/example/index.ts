import { readFileSync } from "fs";
import OpenAI from "openai";
import { join } from "path";
import { parse, preparePrompt } from "promptmd";

async function main() {
  const promptContent = readFileSync(
    join(__dirname, "recipe.prompt.md"),
    "utf-8"
  );
  const prompt = parse(promptContent);

  // Use the parsed prompt
  const pp = preparePrompt(prompt);
  const msgs = pp({
    ingredient: "chicken",
  });

  // Use GitHub Models to prompt the model
  const token = process.env["GITHUB_TOKEN"];
  const endpoint = "https://models.inference.ai.azure.com";
  const modelName = prompt.metadata["model"] as string;

  const client = new OpenAI({ baseURL: endpoint, apiKey: token });

  const response = await client.chat.completions.create({
    messages: msgs.map((m) => ({
      role: m.role,
      content: m.message,
    })),
    temperature: 1.0,
    top_p: 1.0,
    max_tokens: 1000,
    model: modelName,
  });

  console.log(response.choices[0].message.content);
}

main().catch((err) => {
  console.error("Error:", err);
  process.exit(1);
});
