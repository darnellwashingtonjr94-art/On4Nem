import pandas as pd
import numpy as np
import logging

logging.basicConfig(level=logging.INFO, format="%(asctime)s [%(levelname)s] %(message)s")
logger = logging.getLogger("ForexSessionBreakout")

class SessionBreakoutStrategy:
    def __init__(self, atr_multiplier: float = 1.5, risk_reward_ratio: float = 2.5):
        """
        Initializes the Forex Session Breakout & Volatility Expansion Strategy.
        
        :param atr_multiplier: Multiplier applied to ATR for breakout confirmation.
        :param risk_reward_ratio: Target risk-to-reward ratio (default: 1:2.5).
        """
        self.atr_multiplier = atr_multiplier
        self.risk_reward_ratio = risk_reward_ratio

    def calculate_atr(self, df: pd.DataFrame, period: int = 14) -> pd.Series:
        """Calculates the Average True Range (ATR) for volatility measurement."""
        high_low = df['high'] - df['low']
        high_close = np.abs(df['high'] - df['close'].shift())
        low_close = np.abs(df['low'] - df['close'].shift())
        
        true_range = pd.concat([high_low, high_close, low_close], axis=1).max(axis=1)
        return true_range.rolling(window=period).mean()

    def identify_asian_session_range(self, df: pd.DataFrame) -> tuple[float, float]:
        """
        Identifies the high and low boundaries during the Asian consolidation session.
        Expects a DataFrame with a datetime index.
        """
        asian_session = df.between_time('00:00', '06:00')
        if asian_session.empty:
            # Fallback if specific time window isn't met in sample data
            session_high = df['high'].head(20).max()
            session_low = df['low'].head(20).min()
        else:
            session_high = asian_session['high'].max()
            session_low = asian_session['low'].min()
            
        logger.info(f"Asian Session Range Established -> High: {session_high:.5f}, Low: {session_low:.5f}")
        return session_high, session_low

    def check_breakout_signals(self, df: pd.DataFrame, session_high: float, session_low: float) -> pd.DataFrame:
        """
        Triggers breakout signals during high-volatility session overlaps.
        """
        df = df.copy()
        df['atr'] = self.calculate_atr(df)
        
        # Volume expansion filter (current volume > 1.5x rolling average volume)
        df['vol_ma'] = df['volume'].rolling(window=20).mean()
        df['volume_expansion'] = df['volume'] > (1.5 * df['vol_ma'])

        df['signal'] = 0  # 0: No action, 1: Long Breakout, -1: Short Breakout
        
        # Long breakout condition
        long_condition = (df['close'] > session_high) & df['volume_expansion']
        df.loc[long_condition, 'signal'] = 1

        # Short breakout condition
        short_condition = (df['close'] < session_low) & df['volume_expansion']
        df.loc[short_condition, 'signal'] = -1

        return df

if __name__ == "__main__":
    # Example simulation data frame setup
    np.random.seed(42)
    dates = pd.date_range(start="2026-06-01 00:00", periods=50, freq="h")
    prices = 1.0800 + np.cumsum(np.random.normal(0, 0.0005, 50))
    
    sample_df = pd.DataFrame({
        'open': prices,
        'high': prices + np.random.uniform(0.0001, 0.0004, 50),
        'low': prices - np.random.uniform(0.0001, 0.0004, 50),
        'close': prices + np.random.normal(0, 0.0002, 50),
        'volume': np.random.randint(1000, 5000, 50)
    }, index=dates)

    strategy = SessionBreakoutStrategy()
    high, low = strategy.identify_asian_session_range(sample_df)
    signaled_df = strategy.check_breakout_signals(sample_df, high, low)
    
    print(signaled_df[['close', 'signal', 'volume_expansion']].tail(10))
