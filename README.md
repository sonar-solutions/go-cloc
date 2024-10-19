# Go-Cloc

## Overview

This tool simplifies the process of obtaining an accurate Lines of Code (LOC) count for an organization's DevOps platform. It can automatically discover repositories and quickly calculate the total LOC with a single executable. 

It is also **significantly more performant** than the [cloc](https://github.com/AlDanial/cloc) tool. See [performance benchmark](#performance-benchmarks) for comparisons.


### Usage

Please download the appropriate [artifact](https://github.com/cole-gannaway/go-cloc/releases) for your platform.

Simply run the below command to discover all repositories in your **DevOps Organization**.
```sh
./go-cloc --devops <DevOpsPlatform>  --organization <YourOrganizationName>  --accessToken <YourPersonalAccessToken>
```
This will output the total Lines of Code (LOC) count for the entire organization. See example below.
```
2024/09/29 17:37:04 [INFO] Discovering repositories in  MyExampleOrganization
2024/09/29 17:37:04 [INFO] Discovered  50  repositories in  MyExampleOrganization
2024/09/29 17:37:04 [INFO] 1 / 50  cloning respository  example-repo ...
2024/09/29 17:37:05 [INFO] Scanning  example-repo ...
2024/09/29 17:37:05 [INFO] Done! Results for  example-repo  can be found  MyExampleOrganization-example-repo.csv
...
...
...
2024/09/29 17:37:05 [INFO] 0 repos failed to scan.
2024/09/29 17:37:05 [INFO] Total LOC results can be found  AAA-combined-total-lines.csv
2024/09/29 17:37:05 [INFO] Total LOC for  MyExampleOrganization  is  23005
23005
```

## Examples
Using **Github** as the DevOps platform
```sh
./go-cloc --devops GitHub --organization MyExampleOrganization --accessToken abcdefg1234 
```
Using a self-hosted **AzureDevOps** as the DevOps platform without HTTPS
```sh
./go-cloc --devops AzureDevOps --organization MyExampleOrganization --accessToken abcdefg1234 --use-https=false --devops-base-url-override devops.internal.company.com
```
**Local** scan of a single file
```sh
./go-cloc main.js 
```

## Requirements
1. An **Access Token** for your appropriate DevOps platform (GitHub, Azure DevOps, GitLab, or Bitbucket) with **read** access for each of the repositories within the organization. See [below](#personal-access-tokens) for more details.

## Options
```sh
./go-cloc --help
```
-  `--accessToken`
       Your DevOps personal access token used for discovering and downloading repositories in your organization
-  `--clone-repo-using-zip`
       When true, repositories are downloaded as zip files instead of git clone to drastically improve performance. Default is false. This is a BETA feature and has not been extensively tested
-  `--devops`
       GitHub, AzureDevOps, Bitbucket, GitLab, or Local (default "Local")
-  `--devops-base-url-override`
       Overrides the base URL for the DevOps provider. Defaults will be "github.com", "dev.azure.com", "bitbucket.org", or "gitlab.com". However, you can override this with your own self-hosted ip or domain
-  `--dump-csvs`
       When false, disables csv file dumps. DEBUG logging available to still see csv results in logs. (default true)
-  `--exclude-repositories-file`
       Path to your exclude repositories file. Defines repositories to exclude, all others will be included. Please see the README.md for how to format your exclude repositories configuration
-  `--ignore-file`
       Path to your ignore file. Defines directories and files to exclude when scanning. Please see the README.md for how to format your ignore configuration
-  `--include-repositories-file`
       Path to your include repositories file. Defines repositories to include, all others will be excluded. Please see the README.md for how to format your include repositories configuration
-  `--local-file-path`
       Path to youthe local file or directory that you wish to scan
-  `--log-level`
       Log level - DEBUG, INFO, WARN, ERROR (default "INFO")
-  `--organization`
       Your DevOps organization name
-  `--print-languages`
       Prints out the supported languages, file suffixes, and comment configurations. Does not run the tool.
-  `--results-directory-path`
       Path to a new directory for storing the results. Default the tool will create one based on the start time
-  `--use-https`
       When false, uses http instead of https for all HTTP calls. (default true)

## Ignore Files

The ignore file is a simple text file used to exclude certain directories and files from processing. You can use a wildcard (`*`) to match patterns, similar to regular expressions. However, you can only use one `*` wildcard at a time. Make sure to place your ignore patterns in the ignore file, one per line, to apply them effectively.

This same configuration format applies to ***exclude*** or ***include*** repositories when using the `--devops` flag. Note: if using the `--devops` flag, these patterns will apply to all repositories.

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
# Local scan with ignoring certain files or directoreis
$ ./go-cloc src/ --ignore-file ignore.txt

# DevOps scan ignoring certain repositores 
$ ./go-cloc --devops GitHub \
      --organization MyExampleOrganization \
      --accessToken abcdefg1234 \
      --exclude-repositories-file github_repos_to_ignore.txt

# DevOps scan only including certain repositories
$ ./go-cloc --devops GitHub \
      --organization MyExampleOrganization \
      --accessToken abcdefg1234 \
      --include-repositories-file github_repos_to_include.txt
```

## Personal Access Tokens

Personal Access Tokens (PATs) are used to authenticate and authorize access to your DevOps platform. They are necessary for the tool to discover and clone repositories within your organization. Below are the steps to generate a PAT for different DevOps platforms:

### GitHub
1. Navigate to [GitHub Settings](https://github.com/settings/tokens).
2. Click on **Generate new token**.
3. Under **Select Scopes**, select **repo**.
5. Click **Generate token** and copy the token for use.

### Azure DevOps
1. Navigate to [Azure DevOps](https://dev.azure.com).
2. Click on **User Settings** and select **Personal Access Token**.
3. Click on **New Token**.
4. Set the name, organization, and scopes for the token.
5. Ensure that **Code -> Read** is selected as a scope.
6. Click **Create** and copy the token for use.

### GitLab
1. Navigate to [GitLab](https://gitlab.com).
2. Select your **Organization**.
3. Click on **Settings** in the left sidebar and select **Access Tokens**.
4. Provide a name and expiration date for the token.
5. Select the scopes `read_api` and `read_repository` to grant the necessary permissions.
6. Click **Create personal access token** and copy the token for use.

### Bitbucket
1. Navigate to [Bitbucket](https://bitbucket.org).
2. Select your **Organization**.
3. In the left sidebar and select **Access Tokens**.
4. Provide a name and expiration date for the token.
5. Select the scopes **Repository** to **Read**.to grant the necessary permissions.
6. Click **Create** and copy the token for use.

## Language Support
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

## Extensibility
If successful, the tool will print the total lines of code (LOC) count on its own line. See below for an example. If it fails, it will return a non-zero exit code for easy integration with scripts or other 3rd party tools.
```sh
# Below shows the final LOC outputted on its own line for ease of use
2024/09/29 17:37:05 [INFO] Total LOC results can be found  AAA-combined-total-lines.csv
2024/09/29 17:37:05 [INFO] Total LOC for  MyExampleOrganization  is  23005
# Example final line below
23005
```

## Performance Benchmarks

```sh
# Scanning 1 Billion Lines of Code

# go-cloc finished in < 5s
time ./go-cloc --local-file-path one-billion-loc-test 
3.9s user 0.72s system 93% cpu 4.976 total

# cloc finished in ~2.5 minutes
time cloc one-billion-loc-test
128.48s user 4.22s system 96% cpu 2:17.72 total

```