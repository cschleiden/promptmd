const VariableRegex = /\{\{([a-zA-Z0-9_]+)\}\}/g;

export type args = { [variableName: string]: unknown };

export function prepare(template: string): (args: args) => string {
  const segments: (string | ((a: args) => string))[] = [];
  let lastIndex = 0;
  let match;

  while ((match = VariableRegex.exec(template)) !== null) {
    if (match.index > lastIndex) {
      segments.push(template.substring(lastIndex, match.index));
    }

    const variableName = match[1];
    segments.push((args) => formatVar(args[variableName]));
    lastIndex = match.index + match[0].length;
  }

  if (lastIndex < template.length) {
    segments.push(template.substring(lastIndex));
  }

  return (args: args) => {
    let r = "";

    for (const segment of segments) {
      if (typeof segment === "string") {
        r += segment;
      } else {
        r += segment(args);
      }
    }
    return r;
  };
}

function formatVar(a: unknown): string {
  if (a === null || a === undefined) {
    return "";
  }

  return String(a);
}
