# mankka
Go TUI application to play media files with MPV.

Requirements: <br>
  - [mpv](https://mpv.io/)
  - [Golang](https://go.dev/doc/install)

<br>
<details>
<summary> To install application on Linux </summary>

  <br>

  ```
  .install/install.sh
  ```
  This script tries to install mpv media player if user don't have it.
  To use the application follow [Using the application instructions](https://github.com/MikkoVasankari/mankka?tab=readme-ov-file#using-the-application).

  alternatively you can download the executable as zip from repo [mankka.zip](https://github.com/MikkoVasankari/mankka/raw/main/mankka.zip).
_________________
</details>
<details>
<summary> To install application on Windows </summary>

  <br>

  Download the executable as zip from repo [mankka.zip](https://github.com/MikkoVasankari/mankka/raw/main/mankka.zip).

_________________
</details>
<details>
<summary> Build from source </summary>

  <br>

  ```
  git clone https://github.com/MikkoVasankari/mankka.git
  ```
  Navigate to cloned repo and run go build or installation script suitable for you from /install/ directory


  ```
  go build
  ```

</details>

_________________
### Using the application
```
mankka .
```

You can also use:
```
mankka
```
or to use your path
```
mankka /path/to/your/dir
```

### TODO
- [ ] Create a sqlite db file for saving regularly used filepaths.
- [ ] Make setting for no-video flag / In general make setting how mpv should act
- [ ] Make that "shift + enter" updates the list of playable files if the entry is a directory
- [ ] Add Current song to terminal title
