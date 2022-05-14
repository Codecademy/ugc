---
Title: "An Introduction to Abstract Data Types in JavaScript"
Description: "A simple introduction to ADTs in JavaScript."
DatePublished: "2022-04-25"
Categories:
  - "code-foundations"
  - "computer-science"
Tags:
  - "Data Structures"
  - "Classes"
  - "OOP"
CatalogContent:
  - "introduction-to-javascript"
  - "paths/computer-science"
---

![Introduction to Abstract Data Type](https://raw.githubusercontent.com/Codecademy/ugc/main/content/goku-kun/adt.png)

## Introduction

An **Abstract Data Type** (ADT), as the name suggests, is an abstract understanding of a data structure. An ADT is defined through its behavior and characteristics, particularly in terms of what data can be stored into it, the operations that can be performed on this data, and the behavior of these operations. For example, stacks and queues can be internally implemented using linked-lists made up of nodes or [arrays](https://www.codecademy.com/resources/docs/javascript/arrays). However, the primary function of a stack is to be a last in, first out (LIFO) data structure and the primary function of a queue is to be a first in, first out (FIFO) data structure. The behavior, from the point of the user, remains intact, regardless of the internal implementation either using linked-lists or arrays. If the user was interacting with a stack, the user will simply worry about pushing data onto the stack or popping data off the stack. The user won't need to have the knowledge of how that stack is working internally.

In contrast to the data structures, which are specific and detailed implementations that deal with how the data structure does its job, an ADT focuses on _what it does_ and not how it does its job. In short, the ADT defines what that particular data construct must do and the data structure is the concrete implementation of that construct.

An analogy to explain ADTs in terms of web development would be [CRUD](https://www.codecademy.com/resources/docs/general/http) (abbreviated as create, read, update and delete) APIs. The user of any CRUD API has to simply know what request method (GET, POST, PUT/PATCH, or DELETE) should they send, and if they followed the rules of the API, the API server would send data back. The user didn't have to worry about the internal workings of the API server. They simply had to know the rules of interactions and behavior of a CRUD API. In this case, the CRUD API is functioning as an ADT _from the perspective of the user_.

There are no specific rules which force the implementation of particular methods and operations in a particular ADT. This is decided based on the requirements in a use-case scenario and ultimately by design choice.

## Why use ADTs?

There are 3 general advantages of using ADTs, listed as follows:

### Encapsulation

An ADT will provide certain methods and properties. And the knowledge of these methods and properties is all the user will need to successfully operate with the ADT.

### Compartmentalization

The code that is using the ADT will not have to be changed even if the internal workings of the ADT have been changed. The change in the ADT is isolated and compartmentalized.

### Adaptability

Real world programs continue to evolve with ever changing requirements and new constraints. Differently implemented ADTs, with all the same properties and methods, can thus be used interchangeably. For example, consider a linked list created using arrays that contained names of the patients in a hospital. It's later decided to include all the information about the patient in the linked list. Then, a linked list implemented using class based nodes with all the necessary fields would serve as a much better replacement as compared to the linked list that is simply using arrays. Therefore, ADTs can adapt to the situation in which they're used.

## General operations supported by ADTs

ADTs support the follow operations:

- **Traversing**, which allows each element in the ADT to be accessed once for processing.
- **Searching**, which allows the user to look for a specific element in the ADT.
- **Inserting**, which allows the user to insert an element at a particular index/space in the ADT.
- **Deleting**, which allows the user to either delete a particular element or delete an element at a particular location.
- **Sorting**, which allows the elements to be ordered in ascending or descending order, depending on the preference.

## How ADTs coexist with other programs

Each ADT supports specific operations which can be leveraged in a particular situation. Some ADTs can provide better speeds at looking up data while others can save space. But, these ADTs work in conjunction with other programs to track, store, retrieve, and manipulate the data. The user is the one that decides which particular ADT would best serve their requirements.

![Abstract Data Type Usage](https://raw.githubusercontent.com/Codecademy/ugc/main/content/goku-kun/abstract-data-type-usage.png)

## Linked-lists as an ADT

Linked-lists are made up of a sequence of elements (called nodes, refer the left note below), which may or may not be stored sequentially in memory. For a simple linked-list, each node has the ability to store some data and the link to the next node in the linked-list. Every linked-list begins at the head node which is then linked to the next node.

![node and linked list](https://raw.githubusercontent.com/Codecademy/ugc/main/content/goku-kun/node-linked-list.png)

Linked-list ADTs would support the following operations:

```js
// General operations
getHead(); // Returns back head node
getSize(); // Return the current size of the linked-list
isEmpty(); // Returns true if the linked-list is empty

// Insert and replace operations
insertBeginning(element); // Inserts new element at the beginning of the linked-list
insertEnd(element); // Inserts new element at the end of the linked-list
insertAtPosition(element, index); // Inserts new element at the given positional index
replaceAtPosition(element, index); // Replaces the element at give index with the new element

// Delete operations
deleteBeginning(); // Removes the first element and returns its element
deleteEnd(); // Removes the last element and returns its element
deleteAtPosition(index); // Removes node from the given positional index and returns its element

// Traverse, sort, and search operations
traverse(); // Goes through all the elements once in the linked list and prints them
search(element); // Searches given element and returns true if the element is found in linked-list
sort(order); // Sorts the linked-list in the given order (ascending/descending)
retrieve(index); // Returns the element stored at the given index location
```

## Stack ADTs

**Stacks** are linear data structures in which data is entered and removed from _only a single point_. This point is called the **top** of the stack. It follows the last-in, first-out (L I F O) format for storing and discarding data. This means that the last element added to the top of the stack is the first element that will be removed from the stack. There is no other way to access other elements in the stack but the element that is at the top of the stack.

![Stack A D T](https://raw.githubusercontent.com/Codecademy/ugc/main/content/goku-kun/stack-adt.png)

A stack ADT supports the following operations:

```js
push(element); // Inserts a new element at the top of the stack
pop(); // Removes the element stored at the top of the stack and returns it
peek(); // Returns the top element without removing it from the stack
```

## Queue ADTs

**Queues** are linear data structures, in which data is inserted from one end and removed from the other end. The place where the data is inserted in the queue is called the **rear** end of the queue and this insertion operation is called **enqueue**. Data can be removed from the **front** of the queue and this deletion operation is called **dequeue**. It follows the first-in, first-out (F I F O) configuration for storing and removing data. This means that the data that was first to enter the queue is also first to leave the queue.

![Queue ADT](https://raw.githubusercontent.com/Codecademy/ugc/main/content/goku-kun/queue-adt.png)

A Queue ADT supports the following operations:

```js
enqueue(element); // Inserts a new element at the rear of the queue
dequeue(); // Removes the element at the front of the queue and returns it
```

## Conclusion

ADTs fulfill a very important role in everyday programming tasks. This article talked about how ADTs are used in application development. It also showed how ADTs offer encapsulation, compartmentalization, and adaptability. Finally, the article discussed about some basic ADTs such as linked-lists, stacks and queues.
