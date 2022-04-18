---
Title: "How To Use Git For Beginners"
Description: "Beginner friendly content for using Git"
Categories:
  - "computer-science"
  - "developer-tools"
Tags:
  - "Git"
  - "GitHub"
CatalogContent:
  - "learn-git"
  - "learn-the-command-line"
  - "paths/computer-science"
---
[the Git Website]: https://git-scm.com/downloads
[Add]: https://www.codecademy.com/resources/docs/git/add
[Reset]: https://www.codecademy.com/resources/docs/git/reset
[here]: https://www.theserverside.com/feature/Why-GitHub-renamed-its-master-branch-to-main
[GitHub]: https://github.com/

_**Prerequisites:** the Command Line_

## Introduction
Git is a version control system used for tracking changes to files over a period of time. Understanding Git is a fundamental skill for all developers as it is the modern standard in the software development industry for tracking and coordinating work within a dev team.

## Getting Started
The easiest way to get started with using Git is to download and install the software following the instructions on [the Git Website] for your operating system. Once installed run the command below from a terminal (for Mac and Linux) or command prompt (for Windows) to check that installation was successful:
```
git --version
``` 
If Git has been successfully installed, the terminal/command prompt should output the current version of Git, which at the time of writing was ```2.35.3```.

## Git Workflow
Before jumping in to using Git, it will be helpful to have an overview of a straightforward Git workflow. 

A Git project generally has:

- **A Working Directory** - the project folder on your computer where files are created, edited and deleted i.e. where work is actually done.

- **A Staging Area** - contains the changes made to files in the Working Directory ready to be committed.

- **A Repository** - also known as the repo, the place where changes are saved by Git as different versions of the project. There will always be a local repo but there will also often be a remote repository hosted by a provider such as GitHub, GitLab or Bitbucket.

A typical Git workflow involves:

1. Initialising Git in the Working Directory;
2. Working on files in the Working Directory, then adding these to the Staging Area;
3. Commmiting the changes from the Staging Area into the Repository.

Let's unpack each of these steps.

## Initialising Git
To turn a Working Directory into a Git project, navigate to the directory in the terminal/command prompt and run the following command:
```
git init
```
This initialises a Git repository (which is just a name for a special folder called ```.git``` inside a project) on your local machine, and is required in order for Git to start tracking changes. This step might look something like this:

![git-init](git-init.png)

When ```git init``` is run inside the Git Tutorial folder, an empty Git repository is created inside of that directory. The repository is currently empty because Git doesn't know at this point which files it should track. Any untracked files can be viewed using the command:
```
git status
```
This will display untracked files in <span style="color:red">red</span>, and will show some useful output confirming that nothing has been added. For example, if ```git status``` is run in the Git Tutorial folder, which now contains a file called ```learning-git.txt``` the output looks like this:

![git-status-untracked](git-status-untracked-files.png)

You will see the output says ```On branch main``` which you can ignore at the moment. You'll also see a mysterious file called ```.DS_Store``` which is a file that is automatically created by Mac OS X. This file contains information about system configurations, so should not be committed as part of the Git workflow. To make sure this file is not added to the commit, a special file called ```.gitignore``` must be made in the Working Directory.

### .gitignore
The ```.gitignore``` file is a special file which generally lives in the root of the Working Directory. It contains a list of files and/or directories which a developer has added that Git should not track, and so will not be committed. These files may be log files, or particular modules in a project that shouldn't be tracked, or (for Mac users) files like ```.DS_Store```. Make sure to try and remember to add files which should not be tracked ***before*** they are committed.

Now, with the ```.gitignore file``` made and the ```.DS_Store``` file added, running ```git status``` in the terminal shows the following output:

![git-status-untracked-with-gitignore](git-status-untracked-with-ignore.png)

## Making Changes and Adding Files to the Staging Area
In order for Git to start tracking changes, files need to be added to the Staging Area. This can be done with the command:
```
git add <filename>
```
where filename is the name and extension of the file to be added to the staging area. Multiple files can also be added with the syntax:

```
git add <filename1> <filename2> <filename3>
```

If all files in the Working Directory need to be tracked, the syntax

```
git add .
```
can be used.

Once a file has been added, ```git status``` can be run in the terminal/command prompt to check that the file is in the Staging Area. Git will show the changed files to be committed in <span style="color:green">green</span> text like so:

![git-status-untracked-with-gitignore](git-add-and-status.png)

You can find out more about the additional extensions to ```git add``` in the [Add] section of the Codecademy Docs for Git.

### Oops! I didn't want to add that!
There may be occasions where a file has been incorrectly added to the Staging Area. This can be undone with the command ```git reset <filename>```. This will not affect the changes done to the file in any way, it will simply remove it from the staging area. 

In the case where all files needs to be unstaged, the command ```git reset``` can be run without adding further arguments.

Note that as of ```git version 2.24``` there is alternative syntax to unstage files, which is ```git restore --staged <filename>``` where ```filename``` is the name of the file to be unstaged.

## Committing Changes
Committing is generally the last step of the Git workflow and is generally thought of being a snapshot of a project at a particular time. In this stage, changes in the Staging Area are saved inside the local repository, which is achieved with the command ```git commit -m``` followed by a space and a short message explaining the commit in quotes, for example:
```
git commit -m "Add title and description to intro file"
```
There are general conventions for writing a commit message including making sure they are no more than 50 characters long and writing them in the imperative or present tense. Messages should always be clear and confirm what the change made is so that others can easily see what has changed. 

Committing is an important step in the Git workflow, as once a change has been committed, it can be recalled at a later date or the project can be rewound to that particular version. Commit history can be viewed with ```git log```, which will display a list of commits in chronological order (the most recent commits being at the top), along with information such as the author of the commit, the date and time of the commit and the commit message.

### Oops! I didn't want to commit that!
In the situation where the last commit needs to be undone, the ```git reset``` command can again be used, wtih some modifications. 

The current commit is called the ```HEAD``` commit, which is generally the most recently made commit. To find out which commit this is, the command ```git show HEAD``` can be executed in the terminal/command prompt. This will display information about the most recent commit, including a unique 40 character SHA hash, which Git uses like an id to identify revisions in the repo.

To undo the immediately previous commit, the command ```git reset --soft HEAD~1``` can be run, which will rewind the current ```HEAD``` commit to the immediately previous commit. The ```--soft``` flag ensures that that any changes made to the files are preserved.

If the rewind needs to go beyond the most recent commit, the command ```git reset SHA```, where ```SHA``` is the first 7 charcters of the SHA of the commit, can be used. Note that the SHA of all previous commits can be viewed using ```git log```.

More about the ```reset``` command can be viewed in the [Reset] section of the Codecademy docs for Git.

## Git Branching
In the Initialising Git section of this article, the output from running ```git status``` referred to being on ```branch main```. In Git, branches are a core feature. They can be created, deleted, compared, merged and tracked. 

The first branch to be aware of is the ```HEAD``` branch, which is the currently active branch. So, the ```main``` branch referenced above was the currently active branch and in GitHub is the name of the default branch. In Git, the default branch name is ```master```, which was also originally the case in GitHub however this was renamed for the reasons outlined [here]. The main take away though is that both ```master``` and ```main``` are used as default branch names.

The idea of Git branches is that different branches can be created and worked on, and changes can be merged into the ```main``` branch. Branches are in effect an independent line for the adding/staging/committing process, forked from the ```main``` branch. Once on a branch, commits are recorded in that branch's history and when ready a branch (including the changes made on that branch) can be merged i.e. combined into ```main```.

A new branch can be made with the following command:
```
git branch <name-of-new-branch>
```
To switch over to the new branch in a project the command:
```
git checkout <name-of-the-branch>
```
is used. These two steps can however be neatly combined into one using:
```
git checkout -b <name-of-new-branch>
```
Once on a new branch, files in the Working Directory can be worked on as usual, and changes can be added and committed. Those changes will however only be committed to the current branch and will not affect anything on ```main```.

### Merging
Merging combines multiple commits into one history, and is generally used to join or combine two branches. So, when ready for work on a branch to be merged into ```main```, the command ```git merge branch-name``` is used, where ```branch-name``` is the name of the branch to be merged into the ```main``` branch. This command is most often used when working on a project with others (more on this below).

Once a branch has been merged, it is no longer needed so can be deleted with:

```
git branch -d <branch-name>
``` 
This will though only delete the branch once it has been merged into ```main```. Replacing the ```-d``` flag with ```-D``` will delete the branch even if it hasn't been merged with ```main```.

## Collaboration
When collaborating with others on a project, it is generally the case that a shared remote Git repository will exist so that multiple people can work on the same project from different locations. There are a wide range of providers of remote repositories but the most popular and well-known is [GitHub].

### Git != GitHub
When starting out, it is easy to get confused between Git and GitHub but they are not the same. Git is a version control tool whereas GitHub is an onlne hosting service for Git repositories that has tools which integrate with Git. Git does not require GitHub to work and a remote repository is not needed for using Git however, when working with others on a project, GitHub (and any similar provider) makes doing so much easier.

### Workflow
When working with others, the general workflow is as follows:
1. **Create a Remote Repository** - where a remote repo does not already exist, one is generally created by one of the people working on the project using a hosting provider;
2. **Pull Changes to a Local Repo** - combines changes from the Remote Repo into a local branch
3. **Copy/Clone the Remote Repository to a Local Repo** - this replicates and copies everything in the remote repository to a local copy of the Git project.
4. **Create a new Local Branch** - a new branch will generally be made, forked from ```main```, and the new feature or fix will be worked on;
5. **Add and Commit Changes** - the normal Git workflow is followed for adding and committing changes;
6. **Push Changes to the Remote and Create a Pull Request** - the changes made on the local branch are pushed up to the remote branch to be reviewed by team members;
7. **Merge the Local Branch with Main** - once the changes have been approved by team members, the local branch is merged into the main branch.

Let's take a closer look at steps 2 and 5.

### Fetching/Merging/Pulling
To make sure the local repo is up to date with the remote, changes need to be fetched and merged from the remote and integrated into the local branch. This can be done in two stages using ```git fetch``` and ```git merge``` or in one step using ```git pull```.

```git fetch``` downloads files and commits from a remote repository into the local repo. It is used to see what changes have been made to the remote before they are merged into the local repository.

Fetching all branches from the remote can be achieved with ```git fetch <remote-name>```.

Alternatively, a specific branch can be fetched with ```git fetch <remote-name><branch-name>```.

Lastly, all registered remotes and their associated branches can be achieved with ```git fetch --all```.

```git fetch``` does not automatically merge changes made in the remote repository to the local repo, so is useful in cases where code needs to be reviewed before being locally merged. In order to merge changes into the local repo, the usual ```git merge``` command is required.

The ```git fetch <remote>``` and ```git merge origin <local-branch>``` command is combined as one into the ```git pull``` command, which automatically fetches and merges changes from a remote into the local branch. ```git pull``` can be run as a standalone command or with options such as ```git pull <remote-name>``` which will fetch and merge a specified remote with the local branch.

It is important to form a habit of fetching and merging or pulling when working in a team to make sure the codebase being worked on locally is always up to date with the remote.

### Pushing
When collaborating on a project, it is often the case that to (amongst other reasons) prevent broken code from being added to the main development branch and/or to ensure code is correctly formatted and follows whatever code conventions have been set by the team, those who have not been working on the new feature or fix will be asked to review the code. 

Code reviews are an important part of quality assurance in software development and not only helps protect against broken code from being mistakenly included in a codebase but also helps share knowledge amongst a team and can help develop the skills of both the reviewer and author.

One of the advantages of creating branches, rather than working off ```main``` is that any changes made to that branch do not affect the work of any other developers in the team. When a branch is pushed up to the remote, the team members can also take the time to review the code properly without needing to worry that there may be broken code included in ```main```.

To push code to the remote repo, the ```git push``` command is used as follows:
```
git push <name-of-remote-branch><local-branch-name>
```
This command pushes the local branch, along with all of the changes which have been committed to that branch, to the remote destination repository. If the remote branch does not already exist, one will be automatically created with the name of the local branch. An alternative syntax to this is:
```
git push origin <local-branch-name>
```
This also pushes the current local branch to a remote branch of the same name.

Once pushed, a pull request (sometimes called a merge request) can be created, which is generally a request to merge a branch into ```main```. Although the command line has been used throughout this article, there are some cases where using a GUI (Graphical User Interface) can be easier. This is definitely the case when creating a pull request, which can be done through the GUI of all of the major hosting services. 

After the pull request has been made, other team members can review the code and make suggested changes. Any changes required can be made locally and committed with the usual ```add/commit``` commands and pushed to the remote as described above. 

## Conclusion
In this article, we have looked at how to create a new Git project and how to add and commit changes. We have also looked at important concepts, like Git branching and have also looked at a straightforward workflow for collaborating on projects.

Git can be quite the minefield for beginners but as is the case when learning any new subject, it just needs time, practice and patience to get used to. As Git is the industry standard for version control, it is important as a beginner to get used to using Git and for those self-taught programmers amongst us, using hosting providers such as GitHub, which is an important tool not just for showcasing projects we have worked on but for showing an understanding of how Git works.