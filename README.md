# 🎬 go-tmdb-cli

A simple and fast command-line tool written in Go to fetch and display movie data from **The Movie Database (TMDB)** API.

This tool lets you quickly browse movies like:

* 🎥 Now Playing
* 🔥 Popular
* ⭐ Top Rated
* 📅 Upcoming

All directly from your terminal, with clean and readable table output.

---

# 🚀 Features

* Fetch real-time movie data from TMDB
* Clean CLI table output using Go’s `tabwriter`
* Optional full title display (`--full`)
* Lightweight and fast (no external dependencies)

---

# 📦 Installation

## 1. Clone the repository

```bash
git clone https://github.com/JanisJuska/go-tmdb-cli.git
cd go-tmdb-cli
```

## 2. Build the binary

You can name the build however you like:

### Option A: `tmdb`

```bash
go build -o tmdb
```

### Option B: `tmdb-cli`

```bash
go build -o tmdb-cli
```

---

## 3. (Optional) Move to PATH

To use it globally:

```bash
mv tmdb /usr/local/bin/
```

Now you can run it from anywhere:

```bash
tmdb --type top
```

---

# 🧑‍💻 Usage

## Basic syntax

```bash
tmdb --type=<category> [--full]
```

---

## 📌 Required flag

### `--type`

Specifies which movies to fetch.

Available values:

* `playing` → Now Playing
* `popular` → Popular movies
* `top` → Top Rated movies
* `upcoming` → Upcoming releases

---

## 📌 Optional flag

### `--full`

Displays full movie titles (no truncation).

---

# 💡 Examples

### Show top rated movies

```bash
tmdb --type=top
```

### Show popular movies with full titles

```bash
tmdb --type=popular --full
```

### Show currently playing movies

```bash
tmdb --type=playing
```

---

# 🖥 Example Output

```
Showing Top Rated movies from TMDB:

#   Title                          Release Date   Lang   Popularity Score   Vote Avg
1   The Shawshank Redemption       1994-09-23     en     46.37              8.7
2   The Godfather                  1972-03-14     en     42.00              8.7
...
```

---

# ⚠️ Notes

* This tool uses the TMDB API — you may need your own API key/token.
* Keep your API token secure (avoid committing it to public repos).

---

# 🛠 Tech Stack

* Go (standard library)
* `net/http` for API requests
* `encoding/json` for parsing
* `text/tabwriter` for CLI formatting

---

# 📈 Future Improvements (ideas)

* Pagination support (`--page`)
* Limit results (`--limit`)
* Genre mapping (IDs → names)
* Colored CLI output
* Subcommands (`tmdb top`, `tmdb popular`)

---

# 📄 License

MIT License

---

# 🙌 Acknowledgements

* [TMDB API](https://www.themoviedb.org/documentation/api)

---

Enjoy browsing movies from your terminal 🍿
