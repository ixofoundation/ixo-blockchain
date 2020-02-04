# Functions Library
The Bonds Module is deployed with a built-in set of libraries for commonly-used algorithmic pricing and reserve functions. It also includes algorithmic application logic and features, such as *Augmented Bonding*. Additional functions can be added to the Library through SDK updates. This requires a formal process of governance to approve updates, to assure the integrity of these functions.

## Function Types
The following function types will be included in the standard Bonds SDK Module:
* Power (exponential)
* Logistic (sigmoidal)
* Constant Product (swapper)
Algorithmic Applications include:
* Alpha Bonds (Risk-adjusted bonding)
* Innovation Bonds (offers bond shareholders contingent rights to future IP rights and/or revenues)
* Impact Bonds (offers bond shareholders contingent rights to success-based outcomes payments and/or rewards)

### Power Function (exponential)

Pricing function:

<img alt="drawing" src="./img/power1.png" height="20"/>

Integral:

<img alt="drawing" src="./img/power2.png" height="40"/>

### Logistic Function (sigmoidal)

Pricing function:

<img alt="drawing" src="./img/sigmoid1.png" height="80"/>

Integral:

<img alt="drawing" src="./img/sigmoid2.png" height="55"/>

### Constant Product Function (swapper)

Reserve function:

<img alt="drawing" src="./img/swapper.png" height="20"/>
