import os
import logging
import httpx
from typing import Dict, Any, List

logging.basicConfig(level=logging.INFO, format="%(asctime)s [%(levelname)s] [%(name)s] %(message)s")
logger = logging.getLogger("PolymarketService")

class PolymarketClient:
    def __init__(self, host: str = "https://clob.polymarket.com"):
        """
        Initializes the Polymarket prediction market API client.
        """
        self.host = host
        self.gamma_url = "https://gamma-api.polymarket.com"
        self.client = httpx.Client(timeout=10.0)

    def get_active_markets(self, limit: int = 10) -> List[Dict[str, Any]]:
        """Fetches active prediction markets and sports betting outcomes from Gamma API."""
        try:
            response = self.client.get(f"{self.gamma_url}/markets", params={"active": "true", "limit": limit})
            response.raise_for_status()
            markets = response.json()
            logger.info(f"Successfully fetched {len(markets)} active Polymarket prediction feeds.")
            return markets
        except Exception as e:
            logger.error(f"Error fetching Polymarket data: {e}")
            return []

    def get_market_price(self, token_id: str) -> Dict[str, Any]:
        """Queries the current buy price and midpoint for a specific outcome token."""
        try:
            price_resp = self.client.get(f"{self.host}/price", params={"token_id": token_id, "side": "BUY"})
            mid_resp = self.client.get(f"{self.host}/midpoint", params={"token_id": token_id})
            
            return {
                "token_id": token_id,
                "buy_price": float(price_resp.json().get("price", 0.0)) if price_resp.status_code == 200 else 0.0,
                "midpoint": float(mid_resp.json().get("mid", 0.0)) if mid_resp.status_code == 200 else 0.0
            }
        except Exception as e:
            logger.error(f"Failed to query market pricing for token {token_id}: {e}")
            return {}

if __name__ == "__main__":
    poly = PolymarketClient()
    active_markets = poly.get_active_markets(limit=3)
    for m in active_markets:
        print(f"Market: {m.get('question')} | Volume: ${m.get('volume', 0)}")
