# mankka
Go TUI application to play media files with MPV

Requirements: <br>
  - [mpv](https://mpv.io/) 
  - [Golang](https://go.dev/doc/install)

<br>
<details>
<summary> To install project on linux </summary>

  <br>
  
  ```
  .install/install.sh
  ```
  This script tries to install mpv media player if user don't have it.
  
  then follow usage [instructions](https://github.com/MikkoVasankari/mankka?tab=readme-ov-file#using-the-application).
  
  <br>
  
  alternatively you can
  Download the executable as zip from repo [mankka](https://github.com/MikkoVasankari/mankka/raw/main/mankka.zip)
  
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
- [ ] Create a check for file compatability with by mpv.
