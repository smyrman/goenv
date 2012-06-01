Goenv
=====

Goenv lets you create an isolated environment where you install Go packages,
binaries, or even C libraries.

Goenv should work for most Unix variants as long as the terminal emulator is
configured to use *Bash*. For MacOSX, everything except the C library
support is expected to work. You are welcome to implement support for more
shells and operating systems if you want to, e.g. for the
*Windows Command Prompt* or full Bash support on MacOSX or Windows.


Usage
-----

Create or update an environment:

	$ goenv <path to environment>

Activate an environment:

	$ . <path to environment>/sourceme.bash

The shell prompt will now be prefixed with the environment's folder name to
indicate that the environment is active. You may have noted that the prefix has
a certain color.  This color was chosen randomly when you ran the *goenv*
command. To set a new random color, run the goenv command again.

Deactivate an active environment:

	$ deactivate


Installation
------------

Issue:

	$ go get github.com/smyrman/goenv

If you have installed Go as root, and you have not set up a user level GOPATH
and prepended "$GOPATH/bin" to your PATH environment variable, you might need
to issue the command as root.

Feel free to create a proper package for your distro, if you want to.
