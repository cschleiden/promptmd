import { prepare } from "./template";
const fs = require("fs");
const path = require("path");

describe("prepare", () => {
  const template: {
    name: string;
    input: string;
    vars: { [key: string]: string };
    output: string;
  }[] = JSON.parse(
    fs.readFileSync(
      path.resolve(__dirname, "../../testdata/template/template.json"),
      "utf8"
    )
  );

  it.each(template)("$name", ({ name, input, vars, output }) => {
    const r = prepare(input)(vars);

    expect(r).toEqual(output);
  });
});
