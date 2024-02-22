# Event-Driven Web Servers

## Introduction
The purpose of a Web Server is to handle multiple concurrent connections, efficiently handling them while avoiding race conditions and performance issues. 
To maintain multiple connections at the same time, multiples threads can be used where each request will have its own thread. Since each connection ideally is short-lived, we don't need many threads executing in parallel. The problem with this approach is that the programmer might have to deal with synchronization issues, where multiple threads will compete for shared resource, requiring some kind of coordination for this to happen correctly. There is also a problem with how expensive is maintaining threads. Languages like Go have a very unique and effective way to deal with concurrent processing, using goroutines is way more efficient than creating threads, but even so we are not going to use this approach, instead, we are going to manage every request in one single thread where this thread will be handling tiny "slices" of many requests at once.

## Building the Event-Driven Web Server with Go

### The "net" package
Golang has a standard package called `net`. Different from `net/http`, it does not offer http support and is not as simple to create a web server, instead we will create sockets and read byte by byte, building the payload.

With the server running and listening, we can send data with `nc localhost {port}`

