# gb-sublime

Simple app that generates a Sublime Text `.sublime-project` file with the `GOPATH` configured for use by [`gb`](https://github.com/constabulary/gb) and [`GoSublime`](https://github.com/DisposaBoy/GoSublime).

## Installation

```
go get -u github.com/ap0/gb-sublime
```

To ensure it works with `gb`, make sure that the binary is placed in a directory on your $PATH.  If your `$GOPATH`/bin is in your `$PATH`, you're good.

## Running

Navigate to the directory your `gb` project is in, then run:

```
cd /home/adam/code/my-gb-project
# gb plugins are run using `gb <plugin-name>`
gb sublime
```

A `my-gb-project.sublime-project` file will be generated that looks like this:

```
{
    "folders":
    [
        {
            "path": "/home/adam/code/my-gb-project"
        }
    ],
    "settings": {
        "GoSublime": {
            "env": {
                "GOPATH": "/home/adam/code/my-gb-project:/home/adam/code/my-gb-project/vendor"
            }
        }
    }
}
```

You can also override this template.  Running the binary directly (`gb-sublime`, NOT `gb sublime`) with the `-write-template` flag to write out a modifiable template to your home folder:

```
$ ./gb-sublime -write-template
writing template to /home/adam/.gb-sublime-project.template...
```

Modify this to your liking, then run `gb project` again and it will use that template.