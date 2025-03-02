Gator
======

Warning
--------
This software has been created for an online backend course on boot.dev website.
It has not been thouroghly tested and may not work as intended.

USE AT YOUR OWN RISK !!!

Prerequisites
--------
Following software need to be installed on your system
- Postgres (version 15+): [link](https://www.postgresql.org)
- Go language (version 1.24.0 or superior): [link](https://go.dev)

Please refer to official website for installation instructions, or to your distribution package manager if you're on Linux for example

Installation
--------
- Clone this repository in your user workspace folder
- Open a terminal and change directory (`cd`) to is folder
- Run `go install` to compile the code and put resuting executable in GOPATH folder

If GOPATH is in your PATH environment variable, you can run it with `gator` command

Configuration
--------
In order to use this software you will have to create `.gatorconfig.json` file in your HOME folder
This file should contain a JSON object (with `{}` delimiter) with following key(s):
- db_url (string): the value is the connection string `gator` will use to connect to postgres database

Usage
--------
Command line: `gator <cmd> [Arguments...]`

list of commands (non exhaustive):
- `register <username>`: create user with given name and login with newly created user
- `login <username`: login with given username. User must have been previously created with `register` command
- `reset`: clean database. Deletes all users and associated data
- ....

