# Deprecated

Use this one instead -> [DumbMahreeo/rockfetch](https://github.com/DumbMahreeo/rockfetch)

---

# Gofetch
Just my little fetch scripts written in Go. Mostly made to replace ufetch

# Many scripts
Just like ufetch, gofetch has a script for every distro.

I'll add more scripts as soon as I need them.

Feel free to contribute if you want to add yours

# Theme
Feel free to modify it and upload more themes, the theme you can find right now is just the theme I'm using right now (as this is something I made for myself and published because why not)

---

# WM env var
I think this is the cleanest and less bloated way to detect for an environment. No need for fancy stuff like ufetch does, an env var is more than enough.

---

# Fedora
The Fedora script got many versions

* Standard: The one where every time `dnf list installed` is called every time
* Cached: The one where the number of packages is read from a cache file (described in the source file)
* DBversion: The one where to get the list of installed packages, `/var/cache/dnf/packages.db` is queried

(I suggest the dbversion)

---

# Installation

example of an arch installation

```bash
git clone https://DumbMahreeo/gofetch
cd arch # change directory accordingly
go build -ldflags="-s -w" gofetch.go
sudo mv gofetch /usr/bin/
```
