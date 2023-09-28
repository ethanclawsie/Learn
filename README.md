## Installation

### Prerequisites

Before you can use the Learn CLI, ensure that you have the following prerequisites installed:

- [Go](https://golang.org/doc/install) (only required if you want to build from source)

### Download and Run

1. **Download the Executable**

   Visit the [Releases](https://github.com/yourusername/learn-cli/releases) section of this repository and download the appropriate binary for your operating system (e.g., `learn` for Unix-like systems or `learn.exe` for Windows).

2. **Move the Executable**

   Place the downloaded `learn` (or `learn.exe`) executable in a directory that is included in your system's `PATH`. This step is necessary to run the CLI from any location in the terminal.

   - For Unix-like systems (Linux and macOS), you can add the following line to your shell profile file (e.g., `~/.bashrc`, `~/.zshrc`, or `~/.bash_profile`):

     ```bash
     export PATH=$PATH:/path/to/learn-cli-directory
     ```

   - On Windows, you can add the directory containing `learn.exe` to your system's `PATH` via the system settings.

3. **Run the CLI**

   You can now run the Learn CLI from anywhere in the terminal by typing:

   ```bash
   learn
   ```

### Usage

Here are some basic usage instructions for the Learn CLI:

1. When prompted, enter the full path to the text file containing your terms and definitions.
2. Choose whether you want to learn terms in random order (type "yes" or "no").
3. Start interactive learning by answering questions. The CLI will provide feedback on your answers.
