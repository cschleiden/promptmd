import { Prompt, PromptTemplate, Role, Roles } from "./types";
import { parse as parseYaml } from "yaml";

const frontmatterRegex = /^---\n([\s\S]*?)\n---\n/;

/**
 * parse parses the given prompt file into a Prompt object representing (optional) metadata and
 * message templates.
 *
 * @param content Input file
 * @returns Parsed prompt
 */
export function parse(content: string): Prompt {
  // Check if there is any frontmatter. Frontmatter is separated from the markdown content by `---`
  // before and after the content
  const frontMatterMatch = frontmatterRegex.exec(content);
  let metadata: { [key: string]: unknown } | undefined;

  if (frontMatterMatch) {
    const frontmatterContent = frontMatterMatch[1];
    metadata = parseYaml(frontmatterContent);

    // Remove frontmatter from prompt content
    const frontMatterEndIndex =
      frontMatterMatch.index + frontMatterMatch[0].length;
    content = content.substring(frontMatterEndIndex).trim();
  }

  const messages = parseMessages(content);

  return {
    metadata,
    messages,
  };
}

const roleRegex = new RegExp(
  `\\s*#?\\s*(${Roles.join("|")})\\s*:\\s*\\n`,
  "im"
);

function parseMessages(content: string): PromptTemplate[] {
  // Parse the template(s)
  const messages: PromptTemplate[] = [];

  // Trim and remove empty chunks
  const chunks = content.split(roleRegex).filter((x) => x.trim());

  // If it's just one chunk, default to a system message
  if (chunks.length === 1) {
    return [
      {
        role: "system",
        message: content.trim(),
      },
    ];
  }

  // Chunks should be even, as they are split by the regex
  if (chunks.length % 2 !== 0) {
    throw new Error("Invalid format");
  }

  for (let i = 0; i < chunks.length; i += 2) {
    const chunk = chunks[i];

    const role = chunk.trim() as Role;
    const message = chunks[i + 1].trim();

    messages.push({
      role,
      message,
    });
  }

  return messages;
}
