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

# trace allocations:
go build -gcflags "-m -m"

# benchmark:
go test -bench=.

#benchmark the memory as well:
go test -bench=. -benchmem

# run the benchmark for 3 seconds and do it 5 times:
go test -bench=BenchmarkMultiply -benchtime=3s -count=5

go install golang.org/x/perf/cmd/benchstat@latest
# go test -bench=. > old.txt
# go test -bench=. > new.txt
# benchstat old.txt new.txt


# cpu profiling requires changes to the program:
import (
	"runtime/pprof"
)
func main() {
	// ...
	f, err := os.Create("cpuprofile.out")
	if err != nil {
	// Handle error
	}
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	// ... (Rest of your code)
}
go build monitor.go
./monitor

# analyze the output:
go tool pprof cpuprofile.out

# create a flame graph:
go tool pprof -web cpuprofile.out

# setup memory profiling:

f, err := os.Create("memprofile.out")
if err != nil {
// Handle error
}
defer f.Close()
runtime.GC()
pprof.WriteHeapProfile(f)



# analyze the memory profiling output
go tool pprof memprofile.out`
go tool pprof -web memprofile.out

# sync workspace code:
go work sync
```

## Notes on packages/modules

- a package must go into its own folder
- the folder must have the same name as the package
- a package can be spread across multiple files
- those files can have any name, `package main` for instance does not need to be `main.go`
- imports for local packages have to include the module path of the thing you are working on


Example on creating a package and organizing some code in a sub-package:
```
spewg-cache/
├── main.go
├── go.mod                  # go mod init example.com/spewg-cache
└── spewg/
    └── server.go
└── example/
    └── ex.go	
```

In the `main.go` file:

```go
package main
import (
	"example.com/spewg-cache/example"
	"example.com/spewg-cache/spewg"
)
```

Items you want to export need to start with a capitalized letter:
```go
func Example()
struct Example{
	
}
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


## Memory management

The garbage collector in Go has some jobs to avoid common mistakes and accidents: 
- it tracks allocations on the heap
- frees unneeded allocations
- keeps the allocations in use. 

These jobs are sometimes referrred to as memory inference, or “What memory should I free?”. The two
main strategies for dealing with memory inference are tracing and reference counting.

Go uses a tracing garbage collector (GC for short), which means the GC will trace objects reachable by a chain of references from “root” objects, consider the rest as “garbage,” and collect them. 

You must have heard this at least once in the tech community: “Garbage collection in Go is automatic, so you can forget about memory management.” Yeah, and I’ve got some prime real estate on the moon to sell you. Believing this is like thinking your house cleans itself because you’ve got a Roomba. In Go, understanding garbage collection is not just a nice-to-have; it’s your ticket to writing efficient, high-performance code. So, buckle up, we’re diving into a world where “automatic” doesn’t mean “magical.”


### Stack and heap allocation

Stack allocation in Go is for variables whose lifetime are predictable and tied to the function calls that create them. These are your local variables, function parameters and return values.

The stack is very efficient because of the LIFO nature. Allocating and deallocating is a matter of moving the stack pointer up or down. This simplicity makes it fast but it also introduces limitations. The stack size is small and if you put too much on it, you will get 'stack overflow'.

Heap allocations are for variables whose lifetime is less predictable and not strictly tied to where they were created. Variables are put here typically in case they need to outlive a function. The heap is more flexible (it can grow beyond the stack size) and dynamic. Variables here can be accessed globaly. The cost of this flexibility is that allocating on the heap is slower due to the need for more complex bookkeeping and the responsibility of managing this memory falls to the garbage collector, which adds overhead.

The Go compiler performs 'escape analysis' to decide if a variable should live on the stack or on the heap. If the compiler determines that the lifetime of a variable doesn’t escape the function it’s in, to the stack it goes. But if the variable’s reference is passed around or returned from the function, then it “escapes” to the heap.

While Go abstracts much of the memory management complexity, having a good understanding of how heap and stack allocations work can greatly impact the performance of your applications.

As a rule of thumb, keep your variables in the scope as narrow as possible, and be cautious with pointers and references that might cause unnecessary heap allocations.

### Go's GC

Go’s garbage collection is based on a concurrent, tri-color mark-and-sweep algorithm.

The GC runs alongside your program and does not need to halt everything to clean it up.

Tri-color refers to how the GC views objects:
- green: in use
- read: ready to delete
- yellow: maybe in use, maybe not

The mark-and-sweep part is the definition of the two main phases of the process. In short, during the `mark` phase, the GC scans your objects, flipping their colors based on accessibility. In the `sweep` phase, it takes out the trash – the red objects.

During the marking phase, there is a 0.3 milliseconds stop the world event to identify the root set. Here, the GC identifies what is in use and what not.

After identifying the root set, the marking happens. This happens concurrently with the program.

During the sweeping phase, actual memory is reclaimed by deleting objects.


### GOGC

The GOGC is the environmental variable that is tuning knob for the garbage collector. The default value is 100, which means that the GC tries to leave at least 100% of the initial heap memory available after a new GC cycle. Adjusting the GOGC value allows you to tailor the garbage collection to the specific needs of your application.

For instance, when set to 50, the GC runs more frequently, keeping the heap size smaller at the expense of CPU. When set to 200, the GC runs less frequently, costing more memory but saving on the CPU.

### GODEBUG

GODEBUG environment variable in Go is a powerful tool for developers, offering insights into the
inner workings of the Go runtime. Specifically, the GODEBUG=gctrace=1 setting is often used to
gain detailed information about garbage collection processes.

### Memory ballast

Allocate a sizeable amount of memory that is always referenced to ensure the GC spends less cycles.

For instance, start off allocating a big array that can handle thousands of bytes for storing session information to mitigate 'refresh storms' ( [Twitch blog](https://blog.twitch.tv/en/2019/04/10/go-memory-ballast-how-i-learnt-to-stop-worrying-and-love-the-heap/)).

Note that this is a strategy that can mask underlying performance issues.

### GOMEMLIMIT

With GOMEMLIMIT, you set a soft cap on the memory usage of the Go runtime, encompassing the
heap and other runtime-managed memory. This cap is like telling your application, 'Here’s your
memory budget; spend it wisely.'

By default, GOMEMLIMIT is set to math.MaxInt64, effectively disabling the memory limit.

### Memory arenas

Go 1.20 introduced an expirimental arena package that offers memory arenas. These arenas can enhance performance by decreasing the number of allocations and deallocations that need to be done during runtime.

Memory arenas are a useful tool for allocating objects from a contiguous region of memory and freeing them all at once with minimal memory management or garbage collection overhead. They are especially helpful in functions that require the allocation of many objects, processing them for a significant amount of time, and then freeing all the objects at the end.

## Analyzing performance

### Escape analysis

Escape analysis is a compiler optimization technique that’s used to determine whether a variable can
be safely allocated on the stack or if it must “escape” to the heap. The primary goal of escape analysis is to improve memory usage and performance by allocating variables on the stack whenever possible since stack allocations are faster and more CPU cache-friendly than heap allocations.

### Pointers

Imagine that you’re at a huge music festival. A pointer is not the stage where the band is playing; it’s the map that shows you where the stage is.

To declare a pointer in Go, you use an asterisk (*) before the type. This tells Go, “This variable is going to hold a memory address, not a direct value.” Here’s how it looks:

```go
var p *int
```

This line declares a pointer, p, that will point to an integer. But right now, p doesn’t point to anything. It’s like having a map with no marked locations. To point it at an actual integer, you must use the address-of operator (&):

```go
var x int = 10
p = &x
```

Now, p holds the address of x. You’ve marked your stage on the festival map.

Dereferencing is how you access the value at the memory address the pointer is holding. You can do
this with the same asterisk (*) you used to declare a pointer, but in a different context:

```go
fmt.Println(*p)
```
This line doesn’t print the memory address stored in p; it prints the value of x that p points to, thanks to dereferencing. You’ve gone from looking at the map to standing in front of the stage, enjoying the music.


### Stack and heap allocation

Here are some best practices concerning allocation:
• Minimize large local variables: Consider using the heap for large data structures to avoid consuming too much stack space
• Be cautious with recursion: Ensure recursive functions have a clear termination condition to prevent stack overflow
• Understand stack versus heap allocation: Use the stack for short-lived variables and the heap for variables that need to outlive the function call


Tracing allocations:
```
go build -gcflags "-m -m"
```


### Benchmarking

Benchmarking is a systematic method of measuring and comparing the performance of software. It’s not just about running a piece of code and seeing how fast it goes; it’s about creating a controlled environment where you can understand the impact of changes in code, algorithms, or system architecture. The goal is to provide actionable insights that guide optimization efforts, ensuring that they’re not just shots in the dark.

### Memory profiling

Memory profiling helps you analyze how your Go program allocates and uses memory. It’s critical in systems programming. where you frequently deal with constrained resources and performance-sensitive operations. Here are some key questions it helps answer:
• Memory leaks: Are you unintentionally holding on to memory that’s no longer needed?
• Allocation hotspots: Which functions or code blocks are responsible for most allocations?
• Memory usage patterns: How does memory use change over time, especially under different
load conditions?
• Object sizes: How can you understand the memory footprint of key data structures?

Memory profiling is setup like this:
```go
f, err := os.Create("memprofile.out")
if err != nil {
// Handle error
}
defer f.Close()
runtime.GC()
pprof.WriteHeapProfile(f)
```

We analyze it using the following:

`go tool pprof memprofile.out`


## Networking

Go lets you communicate over TCP connections using two primary abstractions: 
- `net.Conn`: represents a single TCP connection
- `net.Listener`: waits around for incoming connection requests



HTTP verbs:
- GET: Requests data from a specified resource
- POST: Submits data to be processed to a specified resource
- PUT: Updates a specified resource with provided data
- DELETE: Deletes a specified resource
- PATCH: Applies partial modifications to a resource


HTTP status codes
HTTP status codes are issued by a server in response to a client’s request. They are grouped into five classes:
• 1xx (Informational): The request was received, continuing process
• 2xx (Success): The request was successfully received, understood, and accepted
• 3xx (Redirection): Further action needs to be taken in order to complete the request
• 4xx (Client Error): The request contains bad syntax or cannot be fulfilled
• 5xx (Server Error): The server failed to fulfill an apparently valid request

## Certificates

TLS certificates are a fundamental aspect of secure communication over the internet, providing
encryption, authentication, and integrity. In the context of Go, TLS certificates are used to secure
communication between clients and servers, such as in HTTPS servers or clients that need to securely
connect to other services.

A TLS certificate, often simply called a Secure Sockets Layer (SSL) certificate, serves two main purposes:
- Encryption: Ensures that the data exchanged between the client and server is encrypted, protecting it from eavesdroppers
- Authentication: Verifies the identity of the server to the client, ensuring that the client is talking to the legitimate server and not an imposter

A TLS certificate contains the certificate holder’s public key and identity (domain name), and it is signed by a trusted Certificate Authority (CA). When a client connects to a TLS/SSL-secured server, the server presents its certificate. The client verifies the certificate’s validity, including the CA’s signature, the certificate’s expiration date, and the domain name.


```
openssl version
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365
```

### TLS pitfalls

There is a list of pitfalls and things to keep in mind when we’re dealing with TLS in general.

Let’s look at some of them:
- Validity: Ensure your certificates are valid (not expired) and renew them as necessary. Using expired certificates can lead to service outages.
- Security: Keep your private keys secure. If a private key is compromised, the corresponding certificate can be misused to intercept or tamper with secure communications.
- Trust: For production environments, use certificates issued by a trusted CA. Browsers and clients trust these CAs and will show warnings or block connections to sites with self-signed or untrusted certificates.
- Domain matching: The domain name on the certificate must match the domain name that clients use to connect to your server. Mismatches can lead to security warnings.
- Certificate chains: Understand how to serve the full certificate chain (not just your server’s certificate) to ensure compatibility with clients.
- Performance: TLS/SSL has a performance impact due to the encryption and decryption process. Use efficient cipher suites and consider server and client capabilities.

### UPD in Go

Golang’s net package provides excellent support for UDP programming. Key functions/types include the following:
- net.DialUDP(): Establishes a UDP “connection” (more of a communication channel)
- net.ListenUDP(): Creates a UDP listener to receive incoming packets
- UDPConn: Represents a UDP connection, providing methods such as the following:
  - ReadFromUDP()
  - WriteToUDP()

In UDP, we can apply a technique called Selective Retransmissions (also known as Selective Acknowledgments, or SACK).

### Websockets

The connection is established through a handshake over HTTP but then upgraded to a long-lived TCP connection. Once established, it has minimal message framing overhead, making it suitable for real-time scenarios.

## Telemetry

### Logging

The general guideline is as follows:
- Log consumption tools: Choose JSON for advanced processing tools; choose structured text for simplicity or direct consumption.
- Data complexity: Use JSON for complex, nested data; structured text for simpler data.
- Performance considerations: Opt for structured text when performance is critical; use JSON with performance impact in mind.
- Analysis and troubleshooting: Select JSON for in-depth analysis needs; structured text for basic troubleshooting.
- Team and infrastructure: Consider team familiarity and infrastructure capabilities.


The best practices can be summarized here:
- Use structured logging: Structured logs make it easier to search and analyze data. Use a consistent format such as JSON across your logs
- Implement log rotation and retention policies: Automatically rotate logs and define retention policies to manage disk space and comply with data retention requirements
- Secure log data: Ensure that logs are stored securely, access is controlled, and transmission of log data is encrypted
- Monitor log files for anomalies: Regularly review log files for unusual activity or errors that could indicate operational or security issues

### Tracing

At its core, Golang’s tracing framework leverages the runtime/trace package to let you peer into the running soul of your application. By collecting a wide range of events related to goroutines, heap allocation, garbage collection, and more, it sets the stage for a deep dive into the inner workings of your code.


## Distributing Go code

### Modules

A module is a collection of related Go packages. It serves as a `versionable` and interchangeable unit of source code.

Modules have two main objectives: to maintain the specific requirements of dependencies and to create reproducible builds.

A repository is like a section in a library dedicated to a specific series or collection. Each module represents a book series within this section. Each book series (module) consists of individual books (packages). Finally, each book (package) contains chapters (Go source files), all within the covers (directory) of that book.


### Module workspaces

A Go module workspace is a way to group multiple Go modules that belong to the same project. This feature, introduced to tackle the very beast of dependency management, allows developers to work with multiple modules simultaneously. They aren’t just about neatness. It fundamentally changes how the Go toolchain resolves dependencies.

A Go workspace is a directory containing a unique go.work file referencing one or more go.mod files, each representing a module. This setup permits us to build, test, and manage multiple interrelated modules without the usual headaches of version conflicts.

Within a workspace, the Go compiler treats them as peers instead of relying on an external go.mod file for each module. It looks at the workspace’s go.work file, which lists all modules within the project, making sure everyone plays nicely together.

In other words, workspaces create a self-contained ecosystem for your project. Any changes you make within one module immediately ripple across the others. This streamlines development, particularly when juggling interconnected components of a larger application.

To sync the code in a workspace, use `go work sync`.

Followup:

Learn C Programming - Second Edition: A beginner's guide to learning the most powerful and general-purpose programming 