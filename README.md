# Balance Timer

## About

This is a desktop Timer app that can help the user manage Productivity & Rest periods during work sessions.

Use it to find time for taking breaks. Which in turn will help:  
🧘 connect with reality and stay mindful throughout the day,  
💤 battle fatigue,  
🧠 focus better and get more stuff done,  
🌱 stand up more often -> health benefits,  
🙂 improve the mood.  

## Technologies used

It is written in Go + Wails on the backend and Vue + TS on the frontend.

## Usage

The executable can be acquired in two ways:

### Option 1. Download

[📦 Latest release](https://github.com/gVguy/balance-timer/releases/)

### Option 2. Build locally

Following tools are required to be installed on your system:
- go 1.21+
- npm (node 15+)
- (mac) xcode command line tools
- wails cli

You can follow the official wails [installation guide](https://wails.io/docs/gettingstarted/installation) for detailed instructions.

1. Clone this repository and `cd` into the project folder
2. *(optional)* Run `wails doctor` to check for missing dependencies
3. Run `wails build`
4. Find your executable in `./build/bin/`
