Sometimes, a module might be coupled with the database or external parties.

In the end, we should do integration or end-to-end testing, but doing unit tests on these modules have their benefits. Mainly, isolation and faster feedback to see if the logic flows are correct.

We will `mock` the behavior of the dependency-packed modules.