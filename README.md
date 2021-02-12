# MeowTube

### YouTube client on your terminal
![Example](./termtosvg_k3uvtc_0.svg)

## Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Usage](#usage)
- [Contributing](./CONTRIBUTING.md)

## About <a name = "about"></a>

MeowTube is a CLI (Command Line Interface) to interact with youtube videos or audios and easy to play it via VLC. No need any account to use it. 

## Getting Started <a name = "getting_started"></a>

First, make sure VLC already installed on your machine.

### Prerequisites

1. [VLC Media Player](https://www.videolan.org/vlc/) Installed
2. Register VLC (location where VLC installed) to your `PATH` variable
3. for windows user, better to use [cmder](https://cmder.net), ANSI color supported by default
4. Check everything is good, type `vlc` on your terminal / cmd.

   ```
   vlc
   ```

#### Register VLC to the PATH

- Linux, no need extra step
- Mac OS / OSX, See [Official Doc](https://wiki.videolan.org/MacOS/#Command_line) and then export that location to your `.zshrc` or `.bashrc`
  ```
  echo 'export PATH="$PATH:/Applications/VLC.app/Contents/MacOS"' >> ~/.zshrc && source ~/.zshrc
  ```
- Windows, see [Official Doc](https://wiki.videolan.org/Windows/#Step_2:_Command_Line_Startup) and then copy that directory location to your PATH on your environment variable ([see reference](https://stackoverflow.com/a/44272417)).

### Installing

1. Download MeowTube via [release page](https://github.com/ikhsanalatsary/MeowTube/releases)
2. Choose which target OS do you use
3. Extract the downloaded file
4. Move the file to any folder what you want (OPTIONAL)
5. Register MeowTube to your `PATH` variable (choose where meowtube is located)
6. Create an alias for MeowTube (OPTIONAL)

#### Register MeowTube to the PATH

- unix based(linux & macos). export to `.zshrc` or `.bashrc`
  ```
  echo 'export PATH="$PATH:$HOME/MEOWTUBE_LOCATION_FOLDER"' >> ~/.zshrc && source ~/.zshrc
  ```
- Windows, copy the directory location to your PATH on your environment variable ([see reference](https://stackoverflow.com/a/44272417))

## Usage <a name = "usage"></a>

Check everything is good. Type on your terminal

```
meowtube
```

or

```
meowtube --help
```

### Command Line Arguments

| Argument | Description                                          |
| -------- | ---------------------------------------------------- |
| help     | Help about any command                               |
| play     | To play YouTube video                                |
| popular  | To see popular videos on YouTube                     |
| search   | To search for videos according to certain characters |
| trending | To see trending videos on YouTube                    |

#### Play Arguments

| Argument    | Description                                                                         |
| ----------- | ----------------------------------------------------------------------------------- |
| :YoutubeURL | Valid YouTube video url e.g: `https://youtu.be/0FZZJHuQMFs`                         |
| :videoId    | Valid Youtube videoId e.g: `"tMzjKjV6r_w"`                                          |
| audio       | To play audio only                                                                  |
| playlist    | To play all videos from YouTube playlist                                            |
| list        | shorthand for playlist. To play all videos from YouTube playlist                    |
| video       | To play YouTube video                                                               |

**NOTE:** Every argument has `--help` flag to see their specific usage

### Global config
This is to use excluding or including invidious instances. You can add this config on your `$HOME` PATH and named it with `.meowtube.yaml`. For excluding, you can set it as false. Example:

```
invidious.fdn.fr: false
invidious.kavin.rocks: false
invidious.snopyta.org: false
yewtu.be: true
ytprivate.com: false
```

## ✍️ Authors <a name = "authors"></a>

- [@ikhsanalatsary](https://github.com/ikhsanalatsary) - Idea & Initial work

## 🎉 Acknowledgements <a name = "acknowledgement"></a>

- Inspired by [ohmyzsh/spotify](https://github.com/ohmyzsh/ohmyzsh/blob/master/plugins/osx/spotify)
- Using [Invidious APIs](https://github.com/iv-org/invidious)
