# Support tools

A collection of CLI tools meant to support and automate various aspects of
developing Envoy, particularly those related to code review. For example,
automatic DCO signoff and pre-commit format checking.

## Usage

To get started, you need only navigate to the Envoy project root and run:

```bash
./support/bootstrap
```

This will set up the development support toolchain automatically. The toolchain
uses git hooks extensively, copying them from `support/hooks` to the `.git`
folder.

The commit hook checks can be skipped using the `-n` / `--no-verify` flags, as
so:

```bash
git commit --no-verify
```

## Functionality

Currently the development support toolchain exposes two main pieces of
functionality:

* Automatically appending DCO signoff to the end of a commit message if it
  doesn't exist yet. Correctly covers edge cases like `commit --amend` and
  `rebase`.
* Automatically running DCO and format checks on all files in the diff, before
  push.