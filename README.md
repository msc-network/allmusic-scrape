## Allmusic album scraper

Run the executable from the command line with the flag -i ALBUM-ID

Album ID is found at the end of the URL beginning  with 'mw'

e.g. on Linux

```./builds/allmusic-scraper-linux-amd64-0.0.2-alpha -i mw0000275621```

Or with optional filename

```./builds/allmusic-scraper-linux-amd64-0.0.2-alpha -i mw0000275621 -f scrape.txt```

Run without arguments to get a list of options that can be passed to the binary

This will write a file called data.txt in the format

Artist - Album Title

Song Names [ Composers ]

### Building Binaries

- Clone the repo

- Run ``` ./build-allmusic-scraper.sh```

It is setup to build for the following platforms

Linux 64bit - allmusic-scrape-linux-amd64-VERSION

Windows 32bit (untested) - allmusic-scrape-windows-386-VERSION.exe

Windows 64bit (untested) - allmusic-scrape-windows-amd64-VERSION.exe

MacOS (untested) - allmusic-scrape-darwin-amd64-VERSION

#### For Linux

Run ```go build allmusic-scrape```
