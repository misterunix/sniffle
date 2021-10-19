# sniffle

What a name right? When creating this repo Github showed. 
```Great repository names are short and memorable. Need inspiration? How about shiny-sniffle?```
And that cracked me up, so I chose sniffle. What the heck right?

## What is it?
It is a set of routines I end up rewriting or grabbing from other code on my internal server. 

## Why?
I have never made a public module before and wanted to give it a shot and see what happens.

## What is there?
- Hashing functions.
  - file hash
  - string hash
- JSON functions - this is the big one for me. I use JSON a lot!
  - write a slice of a struct to a file
  - read a file into a slice of a struct
- CPU and system functions.
  - uses on go only packages
  - gets CPU cores and type
  - gets os and pid
  - gets Go version and number of go routines. 
  - This will get more as I need them.
- String 
  - Left justify with padding.
  - Right justify with padding.
  - Center with padded, optional padding on the right.
  - Split lines at width and/or word boundry.

## Why in the world did you do this. 
Well, as I said, it's stuff I use over and over again. We all know there are 100 different ways to do something in go and this is just my way. :beer:

## Contribs? 
Sure, if they only require the standard Go library. Let's do it. 

## Formatting
Clarity over cleverness!

## Conduct
Treat this as if your grandma is in the room. Keep it civil.