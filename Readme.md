This repository shows, that the GORM hook `AfterCreate` is not executed while adding records
into the mailboxes table (each entry into the mailboxes table should increase a the counter MailboxCount in the Domains table)

# Setup
On windows you need CGO support for go, you can install it using scoop:
```shell
scoop install mingw-winlibs
```

# Run
Execute it with:

```shell
go run ./...
```