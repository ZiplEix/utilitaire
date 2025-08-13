# GITIGN

gitign is a command-line tool designed to help you easily generate .gitignore files based on the programming languages or frameworks in your project. It can automatically detect the languages you're using or allow you to specify which ones to include, append rules to an existing .gitignore, and even optimize the file by removing duplicate rules.

## Features

- Automatically detect programming languages in your project and generate a `.gitignore` file.
- Specify multiple language extensions to generate custom `.gitignore` files.
- Append rules to an existing `.gitignore` file.
- Optimize the `.gitignore` file by removing duplicate rules.
- Ignore specific folders or file types during the generation process.

## Usage

The basic usage of gitign is to generate a `.gitignore` file. If no language extensions are provided, it will automatically detect the languages in your project by walking through the directory tree.

```bash
gitign
```

you can also specify the language extensions you want to include in the `.gitignore` file:

```bash
gitign .go .py .ts
```

if extensions are provided, gitign will only include the specified languages in the `.gitignore` file and not walk through the directory tree.

### Available options

- **`-h`**: Display the help message.

    ```bash
    gitign -h
    ```

- **`--ignore`**: Specify a comma-separated list of folders or file extensions to ignore during language detection.

    ```bash
    gitign --ignore=node_modules,build,.java
    ```

- **`--append:`** Append the generated rules to an existing .gitignore file.

    ```bash
    gitign go --append
    ```

- **`--optimize`**: Optimize the .gitignore file by removing duplicate rules. This can be run on its own:

    ```bash
    gitign --optimize
    ```

    or combined with other options:

    ```bash
    gitign go --append --optimize
    ```

## Adding templates

If you want to add or modify the language templates, you need to update the [generator/langage.go](./generator/langage.go) file. This file contains the mapping between file extensions and their corresponding .gitignore templates. To add a new template:

1. Open the generator/langage.go file.
2. Add your new language or framework in the langagesExtensions map.
3. Specify the URL of the corresponding .gitignore template (or leave it blank if not needed).

Exemple:

```go
".newlang": {"NewLang", "NewLang", true, "https://raw.githubusercontent.com/github/gitignore/refs/heads/main/NewLang.gitignore"},
```

## Contributing

If you have any suggestions, bug reports, or feature requests, feel free to open an issue or submit a pull request. Your feedback is highly appreciated!
