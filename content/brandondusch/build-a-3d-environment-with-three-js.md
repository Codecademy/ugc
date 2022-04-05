---
Title: "Build a 3D Environment with Three.js"
Description: "Step-by-step tutorial about how to build a 3D environment with Three.js and render/move 3D objects."
DatePublished: "2022-02-28"
Categories:
  - "web-development"
  - "game-development"
  - "javascript"
  - "html-css"
Tags:
  - "Three.js"
  - "Animation"
  - "Node"
CatalogContent:
  - "learn-node-js"
  - "learn-a-frame"
---

[three.js]: https://threejs.org
[webgl]: https://get.webgl.org/
[github homepage]: https://github.com
[frustum]: https://en.wikipedia.org/wiki/frustum
[field of view]: https://en.wikipedia.org/wiki/Field_of_view
[width:height ratio]: https://en.wikipedia.org/wiki/Aspect_ratio_(image)
[range of viewable space]: https://en.wikipedia.org/wiki/Viewing_frustum
[terminal]: https://www.codecademy.com/resources/docs/general/terminal
[installing three.js]: https://threejs.org/docs/index.html#manual/en/introduction/Installation
[intall it with npm]: https://www.codecademy.com/resources/docs/javascript/npm
[cdn link]: https://www.codecademy.com/resources/docs/general/cdn
[via cdn]: https://cdnjs.com/libraries/three.js/r128
[perspective projection]: https://en.wikipedia.org/wiki/Perspective_(graphical)
[official website]: https://threejs.org/
[three.js documentation]: https://threejs.org/docs/index.html#manual/en/introduction/Creating-a-scene
[learn a-frame]: https://www.codecademy.com/learn/learn-a-frame
[source code for this article]: https://github.com/Codecademy/articles/tree/main/build-3d-environment-with-three-js
[github gif of rotating three.js globe]: https://raw.githubusercontent.com/Codecademy/articles/main/build-3d-environment-with-three-js/github.gif?token=ABGEO3RHPAZJER45E463TLDBCK6KA
[completed 3d environment with rotating cube]: https://media.giphy.com/media/STFCkLhIXzEEEqMXko/giphy.gif?cid=790b761181498784867ee636f3df2887406daf1b2227a08a&rid=giphy.gif&ct=g
[rendered page with full-sized body]: https://raw.githubusercontent.com/Codecademy/articles/main/build-3d-environment-with-three-js/rendered_page_full_body.png?token=ABGEO3RHPAZJER45E463TLDBCK6KA
[rendered page with cube]: https://raw.githubusercontent.com/Codecademy/articles/main/build-3d-environment-with-three-js/rendered_page_cube.png?token=ABGEO3TNG23SPGEGLFBXJKTBCK6LK

## Introduction

[Three.js] is a JavaScript library that features 3D objects and views rendered on a web page. It builds on top of [WebGL] by adding functionality for visual aesthetics including:

- Effects such as lights and shadows.
- Materials for building shapes.
- Textures for adding details to those shapes.

Since its release in 2010, Three.js has been used by many developers and companies alike. Below is the [GitHub homepage], which uses Three.js to render a globe. It rotates and emits interesting connections between different points. We can even change it's rotation with our mouse!

![GitHub GIF of rotating Three.js globe]

In this article, we are going to learn how to build a 3D environment with Three.js. Inside of this environment, we will render a cube that rotates at a modest speed. Below is what our completed environment will look like:

![Completed 3D environment with rotating cube]

### The `renderer`, `Scene`, and `Camera`

It should be noted that the Three.js API uses a considerable amount of stage and camera projection terms to name classes, functions and parameters.

The `renderer` object is the root of a Three.js program and carries two parameters:

- A `Scene` object that contains 3D objects, lights, cameras, etc.
- A `Camera`, which is an object-based abstraction of a camera view that exists both inside and outside of scenes.

Scenes and their child elements make up the _scenegraph_, a tree-like representation of the parent/child-relationship between objects in the scene. Scenegraphs can also contain zero or more cameras.

Cameras use methods that utilize parameters named after terms in camera projection. The following terms define the observable "shape" (or [frustum]) of the camera:

- `fov` stands for [field of view], which is the range of the observable world from a (camera’s) perspective at a given moment in time, measured in degrees.
- `aspect` describes the [width:height ratio](<https://en.wikipedia.org/wiki/Aspect_ratio_(image)>) of the to-be-rendered `<canvas>` element.
- `near` and `far` define the [range of viewable space](https://en.wikipedia.org/wiki/Viewing_frustum) between the camera lens and the drawn objects.

With some terms out of the way, let’s begin building a 3D environment.

## Step 1: setting up and installing Three.js

We’ll begin by opening the [terminal], creating a directory called **/helloCube**, and changing into it.

```bash
$ mkdir helloCube
$ cd helloCube
```

Next, we will create the following files:

```bash
$ touch helloCube.html
$ touch helloCube.js
```

Then, let’s install Three.js. There are two primary options for [installing Three.js](https://threejs.org/docs/index.html#manual/en/introduction/Installation) in a project:

- We could [install it with npm] and import as a Node module with `import` or `require()`.
- We could use `<script>` elements to import the package source code via a [CDN link].

For this article, we are going to use a CDN link to install Three.js. Let’s head to the next step to add markup.

## Step 2: adding the HTML and connecting to our JS

Let’s begin this step by opening **helloCube.html** and add the following markup:

```html
<!DOCTYPE html>
<html lang="en">
  <!-- HTML head -->
  <head>
    <title>HelloCube Three.js!</title>
  </head>
  <!-- HTML body -->
  <body></body>
</html>
```

Our 3D environment will eventually render a `<canvas>` element to the body. By default, most HTML elements are displayed at the "block"-level where they begin on a new line and take up as much width as possible. We should add some styles to the page so that the body is as tall as the screen being used and the canvas is horizontally and vertically centered. Next, we’ll add a pair of `<style>` tags to the `<head>` element and do the following:

1. Set the `margin` and `padding` of `html` and `body` to 0.
2. Set the `min-height` of the `body` to `100vh` to ensure the canvas is centered.
3. Set the `display` of the `body` to `flex`.
4. Set `align-items` and `justify-content` to `center`.

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <style>
      html,
      body {
        margin: 0;
        padding: 0;
      }

      body {
        min-height: 100vh;
        display: flex;
        align-items: center;
        justify-content: center;
      }
    </style>
    <title>HelloCube Three.js!</title>
  </head>

  ...
</html>
```

Let's go ahead and save the file. Next, we will render the page on a browser with the following command:

```bash
$ open helloCube.html
```

When our page loads, it should look like this:

![Rendered page with full-sized body]

Lastly, in preparation for the next step, we’re going to connect our HTML to the Three.js library as well as our own **helloCube.js** file. Let’s add two pairs of `<script>` tags inside the `<head>` element. One of the scripts will import Three.js [via CDN](https://cdnjs.com/libraries/three.js/r128) with the following url: `https://cdnjs.com/libraries/three.js/r128`. In order to ensure that our Three.js connection can access elements in the `<body>` _after_ the DOM loads, lets add a `defer` attribute to both `<script>` tags.

The other will link the **helloCube.html** file with the **helloCube.js** file:

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <style>
      html,
      body {
        margin: 0;
        padding: 0;
      }

      body {
        min-height: 100vh;
        display: flex;
        align-items: center;
        justify-content: center;
      }
    </style>
    <script src="https://cdnjs.com/libraries/three.js/r128" defer></script>
    <script src="helloCube.js" defer></script>
    <title>HelloCube Three.js!</title>
  </head>

  <body></body>
</html>
```

Our markup is now connected with Three.js and with our local **helloCube.js** file.

Let’s save and close the **helloCube.html** file. Next, we’ll proceed to build the actual 3D Three.js environment!

## Step 3: establishing the 3D environment

For the remaining steps, we will finish building our 3D environment in the **helloCube.js** file. Let’s open the file and define a function called `create3DEnvironment()`. Then, we will execute it directly afterwards.

```js
const create3DEnvironment = () => {};

create3DEnvironment();
```

When the **helloCube.html** loads at the `<script>` tag linking to the **helloCube.js** file, the `create3DEnvironment()` method will be invoked. To access functions and class constructors from Three.js, we’re going to use the `THREE` constant. We are going to create our root object &mdash; the `renderer` &mdash; with the `THREE.WebGLRenderer()` function:

```js
const create3DEnvironment = () => {
  const renderer = new THREE.WebGLRenderer();
};

create3DEnvironment();
```

By passing nothing into `THREE.WebGLRenderer()`, the `renderer` will create a new `<canvas>` element when in use. This is effectively our 3D environment! Although our 3D environment is technically now created, we still would like a `<canvas>` in **helloCube.html** to communicate with the `renderer`. There is more to the 3D environment than building the renderer.

Let’s move on to the next step where we will set up a `Camera` object.

## Step 4: setting up a new camera

In order to "see" the objects we render in our environment, we need to create a `Camera` object. We will be using a [perspective projection](<https://en.wikipedia.org/wiki/Perspective_(graphical)>) by creating a `new THREE.PerspectiveCamera()` and passing in the following values:

- A `fieldOfView` (integer) to set the range for what the camera can observe in the environment.
- An `aspect` ratio (integer) of the camera's height and width.
- A space range that describes how `near` and `far` objects can be viewed by the camera (both integers).

```js
  ...
  const renderer = new THREE.WebGLRenderer();

  const fieldOfView = 75;
  // Measured in degrees, not radians

  const aspect = 2;
  // The canvas default (300px-wide: 150px tall --> 2:1 -->  2)

  const near = 0.1;

  const far = 5;

  const camera = new THREE.PerspectiveCamera(fieldOfView, aspect, near, far);
};

create3DEnvironment();
```

We’ve got a new `camera` going! Let’s move on to the next step to build a scene for our camera to see _into_.

## Step 5: creating a scene with 3D objects

The canvas inside the `renderer` constant is where a `Scene` object is set. Scenes are areas where 3D objects and effects, such as light effects, are stored. These objects are composed of geometric and material properties meshed together into one cohesive, 3-dimensional "shape", like a cube or a sphere.

Let’s first create our `scene` with the `THREE.Scene()` method:

```js
  …

  const camera = new THREE.PerspectiveCamera(fieldOfView, aspect, near, far);

  const scene = new THREE.Scene();
};

create3DEnvironment();
```

We are now ready to add some 3D objects to our scene! The core pieces of a 3D object are:

- A _geometry_ that defines the size and dimensions of the object.
- A _material_ that defines the overall appearance of the object.

Let’s begin building the geometry for our cube.

The first thing to do is define a `new THREE.BoxGeometry()` by passing in a `width`, `height`, and `depth`. We’ll store it in a `geometry` variable.

```js
  …

  const scene = new THREE.Scene();

  const width = 1;
  const height = 1;
  const depth = 1;
  const geometry = new THREE.BoxGeometry(width, height, depth);
};

create3DEnvironment();
```

Next, we’ll need to create a `material` for the cube’s appearance. We can create it by using the `THREE.MeshBasicMaterial()` method and passing in an object with a `color` property and a value.

```js
  const geometry = new THREE.BoxGeometry(width, height, depth);

  const material = new THREE.MeshBasicMaterial({ color: 0xc2c5cc });
};

create3DEnvironment();
```

Lastly, we will create the actual `cube` object. Let’s combine the `geometry` with the `material` by using the `THREE.Mesh()` method. Then, we’ll use `.add()` to add the `cube` to the `scene`:

```js
  const material = new THREE.MeshBasicMaterial({ color: 0xc2c5cc });

  const cube = new THREE.Mesh(geometry, material);
  scene.add(cube);
};

create3DEnvironment();
```

Let’s go ahead and save the **helloCube.js** file. It’s now time to see what our rendered cube finally looks like in the next step!

## Step 6: rendering the scene and camera

Let’s take a step back and look at what we’ve done so far with our 3D environment:

- We created a `scene` that contains 3D objects and effects.
- We created a `camera` that "views" the 3D objects and effects.
- We created a `renderer` to facilitate a renderable `<canvas>`.

Next, we will use the `renderer`’s `.render()` method to create a 3D environment. We will pass in the `scene` and `camera` objects we built in the previous steps. They will be rendered through the returned `<canvas>` element.

```js
  const cube = new THREE.Mesh(geometry, material);
  scene.add(cube);

  renderer.render(scene, camera)
};

create3DEnvironment();
```

We can then append the `domElement` of the `renderer` to the DOM:

```js
  renderer.render(scene, camera)
  document.body.appendChild(renderer.domElement);
};

create3DEnvironment();
```

And our rendered page should look like this:

![Rendered page with cube]

Wait?! That looks more like a square than a cube! Let’s find out for sure by trying to move it in the final step.

## Step 7: animating the cube

In this last step, we are going to write an `.animate()` method that will move the cube in the 3D environment that just build in the previous step:

```js
  …

  const animate = (time, speed=1) => {
    time *= 0.001; // converted to seconds

    const rotation = time * speed;
    cube.rotation.x = rotation;
    cube.rotation.y = rotation;

    renderer.render(scene, camera)
    document.body.appendChild(renderer.domElement);
    requestAnimationFrame(animate)
  };

  requestAnimationFrame(animate)
};

create3DEnvironment();
```

Some of the code shown above is new and some is refactored. Here is a breakdown of what we just did:

- First, we defined an `.animate()` method that accepts two integers, a `time` and a `speed`, which is defined with 1.
- Next, we converted the `time` into seconds.
- Then, we defined a `rotation` by multiplying the `time` by the `speed`.
- We then assigned the `rotation` to the `cube.rotation.x` and `cube.rotation.y` coordinates to "rotate" the cube.
- Next, we rendered the `scene` and `camera`
- Then we appended the `renderer.domElement` to the DOM.
- Right after that, we used Three.js’s `requestAnimationFrame()` method to make a recursive call to our `.animate()` method to keep the cube constantly rotating.
- `requestAnimationFrame()` is also used outside of `.animate()` meant to start the rotation.

Our rendered page should look something like this:

![Completed 3D environment with rotating cube]

## Conclusion

There we have it! We just learned how to use the Three.js library to build a 3D environment. More specifically:

- We built a `renderer` that created a `scene` that contains 3D objects.
- We built a `camera` that "views" the objects in the `scene`.
- We created an `.animate()` method that renders the `scene` and `camera` in the returned `<canvas>` element.

This is only scratching the surface of what Three.js can do. Visit their [official website] to view other projects and learn more!

### Resources

- [Source code for this article]
- [Three.js documentation]
- [Learn A-Frame]
