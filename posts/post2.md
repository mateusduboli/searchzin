* Preface

Changing companies and continent

* Benchmark why?

Since this is a study project is good to be able to understand why you implement things in the engine, so to see what are the benefits of X over Y.
This post will come for the different benchmarks scenarious and what measurments are important considering my current storage strategy.

* Setup

Using google compute engine for the tests, not much reasoning besides want to learnign better the platform and good cost.

** Created a project with a fixed budget as for me not to lose to much money

I'll be using a 2 core instance as my current code is not multi-thread yet. And another client machine wuth a more powerful core system as to really push our current code.

I'm also doing local az communication to not spend a lot of money on this, since it can be quite expensive depending on the amount of data I'm using.

For the current budget I'm using first 1 € to be very controlled with my expenses


** Automation

As this is not a production setup, I'll spend sometime with automation of the setup, first create a instance with an arbitrary release of my code, so I can compare different versions more easily.

- Automate the release creation with a free CI platform for it

