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
* DBversion: The one where to get the list of installed packages `/var/cache/dnf/packages.db` is queried

(I suggest the dbversion or the standard one. When I was on Fedora I used the dbversion)

---

# Rufetch
Right now I'm doing [this but in Rust](https://github.com/DumbMahreeo/rufetch), so probably I won't make any more of these ones.
