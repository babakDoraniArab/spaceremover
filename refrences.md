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