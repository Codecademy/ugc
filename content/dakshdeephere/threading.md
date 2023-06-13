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

_**Prerequisites:** Understanding of C, Operating System, Visual studio code_  
_**Version:** Standard C language versions C89/C90, C99, C11, and C18_

[oneToOne]:https://raw.githubusercontent.com/Codecademy/ugc/main/content/dakshdeephere/one-to-one.png

![oneToOne]

## Introduction

A thread is a path of execution within a process. A process can contain multiple threads. A thread is also known as lightweight process. The idea is to achieve parallelism by dividing a process
into multiple threads. For example, in a browser, multiple tabs can be different threads. MS Word uses multiple threads: one thread to format the text, another thread to process inputs, etc.

**`Thread`** is an execution unit which consists of its own program `counter`, a `stack`, and a set of `registers`. `Program counter` keeps track of which instruction to execute next, system registers which hold its current working variables, and a stack which contains the execution history.

<div>
  <div style="display: inline-block; width: 45%;">
    <img src="single-thread-process.png" alt="Image 1" style="width: 100%;">
    <p style="font-weight: bold;">Single Thread</p>
  </div>
  <div style="display: inline-block; width: 45%;">
    <img src="Multi-threaded-process.png" alt="Image 2" style="width: 100%;">
    <p style="font-weight: bold;">Multi Thread</p>
  </div>
</div>

## Multithreading

- A thread is a path which is followed during a program’s execution.
- Let's consider an instance where a program cannot simultaneously read keystrokes and create drawings. These activities pose a challenge for the program as they cannot be performed concurrently. However, this issue can be resolved through multitasking, enabling the program to execute multiple tasks simultaneously.
- Multitasking is of two types: `Processor based` and `thread based`.
- `Processor based` multitasking is totally managed by the OS, however multitasking through multithreading can be controlled by the programmer to some extent.
- The concept of multi-threading needs proper understanding of these two terms – a `process` and a `thread`.
- A process is a program being executed. A process can be further divided into independent units known as threads.
- A `thread` is like a small light-weight process within a process Or we can say a collection of threads is what is known as a process.

## Process vs Thread

| S.N.      | Process | Thread |
| -- | ----------- | ----------- |
| 1      | Process is heavy weight or resource intensive.| Thread is light weight, taking lesser resources than a process.      |
| 2   | Process switching needs interaction with operating system.        | Thread switching does not need to interact with operating system. |
| 3 | In multiple processing environments, each process executes the same code but has its own memory and file resources. | All threads can share same set of open files, child processes. |
| 4 | If one process is blocked, then no other process can execute until the first process is unblocked. | While one thread is blocked and waiting, a second thread in the same task can run. |
| 5 | Multiple processes without using threads use more resources. | Multiple threaded processes use fewer resources. |
| 6 | In multiple processes each process operates independently of the others. | One thread can read, write or change another thread's data. |

## Benefits of creating threads in Operating System

- `Responsiveness` – multi-threading increase the responsiveness of the process. For example, in MSWord while one thread does the spelling check the other thread allows you to keep tying the input. Therefore, you feel that Word is always responding.
- `Resource sharing` – All the threads share the code and data of the process. Therefore, this allows several threads to exist within the same address space
- `Economy` – For the same reason as mentioned above it is convenient to create threads. Since they share resources they are less costly
- `Scalability` – Having a multiprocessor system greatly increases the benefits of multithreading. As a result, each thread can run in a separate processor in parallel.

## Challenges for Programmers while creating Threads

- `Dividing activities` – It involves finding the functions within the job that can be run in parallel on separate processors.
- `Balance` – The task assigned to each processor must also be equal. Now there can be different parameters for that. One parameter can be, assign equal tasks to each processor. But, tasks assigned to more processor may require higher execution time thus overloading one processor. Thus, simply assigning equal tasks to each processor may not work.
- `Data splitting` – Another challenge is to split the data required for each task.
- `Data dependency` – sometimes the data required by one thread (T1) might be produced by another (T2). Thus, T1 can not run before T2. Therefore, it becomes difficult for programmers to code.
- `Testing and debugging` – Multiple threads running in parallel on multiple cores poses another challenge in the testing of applications.

## Example

Now see how to create and join threads with code sample in `C`:

```c
#include <stdio.h>
#include <pthread.h>

void *factorial(void *arg) {
    int num, fact = 1;
    printf("Enter a number to find its factorial: ");
    scanf("%d", &num);

    for (int i = 1; i <= num; i++) {
        fact *= i;
    }

    printf("Factorial of %d is %d\n", num, fact);
    pthread_exit(NULL);
}

void *hcf(void *arg) {
    int num1, num2, hcf;
    printf("Enter two numbers to find their HCF: ");
    scanf("%d %d", &num1, &num2);

    int temp1 = num1, temp2 = num2;

    while (num1 != num2) {
        if (num1 > num2) {
            num1 -= num2;
        } else {
            num2 -= num1;
        }
    }

    hcf = num1;
    printf("HCF of %d and %d is %d\n", temp1, temp2, hcf);
    pthread_exit(NULL);
}


int main() {
    pthread_t thread_factorial, thread_hcf;
    int choice;

    printf("Choose an option:\n");
    printf("1. Find factorial of a number\n");
    printf("2. Find HCF of two numbers\n");
    printf("Enter your choice (1 or 2): ");
    scanf("%d", &choice);

    if (choice == 1) {
        pthread_create(&thread_factorial, NULL, factorial, NULL);
        pthread_join(thread_factorial, NULL);
    } else if (choice == 2) {
        pthread_create(&thread_hcf, NULL, hcf, NULL);
        pthread_join(thread_hcf, NULL);
    } else {
        printf("Invalid choice. Exiting.\n");
        return 1;
    }

    printf("Program completed. Exiting.\n");
    return 0;
}
```

To compile the above code, you have to use following command:

```shell
chmod u+x filename.c
gcc -o multithreading multithreading.c -pthread
./multithreading 
```

The output for the above code will be:

```shell

