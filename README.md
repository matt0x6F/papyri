# Papyri

Papyri is an in-development bookmark manager. 

A high level list of goals:

- [ ] Import HTML formatted bookmarks from Firefox and Chrome
- [ ] Automatic imports from HackerNews favorites
- [ ] Automatic imports from GitHub stars
- [ ] Backup to remote location
- [ ] Deduplicate imports

## Hacking

Papyri requires Go >=1.17, [Wails v2](https://wails.io), and NodeJS. Go modules and NPM will retrieve all relevant dependencies.

### Live Development

To run in live development mode, run `wails dev` in the project directory. The development application will perform hot code reloading, 
however, the frontend dev server will also run at http://localhost:3000. The browser allows use of browser and Vue dev tools.

### Building

To build a redistributable, production mode package, use `wails build`.
