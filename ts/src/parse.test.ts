import { parse } from "./parse";
import * as fs from "fs";
import { basename, extname, join } from "path";

describe("parse", () => {
  const testDataDir = join(__dirname, "../../testdata/parse");
  const testFiles = fs
    .readdirSync(testDataDir)
    .filter((file) => file.endsWith(".prompt.md"))
    .map((file) => ({
      name: basename(file, extname(file)),
      file,
    }));

  it.each(testFiles)("$name", ({ name, file }) => {
    const promptFile = join(testDataDir, file);
    const expectedFile = join(
      testDataDir,
      file.replace(".prompt.md", ".expected.json")
    );

    const promptContent = fs.readFileSync(promptFile, "utf-8");
    const expectedContent = fs.readFileSync(expectedFile, "utf-8");
    const expected = JSON.parse(expectedContent);

    const result = parse(promptContent);
    expect(result).toEqual(expected);
  });
});
