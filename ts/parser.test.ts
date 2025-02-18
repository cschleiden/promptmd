import { parsePromptFile } from './parser';
import * as fs from 'fs';
import * as path from 'path';

describe('parsePromptFile', () => {
  const testDir = path.join(__dirname, 'test_files');

  beforeAll(() => {
    if (!fs.existsSync(testDir)) {
      fs.mkdirSync(testDir);
    }
  });

  afterAll(() => {
    fs.rmdirSync(testDir, { recursive: true });
  });

  function createTempFile(content: string): string {
    const filePath = path.join(testDir, `${Date.now()}.prompt.md`);
    fs.writeFileSync(filePath, content);
    return filePath;
  }

  test('empty file', async () => {
    const filePath = createTempFile('');
    const promptFile = await parsePromptFile(filePath);

    expect(promptFile.frontMatter).toEqual({});
    expect(promptFile.prompts).toEqual([]);
  });

  test('only front matter', async () => {
    const content = `---
title: Example Prompt
description: This is an example prompt file.
---`;
    const filePath = createTempFile(content);
    const promptFile = await parsePromptFile(filePath);

    expect(promptFile.frontMatter).toEqual({
      title: 'Example Prompt',
      description: 'This is an example prompt file.',
    });
    expect(promptFile.prompts).toEqual([]);
  });

  test('only prompts', async () => {
    const content = `system:
You are a helpful assistant.

user:
What is the weather like today?

assistant:
The weather is sunny with a high of 75 degrees.`;
    const filePath = createTempFile(content);
    const promptFile = await parsePromptFile(filePath);

    expect(promptFile.frontMatter).toEqual({});
    expect(promptFile.prompts).toEqual([
      { role: 'system', content: 'You are a helpful assistant.\n' },
      { role: 'user', content: 'What is the weather like today?\n' },
      { role: 'assistant', content: 'The weather is sunny with a high of 75 degrees.\n' },
    ]);
  });

  test('front matter and prompts', async () => {
    const content = `---
title: Example Prompt
description: This is an example prompt file.
---

system:
You are a helpful assistant.

user:
What is the weather like today?

assistant:
The weather is sunny with a high of 75 degrees.`;
    const filePath = createTempFile(content);
    const promptFile = await parsePromptFile(filePath);

    expect(promptFile.frontMatter).toEqual({
      title: 'Example Prompt',
      description: 'This is an example prompt file.',
    });
    expect(promptFile.prompts).toEqual([
      { role: 'system', content: 'You are a helpful assistant.\n' },
      { role: 'user', content: 'What is the weather like today?\n' },
      { role: 'assistant', content: 'The weather is sunny with a high of 75 degrees.\n' },
    ]);
  });
});
