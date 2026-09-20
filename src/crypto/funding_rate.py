import logging

logging.basicConfig(level=logging.INFO, format="%(asctime)s [%(levelname)s] %(message)s")
logger = logging.getLogger("PerpFundingHarvest")

class PerpFundingHarvestStrategy:
    def __init__(self, target_apr: float = 0.12):
        self.target_apr = target_apr

    def evaluate_funding_rate(self, current_annualized_rate: float) -> bool:
        """Trigger Condition: Annualized funding rate yield > 12% APR."""
        logger.info(f"Evaluating funding rate: {current_annualized_rate * 100:.2f}% APR (Target: >{self.target_apr * 100}%)")
        return current_annualized_rate > self.target_apr

    def execute_harvest(self, asset: str, size_usd: float):
        logger.info(f"⚖️ Executing Delta-Neutral Funding Harvest for {asset} with ${size_usd:.2f} (Buy Spot, Short Perp).")
