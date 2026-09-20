import logging

logging.basicConfig(level=logging.INFO, format="%(asctime)s [%(levelname)s] %(message)s")
logger = logging.getLogger("PEADMomentum")

class PEADMomentumStrategy:
    def __init__(self, sue_threshold: float = 1.5):
        self.sue_threshold = sue_threshold

    def evaluate_earnings(self, sue_score: float, guidance_revised_up: bool) -> bool:
        """Trigger Condition: SUE score > +1.5 with upward revenue guidance revisions."""
        logger.info(f"Evaluating PEAD -> SUE Score: {sue_score} | Guidance Positive: {guidance_revised_up}")
        return sue_score > self.sue_threshold and guidance_revised_up
