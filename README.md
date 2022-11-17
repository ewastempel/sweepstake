# sweepstake
Simple GO program that randomly assigns countries to the sweepstake participants.

## How it works
The sweepstake program uses two json data inputs:
* `data/countries.json` which holds information on countries that participate in the World Cup 2022.
* `data/participants.json` which holds information on people that participate in the sweepstake.

The sweepstake then produces a map of participants and the countries that have been randomly assigned to them.
NOTE: it is possible to buy multiple bets and therefore a number of countries (equal to a number of bets) can be assigned to each participant.
