# gb-sublime

Simple app that generates a Sublime Text `.sublime-project` file with the `GOPATH` configured for use by [`gb`](https://github.com/constabulary/gb) and [`GoSublime`](https://github.com/DisposaBoy/GoSublime).

## Usage

Navigate to the directory your `gb` project is in, then run:

```
cd /home/adam/code/my-gb-project
go get github.com/ap0/gb-sublime
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
                "GOPATH": "/home/adam/code/my-gb-project:/home/adam/code/my-gb-project/vendor",
                "PATH": "$PATH:$HOME/src/go/bin"
            }
        }
    }
}
```