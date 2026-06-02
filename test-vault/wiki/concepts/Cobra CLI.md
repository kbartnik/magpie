---
title: "Cobra CLI"
type: concept
status: active
created: 2026-06-02
updated: 2026-06-02
sources:
  - dev/learning/magpie-go/2026-06-02.md
related:
  - Go Modules and Packages
  - Go Error Handling
tags:
  - go
  - cli
  - cobra
---

# Cobra CLI

Cobra is the standard Go library for building CLI tools with subcommands, flags, and help generation. Used by kubectl, Hugo, GitHub CLI, and most serious Go CLIs.

## Command Lifecycle

Each command runs through: `PersistentPreRun → PreRun → Run → PostRun → PersistentPostRun`.

The `Persistent` variants propagate **down the entire subcommand tree**. Non-persistent variants run only for the specific command they're defined on.

## PersistentPreRunE Propagation

`PersistentPreRunE` on the root command runs before every subcommand — making it the right place for shared setup (vault resolution, auth, config loading):

```go
rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
    // runs before magpie inbox, magpie archive, magpie log, etc.
    return nil
}
```

**Critical gotcha:** if a subcommand defines its own `PersistentPreRunE`, it **silently overrides the parent's** — the root hook won't run. Fix by calling the parent explicitly:

```go
subCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
    if err := rootCmd.PersistentPreRunE(cmd, args); err != nil {
        return err
    }
    // subcommand-specific setup
    return nil
}
```

This is a well-known cobra trap. If shared setup mysteriously stops running, check for a subcommand overriding `PersistentPreRunE`.

## RunE vs Run

| | `Run` | `RunE` |
|--|-------|--------|
| Signature | `func(cmd, args)` | `func(cmd, args) error` |
| Error handling | Manual — print + os.Exit yourself | Cobra prints error, exits with code 1 |
| Testing | Capture stderr to inspect errors | Inspect returned error directly |

Prefer `RunE` in almost every case. Combined with `SilenceUsage: true`, cobra won't print the full usage block on every non-usage error — only when the error is actually a usage problem.

## Flags

```go
// Persistent flag — available to command and all subcommands
rootCmd.PersistentFlags().StringVar(&vaultFlag, "vault", "", "vault path")

// Local flag — available only to this command
cmd.Flags().BoolVar(&verbose, "verbose", false, "verbose output")
```

## Unknown Subcommand Handling

To catch unknown subcommands (e.g. for plugin dispatch):

```go
rootCmd.Args = cobra.ArbitraryArgs
rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
    if len(args) > 0 {
        fmt.Fprintf(os.Stderr, "unknown subcommand %q\n", args[0])
        return nil
    }
    return cmd.Help()
}
```

## See Also

- [[Go Error Handling]] — error wrapping, sentinel errors
- [[Go Modules and Packages]] — module structure for CLI projects
