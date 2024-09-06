
# File Encryption Tool [GlimCrypt]

![Go Version](https://img.shields.io/github/go-mod/go-version/Glimor/glimcrypt) 

## 🛡️ Introduction

This project is an **AES-256 CBC File Encryption/Decryption tool** written in Go. The purpose of this project is to demonstrate a simple yet powerful file encryption tool which encrypts files using a password based key generation mechanism. This project is **educational** and aims to provide an understanding of cryptography in Go. It is not intended for production use.

Users can encrypt files by providing **8 distinct keywords** which are used to generate a secure key for AES-256 encryption. Similarly, files can be decrypted using the same keywords.

## ✨ Features

- **AES-256 CBC Encryption**: Ensures that your files are encrypted securely using a widely trusted cryptography standard.
- **8-Keyword Based Key Generation**: Users generate keys dynamically based on inputted keywords, ensuring unique and user-specific keys.
- **File Explorer Support**: Allows users to select files using a GUI-based file explorer.
- **Platform Agnostic**: Works on **Linux**, **macOS**, and **Windows**.
- **Secure Padding (PKCS7)**: Handles padding securely for file encryption/decryption.

## 🧑‍🏫 Educational Use

This project is designed for **educational purposes**, demonstrating:
- Key generation from multiple keywords.
- AES-256 CBC encryption and decryption in Go.
- Handling file I/O operations securely.
- Error handling and cryptography best practices.

### ⚠️ Disclaimer

This project is **not suitable for production use**. It should be used solely as an educational reference for understanding file encryption in Go. 

---

## 🚀 Getting Started

### Prerequisites

- **Go** (1.18 or higher)
- **Git**

### Installation

1. **Clone the repository:**

```bash
git clone https://github.com/Glimor/glimcrypt.git
cd glimcrypt
```

2. **Install dependencies:**

The project uses some third-party packages like Cobra (for CLI) and Zenity (for file explorer).

```bash
go mod tidy
```

### Usage

This project provides a simple CLI interface to encrypt or decrypt files.

#### 1. **Encrypt a file**

Run the following command to encrypt a file:

```bash
go run main.go encrypt
```

You will be prompted to input **8 distinct keywords** which will be used to generate the encryption key. Then, a file explorer window will open for you to select the file you want to encrypt.

#### 2. **Decrypt a file**

To decrypt a file use the following command:

```bash
go run main.go decrypt
```

Again you will be prompted to input the same **8 keywords** used for encryption and a file explorer window will open to select the encrypted `.enc` file. If the keywords match the file will be decrypted and saved with a `.dec` extension.

---

## 📂 Project Structure

```
myapp/
├── cmd/
│   ├── root.go          # Main command handling
│   ├── encrypt.go       # Encrypt command logic
│   └── decrypt.go       # Decrypt command logic
├── internal/
│   ├── cryptography.go  # AES-256 CBC encryption/decryption logic
│   ├── keygen.go        # Key generation from 8 keywords
│   ├── fileexplorer.go  # File explorer integration using Zenity
├── go.mod               # Go module definition
├── go.sum               # Dependency checksums
├── main.go              # Application entry point
```

### Key Files:

- **`main.go`**: Application entry point, handles command execution.
- **`cmd/`**: Contains CLI commands (`encrypt` and `decrypt`).
- **`internal/cryptography.go`**: Contains AES-256 CBC encryption and decryption logic, as well as padding/unpadding mechanisms.
- **`internal/keygen.go`**: Handles dynamic key generation based on user keywords.
- **`internal/fileexplorer.go`**: Uses Zenity for file selection via a file explorer.

---

## 🔧 Dependencies

This project uses a few key dependencies:

- **[Cobra](https://github.com/spf13/cobra)**: CLI framework used to create command-line interfaces.
- **[Zenity](https://github.com/ncruces/zenity)**: Allows the file explorer functionality across different platforms.
- **[crypto/pbkdf2](https://pkg.go.dev/golang.org/x/crypto/pbkdf2)**: Used for secure password-based key derivation.

You can install these dependencies using:

```bash
go mod tidy
```

---

## 🚨 Error Handling

This project includes proper error handling for common scenarios:
- Incorrect password input during decryption results in an informative error message.
- Missing or corrupted files are caught with user-friendly error handling.
- Padding issues during decryption (e.g., due to incorrect keys) are handled gracefully, informing the user about the problem.

---

## 🛠️ Contributing

Contributions to this educational project are welcome. Please open an issue or submit a pull request for improvements.

## Support the Project

If you like this project, consider buying me a coffee!

[![Buy Me A Coffee](https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png)](https://www.buymeacoffee.com/glimor)