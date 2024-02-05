# Consensus Breaking Mechanism

Consensus in blockchain refers to the process on how nodes across the system agree on the state and validity of transactions. A consensus-breaking change would require all nodes to upgrade their software to the latest version, otherwise they would not be able to communicate with each other and attach new blocks.

In this case, we initially created a blockchain called match that contained information about football games:
1. Date
2. Home (Home Team)
3. Away (Away Team)

Then as for the consensus-breaking portion, I decided to add more relevant information with regards to a football game: 
- Stadium (Name of the staidum)
- Specatator (Number of spectators to attend the game)

As a result, when we try to add a new block onto the chain, the transaction would not occur. This is becasue the nodes running the older software would not agree to the new fields and reject the changes.
