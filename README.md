TS6
===

This is a generic package implementing support for the TS6 server linking 
protocol as implemented in Elemental-IRCd. 

Things it will handle for you:
 - Server login and password assurance
 - Keeping track of network state
 - Sending events to a common channel

This will function without the use of global variables so many instances may be 
made into a Go application.
