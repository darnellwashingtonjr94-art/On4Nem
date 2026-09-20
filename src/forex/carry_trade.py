import logging

logging.basicConfig(level=logging.INFO, format="%(asctime)s [%(levelname)s] %(message)s")
logger = logging.getLogger("CarryTradeStrategy")

class MacroCarryTradeStrategy:
    def __init__(self, min_rate_diff_pct: float = 2.5):
        self.min_rate_diff_pct = min_rate_diff_pct

    def check_carry_condition(self, rate_differential: float, vix: float) -> bool:
        """Trigger Condition: Central bank rate diff >= 2.5% and VIX < 20."""
        return rate_differential >= self.min_rate_diff_pct and vix < 20
