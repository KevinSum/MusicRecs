# MusicRecs

Music streaming services have a tendency to suggest songs you’ve listened to before. This is a project that integrates with the LastFM api, intended to provide music recommendations that aren’t based on your listening history. Instead the system give recommondations for a selected song or artists based off what other listeners also like to listen to. The project also features a filter list which you can add to to filter out specific artists from song/artists recommendations.

## Features
- Server set up to communicate with the lastFM API over HTTP requests to fetch artist/track recommendations for specific songs/artists.
- Server storing a list of filtered artists in an SQLite database to avoid recommending music from certain artists.
- Server providing RESTful API endpoints for adding/removing artists to filter list and getting song/artists recommendations for certain songs/artists over HTTP.
- Server deployed to AWS/ECS as a Docker image.
- Command line interface implemented to serve as a front-end for communicating with the server over HTTP.

## Setup
Run command `go run cli/main.go` and run the following commands

- `help` Display help message detailing all commands
- `getSimilarArtists <artist>` Retrive list of similar artists for selected artist
- `getSimilarTracks <track>, <artist>` Retrive list of similar tracts for selected track
- `addToBlacklist <artist>` Add an artist to filter list so that they won't be recommended
- `removeFromBlacklist <artist>` Remove an artist from filter list
- `getBlacklist <artist>` Retrieve filter list
