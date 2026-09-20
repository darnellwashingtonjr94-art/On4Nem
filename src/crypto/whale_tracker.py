import logging

logging.basicConfig(level=logging.INFO, format="%(asctime)s [%(levelname)s] %(message)s")
logger = logging.getLogger("WhaleTracker")

class WhaleTrackingStrategy:
    def __init__(self, inflow_threshold_usd: float = 250000.0):
        self.inflow_threshold_usd = inflow_threshold_usd

    def check_wallet_flow(self, net_inflow_usd: float) -> bool:
        """Trigger Condition: Consolidated net inflow > $250,000 within a 5-minute window."""
        logger.info(f"Tracked Wallet Inflow: ${net_inflow_usd:,.2f} | Threshold: ${self.inflow_threshold_usd:,.2f}")
        return net_inflow_usd >= self.inflow_threshold_usd
