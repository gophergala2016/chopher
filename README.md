# Chopher (Chopin + Gopher) [![Build Status](https://travis-ci.org/gophergala2016/chopher.svg?branch=master)](https://travis-ci.org/gophergala2016/chopher)
## [You must hear it to believe it! Click here!](http://chopher.herokuapp.com/)

Chopher is a for fun experimental project that hashes files as music.

For a given file Chopher will go through it's bytes and create a song. The first 16 bytes of a files determine the scale and the key note of the song, and the rest are used to determine the notes from the scale.

Chopher automatically determines whether a note can be part of a chord that fits the current scale, or if it should be a standalone note.

To generate the sound the [Karsplus-Strong algorithm](https://en.wikipedia.org/wiki/Karplus%E2%80%93Strong_string_synthesis) is used, but I think the implementation can be much faster. Due to the speed of the implementation the songs are limited to 200 notes.
