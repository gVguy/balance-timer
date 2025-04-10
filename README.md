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

## Installation

The executable can be acquired in two ways:

### Option 1. Download

[📦 Latest release](https://github.com/gVguy/balance-timer/releases/)

The app is released cross-platform. You can download a build for your os from the latest release assets.

### Option 2. Build locally

Following tools are required to be installed on your system:
- go 1.21+
- npm (node 15+)
- wails cli

You can follow the official Wails [installation guide](https://wails.io/docs/gettingstarted/installation) for detailed instructions and platform-specific dependencies.

1. Clone this repository and `cd` into the project folder
2. *(optional)* Run `wails doctor` to check for missing dependencies
3. Run `wails build`
4. Find your executable in `./build/bin/`

## Troubleshooting

### “Balance Timer” is damaged and can’t be opened.

If you're on a macOS and have downloaded the app, you're likely to see the message saying it's damaged. It's not. It's just apple trying to make me pay a hundred bucks for signing my free open-source app.


<details>
  <summary>More info</summary>
  Apple's policy is that every application must be signed with a Developer ID, which at the time of writing costs $99/year. Balance Timer is not sponsored and distributed open-source, so I don't expect to earn anything from it. Therefore it is not signed with such ID.

  When an unsigned app is downloaded from the internet, macOS marks it as "damaged" to prevent unidentified software from executing.
</details>

To **workaround** this warning follow these steps:
1. Launch the terminal
2. `cd` into the folder containing the executable
3. Run `xattr -cr "Balance Timer.app/"` (assuming you didn't rename the app).

> The release files are automatically built in a github workflow from the source code in this repository.
>
> Still, if you don't trust in the executable's integrity and don't want to run it, you have an option of building your own executable from source (described above).
