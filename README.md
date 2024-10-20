# Go-Cloc

## Overview

This tool quickly calculates an accurate Lines of Code (LOC) count for a file or directory with a single executable. 

It is **significantly more performant** than the [cloc](https://github.com/AlDanial/cloc) tool. See [performance benchmark](#performance-benchmarks) for comparisons.


### Usage

Please download the appropriate [artifact](https://github.com/cole-gannaway/go-cloc/releases) for your platform.

Simply run the below command to calculate the Lines of Code (LOC) for the file or directory.
```sh
# Single file
./go-cloc --path test.js
# Directory or Folder
./go-cloc --path src/main
# Output results to a CSV file
./go-cloc --path folder --scan-id folder --dump-csv=true
```
This will output the total Lines of Code (LOC) count for the entire organization. See example below.
```
2024/09/29 17:37:05 [INFO] Setting Log Level to INFO
2024/09/29 17:37:05 [INFO] Parsing CLI arguments
2024/09/29 17:37:05 [INFO] Scanning  src/main ...
2024/09/29 17:37:05 [INFO] Results by file for  src/main :
2024/09/29 17:37:05 [INFO] filePath,blank,comment,code
2024/09/29 17:37:05 [INFO] file1.cpp,100,0,1300
2024/09/29 17:37:05 [INFO] file2.js,100,0,150
...
...
2024/10/20 01:54:22 [INFO] total,200,0,1450
2024/09/29 17:37:05 [INFO] Total LOC for  src/main  is  1450
1450
```

## Requirements
1. Please download the appropriate [artifact](https://github.com/cole-gannaway/go-cloc/releases) for your platform and simply run the single exectuable

## Options
```sh
./go-cloc --help
```
-  `--dump-csv`
       When true, dumps results to a csv file, otherwise gives results in logs
-  `--ignore-file`
       Path to your ignore file. Defines directories and files to exclude when scanning. Please see the README.md for how to format your ignore configuration
-  `--log-level`
       Log level - DEBUG, INFO, WARN, ERROR (default "INFO")
-  `--override-languages-path`
       Path to languages configuration to override the default configuration
-  `--print-languages`
       Prints out the supported languages, file suffixes, and comment configurations. Does not run the tool.
-  `--results-directory-path`
       Path to a new directory for storing the results. Default the tool will create one based on the start time
-  `--scan-id`
       Identifier for the scan. For reference in a csv file later

## Ignore Files

The ignore file is a simple text file used to exclude certain directories and files from processing. You can use a wildcard (`*`) to match patterns, similar to regular expressions. However, you can only use one `*` wildcard at a time. Make sure to place your ignore patterns in the ignore file, one per line, to apply them effectively.

- To ignore all files in a specific directory:

```sh
/path/to/directory/*
```

- To ignore all files ending in `.log` or `.js`:
```sh
*.log
*.js
```

* Combined examples
```sh
# Local scan with ignoring certain files or directories
$ ./go-cloc --path src/main --ignore-file ignore.txt
```

## Extensibility
If successful, the tool will print the total lines of code (LOC) count on its own line. See below for an example. If it fails, it will return a non-zero exit code for easy integration with scripts or other 3rd party tools.
```sh
# Below shows the final LOC outputted on its own line for ease of use
2024/10/20 01:54:22 [INFO] total,200,0,1450
2024/09/29 17:37:05 [INFO] Total LOC for  src/main  is  1450
# Example final line below
1450
```
## Performance Benchmarks

```sh
# Scanning 1 Billion Lines of Code

# go-cloc finished in < 5s
time ./go-cloc --path one-billion-loc-test
3.9s user 0.72s system 93% cpu 4.976 total

# cloc finished in ~2.5 minutes
time cloc one-billion-loc-test
128.48s user 4.22s system 96% cpu 2:17.72 total
```

## Language Support
Below is the default language configuration.

```json
{
  "Abap": {
    "LineComments": ["\""],
    "MultiLineComments": [["/*", "*/"]],
    "Extensions": [".abap", ".ab4", ".flow"]
  },
  "ActionScript": {
    "LineComments": ["//"],
    "MultiLineComments": [["/*", "*/"]],
    "Extensions": [".as"]
  },
  "Apex": {
    "LineComments": ["//"],
    "MultiLineComments": [["/*", "*/"]],
    "Extensions": [".cls", ".trigger"]
  },
  "C": {
    "LineComments": ["//"],
    "MultiLineComments": [["/*", "*/"]],
    "Extensions": [".c"]
  },
  "C Header": {
    "LineComments": ["//"],
    "MultiLineComments": [["/*", "*/"]],
    "Extensions": [".h"]
  },
  "C#": {
    "LineComments": ["//"],
    "MultiLineComments": [["/*", "*/"]],
    "Extensions": [".cs"]
  },
  "C++": {
    "LineComments": ["//"],
    "MultiLineComments": [["/*", "*/"]],
    "Extensions": [".cpp", ".cc", ".cxx", ".c++"]
  },
  "C++ Header": {
    "LineComments": ["//"],
    "MultiLineComments": [["/*", "*/"]],
    "Extensions": [".hh", ".hpp", ".hxx", ".h++", ".ipp"]
  },
  "COBOL": {
    "LineComments": ["*", "/"],
    "MultiLineComments": [],
    "Extensions": [".cbl", ".ccp", ".cob", ".cobol", ".cpy"]
  },
  "CSS": {
    "LineComments": ["//"],
    "MultiLineComments": [["/*", "*/"]],
    "Extensions": [".css"]
  },
  "Flex": {
    "LineComments": ["//"],
    "MultiLineComments": [["/*", "*/"]],
    "Extensions": [".as"]
  },
  "Golang": {
    "LineComments": ["//"],
    "MultiLineComments": [["/*", "*/"]],
    "Extensions": [".go"]
  },
  "HTML": {
    "LineComments": [],
    "MultiLineComments": [["<!--", "-->"]],
    "Extensions": [
      ".html",
      ".htm",
      ".cshtml",
      ".vbhtml",
      ".aspx",
      ".ascx",
      ".rhtml",
      ".erb",
      ".shtml",
      ".shtm",
      ".cmp"
    ]
  },
  "JCL": {
    "LineComments": ["//"],
    "MultiLineComments": [["/*", "*/"]],
    "Extensions": [".jcl", ".JCL"]
  },
  "Java": {
    "LineComments": ["//"],
    "MultiLineComments": [["/*", "*/"]],
    "Extensions": [".java", ".jav"]
  },
  "JavaScript": {
    "LineComments": ["//"],
    "MultiLineComments": [["/*", "*/"]],
    "Extensions": [".js", ".jsx", ".jsp", ".jspx", ".jspf", ".mjs"]
  },
  "Kotlin": {
    "LineComments": ["//"],
    "MultiLineComments": [["/*", "*/"]],
    "Extensions": [".kt", ".kts"]
  },
  "Objective-C": {
    "LineComments": ["//"],
    "MultiLineComments": [["/*", "*/"]],
    "Extensions": [".m"]
  },
  "Oracle PL/SQL": {
    "LineComments": ["--"],
    "MultiLineComments": [["/*", "*/"]],
    "Extensions": [".pkb"]
  },
  "PHP": {
    "LineComments": ["//", "#"],
    "MultiLineComments": [["/*", "*/"]],
    "Extensions": [".php", ".php3", ".php4", ".php5", ".phtml", ".inc"]
  },
  "PL/I": {
    "LineComments": ["--"],
    "MultiLineComments": [["/*", "*/"]],
    "Extensions": [".pl1"]
  },
  "Python": {
    "LineComments": ["#"],
    "MultiLineComments": [["\"\"\"", "\"\"\""]],
    "Extensions": [".py", ".python", ".ipynb"]
  },
  "RPG": {
    "LineComments": ["#"],
    "MultiLineComments": [],
    "Extensions": [".rpg"]
  },
  "Ruby": {
    "LineComments": ["#"],
    "MultiLineComments": [["=begin", "=end"]],
    "Extensions": [".rb"]
  },
  "SQL": {
    "LineComments": ["--"],
    "MultiLineComments": [["/*", "*/"]],
    "Extensions": [".sql"]
  },
  "Scala": {
    "LineComments": ["//"],
    "MultiLineComments": [["/*", "*/"]],
    "Extensions": [".scala"]
  },
  "Scss": {
    "LineComments": ["//"],
    "MultiLineComments": [["/*", "*/"]],
    "Extensions": [".scss"]
  },
  "Swift": {
    "LineComments": ["//"],
    "MultiLineComments": [["/*", "*/"]],
    "Extensions": [".swift"]
  },
  "T-SQL": {
    "LineComments": ["--"],
    "MultiLineComments": [],
    "Extensions": [".tsql"]
  },
  "Terraform": {
    "LineComments": [],
    "MultiLineComments": [],
    "Extensions": [".tf"]
  },
  "TypeScript": {
    "LineComments": ["//"],
    "MultiLineComments": [["/*", "*/"]],
    "Extensions": [".ts", ".tsx"]
  },
  "Visual Basic .NET": {
    "LineComments": ["'"],
    "MultiLineComments": [],
    "Extensions": [".vb"]
  },
  "Vue": {
    "LineComments": ["<!--"],
    "MultiLineComments": [["<!--", "-->"]],
    "Extensions": [".vue"]
  },
  "XHTML": {
    "LineComments": ["<!--"],
    "MultiLineComments": [["<!--", "-->"]],
    "Extensions": [".xhtml"]
  },
  "XML": {
    "LineComments": ["<!--"],
    "MultiLineComments": [["<!--", "-->"]],
    "Extensions": [".xml", ".XML", ".xsd", ".xsl"]
  },
  "YAML": {
    "LineComments": ["#"],
    "MultiLineComments": [],
    "Extensions": [".yaml", ".yml"]
  }
}

```
### Customization
To customize this configuration, copy the above JSON, customize it to your needs, and pass in the file path as `--override-languages-path`. See [options](#options) for more details.