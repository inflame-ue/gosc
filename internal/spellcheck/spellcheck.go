package spellcheck

// TODO: Implement the spellcheck package, consult the notes below for architectural decisions.
//
// Rather obviously, the spellcheck package should not concern itself with the source of the text
// that is provided to it. Here we should adopt the UNIX philosophy of abstracting the particulars
// through the common file descriptor(in our case this will probably be an array of bytes, for efficiency)
//
// Other question then arises: How do we treat text that pre-supposes some sort of grammar?
// For instance, the example specification sends a markdown file to the program and we should
// not mark correct markdown syntax as being erroneous. Should we simply ignore all non-letter
// characters when looking through the file? Should we support non-ASCII languages through
// a parameter?
//
// What the project specification proposes is to "leverage an existing dictionary and then
// reading files, checking them, and letting the user know about the issue and where it is."
//
// Sample output to stderr is the following:
// Typos found:
// - Line 4, Col 45: "teh" appears to be a typo
// - Line 8, Col 4: "Borwsre" appears to be a typo
//
// How do we know that if we see the file in a dictionary that looks like "the" and then see a word that looks
// like "teh" that it's actually a typo?
