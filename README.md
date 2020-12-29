
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
