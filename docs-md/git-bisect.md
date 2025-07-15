# Git Bisect

Ellie provides subcommands to help you use git bisect to find the commit that introduced a bug.

## Start Bisect

```
ellie git bisect
```

Starts a bisect session.

## Mark Good Commit

```
ellie git bisect good
```

Marks the current commit as good.

## Mark Bad Commit

```
ellie git bisect bad
```

Marks the current commit as bad.

## Reset Bisect

```
ellie git bisect reset
```

Ends the bisect session and resets the state. 