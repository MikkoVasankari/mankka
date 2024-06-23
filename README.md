# mankka
Go TUI application to play media files with MPV

Requirements: <br>
- [mpv](https://mpv.io/) <br>
- [Golang](https://go.dev/doc/install)

### To install project on linux
!!! This command tries to install mpv media player if user don't have it.
``` 
.install/install.sh
```

### Running the application
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
