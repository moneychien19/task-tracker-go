# Task Tracker

## Usage

```bash
go build -o main
```

### Adding a new task

```bash
./main add "Buy groceries"
```

### Updating and deleting tasks

```bash
./main update 1 "Buy groceries and cook dinner"
./main delete 1
```

### Marking a task as in progress or done

```bash
./main mark-in-progress 1
./main mark-done 1
```

### Listing all tasks

```bash
./main list
```

### Listing tasks by status

```bash
./main list done
./main list todo
./main list in-progress
```
