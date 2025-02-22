import { defineConfig } from "tsup";

export default defineConfig({
  entry: {
    index: "./src/index.ts",
    parse: "./src/parse.ts",
    template: "./src/template.ts",
  },
  format: ["cjs", "esm"],
  dts: true,
  splitting: false,
  clean: true,
  minify: process.env.NODE_ENV === "production",
  sourcemap: true,
  outDir: "dist",
  target: "es2020",
  external: ["jest"],
});
