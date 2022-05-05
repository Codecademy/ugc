---
Title: "Creating Queues Using JavaScript"
Description: "Creating a queue and implementing several of its operations using JavaScript."
DatePublished: "2022-04-28"
Categories:
  - "computer-science"
Tags:
  - "Algorithms"
  - "Arrays"
  - "ES6"
  - "Classes"
CatalogContent:
  - "introduction-to-javascript"
  - "linear-data-structures"
---

_**Prerequisites:** JavaScript_  
_**Versions:** ECMAScript 2015 (ES6)_

## Introduction

During conversations about data structures in the programming world, queues are heard of quite often.
A queue is a linear data structure that follows the first-in-first-out (FIFO) pattern; i.e., removal takes place at the front, and addition takes place at the end. Think of it as the checkout point of a grocery store. Customers who get to the checkout counter first get attended to first. They can only join the queue from the back of the line, and after they get attended to, they leave from the front of the line.

## Linear Data Structures

Linear data structures are types of data structures in which elements are stored sequentially. The elements are organized in such a way that each element is directly linked to its previous and next element, except for the first element that is only linked to its previous element and the last element that is only linked to the next element.

Queues follow this arrangement pattern, and we can see that by referring to our grocery store checkout counter analogy. Each customer that walks to the checkout counter represents an element, the first customer has a link to the previous customer, the last customer has a link to the next customer and every other customer in-between has a link to a previous and next customer. 

## Use cases for queues

There are several use cases for queues in the programming and real world. Some of them include:

- A desk printer: When you send documents to the printer they are printed in the same order in which they are sent. This is very useful when you're trying to print documents that follow a specific order. The printer makes sure the pages don't get mixed up.
- Adding songs to a queue in your music player: Sometimes, while we are working on a task, we have our go-to songs to keep the momentum going. The queue function of your music player makes sure you listen to your song selections just the way it is arranged.
- Customer service wait-lines: if you go to a bank to file a complaint, they make you wait in line and the next person in line won't be called unless the issue of the current customers has been resolved.
- File sharing/data transfer between two processes: When transferring files from one device to another, the files are received on the other end in the same order they were sent, irrespective of the time it takes for any of them to get completed.

Queues can be implemented in any programming language but our focus is on how to create them using JavaScript. In JavaScript, it can also be implemented in two ways: arrays and linked lists. For this article, we will implement queues using arrays.

## Operations in queues

These are some basic operations that are performed on queues:

- Enqueue
- Dequeue
- Peek
- Reversing a queue

For the implementation of the above operations, we’d use an [ES6 class](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Classes) and have the various operations as methods.

## Creating a queue with its operations 

The first step is to initialize our class with its own storage (an array where our queue elements would be stored):

```js
class Queue {
  constructor() {
    this.queue = [];
  }
}
```

The code block above represents the creation of a class, the [`constructor()`](https://www.codecademy.com/resources/docs/javascript/constructors) method is a special method of a class that is used to create an object instance of that class. It can be used to initialize an object that would be used across the class methods. The [`this`](https://www.codecademy.com/resources/docs/javascript/this) keyword serves as a regular object within this context. It can be used to initialize any value.

Next up, we add each queue operation as a method of the class.

### Enqueue

This term refers to adding a new element to the queue. Since we're implementing our queue with an array, we can use the `.push()` array method to add new elements to the queue.

```js
class Queue {
  constructor() {
    this.queue = [];
  }

  enqueue(element) {
    this.queue.push(element);
  }
}
```

The added lines to our code block represent a method of class `Queue`. It handles one operation, adding a new item to the array object initialized in our constructor.
The method accepts one argument `element` and then adds it to `this.queue` using the `.push()` array method.

### Dequeue

This term refers to removing a new element from the queue. Again, the `.shift()` array method takes care of this for us easily.

```js
class Queue {
  constructor() {
    this.queue = [];
  }

  enqueue(element) {
    this.queue.push(element);
  }

  dequeue() {
    return this.queue.shift();
  }
}
```

We added a new method `dequeue()`. This method doesn't accept any arguments, unlike the previous method. It returns an element from the front of the queue and also removes it using the `.shift()` array method.

### Peek

We use the peek method to check for the element at the front of the queue, without removing it:

```js
class Queue {
  constructor() {
    this.queue = [];
  }

  enqueue(element) {
    this.queue.push(element);
  }

  dequeue() {
    return this.queue.shift();
  }

  peek() {
    return this.queue[0];
  }
}
```

The `peek()` method checks for the value at the front of the queue by accessing the first index of the queue.

### Reversing a Queue

As the title implies, we are simply trying to change the order of the queue from back to front. The `reverse()` method is handled using a [while loop](https://www.codecademy.com/resources/docs/javascript/loops).

```js
class Queue {
  constructor() {
    this.queue = [];
  }

  enqueue(element) {
    this.queue.push(element);
    return this.queue;
  }

  dequeue() {
    return this.queue.shift();
  }

  peek() {
    return this.queue[0];
  }

  reverse() {
    // declare an empty array
    const reversed = []

    // iterate through the array using a while loop
    while (this.queue.length > 0) {
      // remove the last element and push it to the new array
      reversed.push(this.queue.pop());
    }
    // set queue using the new array
    this.queue = reversed;
    return this.queue;
  }
}
```

## Usage

To see how our code works, we are going to test out each method of our class.
First, we need to create a new instance of our class, then use that instance to access our methods.

```js
const result = new Queue // creating a new instance of our class

// adding elements
result.enqueue(5) // returns [5]
result.enqueue(3) // returns [5, 3]
result.enqueue(4) // returns [5, 3, 4]
result.enqueue(7) // returns [5, 3, 4, 7]

// removing an element
result.dequeue() // returns [3, 4, 7]

// checking the first element in the queue
result.peek() // returns 3

// reversing the queue
result.reverse() // returns [7, 4, 3]
```

## Conclusion

Summary:
- We explained the concept of queues (FIFO) and how they are applied in the real world.
- We created a class and performed various queue operations (enqueue, dequeue, peek, reverse).

That’s it, folks. We’ve successfully implemented a queue and its basic operations using JavaScript.