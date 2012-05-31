Goenv
=====

A tool similar to virtualenv for Python, but for Go. It lets you creare a
"virtual environment" where you install go packages.

Goenv should work for terminal emulators that are configured to use Bash.  You
are welcome to implement support for more shells if you want to, e.g. for the
*Windows Command Prompt*.

Usage
-----

Create or update an environemnt:

	$ goenv <path to environment>

Avctivate an environment:

	$ . <path to environment>/sourceme.bash

The shell prompt will now be prefixed with the goenv folder name to indicate that
the goenv is active. You may have noted that the prefix has a certain color.
This color was chosen randomly when you ran the goenv command. To set a new random
color, run the goenv command again.

Deactivate an active environment:

	$ deactivate


Installation
------------

Issue:

	$ go get github.com/smyrman/goenv

If you have not set up a user level GOPATH, and prepended "$GOPATH/bin" to your
PATH environment variable, you might need to issue the command as root.

Feel free to create a proper package for your distro, if you want to.
