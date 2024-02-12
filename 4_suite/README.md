Sometimes, a module would need some setups to be able to test.

Instead of setting up the same in all test sets, we'll set up a `suite`.

For a suite, we'll define
1. Setup (Arrange)
2. Call functions we want to test (Act)
3. See if result is as expected (Assert)
4. Teardown the setup for performance and repeatability (De-Arrange)

We usually call this, the `AAA` of testing.