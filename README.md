[![Build](https://github.com/kirides/DaedalusLanguageServer/actions/workflows/build.yml/badge.svg)](https://github.com/kirides/DaedalusLanguageServer/actions/workflows/build.yml)

## Usage

see: https://github.com/kirides/vscode-daedalus

## Custom externals

The server looks for a `_externals\` directory located in the workspace root.  
When there is a `_externals\externals.src` it will try to parse it and all referenced files right after parsing the built-in externals and before parsing user scripts.  
- If there is no `_externals\externals.src` we look for a `_externals\externals.d` and try to parse that.

This externals should be provided from Union plugins such as zParserExtender.
