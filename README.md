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
  
</details>

<details>
<summary> To install application on Windows </summary>

  <br>
  
  Download the executable as zip from repo [mankka.zip](https://github.com/MikkoVasankari/mankka/raw/main/mankka.zip).
  
</details>

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
- [ ] Make Windows installation script.
