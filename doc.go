/*

Package dsupdate is a library for updating DS records with DK
Hostmasters proprietary DS Update protocol
(https://github.com/DK-Hostmaster/dsu-service-specification).

It is work in progress and not in a functional state (it doesn't even
try to update anything yet).

Besides eventually solving a use case for me the purpose of this
project is also me getting more familiar with Go, train TDD (which
I've always postponed) and try out some new techniques inspired by
among others:

- Functional options for friendly APIs by Dave Cheney (https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis)

- Vanity Go Import Paths by Andrew Brampton (https://blog.bramp.net/post/2017/10/02/vanity-go-import-paths/)

*/
package dsupdate
