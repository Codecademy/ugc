---
Title: 'How to Convert CSS to SCSS'
Description: 'Guide on converting CSS to SCSS syntax.'
DatePublished: "2022-02-28"
Categories:
  - 'web-development'
  - 'web-design'
  - 'html-css'
Tags:
  - 'Sass'
  - 'Style'
CatalogContent:
  - 'learn-css'
  - 'paths/front-end-engineer-career-path'
---

<img src="https://upload.wikimedia.org/wikipedia/commons/9/96/Sass_Logo_Color.svg" width="20%">
<br>

_**Prerequisites:** HTML, CSS_  
_**Versions:** SASS 1.38_

[SCSS](https://sass-lang.com) is the syntax used for the scripting language SASS, or Syntactically Awesome Style Sheet. This syntax can be used to significantly improve the readability of CSS code. It offers many advanced features that will make it easier for you to shorten your code. Since it is more advanced than CSS, it is sometimes coined as Sassy CSS. In this article, we’re going to learn more about what makes this style sheet so sassy.

## Getting Started With SASS

Depending on your preference, SASS can be installed in many different ways. There are several free applications that allow you to have SASS up and running in no time. It can also be installed directly from the command line. If you don’t have SASS already installed, then take some time to explore your options here: [SASS Install Guide](https://sass-lang.com/install).

## Variables

SCSS makes use of variables. Unlike CSS, where you have to call a `var()` function to make a variable, SCSS allows you to make variables directly. This is great for keeping track of things like fonts, colors, and sizing that you know you’re going to use over and over again.

The syntax for SASS variables is as follows:

```css
$variableName: value;
```

Let’s take this piece of code for example:

```css
body {
  color: #000000;
  font: 100% Helvetica, sans-serif;
  font-size: 50px;
  font-weight: lighter;
}
```

If we know we’re going to be using the color black and font Helvetica often in our code, then we can set them to variables using SCSS. We can do this as follows:

```css
$black: #000000;
$font-type: Helvetica, sans-serif;

body {
  color: $black;
  font: 100% $font-type;
  font-size: 50px;
  font-weight: lighter;
}
```

We initiated a variable named black and font-type using the symbol, `$`, and defined each with the desired output value. We then were able to call each variable by calling its name starting with `$`. 

When our code becomes more lengthy, it can become tedious to keep track of things. Variables are a great way to store items that we would like to have for later use. They can be a container for many things including strings, booleans, numbers, colors, and more. Storing these commonly used items in variables can make your code shorter and easier to read.

## Nesting

When defining rules in CSS, they must be defined one after another. CSS does not allow nesting. However, this can be done in SCSS.

Take this CSS code for a navigation bar as an example:

```css
nav ul {
  margin: 2;
  padding: 2;
  list-style: none;
}

nav li {
  display: inline-block;
}

nav a {
  display: block;
  padding: 12px 24px;
  text-decoration: none;
}
```

We can nest this in SCSS:

```css
nav {
  ul {
    margin: 2;
    padding: 2;
    list-style: none;
  }
  li {
    display: inline-block;
  }
  a {
    display: block;
    padding: 12px 24px;
    text-decoration: none;
  }
}
```

Each child element is nested inside the parent element of `nav`. The hierarchical structure in SCSS makes finding and changing elements much easier.

## Importing Files

SCSS has a major upgrade for importing files. In CSS, when a file is imported, an HTTP request is made each time the file is called. SCSS eliminates this by directly including the file into the CSS code. This improves the runtime and performance of your code.

The syntax for importing files is as follows:

```css
@import "filename";
```

When using SASS, there is no need to include the file extension in the file name. SASS automatically assumes you’re importing a file of **.sass** or **.scss**.

Let’s say you have a file called **default.scss** that contains the following code:

```css
html,
body,
ul,
li {
  margin: 2;
  padding: 2;
}
```

You can import this file into another file by including the import line at the very top of your file, like so:

```css
@import "default";

p {
  text-align: center;
  line-height: 1.8;
  font-size: 25px;
  display: block;
  margin: 30px 0 10px;
}
```

## Mixins

SASS includes a feature called mixins that allows you to reuse snippets of CSS code wherever you want. Once you create a mixin, you can use it by calling it. The syntax for creating a mixin is as follows:

```css
@mixin name {
  property: value;
  property: value;
  property: value;
  …
}
```

If we wanted to create a mixin to highlight a piece of text we think is important, we can make one like this:

```css
@mixin highlight-text{
  color: blue;
  font-size: 50px;
  font-weight: bold;
}
```

Then we can call it anywhere we want. Let’s try this by pretending we have a class called `.highlight` where we want to use our mixin, we can write the following:

```css
.highlight {
  @include highlight-text;
  background: yellow;
}
```

The CSS will compile like this:

```css
.highlight {
  color: blue;
  font-size: 50px;
  font-weight: bold;
  background: yellow;
}
```

The use of mixins eliminates the need to repeat yourself, resulting in cleaner code.

## Extend/Inheritance

With the built-in feature `@extend`, SASS allows you to share CSS properties to multiple selectors. For example, if we have a basic button with the following properties:

```css
.button-basic {
  border: none;
  padding: 25px 35px;
  text-align: center;
  font-size: 28px;
  cursor: pointer;
}
```

We can extend these properties to other buttons we want to create:

```css
.button-back {
  @extend button-basic;
  color: white;
  background: blue;
}

.button-next {
   @extend button-basic;
   color: red;
   background: green;
}
```

The CSS will compile like this:

```css
.button-basic, .button-back, .button-next {
  border: none;
  padding: 25px 35px;
  text-align: center;
  font-size: 28px;
  cursor: pointer;
}

.button-back {
  color: white;
  background: blue;
}

.button-next {
  color: red;
  background: green;
}
```

The `@extend` allows for selectors to inherit properties from each other which eliminates the need to write multiple classes.

## Operators

SASS allows you to make use of math operators like `/`, `*`, `%`, `+`, and `-` to make calculations. Here is an example of calculations in use:

```css
.container {
  width: 520px / 800px * 100%;
}
```

The following will compile in CSS as follows:

```css
.container {
  width: 65.0%;
}
```

Having operators at your fingertips makes it a lot easier to calculate sizing for margins, widths, and padding.

## Functions

To define complex computations that we wish to use multiple times, we can use functions to simplify the process. 

A function in SASS uses the `@function` at-rule and is declared as follows:

```pseudo
@function <name>(<arguments...>) { 
  ... 
}
```

If we needed a function that multiplies a list of numbers we can write the following:

```css
@function mult($numbers...) {
  $mult: 1;
  @each $num in $numbers {
    $mult: $mult * $num;
  }
  @return $mult;
}
```

**Note:** The use of `...` at the end of a function declaration allows for the function to take any number of arguments. Any extra arguments will be passed into the function as a list. This is known as an argument list.
  
We can then use our `mult` function:
```css
.increase {
  size: mult(2px, 4px, 5px, 8px);
}
```
Functions are very useful when creating complex calculations. Try making an `exponent` function that can calculate 3 to the power of 18. 

**Hint:** SASS has a `@for` rule that allows for executing through expressions. It's written as follows:

```css
/* Exclude last expression or value using to  */
@for <variable> from <expression> to <expression> { ... } 

/* Include last value using through */
@for <variable> from <expression> through <expression> { ... }
```

## Conclusion

SASS allows you to make sophisticated style sheets faster. It keeps your code from being repetitive by having features specifically made for code reuse. By incorporating SASS in your code, you’ll find your code to be cleaner and more readable.

- Solution code: [conversion.scss](https://github.com/Codecademy/articles/blob/main/convert-css-to-scss/conversion.scss)
- Documentation of SASS: https://sass-lang.com/documentation
