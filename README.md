# Rename Files
Compiles a simple binary so your technologically inept friends can easily organize their files.

#### For Windows
```
make windows
```
#### For most other Unix-based systems
```
make
```

If you don't have `make`, install it, it's good.

Give them the binary, tell them to put it in the directory of their desire, rename the binary (keeping the `.exe` extension for Windows OS), and execute it. Renames everything, including directories, to the name of the binary. Does so in smart numerical order. Does not crawl into child directories.

Has an alternative usage but you can figure it out from the code.

That's it. Use with caution.