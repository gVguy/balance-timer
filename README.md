# Balance Timer

![App UI preview: productivity session in progress](https://i.imgur.com/Sm5yBxF.jpeg)

## About

**Balance Timer** is a simple desktop app that helps you manage Productivity and Rest periods throughout the day.

When a productivity session ends, the app opens in fullscreen to remind you to take a proper break. Once the rest period is over, it invites you to get back to work.

Taking regular breaks can help you:  
🧘 stay mindful and connected with the present moment,  
💤 battle fatigue,  
🧠 improve focus and productivity,  
🌱 move around more → health benefits,  
🙂 improve your mood.  

## Installation

The executable can be acquired in two ways:

### Option 1. Download

The app is released cross-platform. You can download a build for your os from the latest release assets.

📦 [MacOS build (arm64)](https://github.com/gVguy/balance-timer/releases/latest/download/BalanceTimer-macos.zip)  
📦 [Windows build (amd64)](https://github.com/gVguy/balance-timer/releases/latest/download/BalanceTimer-windows.zip)  
📦 [Linux build (amd64)](https://github.com/gVguy/balance-timer/releases/latest/download/BalanceTimer-linux.zip)

### Option 2. Build locally

If you don't see a build that matches your system above or you prefer not to run an unsigned app downloaded from the internet, you can build your own executable from source code.

#### Perquisites

Following tools are required to be installed on your system:
- go 1.21+
- npm (node 15+)
- wails cli

You can follow the official Wails [installation guide](https://wails.io/docs/gettingstarted/installation) for detailed instructions and platform-specific dependencies.

#### Build steps:

1. Clone this repository and `cd` into the project folder
2. *(optional)* Run `wails doctor` to check for missing dependencies
3. Run `wails build`
4. Find your executable in `./build/bin/`

## Troubleshooting

### App is damaged / unrecognized app

Some systems might block the app on the first launch after downloading. This happens because the app is not signed with a verified certificate.

#### macOS

You may encounter this error:

> **"Balance Timer" is damaged and can’t be opened. You should move it to the Trash.**

It's not really damaged, it's just how macOS treats unsigned apps downloaded from the internet.

**Workaround**:
1. Launch the Terminal
2. `cd` into the folder containing the app
3. Run `xattr -cr "Balance Timer.app/"` (assuming you didn't rename the file)
4. Open the app normally from Finder, this time it should launch without warnings.

You only need to do this once.

#### Windows

Windows may show a warning like:

> **Windows protected your PC. SmartScreen prevented an unrecognized app from starting.**

If you see this warning, click **"Run anyway"**.

If you don't see "Run anyway" button, click on **"More info"** first.

#### Security concerns

Since this app is free, open-source and not sponsored, it's distributed unsigned.

However, the app is safe if you downloaded it with a link above or from this repository's [releases page](https://github.com/gVguy/balance-timer/releases/).

Release build files are automatically generated from the source code in this repo with github-actions. Both the fact that release is created by github-actions and the linked commit are visible for each release.

You are welcome to inspect the source code as well as the action responsible for the release (found in `.github/workflows/`).

Still, if you don't trust in the executable's integrity and don't want to run it, you have an option of building your own executable from source (described above).

## Technologies used

- Wails
- Go
- Vue.js
- TypeScript
- Sass
- Github actions
