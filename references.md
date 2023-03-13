https://pkg.go.dev/io/ioutil

The os.FileInfo interface provides a number of methods that you can use to get information about a file or directory entry. Some of the most commonly used methods include:

`Name()` string: returns the name of the file or directory.
`Size()` int64: returns the size of the file in bytes. For directories, this value is implementation-dependent and may return 0 or another value.
`Mode()` FileMode: returns the file mode bits.
`ModTime()` time.Time: returns the modification time of the file.
`IsDir()` bool: returns true if the file is a directory.
`Sys()` interface{}: returns the underlying data source (can be nil).

```sh
fileInfo, err := os.Stat("/path/to/file")
if err != nil {
    log.Fatal(err)
}

fmt.Println(fileInfo.Name())
fmt.Println(fileInfo.Size())
fmt.Println(fileInfo.Mode())
fmt.Println(fileInfo.ModTime())
fmt.Println(fileInfo.IsDir())


```


## how to make linux executable file build it and package it 

```sh
GOOS=linux GOARCH=amd64 go build -o myapp
tar czvf myapp.tar.gz myapp

```
## how to use it in linux,
 just download it from link below and extract it . then copy this file to your local bin

```sh
tar xzvf myapp.tar.gz
sudo cp myapp /usr/local/bin/
```

## how to make mac executable file build it and package it 

```sh
GOOS=darwin GOARCH=amd64 go build -o myapp
hdiutil create -volname "MyApp" -srcfolder "myapp" -ov -format UDZO "myapp.dmg"


```
## how to use it in linux,
 Download the DMG file (myapp.dmg) from the publicly accessible location.

Double-click the DMG file to mount it and open the disk image.

Drag the myapp executable binary to the Applications folder.

Optionally, create a symlink to the myapp executable binary in a directory in the system's PATH, so that the myapp command can be run from the command line.
```sh
sudo ln -s /Applications/myapp /usr/local/bin/

```