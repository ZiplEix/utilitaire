package gitign

type Langage struct {
	Extension     string
	Name          string
	NeedGitignore bool
	Url           string
}

var langagesExtensions = map[string]Langage{
	".go":    {"Go", "Go", true, "https://raw.githubusercontent.com/github/gitignore/refs/heads/main/Go.gitignore"},
	".js":    {"JavaScript / TypeScript", "JavaScript", true, "https://raw.githubusercontent.com/github/gitignore/refs/heads/main/Node.gitignore"},
	".ts":    {"JavaScript / TypeScript", "TypeScript", true, "https://raw.githubusercontent.com/github/gitignore/refs/heads/main/Node.gitignore"},
	".py":    {"Python", "Python", true, "https://raw.githubusercontent.com/github/gitignore/refs/heads/main/Python.gitignore"},
	".java":  {"Java", "Java", true, "https://raw.githubusercontent.com/github/gitignore/refs/heads/main/Java.gitignore"},
	".rb":    {"Ruby", "Ruby", true, "https://raw.githubusercontent.com/github/gitignore/refs/heads/main/Ruby.gitignore"},
	".php":   {"PHP", "PHP", true, "https://raw.githubusercontent.com/ZiplEix/gitign/refs/heads/master/templates/PHP.gitignore"},
	".cs":    {"C#", "C#", true, "https://raw.githubusercontent.com/github/gitignore/refs/heads/main/VisualStudio.gitignore"},
	".cpp":   {"C++", "C++", true, "https://raw.githubusercontent.com/github/gitignore/refs/heads/main/C%2B%2B.gitignore"},
	".c":     {"C", "C", true, "https://raw.githubusercontent.com/github/gitignore/refs/heads/main/C.gitignore"},
	".rs":    {"Rust", "Rust", true, "https://raw.githubusercontent.com/github/gitignore/refs/heads/main/Rust.gitignore"},
	".kt":    {"Kotlin", "Kotlin", true, "https://raw.githubusercontent.com/github/gitignore/refs/heads/main/Java.gitignore"},
	".swift": {"Swift", "Swift", true, "https://raw.githubusercontent.com/github/gitignore/refs/heads/main/Swift.gitignore"},
	".dart":  {"Dart", "Dart", true, "https://raw.githubusercontent.com/github/gitignore/refs/heads/main/Dart.gitignore"},
	".scala": {"Scala", "Scala", true, "https://raw.githubusercontent.com/github/gitignore/refs/heads/main/Scala.gitignore"},
	// Web languages without ignore needed
	".html": {"HTML", "HTML", false, ""},
	".css":  {"CSS", "CSS", false, ""},
	// Configurations files
	".xml":  {"XML", "XML", false, ""},
	".sh":   {"Shell", "Shell", false, ""},
	".json": {"JSON", "JSON", false, ""},
	".yml":  {"YAML", "YAML", false, ""},
	".toml": {"TOML", "TOML", false, ""},
	".md":   {"Markdown", "Markdown", false, ""},
	".txt":  {"Text", "Text", false, ""},
	".env":  {"Environment", "Environment", false, ""},
}
