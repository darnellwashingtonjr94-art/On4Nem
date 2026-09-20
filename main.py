import asyncio
import logging
from src.risk.circuit_breakers import RiskManager
from src.crypto.arbitrage import CrossDEXArbitrageEngine
from src.stocks.pairs_trading import CointegratedPairsStrategy
from src.forex.session_breakout import SessionBreakoutStrategy

# Configure global system logging
logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s [%(levelname)s] [%(name)s] %(message)s"
)
logger = logging.getLogger("On4NemMain")

class On4NemExecutionEngine:
    def __init__(self):
        logger.info("Initializing On4Nem Multi-Asset Algorithmic Engine...")
        
        # 1. Global Risk Management & Circuit Breakers (5% max drawdown, 4 consecutive losses limit)
        self.risk_manager = RiskManager(
            max_portfolio_drawdown_pct=0.05,
            max_consecutive_losses=4,
            max_single_trade_exposure_usd=25000.0,
            cooldown_period_seconds=900
        )
        
        # 2. Strategy Modules Initialization
        self.crypto_arbitrage = CrossDEXArbitrageEngine(
            rpc_url="https://eth-mainnet.g.alchemy.com/v2/YOUR-API-KEY",
            min_buffer_pct=0.0015
        )
        self.stock_pairs = CointegratedPairsStrategy(
            ticker_a='KO', 
            ticker_b='PEP', 
            window=30, 
            z_threshold=2.0
        )
        self.forex_breakout = SessionBreakoutStrategy(
            atr_multiplier=1.5, 
            risk_reward_ratio=2.5
        )

    async def run_preflight_checks(self) -> bool:
        """Executes automated risk validations before launching trading loops."""
        logger.info("Running pre-flight system health checks...")
        
        # Validate through risk manager circuit breaker
        can_trade = self.risk_manager.check_trade_allowed(proposed_exposure_usd=10000.0)
        if not can_trade:
            logger.warning("Pre-flight check failed: Circuit breaker lock is engaged.")
            return False
            
        logger.info("✅ Pre-flight validation passed. All systems operational.")
        return True

    async def start(self):
        """Starts the main asynchronous multi-asset orchestration loops."""
        is_ready = await self.run_preflight_checks()
        if not is_ready:
            logger.error("Engine startup aborted due to active risk controls.")
            return

        logger.info("🚀 On4Nem framework is live and listening to multi-asset data feeds...")

        try:
            # Run asynchronous crypto scanning loop as the primary engine task
            await self.crypto_arbitrage.run_scan_loop(token_pair="MON/USDT", trade_size_usd=10000.0)
            
        except Exception as e:
            logger.critical(f"Critical runtime exception caught in main execution loop: {e}")
            self.risk_manager.trip_circuit_breaker(f"Runtime exception: {str(e)}")

if __name__ == "__main__":
    engine = On4NemExecutionEngine()
    try:
        asyncio.run(engine.start())
    except KeyboardInterrupt:
        logger.info("On4Nem Engine shut down cleanly by user request.")
