# UGC Articles: Contribution Guide 👩🏻‍💻👨🏾‍💻👩🏼‍💻

Welcome to the Codecademy UGC Articles GitHub repo!

We are an inclusive and passionate team of technologists and life-long learners around the world building free programming resources for a better tomorrow. All the content in UGC articles are written by amazing creative developers like yourself.

If you have some interesting learnings to share with the community, we'd love to have you contribute. 💖

## How do I contribute?

There are many ways to contribute to UGC articles:

- Submit a Pull Request to edit an existing article.
- Submit a Pull Request to create a new article of your choice.
- Take a look in [GitHub Issues](https://github.com/Codecademy/ugc/issues) to get inspirations for your article.
- Join the [#CodecademyUGC](https://twitter.com/search?q=%23CodecademyUGC&src=typed_query&f=live) discussion on Twitter.

If you're new to UGC articles and contributing for the first time, it is recommended that you visit the [Issues](https://github.com/Codecademy/docs/issues) section and ask to be assigned to an open issue that interests you. Otherwise, feel free to submit a [PR](https://www.codecademy.com/resources/docs/git/pull-requests) by creating a new [branch](https://www.codecademy.com/resources/docs/general/git/branch) in your fork to create a new article or edit an existing one.

## What do I need to do before creating a new article?

Before creating your first article, poke around the [/content](https://github.com/Codecademy/ugc/tree/main/content) folder. This is where all the content is stored.

```
.
├── ...
├── content                  # Content files
│   ├── author
|   |   ├── author_meta.json 
|   |   ├── article1.md      
|   |   ├── article2.md
│   └── ...
├── documentation            # Documentation files
└── ...
```

And here, templates for creating your new articles:

| Template                                                                                                 | GitHub Example                                                                                                                                                                                                                      | Article Example                                                            |
| -------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| [Article Template](https://github.com/Codecademy/ugc/blob/main/documentation/authorA/mock-article-1.md)           | [how-to-convert-css-to-scss.md](https://github.com/Codecademy/ugc/blob/main/content/kyrathompson/how-to-convert-css-to-scss.md) ([Raw](https://raw.githubusercontent.com/Codecademy/ugc/main/content/kyrathompson/how-to-convert-css-to-scss.md))                              | Coming soon         |
| [Author Template](https://github.com/Codecademy/ugc/blob/main/documentation/authorA/author_meta.json) | [author_meta.json](https://github.com/Codecademy/ugc/blob/main/content/kyrathompson/author_meta.json) ([Raw](https://raw.githubusercontent.com/Codecademy/ugc/main/content/kyrathompson/author_meta.json)) | Coming soon |

Please read through the following in the [/documentation](https://github.com/Codecademy/docs/tree/main/documentation) folder. In these links, you'll find a write-up of our standards for content and style:

- [Content Standards](https://github.com/Codecademy/ugc/blob/main/documentation/content-standards.md)
- [Categories List](https://github.com/Codecademy/ugc/blob/main/documentation/categories.md)
- [Tags List](https://github.com/Codecademy/ugc/blob/main/documentation/tags.md)


### Codecademy Username and Profile Pic

As a UGC article content creator, you have the opportunity to have your Codecademy username and avatar displayed on the article!

## How do I submit a Pull Request (PR)?

Contributing follows this workflow:

1. Fork [this project repository](https://github.com/codecademy/docs).
2. Clone the forked repository to your computer.
3. Create and switch into a new branch.
4. Edit or create an article and commit the changes.
5. Make a PR to merge your fork with this repo.

If you haven't gone through this workflow before, you can check out [this GitHub tutorial](https://github.com/firstcontributions/first-contributions#readme) (highly recommend) or [this YouTube video](https://www.youtube.com/watch?v=rgbCcBNZcdQ) to learn about how to make a PR from a fork using Git.

Alternatively, if you'd prefer to keep things to the GitHub UI, you can follow the instructions in that video up to 1:18 to fork this repo. After that, you can create your article in your fork using the UI and then make a PR by pressing this handy button:<br>

<img src="https://github.com/Codecademy/docs/blob/main/media/pull-request-ui.png" alt="Code block with Codebyte tags" width="800"/>

If you are uncomfortable using Git, you can also check out [this YouTube video](https://youtu.be/RPagOAUx2SQ) to do this all using the GitHub Desktop app.

## Any tips for a Pull Request?

- Before making a PR, make sure you pushed your changes from a branch other than `main`.
- Name the new branch after the changes being pushed to the PR.
- Keep your PRs byte-sized. 1 article per PR!
- All contributors must sign the Contributor License Agreement (CLA).
- All required [status checks](https://docs.github.com/en/github/collaborating-with-pull-requests/collaborating-on-repositories-with-code-quality-features/about-status-checks) are expected to pass in each PR.
- For Beta, we currently require at least one round of reviews from the [content team members](https://github.com/codecademy/docs#-content-team). Make sure to make the changes within 5 days.
- Requested changes must be resolved before merging.
- Your article will be deployed within the hour after it's merged!

## What do we check for?

- Technical accuracy
- Formatting standards
- Typos/bugs
- Plagiarism

## Additional Resources

Remember, if you ever have any questions at all, we're always here to help in the [Codecademy Forums](https://discuss.codecademy.com/) and [Codecademy Discord](https://discord.com/invite/codecademy).

Tag [@sonnynomnom](https://twitter.com/sonnynomnom) if you find a typo, bug, or issue! 🖖
