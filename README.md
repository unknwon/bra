Brilliant Ridiculous Assistant
==============================

Bra(Brilliant Ridiculous Assistant) is a command line utility tool.

## Usage

```
NAME:
   Bra - Brilliant Ridiculous Assistant is a command line utility tool.

USAGE:
   Bra [global options] command [command options] [arguments...]

VERSION:
   0.3.0.0715

AUTHOR:
  Author - <unknown@email>

COMMANDS:
   init		initialize config template file
   run		start monitoring and notifying
   sync		keep syncing two end points
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version
```

## Quick Start

To work with a new app, you have to have a `.bra.toml` file under the work directory. You can quickly generate a default one by executing following command:

```
$ bra init
```

## Configuration

An example configuration take form [gogsweb](https://github.com/gogits/gogsweb):

```
[run]
init_cmds = [["./gogsweb"]]		# Commands run in start
watch_all = true				# Watch all sub-directories
watch_dirs = [
	"$WORKDIR/conf",
	"$WORKDIR/models",
	"$WORKDIR/modules",
	"$WORKDIR/routers"
]								# Directories to watch
watch_exts = [".go", ".ini"]	# Extensions to watch
ignore = [".git", "node_modules"] # Directories to exclude from watching
build_delay = 1500				# Minimal interval to Trigger build event
cmds = [
	["go", "install"],
	["go", "build"],
	["./gogsweb"]
]								# Commands to run
```

## License

This project is under Apache v2 License. See the [LICENSE](LICENSE) file for the full license text.
