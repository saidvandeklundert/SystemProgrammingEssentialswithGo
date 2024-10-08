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

### Symbolic links

A link or 'pointer' to the place where the actual file is.

In the Linux command line, you can create a symbolic link using the ln command with the -s option:
```
ln -s /home/user/documents/important_document.txt /home/user/desktop/
shortcut_to_document.txt
```

Here’s what’s happening:
• `ln`: This is the command for creating links
• `-s`: This option specifies that we’re creating a symbolic link (symlink)
• `/home/user/documents/important_document.txt`: This is the source file you
want to link to
• `/home/user/desktop/shortcut_to_document.txt`: This is the destination where
you want to create the symbolic link


### Unlinking files

Unlinking a file or symbolic link is removing a file or a symbolic link.

### Memory mapped files

The idea of memory-mapped files was popularized by the UNIX operating system in the 1980s. The mmap system call, introduced in the early versions of UNIX, allowed processes to map files or devices into their address space. This provided a seamless way to work with files as if they were in memory, without the need for explicit file I/O operations.

Good blog [here](https://medium.com/@ZaradarTR/wtf-is-memory-mapped-files-9448c04078a3).


## Signal

Notification to a process that an event has occured. When the kernel generates a signal for a process, it is usually due to an event occurring in one of these
three categories: 
- hardware-triggered events
- user-triggered events
- and software events


In Go, you can turn to the `os/signal` packages to deal with both synchronous as well as asynchronous events.

SIGINT signal is sent to a process in response to the user pressing the interrupt character on the
controlling terminal. The default interrupt character is ^C (Ctrl + C).

Signal handling is crucial for several reasons:
- graceful shutdown: SIGTERM or SIGINT
- resource management: SIGUSR1 and SIGUSR2
- inter-process communication: instruct processes to perform an action, like SIGSTOP or SIGCONT
- emergency stops: SIGKILL or SIGABRT

## Scheduling

Go’s standard library provides several features that can be used to create a job scheduler, such as
goroutines for concurrency and the time package for timing events.

## Pipes

Pipes are fundamental tools in inter-process communication (IPC), allowing for data transfer between system processes.

A pipe is like a conduit within memory designed for transporting data between two or more processes. This conduit adheres to the producer-consumer model: one process, the producer, funnels data into the pipe, while another, the consumer, taps into this stream to read the data. Pipes establish a unidirectional flow of information where the pipe has a write-end and a read-end. If two-way communication is required, 2 pipes have to be used.

Pipes are used for a variety of tasks:
- CLI-utilities
- data streaming
- inter-process data exchange

There are similarities between pipes and Go-channels:
• `Communication mechanisms`: Both pipes and channels are primarily used for communication. Pipes facilitate IPC, while channels are used for communication between goroutines withina Go program.
• `Data transfer`: At a basic level, both pipes and channels transfer data. In pipes, data flows from one process to another, while data is passed between goroutines in channels.
• `Synchronization`: Both provide a level of synchronization. Writing to a full pipe or reading from an empty pipe will block the process until the pipe is read from or written to, respectively. Similarly, sending to a full channel or receiving from an empty channel in Go will block the goroutine until the channel is ready for more data.
• `Buffering`: Pipes and channels can be buffered. A buffered pipe has a defined capacity before
it blocks or overflows, and similarly, Go channels can be created with a capacity, allowing a
certain number of values to be held without immediate receiver readiness.

The following are the differences:
differences:
• Direction of communication: Standard pipes are unidirectional, meaning they only allow data flow in one direction. Channels in Go are bidirectional by default, allowing data to be sent and received on the same channel.
• Ease of use in context: Channels are a native feature of Go, offering integration and ease of use within Go programs that pipes cannot match. As a system-level feature, pipes require more setup and handling when used in Go.


Use pipes in the following scenarios:
• You must facilitate communication between different processes, possibly across different programming languages
• Your application involves separate executables that need to communicate with each other
• You work in a Unix-like environment and can leverage robust IPC mechanisms

Use Go channels when the following applies:
• You are developing concurrent applications in Go and need to synchronize and communicate between goroutines
• You require a straightforward and safe way to handle concurrency within a single Go program
• You must implement complex concurrency patterns, such as fan-in, fan-out, or worker pools, which Go’s channel and goroutine model elegantly handle


Exampel of pipes in use in a CLI tool:
```
cat file.txt | grep "flower"
```

Named pipes are not limited to live processes, unlike anonymous pipes. They can be used between any processes and persist in the filesystem.

You can see named pipes in the filesystem too.
1. `-`: Regular file
2. `d`: Directory
3. `l`: Symbolic link
4. `c`: Character special file
5. `b`: Block special file
6. `p`: Named pipe (FIFO)
7. `s`: Socket

## Unix sockets

Unix sockets, aka Unix domain sockets, allow processes to communicate with each other on the same machine quickly and effeciently, offering an alternative to TCP/IP sockets for IPC. The feature is unique to Unix and Unix-like operating systems, such as Linux.

Unix sockets are ether stream-oriented (such as TCP) or datagram-oriented (such as UDP). They are represented as filesystem nodes, such as files and directories. However, they are not regular files but 'special' IPC mechanisms.

Three key Unix sockets features:
- `efficiency`: no networking overhead.
- `filesystem namespace`: Unix sockets are referenced by filesystem paths. This makes them easy to locate and use but also means they persist in the filesystem until explicitly removed.
- `security`: access to Unix sockets can be controlled using filesytem permissions, providing a level of security based on user and group IDs.

Inspecting sockets is done with `lsof` (list open files). This command offers insights into files accessed by processes. Unix sockets, treated as file, can be examined using `lsof` to gather relevant information.

You can run `lsof` for specific sockets:
```
lsof -Ua /tmp/example.sock
```


Unix domain sockets don’t require the network stack’s overhead, as there’s no need to route data
through the network layers. This reduces the CPU cycles spent on processing network protocols. Unix
domain sockets often allow for more efficient data transfer mechanisms within the kernel, such as
sending a file, which can reduce the amount of data copying between the kernel and user spaces. They
communicate within the same host, so the latency is typically lower than TCP sockets, which may
involve more complex routing even when communicating between processes on the same machine.

It is faster then simply calling the loopback interface because the loopback interface still goes through the TCP/IP stack, even though it doesn’t leave themachine. This involves more processing, such as packaging data into TCP segments and IP packets.

They can be more efficient regarding data copying between the kernel and user spaces. Some Unix
domain socket implementations allow for zero-copy operations, where data is directly passed
between the client and server without redundant copying. This is not possible using TCP/IP since its
communication typically involves more data copying between the kernel and user spaces.

Several systems rely on the benefits of Unix domain sockets, such as D-Bus, Systemd, MySQL/PostgreSQL, Redis, Nginx and Apache.