# Job Search

This is a tool I'm using to keep track of all the positions I've applied to.

I'm doing this to also excercise my go skills and get more used with templ
and htmx.

Currently this is just a toy project, not meant to be deployed anywhere,
but there's already some portability in canse I want to run this again
in another machine.

### Features:

It can toggle the jobs I've applied to and they're persisted on the DB.

### Preview

![Screenshot of the application](https://github.com/didicodethat/jobsearch/blob/main/static/screenshot.png?raw=true)

#### Planned Features

 - (DONE) Adding new positions
 - Editing current positions

#### Not planned but would be good

 - User Authentication / Sessions

### Commands

```sh
$ jobsearch dev
```
runs the development server

```sh
$ jobsearch installdb
```
installs the db, currently this isn't programmable via cli, so it connects to
localhost, port 5432, user postgres, password postgres.

```
$ ./run-dev.sh
```
Runs the server and also runs the templ generate command.

### Tools Used:

 * [htmx](https://htmx.org)
 * [templ](https://templ.guide/)
 * [sqlx](https://github.com/jmoiron/sqlx)
 * [pq](https://github.com/lib/pq)
