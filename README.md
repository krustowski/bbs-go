# bbs-go (go 1.20)

simple implementation of go-telnet server to a BBS service

```
    __    __
   / /_  / /_  _____      ____  ____
  / __ \/ __ \/ ___/_____/ __ \/ __ \
 / /_/ / /_/ (__  )_____/ /_/ / /_/ /
/_.___/_.___/____/      \__, /\____/
                       /____/

savla-dev bbs-go telnet service
telnet bbs.savla.int 5555

```

## build and run

```
# copy example dotenv file
cp .env.example .env

# edit some vars
vim .env

# apply settings, build the binary and run it with env vars exported
make run
```

## system structure

+ message boards
+ TUI forums
+ online games
+ user administration 
