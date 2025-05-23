# !bin/bash

# GitHub latest-download URL (no trailing slash)
REPO_URL="https://github.com/sonar-solutions/go-cloc/releases/latest/download"

# Detect OS and arch
OS="$(uname -s)"
ARCH="$(uname -m)"
echo "Detected: OS=$OS ARCH=$ARCH"

# Determine platform from OS
if [[ "$OS" == "Darwin" ]]; then
    PLATFORM="darwin"
elif [[ "$OS" == "Linux" ]]; then
    PLATFORM="linux"
elif [[ "$OS" == "CYGWIN"* || "$OS" == "MINGW"* || "$OS" == "MSYS"* || "$OS" == "Windows_NT" ]]; then
    PLATFORM="windows"
else
    echo "Unsupported OS: $OS" >&2
    exit 1
fi

# Determine architecture
if [[ "$ARCH" == "x86_64" || "$ARCH" == "amd64" ]]; then
    GOARCH="amd64"
elif [[ "$ARCH" == "aarch64" || "$ARCH" == "arm64" ]]; then
    GOARCH="arm64"
elif [[ "$ARCH" == "386" || "$ARCH" == "i386" ]]; then
    GOARCH="386"
else
    echo "Unsupported architecture: $ARCH" >&2
    exit 1
fi

ZIP_FILE="go-cloc-${PLATFORM}-${GOARCH}.zip"
DOWNLOAD_URL="${REPO_URL}/${ZIP_FILE}"

echo "Downloading ${ZIP_FILE} from ${DOWNLOAD_URL}…"
curl -fsSL -L -o "$ZIP_FILE" "$DOWNLOAD_URL"

echo "Unpacking ${ZIP_FILE}…"
unzip -q "$ZIP_FILE"

# Determine binary name and verify
if [[ "$PLATFORM" == "windows" ]]; then
    BIN="go-cloc.exe"
else
    BIN="go-cloc"
fi

if [[ ! -x "$BIN" ]]; then
    echo "Error: $BIN not found or not executable after unzip" >&2
    exit 1
fi

echo "Installed $BIN successfully."

# Cleanup
rm -f "$ZIP_FILE"