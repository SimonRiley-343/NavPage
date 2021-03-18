module.exports = {
    root: true,
    env: {
        browser: true,
        node: true
    },
    extends: ['plugin:vue/essential', '@vue/standard', 'eslint:recommended', '@vue/typescript/recommended'],
    parserOptions: {
        ecmaVersion: 2020
    },
    plugins: ['@typescript-eslint'],
    rules: {
        '@typescript-eslint/ban-types': 'off',
        '@typescript-eslint/explicit-module-boundary-types': 'off',
        '@typescript-eslint/no-explicit-any': 'off',
        '@typescript-eslint/no-var-requires': 'off',
        '@typescript-eslint/no-this-alias': 'off',
        indent: ['warn', 4, { SwitchCase: 1 }],
        semi: ['error', 'always'],
        curly: ['error', 'all'],
        'default-case': 'error',
        'eol-last': 2,
        'space-before-function-paren': [
            'error',
            {
                anonymous: 'always',
                named: 'never',
                asyncArrow: 'always'
            }
        ],
        'no-multiple-empty-lines': [
            'error',
            {
                max: 1,
                maxBOF: 1
            }
        ],
        'vue/array-bracket-spacing': ['error', 'never'],
        'vue/arrow-spacing': 'error',
        'vue/block-spacing': 'error',
        'vue/brace-style': ['error', '1tbs', { allowSingleLine: true }],
        'vue/camelcase': 'error',
        'vue/comma-dangle': ['error', 'never'],
        'vue/component-name-in-template-casing': ['error', 'kebab-case'],
        'vue/eqeqeq': 'error',
        'vue/key-spacing': 'error',
        'vue/match-component-file-name': 'error',
        'vue/object-curly-spacing': 'error',
        'comma-spacing': ['error', { before: false, after: true }],
        'comma-style': ['error', 'last'],
        'prefer-const': 'off'
    }
};
