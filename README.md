# URLOADER

Just pass `-url` and `-save_path` 
to the program and it will download the file to the specified path.

## Usage

Regular download:

```shell
./urloader -url=https://example.com/file.txt -save_path=/path/to/file.txt
```

Mass download and combine:

```shell
./urloader -combine -save_path /path/to/combine.txt https://example.com/file1.txt https://example.com/file2.txt
```