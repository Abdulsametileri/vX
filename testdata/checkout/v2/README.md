# vX

**vX** is a very basic version control system to understand the idea of event sourcing.

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