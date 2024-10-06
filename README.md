# SystemProgrammingEssentialswithGo

System Programming Essentials with Go - the book

Corresponding repo:
https://github.com/PacktPublishing/System-Programming-Essentials-with-Go

## Prep:

```
# install go
sudo apt update && sudo apt upgrade
sudo apt install golang-go
wget -c https://golang.org/dl/go1.23.2.linux-amd64.tar.gz
sudo tar -C /usr/local -xvzf go1.23.2.linux-amd64.tar.gz
export  PATH=$PATH:/usr/local/go/bin
go env -w GOPROXY=direct
go install -v golang.org/x/tools/gopls@latest

# make the first project
mkdir helloworld
cd .\helloworld
# make the file
go build main.go
go run .

# format the code
go fmt

# run tests:
go test

# check for potential errors or suspicious constructs
go vet

# get another package
go get <package name>
go get golang.org/x/sys/unix
```



## System calls

Sytem calls, aka 'syscalls', are low-level functions provided by the oepration system kernal that allow user-level processes to request services from the kernel.

A processor/CPU has 2 modes of operation:
- user mode: limited access to CPU and memory
- kernel mode: unrestricted access to CPU and memory

From user space, you can use syscalls to cross the border between user and kernel spaces. The syscall API offers a variety of services from creating new processes to handing input and output (I/O) operations. A numerical code uniquely identifies each operation, but we can interact with them through their names.

An example of a syscall is 'open', to open a file. There are more examples listed in a blog [here](https://filippo.io/linux-syscall-table/) and in the Linux git [here](https://filippo.io/linux-syscall-table/).

To make syscalls in Go, there is the older syscall package and the newer x/sys package.

### Brief overview of syscalls

*File operations*

These functions let us interact with general files:
• unix.Create(): Create a new file
• unix.Unlink(): Remove a file
• unix.Mkdir(), unix.Rmdir(), and unix.Link(): Create and remove directories
and links
• unix.Getdents(): Get directory entries

*Signals*

Here are two examples of functions that interact with OS signals:
• unix.Kill(): Send a kill signal to a process
• unix.SIGINT: Interrupt signal (commonly known as Ctrl + C)


*User and group management*

We can manage users and groups using the following calls:
• syscall.Setuid(), syscall.Setgid(), syscall.Setgroups(): Set user and
group IDs

*System information*

We can analyze some statistics on memory and swap usage and the load average using the
Sysinfo() function:
• syscall.Sysinfo(): Get system information

*File descriptors*

While it’s not an everyday task, we can also interact with file descriptors directly:
• unix.FcntlInt(): Perform various operations on file descriptors
• unix.Dup2(): Duplicate a file descriptor

*Memory-mapped files*

Mmap is an acronym for memory-mapped files. It provides a mechanism for reading and writing
files without relying on system calls. When using Mmap(), the operating system allocates a section
of a program’s virtual address space, which is directly “mapped” to a corresponding file section. If
the program accesses data from that part of the address space, it will retrieve the data stored in the
related part of the file:
• syscall.Mmap(): Map files or devices into memory

*Operating system functionality*

The os package in Go provides a rich set of functions for interacting with the operating system. It’s
divided into several subpackages, each focusing on a specific aspect of OS functionality.
The following are file and directory operations:
• os.Create(): Creates or opens a file for writing
• os.Mkdir() and os.MkdirAll(): Create directories
• os.Remove() and os.RemoveAll(): Remove files and directories
• os.Stat(): Get file or directory information (metadata)
• os.IsExist(), os.IsNotExist(), and os.IsPermission(): Check file/directory
existence or permission errors
• os.Open(): Open a file for reading
• os.Rename(): Rename or move a file
• os.Truncate(): Resize a file
• os.Getwd(): Get the current working directory
• os.Chdir(): Change the current working directory
• os.Args: Command-line arguments
• os.Getenv(): Get environment variables
• os.Setenv(): Set environment variables

*The following are for processes and signals:*

• os.Getpid(): Get the current process ID
• os.Getppid(): Get the parent process ID
• os.Getuid() and os.Getgid(): Get the user and group IDs
• os.Geteuid() and os.Getegid(): Get the effective user and group IDs
• os.StartProcess(): Start a new process
• os.Exit(): Exit the current process
• os.Signal: Represents signals (for example, SIGINT, SIGTERM)
• os/signal.Notify(): Notify on signal reception


You can also start a process/cmd using the os package:
```go
    // 	"os/exec"
	cmd := exec.Command("ls", "-ltr")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
```

*Best practices*

As a system programmer using the os and x/sys packages in Go, consider the following best practices:
• Use the os package for most tasks, as it provides a safer and more portable interface
• Reserve the x/sys package for situations where fine-grained control over system calls is necessary
• Pay attention to platform-specific constants and types when using the x/sys package to ensure
cross-platform compatibility
• Handle errors returned by system calls and os package functions diligently to maintain the
reliability of your applications
• Test your system-level code on different operating systems to verify its behavior in
diverse environments


### Checking out syscalls

Use `apt-get install strace -y` and then prefix it for a command, like the following:
```
strace ls
```

You can also trace specific calls by providing additional args, such as the following:
```
strace -e execve ls
```

When you build your app, you can also trace syscalls this way:
```
go build -o app main.go
strace ./app
strace -e write ./app
```

### File descriptors

File descriptors can represent different types of resources:
• Regular files: These are files on disk containing data
• Directories: Representations of directories on disk
• Character devices: Provide access to devices that work with streams of characters, such as
keyboards and serial ports
• Block devices: Used for accessing block-oriented devices, such as hard drives
• Sockets: For network communication between processes
• Pipes: Used for inter-process communication (IPC)

When a shell starts a process, it usually inherits three open file descriptors. Descriptor 0 represents the standard input, the file providing input to the process. Descriptor 1 represents the standard output, the file where the process writes its output. Descriptor 2 represents the standard error, the file where the process writes error messages and notifications regarding abnormal conditions. 

`stdin`, `stderr`, and `stdout` are integral to the development of effective, user-friendly,
and interoperable CLI applications. These standardized streams provide a versatile, flexible, and reliable
means of handling input, output, and errors. By embracing these streams, our CLI applications become
more accessible and valuable to users, enhancing their ability to automate tasks, process data, and
achieve their goals efficiently.

Honoring the streams for example makes the following possible:
```
go run main.go word1 word2 word3 > stdout.txt 2> stderr.txt
```

Everything that is written to `stderr` will be written to `stderr.txt`.

## Files and permissions

In Linux, files are categorized into various types, each serving a unique purpose.

*Regular files*:

Contain data such as text, images or programs.

First char `ls` shows is a `-`.

The `FileMode.IsRegular()` can be checked to see if we are dealing with a regular file.

*Directories*:

Hold other files and directories.

First char ls shows is `d`. 

The `FileMode.IsDir()` can be checked to see if we are dealing with a directory.

*Symbolic links*:

These are pointers to other files. They are denoted by `l` in the first char of `;s`.

The `FileMode` does not tell us if it is a symbolic link, bt we can check if `FileMode` & `os.ModeSymlink` is non-zero.

*Named pipes (FIFOs)*:

Named pipes are mechanisms for inter-process communication, denoted by a `p` in the first char of the file listing. The `os.ModeNamedPipe` bit represents a named pipe.

*Character devices*:

Character devices provide unbuffered, direct access to hardware devices, and are denoted by a c in the first character of the file listing. The `os.ModeCharDevice` bit represents a character device.

*Block devices*:

Provide buffered access to hardware devices and are denoted by a `b` in the first character of the file listing. The `FileMode` does not give you the info, but the os package should allow you to work with block devices.

*Sockets*:

Endpoints for communication, denoted by a `s` in the first char of the file listing. The `os.ModeSocket` but represents a socket.


### Files and permissions:

The FileMode type in Go encapsulates these bits and provides methods and constants for working
with file types and permissions, making it easier to perform file operations in a cross-platform way.

In Linux, the permissions system is a crucial aspect of file and directory security. It determines who
can access, modify, or execute files and directories. Permissions are represented by a combination of
read (r), write (w), and execute (x) permissions for three categories of users: owner, group, and others.

Let’s refresh what these permissions represent:
• `Read (r)`: Allows reading or viewing the file’s contents or listing a directory’s contents
• `Write (w)`: Allows modifying or deleting a file’s contents or adding/removing files in a directory
• `Execute (x)`: Allows executing a file or accessing the contents of a directory (if you have execute
permission on the directory itself)

Linux file permissions are typically displayed in the form of a 9-character string, such as rwxr-xr—,
where the first three characters represent permissions for the owner, the next three for the group, and
the last three for others.

When we combine the file type and its permissions, we form the 10-character string that the ls -l
command returns in the first column of the following example:

```
-rw-r--r-- 1 user group 0 Oct 25 10:00 file1.txt
-rw-r--r-- 1 user group 0 Oct 25 10:01 file2.txt
drwxr-xr-x 2 user group 4096 Oct 25 10:02 directory1
```

Permissions can be returned by `FileInfo.Mode().Perm()` and they are returned in octal value. For example, rwx (read, write,execute) is 7 (4+2+1), r-x (read, no write, execute) is 5 (4+0+1), and so on. So, for example, the permissions -rwxr-xr-- can be succinctly represented as 755 in octal.

### File paths

A file path is a string representation of a file or directory's location within a filesystem. linux example: `/home/klundert/`.

Go offers abstractions over platform-specific implementations n the `path/filepath` package.