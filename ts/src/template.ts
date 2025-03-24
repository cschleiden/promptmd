const VariableRegex = /\{\{([a-zA-Z0-9_]+)\}\}/g;

export type Vars = { [variableName: string]: unknown };

export function prepare(template: string): (args: Vars) => string {
  const segments: (string | ((a: Vars) => string))[] = [];
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

  if (segments.length <= 1) {
    return () => template; // No variables found, return the original template
  }

  return (vars: Vars) => {
    let r = "";

    for (const segment of segments) {
      if (typeof segment === "string") {
        r += segment;
      } else {
        r += segment(vars);
      }
    }
    return r;
  };
}

function formatVar(v: unknown): string {
  if (
    typeof v === "string" ||
    typeof v === "number" ||
    typeof v === "boolean"
  ) {
    return String(v);
  }

  return "";
}
