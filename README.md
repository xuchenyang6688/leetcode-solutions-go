🚀My Go learning journey begins here!

1. Install Go
2. Install Go Extention in VS code
3. Command Palette (Ctrl+Shift+P or Cmd+Shift+P) and typing "Go: Install/Update Tools", install gopls, glv and all.
4. Command Palette (Ctrl+Shift+P or Cmd+Shift+P) and typeing "Preferences: Open User Settings(JSON)"
```
{
    "go.useLanguageServer": true,
    "go.formatTool": "goimports",
    "[go]": {
        "editor.insertSpaces": false,
        "editor.formatOnSave": true,
        "editor.codeActionsOnSave": {
            "source.organizeImports": "explicit"
        }
    }
}

```