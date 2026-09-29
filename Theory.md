what not to do:
- avoid functional decomposition
	- Problems:
		- time decomposition
		- design adds no value
		- perform anti-design effort
		- maximizes the impact of change
what to do:
- Decompose based on volatility
- encapsulate the change to insulate
- do not resonate with change
- implement behavior as interaction between services or subsystems

Axes of Volatility:
- At the same customer over time
- At the same time across customers
Axes should be independent
Volatility decreases top to down
Reusability increases top to down

Within the Business layer, managers encapsulate volatility in the sequence in use cases and workflows, engines encapsulate volatility in business rules and activities, managers may use zero or more engines, and engines may be shared between managers.
