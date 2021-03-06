gong
====

gong is a web git repository viewer written in Go intended to replace [cgit][3]
for my personal use case. I've licensed it under the [AGPLv3][4].

Building
--------

This is a little messy at the moment. Dependencies are currently
[highlight.js][5], [go-ini][6], [blackfriday][9], and [git2go][7]
([libgit2][8] bindings for Go).

For go-ini and blackfriday, you can simply `go get
github.com/go-ini/ini` and `github.com/russross/blackfriday`
respectively, but for git2go, you'll need to `go get
github.com/libgit2/git2go` _and then_ checkout the branch corresponding
to your system's libgit2 version in the cloned repository. For example,
since I have libgit2 version 0.24 installed, I checked out
`$GOPATH/src/github.com/libgit2/git2go` to the `v24` branch.

Once you have the Go dependencies installed, run `go build` in the
`gong` directory (aka the package for the binary).

To download highlight.js, run the `get-highlightjs.sh` shell script. If
you need to add more languages (the default list is short for
performance reasons), you can modify the `languages` file the script
generates and run the script again.


The gong Manifesto
------------------

### Gripes with Other Tools

Many tools for self-hosting git repositories on the web exist, but none
of them perfectly fit my use case.

 * **GitHub**, the "largest host of source code in the world"
   [according to Wikipedia][1].

    - Nonfree, which disappoints me. Free software deserves free tools.
    - Private repositories cost $$$
    - Hosting your own instance costs an ungodly amount of money ($2,500
      a year!)

 * **Gitlab**, a libre alternative to GitHub

    - I've never run it myself honestly, but I've heard it's quite
      resource-intensive. Like, gigabytes-of-ram-required kind of
      intensive.
    - Too many dependencies. According to the [readme][2]:

        * GNU/Linux
        * Ruby
        * Redis
        * MySQL or PostgreSQL
        * SMTP server

      I just want to display a git repository's files in visitors'
      browsers, not send them emails or have pull requests or whatever.
      Having two different databases and an SMTP server is overkill for
      me.

 * **gogs**, a GitHub clone written in Go

    - Close to what I want (e.g., can use SQLite and no SMTP server
      required), but still way too much functionality. I don't want
      logins, pull requests, or databases; I want to display git
      repositories.

 * **cgit**, a cgi C program that I currently use and like

    - Still uses cgi, which requries a FastCGI wrapper on nginx and
      other HTTP servers without cgi support
    - Not mobile-friendly
    - Pages still built around tables
    - No templates of any kind, with HTML instead generated by a bunch
      of `printf()`s. The developers did a great job of making this not
      as bad as it sounds, but it still makes even minor HTML changes
      require a recompile.
    - It's a webapp written in C. Again, it's well-written and I respect
      the developers, but this disturbs me.

### Goals

In essence, provide a minimal, modern Go reboot of cgit for my git
repositories.

 * Rearrange the minimal UI of cgit such that it resembles GitHub's
   interface, which non-{cgit,gitweb} users expect. For example, the
   repository root page should show files and the README
 * Be mobile-friendly
 * Use templates
 * Be fast
 * Don't use a database; instead, read configuration from a file like
   cgit
 * Require no write access to git repositories (users should push over
   SSH instead)
 * Run as a daemon supporting FastCGI or HTTP
 * Perform markdown rendering in Go instead of firing up a Python/Perl
   script (blackfriday is good)
 * Perform syntax highlighting in Go or client-side with JavaScript
   instead of `exec()`ing `highlight` or another program
 * Learn about git!
 * Procrastinate

[1]: https://en.wikipedia.org/wiki/GitHub
[2]: https://gitlab.com/gitlab-org/gitlab-ce/blob/master/README.md
[3]: https://git.zx2c4.com/cgit/about/
[4]: https://www.gnu.org/licenses/agpl-3.0.en.html
[5]: https://highlightjs.org/
[6]: https://github.com/go-ini/ini
[7]: https://github.com/libgit2/git2go
[8]: https://libgit2.github.com/
[9]: https://github.com/russross/blackfriday
