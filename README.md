# On4Nem

![Go](https://img.shields.io/badge/Go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Gemini](https://img.shields.io/badge/Gemini%20Pro-8E75B2?style=for-the-badge&logo=googlebard&logoColor=white)
![Monad](https://img.shields.io/badge/Monad_10k_TPS-8A2BE2?style=for-the-badge&logo=web3&logoColor=white)
![Polymarket](https://img.shields.io/badge/Polymarket-000000?style=for-the-badge&logo=ethereum&logoColor=white)
![GitHub Actions](https://img.shields.io/badge/GitHub_Actions-2088FF?style=for-the-badge&logo=github-actions&logoColor=white)

## What is this?
On4Nem is a smart robot program that automatically makes trades for you. It looks at sports bets and prediction markets to find the best chances to make money. 

## What is this about?
This project is about moving faster than a human ever could. It looks for "arbitrage," which is a fancy word for finding a guaranteed win by betting on all sides across different websites. It watches where the big money (the "whales") is going and copies those winning moves every single day.

## What is each asset?
*   **The Brain (Gemini Pro):** A super-smart artificial intelligence that reads data and makes plans.
*   **The Engine (Monad EVM):** A lightning-fast computer network (blockchain) that processes actions instantly.
*   **The Stores (Polymarket & Sportsbooks):** The websites where people place bets on sports or future events.

## What each asset does?
*   **Gemini Pro** acts like a detective. It looks at all the odds, reads the numbers, and decides exactly *what* trade to make.
*   **Monad EVM** acts like a racecar. Once the brain decides what to do, Monad makes the trade happen in a fraction of a second, before anyone else can beat you to it.
*   **Polymarket & Sportsbooks** are the places the robot goes to buy the "tickets" and collect the profits.

## Why is this cool?
It is cool because it takes human feelings out of trading! Instead of guessing who will win a game, the robot uses pure math to find guaranteed small wins and stacks them up all day long. It does all of this completely on its own, even while you sleep.

## What problems this solves?
*   **Being too slow:** Humans take minutes to click buttons; this robot takes milliseconds.
*   **Losing track:** A human can't watch 100 sports games at once. This robot watches everything on the internet at the exact same time.
*   **Emotional mistakes:** Humans get angry or greedy and make bad bets. The robot never gets emotional and only follows the math.

## Requirements?
To run this robot on your computer, you will need:
*   **Go:** The programming language installed on your computer.
*   **API Keys:** Secret passwords for Gemini Pro and your prediction markets so the robot can log in.
*   **Monad Wallet:** A digital wallet connected to the Monad network to hold your funds.

## How to install this?
Open your computer's terminal (the black typing screen) and follow these simple steps:

1. Download the code to your computer:
   `git clone https://github.com/darnellwashingtonjr94-art/On4Nem.git`
2. Go into the new folder:
   `cd On4Nem`
3. Download the extra pieces the robot needs to work:
   `go mod download`
4. Turn the robot on:
   `go run cmd/agent-core/main.go`

## Project Structure Tree Diagram

```text
On4Nem/
├── .github/workflows/   # Instructions for GitHub to automatically test the code
├── cmd/
│   └── agent-core/      # The main starting button to turn the robot on
├── config/              # The rulebook and saved settings for the robot
├── infrastructure/      # Blueprints for setting up the internet servers
├── pkg/
│   ├── execution/       # The robotic hands that actually place the bets
│   └── (other folders)  # Extra brain pieces and math tools
├── scripts/             # Helpful shortcuts to make building the robot easier
├── test/                # Practice tests to make sure the robot isn't broken
├── .gitignore           # Tells the computer which secret files to hide
├── LICENSE              # The legal rules for sharing this code with others
└── README.md            # The instruction manual you are reading right now!
