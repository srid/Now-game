# Now
Now is a game where success is measured using your concentration.

# Status
I just started working on this project; still in early stages.

# Key ideas
- Uses brain scanner to detect concentration levels
- When concentration goes down, visibility becomes fuzzy
- Concentrate again to re-gain visibility
- Goal: lead your wiggling bot along the narrow pathway

# Technology
- Go: server-side, and for receiving OSC stream from Muse headset
- Elm: client-side game engine and renderer
- muse-io: from Muse SDK (may run on a different machine)

# HACKING

Install the prerequisites:

* Go 1.4+
* [Elm](http://www.elm-lang.org)

Build and run the server:

```
make all run
```

Open the web app at http://localhost:8000

Want to help? Ideas? Get in touch at srid@srid.ca
