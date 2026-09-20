import logging

logging.basicConfig(level=logging.INFO, format="%(asctime)s [%(levelname)s] %(message)s")
logger = logging.getLogger("ForexGridMaker")

class ForexGridMarketMaker:
    def __init__(self, max_spread_pips: float = 0.5):
        self.max_spread_pips = max_spread_pips

    def validate_grid_conditions(self, current_spread_pips: float, is_asian_session: bool) -> bool:
        """Trigger Condition: Low volatility range consolidation during Asian sessions with spread < 0.5 pips."""
        return is_asian_session and current_spread_pips < self.max_spread_pips
