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

[thread]: (https://raw.githubusercontent.com/Codecademy/ugc/main/content/thread.png)

_**Prerequisites:** Understanding of C, Operating System, Visual studio code_
_**Version:** Standard C language versions C89/C90, C99, C11, and C18_

## Introduction

A thread is a path of execution within a process. A process can contain multiple threads. A thread is also known as lightweight process. The idea is to achieve parallelism by dividing a process
into multiple threads. For example, in a browser, multiple tabs can be different threads. MS Word uses multiple threads: one thread to format the text, another thread to process inputs, etc.

**`Thread`** is an execution unit which consists of its own program `counter`, a `stack`, and a set of `registers`. `Program counter` keeps track of which instruction to execute next, system registers which hold its current working variables, and a stack which contains the execution history.
![thread][thread]

