module.exports = {
  env: {
    browser: true,
    es2021: true,
  },
  extends: [
    "eslint:recommended",
    "plugin:react/recommended",
    "plugin:@typescript-eslint/recommended",
    "airbnb",
    "airbnb/hooks",
    "airbnb-typescript",
    "prettier",
  ],
  parser: "@typescript-eslint/parser",
  parserOptions: {
    project: "./tsconfig.json",
    ecmaFeatures: {
      jsx: true,
    },
    ecmaVersion: 13,
    sourceType: "module",
  },
  plugins: ["react", "@typescript-eslint", "prettier"],
  rules: {
    "import/prefer-default-export": "off",
    "react/no-unescaped-entities": "off",
    "@typescript-eslint/no-inferrable-types": "off",
    "@typescript-eslint/return-await": "off",
    "@typescript-eslint/no-empty-interface": "off",
    "no-param-reassign": ["error", { props: true, ignorePropertyModificationsFor: ["draft"] }],
    "default-case": "off",
    "new-cap": "off",
  },
};
