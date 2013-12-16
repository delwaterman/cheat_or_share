Author: V. Orion Delwaterman

### Cheat or Share

This is a personal project with two objectives.

1. Construct a simple simuation to see results of individuals that examines the conflict between idividuals to cooperate and compete against each other

2. Learn the Go language and get familiar with concurrency

### Rules

Initial sketch of rules

1. Initially groups start with 5 individuals

2. Each individual has an upper and lower bound between 1 and 10

3. Start with 5 groups

4. For each turn each group will exert some amount of energy

5. That will be total energy

6. The group that produces the least energy dies

7. The rest of the groups split the total energy left

8. Within each group
  a. the individual who produced the most energy dies
  b. the individual who produced the least energy reproduces

9. The group that has produced the most energy reproduces a whole new group
