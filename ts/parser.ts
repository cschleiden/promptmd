import * as fs from 'fs';
import * as yaml from 'yaml';

interface Prompt {
  role: string;
  content: string;
}

interface PromptFile {
  frontMatter: Record<string, any>;
  prompts: Prompt[];
}

function parsePromptFile(filePath: string): Promise<PromptFile> {
  return new Promise((resolve, reject) => {
    fs.readFile(filePath, 'utf8', (err, data) => {
      if (err) {
        return reject(err);
      }

      const lines = data.split('\n');
      let frontMatterLines: string[] = [];
      let prompts: Prompt[] = [];
      let currentRole: string = '';
      let currentContent: string[] = [];
      let inFrontMatter = false;
      let inPrompt = false;

      for (const line of lines) {
        if (line.trim() === '---') {
          if (!inFrontMatter) {
            inFrontMatter = true;
          } else {
            inFrontMatter = false;
            inPrompt = true;
          }
          continue;
        } else {
          if (!inFrontMatter && !inPrompt) {
            inPrompt = true;
          }
        }

        if (inFrontMatter) {
          frontMatterLines.push(line);
        } else if (inPrompt) {
          if (line.endsWith(':') && (line.startsWith('user') || line.startsWith('system') || line.startsWith('assistant'))) {
            if (currentRole) {
              prompts.push({
                role: currentRole,
                content: currentContent.join('\n'),
              });
              currentContent = [];
            }
            currentRole = line.slice(0, -1);
          } else {
            currentContent.push(line);
          }
        }
      }

      if (currentRole) {
        prompts.push({
          role: currentRole,
          content: currentContent.join('\n'),
        });
      }

      let frontMatter: Record<string, any> = {};
      if (frontMatterLines.length > 0) {
        try {
          frontMatter = yaml.parse(frontMatterLines.join('\n'));
        } catch (err) {
          return reject(err);
        }
      }

      resolve({
        frontMatter,
        prompts,
      });
    });
  });
}

export { Prompt, PromptFile, parsePromptFile };
