# vX

**vX** is a very basic version control system. For details at my medium blog [How was I build a version control system (VCS) using pure Go 🚀](https://abdulsamet-ileri.medium.com/how-was-i-build-a-version-control-system-vcs-using-pure-go-83ec8ec5d4f4)

# Demo

[![asciicast](https://asciinema.org/a/487303.svg)](https://asciinema.org/a/487303)

# vX Commands

All commands: `init`, `add`, `status`, `commit`, `checkout`, `history`

**`vx init`**

**`vx add {file, directory}`**

- vx add a.go
- vx add src/

**`vx status`**

**`vx commit -m "message"`**

- vx commit -m "init"

**`vx checkout {commit_number}`**

- vx checkout v1 -> checkout to first commit
- vx checkout v10 -> checkout to tenth commit

**`vx history`**
