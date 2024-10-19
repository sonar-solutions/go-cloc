import os
import subprocess
import sys
import argparse

GITHUB_ORGANIZATION = os.getenv('GO_CLOC_GITHUB_ORGANIZATION')
GITHUB_ACCESS_TOKEN = os.getenv('GO_CLOC_GITHUB_ACCESS_TOKEN')
AZURE_DEVOPS_ORGANIZATION = os.getenv('GO_CLOC_AZURE_DEVOPS_ORGANIZATION')
AZURE_DEVOPS_ACCESS_TOKEN = os.getenv('GO_CLOC_AZURE_DEVOPS_ACCESS_TOKEN')
GITLAB_ORGANIZATION = os.getenv('GO_CLOC_GITLAB_ORGANIZATION')
GITLAB_ACCESS_TOKEN = os.getenv('GO_CLOC_GITLAB_ACCESS_TOKEN')
BITBUCKET_ORGANIZATION = os.getenv('GO_CLOC_BITBUCKET_ORGANIZATION')
BITBUCKET_ACCESS_TOKEN = os.getenv('GO_CLOC_BITBUCKET_ACCESS_TOKEN')

def execute_go_cloc(go_cloc_path, args):
    
    # Collect all command-line arguments passed to the script
    args = [go_cloc_path] + args

    # Construct the command string
    command = " ".join(args)

    # Run the command and capture the output
    try:
        with os.popen(command) as process:
            last_line = ""
            for line in process:
                print(line, end='')  # Print each line to standard output
                last_line = line.strip()  # Keep track of the last line

        # Parse the desired value from the last line
        if last_line.isdigit():
            totalLoc = int(last_line)
            return totalLoc
        else:
            print("Expected output not found in the last line")
            return None
    except OSError as e:
        print(f"Error executing {go_cloc_path}: {e}")
        sys.exit(1)
    
    
def run_test(name,go_cloc_path,args,expected):
    print(f"--------Running test: {name}---------")
    result = execute_go_cloc(go_cloc_path, args)
    did_pass = (result == expected)
    return {
        "name": name,
        "did_pass": did_pass,
        "expected": expected,
        "actual": result
    }

def print_test_results(test_results):
    for test in test_results:
        print(f"Test: {test['name']}")
        print(f"Expected: {test['expected']}")
        print(f"Actual: {test['actual']}")
        print(f"Pass: {test['did_pass']}")
        print("")
    did_all_pass = all(test['did_pass'] for test in test_results)
    return did_all_pass

if __name__ == "__main__":
    # parse the command-line arguments
    parser = argparse.ArgumentParser(description="Script to take in paths to certain binaries.")
    
    # Add arguments for the paths to the binaries
    parser.add_argument('--go_cloc_path', type=str, required=True, help='Path to the go-cloc binary')
    
    # Parse the arguments
    args = parser.parse_args()
    go_cloc_path = args.go_cloc_path
    
    # Print the parsed arguments
    print(f"Path to the go-cloc binary: {go_cloc_path}")

    # Run the tests
    test_results = []
    test_results.append(
        run_test(name="GitHub", expected=143933,go_cloc_path=go_cloc_path,args=["--devops","GitHub","--organization",GITHUB_ORGANIZATION,"--accessToken",GITHUB_ACCESS_TOKEN,"--log-level","INFO","--dump-csvs=false"])
    )
    test_results.append(
        run_test(name="AzureDevOps", expected=57888,go_cloc_path=go_cloc_path,args=["--devops","AzureDevOps","--organization",AZURE_DEVOPS_ORGANIZATION,"--accessToken",AZURE_DEVOPS_ACCESS_TOKEN,"--log-level","INFO","--dump-csvs=false"])
    )
    test_results.append(
        run_test(name="GitLab", expected=162,go_cloc_path=go_cloc_path,args=["--devops","GitLab","--organization",GITLAB_ORGANIZATION,"--accessToken",GITLAB_ACCESS_TOKEN,"--log-level","INFO","--dump-csvs=false"])
    )
    test_results.append(
        run_test(name="Bitbucket", expected=4317,go_cloc_path=go_cloc_path,args=["--devops","Bitbucket","--organization",BITBUCKET_ORGANIZATION,"--accessToken",BITBUCKET_ACCESS_TOKEN,"--log-level","INFO","--dump-csvs=false"])
    )

    did_all_pass = print_test_results(test_results)
    if did_all_pass:
        print("All tests passed!")
    else:
        print("Some tests failed. See above for details")
        sys.exit(-1)