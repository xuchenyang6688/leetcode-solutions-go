🚀My Go learning journey begins here!

Install Go
Install Go Extention in VS code
Command Palette (Ctrl+Shift+P or Cmd+Shift+P) and typing "Go: Install/Update Tools", install gopls, glv and all.
Command Palette (Ctrl+Shift+P or Cmd+Shift+P) and typeing "Preferences: Open User Settings(JSON)"
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