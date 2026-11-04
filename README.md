# Orion

> An ultra-lightweight personal AI assistant written in Go. Runs on $10 hardware, boots in about a second, and uses only a few MB of RAM.

## About

Orion is a small, fast AI agent you run yourself. It connects to the LLM provider of your choice, chats with you from the terminal, a web console, or chat apps, and can search the web, schedule tasks, and remember context. Because it's a single compact Go binary, it runs almost anywhere, from a laptop to a Raspberry Pi to an old Android phone.

### Key features

- **Ultra-lightweight**: tiny memory footprint and near-instant startup, even on low-end boards.
- **Runs anywhere**: Linux (x86, ARM, RISC-V), macOS, Windows, Android, and Docker.
- **Bring your own model**: configure any supported LLM provider with your own API key.
- **Multiple interfaces**: command line, browser-based web console, terminal UI, and chat-app gateway.
- **Automation tools**: web search, scheduled tasks, logging, planning, and memory.

## Who is it for?

- **Self-hosters and hobbyists** who want a personal AI assistant on a Raspberry Pi, old phone, or cheap Linux board instead of a full computer.
- **Developers** who want a small, hackable Go codebase to build agents on.
- **Privacy-minded users** who want an assistant they run and configure themselves.
- **Edge and IoT builders** who need an AI agent on resource-constrained devices.

### How it helps

- Cuts hardware cost: no need for a Mac Mini or a server, since a $10 to $50 board is enough.
- Starts almost instantly and stays light on RAM, so it can run all day in the background.
- Gives you a simple setup path, from one-click launcher to a single JSON config file.

## Download and Install

### Option 1: Download a prebuilt binary (recommended)

1. Go to the [Releases page](https://github.com/YOUR-USERNAME/orion/releases).
2. Download the file for your OS and architecture.
3. Extract it and (on Linux/macOS) make it executable:

```bash
tar xzf orion_Linux_arm64.tar.gz
chmod +x orion
```

> **macOS:** the first launch may be blocked because the app isn't notarized. Open **System Settings → Privacy & Security** and choose **Open Anyway**.

### Option 2: Docker Compose

```bash
git clone https://github.com/YOUR-USERNAME/orion.git
cd orion
docker compose -f docker/docker-compose.yml --profile launcher up -d
# Open http://localhost:18800
```

### Option 3: Build from source

Requires Go 1.25+ (and Node.js 22+ with pnpm for the web UI).

```bash
git clone https://github.com/YOUR-USERNAME/orion.git
cd orion
make deps
make build
```

### Android

Install the APK from the Releases page, or run the Linux ARM64 build inside [Termux](https://termux.dev).

## Usage

### 1. Web console (easiest)

Double-click `orion-launcher` (or `orion-launcher.exe` on Windows), or run:

```bash
orion-launcher
```

Then open <http://localhost:18800> in your browser to configure a provider and start chatting. Add the `-public` flag if you need access from another machine, Docker, or a VM.

### 2. Terminal UI (headless servers, SSH)

```bash
orion-launcher-tui
```

Follow the menus: **Configure a Provider → Configure a Channel → Start the Gateway → Chat.**

### 3. Command line

Initialize your config and workspace:

```bash
orion onboard
```

This creates `~/.orion/config.json`. Add your model and API key:

```json
{
  "agents": {
    "defaults": {
      "model_name": "your-model-name"
    }
  },
  "model_list": [
    {
      "model_name": "your-model-name"
    }
  ]
}
```

Then run it:

```bash
# Ask a single question
orion agent -m "What is 2+2?"

# Interactive chat
orion agent

# Start the gateway for chat-app integrations
orion gateway
```

## Contributing

Contributions are welcome! Fork the repo, create a branch, and open a pull request. Bugs and ideas go in [Issues](https://github.com/YOUR-USERNAME/orion/issues).

## Credits and License

Orion is based on INDIA based startup by orion. Released under the [MIT License](LICENSE).

## Contact

Maintained by VARUN NEGI(https://github.com/YOUR-USERNAME) 
