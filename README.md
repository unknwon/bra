Brilliant Ridiculous Assistant
==============================

Bra(Brilliant Ridiculous Assistant) is a command line utility tool for Unknwon.

## Usage

```
NAME:
   bra - Brilliant Ridiculous Assistant

USAGE:
   bra [global options] command [command options] [arguments...]

VERSION:
   0.1.0.0930

COMMANDS:
   run		start monitoring and notifying
   sync		keep syncing two end points
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version
```

## Configuration

You have to have a `.bra.toml` file under the work directory. An example configuration take form [gogsweb](https://github.com/gogits/gogsweb):

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
build_delay = 1500				# Minimal interval to Trigger build event
cmds = [
	["go", "install"],
	["go", "build"],
	["./gogsweb"]
]								# Commands to run
```

## License

This project is under Apache v2 License. See the [LICENSE](LICENSE) file for the full license text.