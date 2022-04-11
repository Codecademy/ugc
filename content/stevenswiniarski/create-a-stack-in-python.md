---
Title: "Create a Stack in Python"
Description: "A hands-on tutorial building a stack implementation in Python."
DatePublished: "2022-02-28"
Categories:
  - "computer-science"
  - "python"
Tags:
  - "Data Structures"
  - "Classes"
  - "Lists"
CatalogContent:
  - "learn-python-3"
  - "paths/computer-science"
---

[depth-first search]: https://en.wikipedia.org/wiki/Depth-first_search
[“stack”]: https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
[exception]: https://www.codecademy.com/resources/docs/python/errors
[Python class]: https://www.codecademy.com/resources/docs/python/classes
[`.pop()` method]: https://www.codecademy.com/resources/docs/python/lists/pop
[`str()` function]: https://www.codecademy.com/resources/docs/python/built-in-functions/str
[Stack]: https://raw.githubusercontent.com/Codecademy/ugc/main/content/stevenswiniarski/stack.png

> **Prerequisites:** Python  
> **Versions** Python 3.8

<kbd>Python</kbd> <kbd>Computer Science</kbd>

In computer science, a [“stack”] is a data structure represented by a collection of items that utilizes a last-in-first-out (LIFO) model for access. There are two operations that are fundamental to this data structure:

* A “push()” function that adds an item to the stack.
* A “pop()” function that removes the most recently added item to the stack. 

In this way, this type of collection is analogous to a stack of items such as dinner plates, where the topmost item must be removed to access the items underneath. Stacks are useful in implementing actions such as a [depth-first search]. This article will explore the process of implementing a stack in Python.

![Stack]

### Planning Our Stack Implementation

Before we begin, we need to decide what kind of capabilities we want in our stack implementation. The `.push()` and `.pop()` functions fulfill the minimum requirements. But we also might want the following:

* Python's `len()` function to let us know how many items are in the stack, and to warn us when the stack is empty. It’s good practice to check for an empty stack when using one in a program.
* A `.peek()` function to tell us the value of the top item on the stack without removing it.

Lastly, we want to decide how the `.peek()` or `.pop()` methods behave when they are called on an empty stack. We could return something like `NaN`, but that might lead to subtle errors down the line, especially if a `NaN` value is added to the stack. A better practice in this scenario is to raise an [exception] when we try to use these functions on an empty stack. That way, such an event can get caught during testing and the code using the stack can be updated appropriately.

### Starting To Build the Stack Class

Our stack is going to be a [Python class]. Once we declare our class, the first thing we want to add is a container to hold the items in our stack. To do this we create an internal variable:

```py
class stack:
  def __init__(self):
    self.__index = []
```

Upon initialization of our `stack` class, it will initialize the `__index` variable as an empty list. This list will hold the items in our stack.

### Setting Up the `len()` Function

We’ll set up the `len()` function for our class first, since we’ll want to check it before using our `.pop()` and `.peek()` methods. We’ll do this by implementing a  “magic” method, also called a Dunder (double-underscore) method. Dunder methods allow us to override the behavior of built-in Python operations. For our stack we can leverage the `len()` Dunder method to institute the “length” behavior we need:

```py
class stack:
  def __init__(self):
    self.__index = []

  def __len__(self):
    return len(self.__index)
```

Now, when we call `len(stack_instance)`, it will return the number of items in our `__index` variable.

### Setting Up the `.push()` Method

Next, we want to set up our `.push()` method that will place items in our `__index` variable. Since `__index` is a list, our main decision will be at which “end” of the list we should insert our items.

The first impulse might be to append items to our `__index` list, since we usually think of the highest-indexed item to be the “top”. However, this approach can be problematic for our purposes. This is because our reference, the “top” index, will always be changing as we perform operations on our stack. Additionally, this value would need to be recalculated every time we referenced it.

It is more efficient to add and remove items from the “beginning” of our list, since the index of the “beginning” never changes. It will always be zero. Therefore, our `__index` variable will be ordered with the “top” item as the first item of our list. Since we are working with a Python list, this can be done with the built-in `.insert()` method:

 ```py
class stack:
  def __init__(self):
    self.__index = []

  def __len__(self):
    return len(self.__index)

  def push(self,item):
    self.__index.insert(0,item)
```

### Setting Up the `.peek()` Method

The `.peek()` method is pretty straightforward. It returns the "top" value of the stack, which refers to the first item in our list, `__index[0]`. However, we need to take into account the possibility that our list is empty. We will want to check our stack with the `len()` function and throw an exception if we’re trying to use `.peek()` on an empty stack:

 ```py
class stack:
  def __init__(self):
    self.__index = []

  def __len__(self):
    return len(self.__index)

  def push(self,item):
    self.__index.insert(0,item)

  def peek(self):
    if len(self) == 0:
      raise Exception("peek() called on empty stack.")
    return self.__index[0]
```

### Setting Up the `.pop()` Method

The `.pop()` method is exactly the same as the `.peek()` method with the further step of removing the returned item from the stack. Like `.peek()`, we’ll want to check for an empty list before trying to return a value:

 ```py
class stack:
  def __init__(self):
    self.__index = []

  def __len__(self):
    return len(self.__index)

  def push(self,item):
    self.__index.insert(0,item)

  def peek(self):
    if len(self) == 0:
      raise Exception("peek() called on empty stack.")
    return self.__index[0]

  def pop(self):
    if len(self) == 0:
      raise Exception("pop() called on empty stack.")
    return self.__index.pop(0)
```

It's important to note that a Python list has its own [`.pop()` method], which behaves almost the same as our stack `.pop()` method, except that the list-version can take an index and “pop” an item from anywhere in the list.

### Setting Up the `str()` Function

An additional thing we can do is tell Python how we want our stack printed with the [`str()` function]. At the moment, using it yields the following results:

```py
>>> s = stack()
>>> print(str(s))
'<__main__.stack object at 0x000002296C8ED160>'
```

In order to understand the contents of our stack we’ll want something a little more useful. This is where the `__str__()` Dunder method comes in handy:

 ```py
class stack:
  def __init__(self):
    self.__index = []

  def __len__(self):
    return len(self.__index)

  def push(self,item):
    self.__index.insert(0,item)

  def peek(self):
    if len(self) == 0:
      raise Exception("peek() called on empty stack.")
    return self.__index[0]

  def pop(self):
    if len(self) == 0:
      raise Exception("pop() called on empty stack.")
    return self.__index.pop(0)

  def __str__(self):
    return str(self.__index)
```

This will return the contents of our stack, just like printing out the items of a generic list.

### Using the `stack` Class

We now have a usable `stack` class. The code below highlights all of the functionality we’ve implemented in our custom class:

```py
>>> s = stack()
>>> s.peek()			# stack = []
Exception: peek() called on empty stack.
>>> len(s)	
0
>>> s.push(5)		# stack = [5]
>>> s.peek()
5
>>> s.push('Apple')	# stack = ['Apple',5]
>>> s.push({'A':'B'})	# stack = [{'A':'B'},'Apple',5]
>>> s.push(25)		# stack = [25,{'A':'B'},'Apple',5]
>>> len(s)
4
>>> str(s)
"[25, {'A': 'B'}, 'Apple', 5]"
>>> s.pop()			# stack = [{'A':'B'},'Apple',5]
25
>>> s.pop()			# stack = ['Apple',5]
{'A': 'B'}
>>> str(s)
"['Apple', 5]"
>>> len(s)
2
```

### Conclusion

We’ve now learned how to implement the core functions of a stack class in Python. We could definitely add more functions to this implementation if we wanted. Some examples may include:

* Refactoring `peek()` to look at any item in the stack by index.
* Support for appending the contents of lists as a series of items within our stack.
* Adding a `clear()` method to empty the stack.
* Defining an upper limit to the stack size, which could be valuable in production use to prevent runaway operations from repeatedly adding items to the stack and causing an “Out of Memory” exception. 

With this as a basis, we are well on our way to developing our own stack implementation.
