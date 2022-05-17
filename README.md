

## Custom externals

The server looks for a `_externals\` directory located in the workspace root.  
When there is a `_externals\externals.src` it will try to parse it and all referenced files right after parsing the built-in externals and before parsing user scripts.  
- If there is no `_externals\externals.src` we look for a `_externals\externals.d` and try to parse that.

This externals should be provided from Union plugins such as zParserExtender.


## UPX

The languageserver binary can be compressed using UPX 3.96 with the following arguments:

```
-9                | best compression
--strip-relocs=0  | preserve relocation information
--lzma            | use lzma compression, this saves about 15% of filesize
```

if you want to uncompress the binary you can use the following command:

```
upx -d --strip-relocs=0 DaedalusLanguageServer.exe
```

it is important to specify `--strip-relocs=0`. If you omit this argument, the executable will be damaged.
