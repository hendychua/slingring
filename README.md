slingring
===

Manage feature branches across projects. This program is bias in the sense that it assumes you are using `git` as your version control software.

### Motivation
When working on a new feature/project that requires modifying multiple repositories, it is common to create a new feature branch for each of the repositories, e.g.:
```
cd ~/utils
git checkout master
git pull
git checkout -b new-feature

cd ~/api-server
git checkout master
git pull
git checkout -b new-feature

cd ~/frontend
git checkout master
git pull
git checkout -b new-feature
```

`slingring` makes it easy to manage feature branches by allowing you to group them together.

#### What does slingring mean?
Just a reference I got from Doctor Strange. A slingring helps you open portals and jump to a different place instantaneously.

#### Terminologies
There is only one terminology to remember: `Dimension`. Basically a Dimension is a new feature group. You can add projects to this group a.k.a. group them into a Dimension, so that when you have to switch back onto the feature branch, you can just type `slingring jump <dimension name>`.

#### Usage
Create a new dimension (feature branch):

`slingring create news-feed`


List all dimensions created:

`slingring list`
```
  live-chat
  another-feature
  news-feed
* payments
(* denotes current dimension)
```

Delete a dimension (this only deletes the dimension but does not modify the projects in the dimension):

`slingring delete news-feed`

Describe a dimension:

`slingring describe news-feed`

Add a project to a dimension:

`slingring add-project ~/utils` This command adds to the current dimension. To add to a different dimension, use `--dimension` or `-d` global option, e.g.:

`slingring --dimension live-chat add-project ~/utils`

Remove a project from a dimension:

`slingring remove-project ~/utils` This command removes from the current dimension. To remove from a different dimension, use `--dimension` or `-d` global option, e.g.:

`slingring --dimension live-chat remove-project ~/utils`

Jump dimension, a.k.a. switching git branches:

`slingring jump news-feed` This command will checkout to a branch named `news-feed` in each of the projects that have been added to this dimension. If the branch does not exist, it will be created.

Show config:

`slingring config`
```
BaseBranch: develop // base branch to create new branch from
PullFirst: true // pull latest from remote before checking out new branches
HandleDirtyRepo: AbortAll // How to handle dirty repos
```

Set config:

`slingring config <key> <value>...`

e.g.

`slingring config basebranch master pullfirst false handledirtyrepo stash`


#### Config options
- BaseBranch (default: develop): name of base branch to create new branches from
- PullFirst (default: true): whether to run `git pull` before creating new branches. True or False.
- HandleDirtyRepo: how to handle dirty repos. Valid options:
  - AbortAll (default): abort the whole jump process when one of the repos is dirty.
  - AbortContinue: abort jumping for dirty repos, continue for clean repos.
  - Stash: stash uncommitted changes before switching branches. This will stash uncommited but not untracked changes.

#### Installation
```go get -u github.com/hendychua/slingring
glide install
go install```

#### Known limitations
- This project writes to two files: `~/.slingring/globalData.json` and `~/.slingring/globalSettings.json`. JSON is probably not the most efficient way to write and read data (but it is the simplest).
- Whenever data or config is modified, the whole data/config object is being rewritten into those files. This is fine when the data is small but will take a long time when there are more and more data (this limitation is marked with a `TODO` in the code for future improvement.) It may require changing the data format.
- Not meant to have more than 1 process running `slingring` as this program reads from/writes to those global files and there is no synchronization mechanism in place.
- Only works for git projects.