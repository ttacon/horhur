# horhur

"ruhroh" but backwards.

# What is this for?

Ever had someone uploading files that weren't malware or malicious, but were instead HTML files used to redirect someone to somewhere else malicious? If so, then this library is for you.

It's a late-night, "how could this work" swing at easily identifying this class of annoying but not malware scannable files (really only HTML files at this time, since, you know... "redirects" implies a browser).

# Installation

## For the module

```sh
go get github.com/ttacon/horhur
```

## For the CLI

```sh
go get github.com/ttacon/horhur/cmd
```

# Example usage

```sh
horhur -f https://twitter.com
```

# Redirects handled (w/ examples)

- [ ] Meta-refresh
- [ ] JS history changes
- [ ] JS writing to document
- [ ] Web server 3XX
