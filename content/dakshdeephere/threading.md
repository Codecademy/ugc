---
Title: "Introduction to threading in C"
Description: "Step-by-step explain about what are threads and how threading works with suitable code samples."
DatePublished: "2023-06-04"
Categories:
  - "operating-system"
  - "python"
  - "concurrency"
  - "computer-science"
Tags:
  - "Algorithm"
  - "Node"
CatalogContent:
  - "learn-python"
  - "paths/threading"
---

[threads]: (https://raw.githubusercontent.com/Codecademy/ugc/main/content/thread.png)
![thread](/Users/dakshdeep/Documents/JAVA/Codecademy/ugc/content/dakshdeephere/thread.jpg)
[threads]

_**Prerequisites:** Understanding of C, Operating System, Visual studio code_
_**Version:** Standard C language versions C89/C90, C99, C11, and C18_

## Introduction

A thread is a path of execution within a process. A process can contain multiple threads. A thread is also known as lightweight process. The idea is to achieve parallelism by dividing a process
into multiple threads. For example, in a browser, multiple tabs can be different threads. MS Word uses multiple threads: one thread to format the text, another thread to process inputs, etc.

**`Thread`** is an execution unit which consists of its own program `counter`, a `stack`, and a set of `registers`. `Program counter` keeps track of which instruction to execute next, system registers which hold its current working variables, and a stack which contains the execution history.
![thread][thread]

## Process vs Thread

| S.N.      | Process | Thread |
| -- | ----------- | ----------- |
| 1      | Process is heavy weight or resource intensive.| Thread is light weight, taking lesser resources than a process.      |
| 2   | Process switching needs interaction with operating system.        | Thread switching does not need to interact with operating system. |
| 3 | In multiple processing environments, each process executes the same code but has its own memory and file resources. | All threads can share same set of open files, child processes. |
| 4 | If one process is blocked, then no other process can execute until the first process is unblocked. | While one thread is blocked and waiting, a second thread in the same task can run. |
| 5 | Multiple processes without using threads use more resources. | Multiple threaded processes use fewer resources. |
| 6 | In multiple processes each process operates independently of the others. | One thread can read, write or change another thread's data. |
