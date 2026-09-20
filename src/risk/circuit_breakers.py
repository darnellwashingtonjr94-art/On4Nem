import time
import logging
from typing import Dict, Any

logging.basicConfig(level=logging.INFO, format="%(asctime)s [%(levelname)s] %(message)s")
logger = logging.getLogger("On4NemRiskManager")

class RiskManager:
    def __init__(
        self, 
        max_portfolio_drawdown_pct: float = 0.05, 
        max_consecutive_losses: int = 4,
        max_single_trade_exposure_usd: float = 25000.0,
        cooldown_period_seconds: int = 900
    ):
        """
        Initializes the Risk Management and Circuit Breaker Engine.
        
        :param max_portfolio_drawdown_pct: Hard stop if portfolio drops by this fraction (e.g., 5%)
        :param max_consecutive_losses: Trip circuit breaker after X losing trades in a row
        :param max_single_trade_exposure_usd: Maximum allowed capital allocation per single order
        :param cooldown_period_seconds: Lockout duration after a circuit breaker trip (default: 15 mins)
        """
        self.max_drawdown_pct = max_portfolio_drawdown_pct
        self.max_consecutive_losses = max_consecutive_losses
        self.max_single_trade_exposure_usd = max_single_trade_exposure_usd
        self.cooldown_period_seconds = cooldown_period_seconds

        # State tracking
        self.consecutive_losses = 0
        self.peak_portfolio_value = 0.0
        self.current_portfolio_value = 0.0
        self.is_tripped = False
        self.trip_timestamp = 0.0

    def update_portfolio_value(self, current_value: float):
        """Updates portfolio valuation and monitors peak-to-valley drawdowns."""
        self.current_portfolio_value = current_value
        if current_value > self.peak_portfolio_value:
            self.peak_portfolio_value = current_value

        # Check drawdown condition
        if self.peak_portfolio_value > 0:
            drawdown = (self.peak_portfolio_value - self.current_portfolio_value) / self.peak_portfolio_value
            if drawdown >= self.max_drawdown_pct:
                self.trip_circuit_breaker(f"Max portfolio drawdown reached: {drawdown * 100:.2f}%")

    def record_trade_result(self, pnl: float):
        """Records the outcome of a closed trade to evaluate consecutive loss metrics."""
        if pnl < 0:
            self.consecutive_losses += 1
            logger.warning(f"Trade resulted in a loss. Consecutive losses: {self.consecutive_losses}")
            if self.consecutive_losses >= self.max_consecutive_losses:
                self.trip_circuit_breaker(f"Max consecutive losses limit reached ({self.consecutive_losses})")
        else:
            # Reset consecutive losses on a winning or breakeven trade
            self.consecutive_losses = 0

    def trip_circuit_breaker(self, reason: str):
        """Activates the global safety lockout."""
        self.is_tripped = True
        self.trip_timestamp = time.time()
        logger.critical(f"🚨 CIRCUIT BREAKER TRIPPED: {reason}. Trading halted for {self.cooldown_period_seconds / 60} minutes.")

    def check_trade_allowed(self, proposed_exposure_usd: float) -> bool:
        """
        Validates whether a new trade is permitted based on current system health,
        exposure caps, and active cooling periods.
        """
        # 1. Check if circuit breaker is active
        if self.is_tripped:
            elapsed_time = time.time() - self.trip_timestamp
            if elapsed_time < self.cooldown_period_seconds:
                logger.warning(f"Trade blocked: Circuit breaker active. Cooldown remaining: {(self.cooldown_period_seconds - elapsed_time):.0f}s")
                return False
            else:
                # Cooldown elapsed, reset system
                logger.info("✅ Cooldown period elapsed. Resetting circuit breaker.")
                self.is_tripped = False
                self.consecutive_losses = 0

        # 2. Check single trade exposure limit
        if proposed_exposure_usd > self.max_single_trade_exposure_usd:
            logger.warning(f"Trade blocked: Proposed exposure ${proposed_exposure_usd:.2f} exceeds max cap of ${self.max_single_trade_exposure_usd:.2f}")
            return False

        return True
