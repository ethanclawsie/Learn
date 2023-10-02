# Learn CLI

Learn CLI is a command-line application designed to help users learn and memorize terms and their definitions efficiently. Whether you're studying for an exam, expanding your vocabulary, or simply enhancing your knowledge, Learn CLI provides a convenient way to learn and practice.

## Features

- **Customizable Learning Experience**: Choose between learning terms in random order or sequentially to suit your preferences.
- **Chunked Learning**: Terms are grouped into manageable chunks, making it easier to absorb and retain information.
- **Persistent Configuration**: Learn CLI remembers your last used file and preferences, so you can pick up where you left off.
- **Interactive Learning**: Engage in an interactive learning process where you input your definitions and receive immediate feedback.
- **Easy Setup**: Simply provide the path to a text file containing terms and definitions, and you're ready to start learning.

### Download and Run

1. **Download the Executable**

   Visit the [Releases](https://github.com/ethanclawsie/Learn/releases) section of this repository and download the appropriate binary for your operating system (e.g., `learn` for Unix-like systems or `learn.exe` for Windows).

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

4. **Configuration**:

   - The application will prompt you to provide the path to a text file containing terms and definitions. If you're returning to your previous session, it will remember your last-used file.

5. **Learning**:

   - Choose between learning terms randomly or sequentially as per your preference.
   - Follow the on-screen prompts to answer the definitions and receive immediate feedback.
   - Incorrect answers will be re-asked until you get them right.

6. **Completion**:
   - Upon successfully answering all terms, the application will congratulate you on completing the learning session.

## Example Term File Format

The text file containing terms and definitions should follow this format: Term : Definiton

## Configuration File

Learn CLI uses a configuration file (`learnconfig.json`) to store your preferences and the last-used file path. This file is located in your user's home directory.

## Dependencies

Learn CLI relies on the following Go packages:

- `github.com/urfave/cli`: For creating the command-line interface.
- Standard Go libraries for file handling and randomization.
