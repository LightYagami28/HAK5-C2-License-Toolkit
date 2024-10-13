# Hak5 Cloud C2 Licensing Toolkit 🛠️

## Overview

The **Hak5 Cloud C2 Licensing Toolkit** is a powerful utility designed for interacting with the licensing structure of Hak5 Cloud C2. This toolkit allows users to generate, decode, and manipulate license data easily.

## Features ✨

- **Generate**: Create a test License struct hex string.
- **Decode**: Decode a License / Status struct hex string.
- **Read**: Read values inside `Setup[License]` or `Status[Status]` buckets struct hex string from the database.
- **Crack**: Insert license values into the database.

## Prerequisites 📋

Ensure you have the following installed before using the toolkit:

- [Go programming language](https://golang.org/dl/)
- Git

## Installation 🚀

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/Pwn3rzs/HAK5-C2-License-Toolkit
   cd HAK5-C2-License-Toolkit/
   ```

2. **Initialize Go Module**:
   ```bash
   go mod init pwn3rzs.cloud/hak5-c2-toolkit
   ```

3. **Install Dependencies**:
   ```bash
   go get go.etcd.io/bbolt
   ```

4. **Build the Toolkit**:
   Specify the OS and architecture for your target platform. Here are some examples:

   ```bash
   # macOS ARM64
   GOARCH=arm64 go build -work -ldflags="-s -w" -o HAK5-C2-Toolkit-mac-arm64
   
   # macOS AMD64
   GOARCH=amd64 go build -work -ldflags="-s -w" -o HAK5-C2-Toolkit-mac-amd64
   
   # Windows AMD64
   GOOS=windows GOARCH=amd64 go build -work -ldflags="-s -w" -o HAK5-C2-Toolkit-win-amd64
   
   # Windows i386
   GOOS=windows GOARCH=386 go build -work -ldflags="-s -w" -o HAK5-C2-Toolkit-win-i386
   
   # Linux AMD64
   GOOS=linux GOARCH=amd64 go build -work -ldflags="-s -w" -o HAK5-C2-Toolkit-lin-amd64
   
   # Linux i386
   GOOS=linux GOARCH=386 go build -work -ldflags="-s -w" -o HAK5-C2-Toolkit-lin-386
   ```

5. **Setup the Environment**:
   Move the built binary `./HAK5-C2-License-Toolkit` into the same directory as the `c2-<x.x.x>_<arch>_<os>(.exe)` binary.

## Usage 📝

Run the toolkit and choose one of the following options:

- **Generate**: To generate a test License struct hex string.
- **Decode**: To decode a License / Status struct hex string.
- **Read**: To read the values inside `Setup[License]` or `Status[Status]` buckets struct hex string from the database.
- **Crack**: To start the process of inserting license values into the database.
  > **Important**: Make sure you have replaced the binary, or the license will reset.

## Resources 📚

For more information, check out the following resources:

- **Patched Binaries**:
  - [Telegram](https://t.me/Pwn3rzs/1119)
  - [GitHub Releases](https://github.com/Pwn3rzs/HAK5-C2-License-Toolkit/releases/tag/v3.3.0)

- **Guides and Tutorials**:
  - [CyberArsenal Post](https://cyberarsenal.org/threads/hak5-cloud-c2-analysis-cracking-method.1408/)

## Technical Information ⚙️

- **Database Used**: [BoltDB](https://github.com/etcd-io/bbolt)
- **Encoding Used**: [GOB](https://pkg.go.dev/encoding/gob)
- **Supported Version**: `v3.3.0`
- **Current Version**: [Hak5 Cloud C2 Updates](https://c2.hak5.org/api/v2/feed)

## Contributing 🤝

Contributions are welcome! Please open an issue or submit a pull request if you have improvements or suggestions.

---

**Disclaimer**: This toolkit is intended for educational purposes only. Use it responsibly and in accordance with applicable laws and regulations.
