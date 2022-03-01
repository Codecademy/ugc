# Content Standards

## Components of an Article

All UGC entries are Markdown files and should consist of three parts:

- The **file name**, with the **.md** extension.
- The **metadata** about the content, written in YAML, which appears at the top of the entry.
- The **content**, written in Markdown.

We'll describe the standards for each of these components separately.

### File Name

All entry file names use the **.md** format along with the following standards:

- They must match the title of the entry. (e.g. "Three.js" -> **three-js.md**)
- It must be separated by dashes, where appropriate. (e.g. "Data Types" -> **data-types.md**, "HashTable" -> **hashtable.md**)
- The name should always be lowercase. (e.g. "Array" -> **array.md**, "HTML" -> **html.md**)
- Punctuation must not be included in the file name, even if it exists in the entry name. (e.g. "Moore's Law" -> **moores-law.md**, "Three.js" -> "**three-js.md**)

### Content

- All text should be written in Markdown
  - For more details on Markdown, see Codecademy's [Curriculum Markdown Style Guide](http://curriculum-documentation.codecademy.com/content-guidelines/markdown-style-guide/).
- Each subsection should begin with a heading of size h2 (i.e. preceded by `##`)
- All in-line code should be delineated by single backticks (`)

#### Code Blocks

Code examples are an excellent teaching tool. Many term entries will contain a code example as a standard field. Encyclopedia entries are more flexible, but if your encyclopedia entry doesn't contain any code blocks, take a moment to ask yourself if one would help!

All code blocks should be delineated by triple backticks (```)

- C: ```c
- C#: ```cs
- C++: ```cpp
- CSS: ```css
- Error: ```error
- Emojicode: ```emojic
- Golang: ```go
- HTML: ```html
- Java: ```java
- JavaScript: ```js
- Markdown: ```md
- PHP: ```php
- Plaintext: ```plaintext
- Pseudo: ```pseudo
- Python: ```py
- R: ```r
- Ruby: ```rb
- Sass: ```scss
- Scheme: ```scheme
- Shell/Program Output: ```shell
- SQL: ```sql
- Swift: ```swift

### Metadata

| Variable Name    | Description                                                                                                                                                                                                                                                                                                                                                                                                                                         | Example                                                                                             |
| ---------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `Title`          | The title of the entry. This will be displayed on the page at the top.                                                                                                                                                                                                                                                                                                                                                                              | Title: "OOP"                                                                                        |
| `Description`    | A brief description (ideally under 150 characters) used in search engine results and content previews                                                                                                                                                                                                                                                                                                                                               | Description: "A brief description of OOP"                                                           |
| `Subjects`       | Languages and technical fields that the entry relates to. We're storing all of our Subjects in the [subjects.md file](https://github.com/Codecademy/docs/blob/main/documentation/subjects.md). Please only use Subjects from that list.<br /> - We aim to include an average of 1-2 subjects (and always at least 1!) with every entry.                                                                                                             | Subjects:<br /> - "Computer Science"<br /> - "Game Development"                                     |
| `Tags`           | Key words that are relevant to the entry. We're storing all of our Tags in the [tags.md file](https://github.com/Codecademy/docs/blob/main/documentation/tags.md). Please only use Tags from that list, but if that list feels insufficient, feel free to create a new Tag and add it to tags.md in your PR!<br /> - We aim to include an average of 3-4 tags with every entry.                                                                     | Tags:<br /> - "React"<br /> - "URL"<br /> - "Strings"<br />                                         |
| `CatalogContent` | Slugs of Codecademy course and Path landing pages that relate to the entry. We're storing all of our slugs in the [catalog-content.md file](https://github.com/Codecademy/docs/blob/main/documentation/catalog-content.md).<br /> - Please avoid linking to individual content items, because their URLs may change and some are gated by Pro membership.<br /> - We aim to include 2 slugs with every entry (1 free course and 1 Pro course/Path). | CatalogContent:<br /> - "introduction-to-javascript"<br /> - "paths/front-end-engineer-career-path" |

#### General Writing Tips

- Avoid referencing information that isn't strictly related to the topic of the entry. As a rule, you want to assume as little pre-existing knowledge as possible.
- Avoid using first- and second-person pronouns (e.g. I, we, you) if possible.
- Brevity without sacrificing clarity. Make every word count.
- If the concept is hard, make it easy. If it's dry, make it fun. If it's simple, keep it simple.
- 90% of writing is rewriting.

#### Images and Source Code

If you'd like to include an image, video, GIF, or other file in your entry, please upload it to the same article folder. That way we'll never have any broken links!

### I've read all the content standards. Now what?

Check out the [entry template](https://github.com/Codecademy/docs/blob/main/documentation/entry-template.md) and [term entry template](https://github.com/Codecademy/docs/blob/main/documentation/term-entry-template.md) in this folder. And take a look at [GitHub Issues](https://github.com/Codecademy/docs/issues) to see where help is needed!

For a refresher on how to make a Pull Request, head back to the [Contribution Guide](https://github.com/Codecademy/docs/blob/main/.github/CONTRIBUTING.md). 🎒
