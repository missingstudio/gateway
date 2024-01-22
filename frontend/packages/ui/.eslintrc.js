/** @type {import("eslint").Linter.Config} */
module.exports = {
  root: true,
  extends: ["@missingstudio/eslint-config/react-internal.js"],
  parser: "@typescript-eslint/parser",
  parserOptions: {
    project: "./tsconfig.lint.json",
  },
};
