# do-shell
Do shell (`dosh` on the command line) is an interpreted language for getting things done quickly. It's main strengths are:
- quick tools
- automated integration tests
- system utility scripts

Do shell excells in places where a simple bash script isn't sufficient, but a full blown application is excessive and time consuming.

# Examples
## Instant CLI
WIP

## HTTP
```dosh
#!/usr/bin/env dosh
use "std:http"
use "std:env"

let res = http.get("localhost:3000/get", {
    body: {
        json: here
    },
    headers: {
        Authorization: "..."

    }
})

print(res);
```


## File Manipulation
WIP

## Run commands
WIP

## JSON / YAML Handling
WIP
