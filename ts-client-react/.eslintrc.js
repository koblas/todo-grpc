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
    project: [`${__dirname}/tsconfig.json`, `${__dirname}/tsconfig.eslint.json`],
    ecmaFeatures: {
      jsx: true,
    },
    ecmaVersion: 13,
    sourceType: "module",
  },
  plugins: ["react", "@typescript-eslint", "prettier"],
  ignorePatterns: ["/src/genpb/**"],
  rules: {
    // use-hook-form uses this
    "react/jsx-props-no-spreading": "off",
    // I don't
    "import/prefer-default-export": "off",
    "react/no-unescaped-entities": "off",
    // We're using TS
    "react/require-default-props": "off",
    // I like them tigher
    "@typescript-eslint/lines-between-class-members": "off",
    // Would rather be explicit if possible
    "@typescript-eslint/no-inferrable-types": "off",
    "@typescript-eslint/return-await": "off",
    "@typescript-eslint/no-empty-interface": "off",
    // Immer requires reassignment
    "no-param-reassign": ["error", { props: true, ignorePropertyModificationsFor: ["draft"] }],
    // "default-case": "off",
    // "new-cap": "off",
  },
  // overrides: {
  //   files: ["*.test.*"],
  //   rules: {
  //   }
  // }
};
