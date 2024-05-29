# do shell
The do shell is an interpreted language for doing things. The goal is to eventually work like bash, but with less pain. Here's what an example script would look like:

```do
# everything here is a WIP idea and subject to change
# you can import from other scripts
import { myfunc } from "./other.do"

myfunc()


fn add(a, b) {
    return a + b
}

let myvar = 10

# brackets execute bash commands
# $() escapes back into do shell land
[echo "I have $(10) dollars"]

# commands return an object containing:
# - the stderror
# - the stdout
# - the status code of the exited object
{ out, status, error } = [git status]

```


Do shell is meant for when you need to do a bit more programming than what would be comfortable in a bash script, but you still need to make frequent
