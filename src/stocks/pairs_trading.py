import numpy as np
import pandas as pd
import statsmodels.api as sm
from statsmodels.tsa.stattools import coint
import logging

logging.basicConfig(level=logging.INFO, format="%(asctime)s [%(levelname)s] %(message)s")
logger = logging.getLogger("PairsTradingStrategy")

class CointegratedPairsStrategy:
    def __init__(self, ticker_a: str, ticker_b: str, window: int = 30, z_threshold: float = 2.0):
        """
        Initializes the Cointegrated Pairs Trading Strategy.
        
        :param ticker_a: First stock ticker (Dependent variable Y, e.g., 'KO')
        :param ticker_b: Second stock ticker (Independent variable X, e.g., 'PEP')
        :param window: Rolling window for moving average and standard deviation (default: 30 days)
        :param z_threshold: Z-score threshold to trigger mean-reversion trade (default: 2.0)
        """
        self.ticker_a = ticker_a
        self.ticker_b = ticker_b
        self.window = window
        self.z_threshold = z_threshold

    def test_cointegration(self, price_a: pd.Series, price_b: pd.Series) -> tuple[bool, float]:
        """
        Performs the Engle-Granger two-step cointegration test.
        :return: (is_cointegrated (bool), p_value (float))
        """
        score, p_value, _ = coint(price_a, price_b)
        is_coint = p_value < 0.05
        logger.info(f"Cointegration Test ({self.ticker_a} vs {self.ticker_b}): p-value = {p_value:.5f} | Cointegrated: {is_coint}")
        return is_coint, p_value

    def calculate_half_life(self, spread: pd.Series) -> float:
        """
        Calculates the half-life of mean reversion using an AR(1) autoregression model on the spread.
        """
        spread_lag = spread.shift(1).dropna()
        spread_diff = spread.diff().dropna()
        
        # Align series lengths
        df = pd.DataFrame({'diff': spread_diff, 'lag': spread_lag}).dropna()
        
        # OLS regression of spread diff on lagged spread
        model = sm.OLS(df['diff'], sm.add_constant(df['lag']))
        results = model.fit()
        
        # Lambda coefficient is the slope
        lam = results.params['lag']
        half_life = -np.log(2) / lam if lam < 0 else np.inf
        logger.info(f"Calculated Mean-Reversion Half-Life: {half_life:.2f} periods")
        return half_life

    def generate_signals(self, df_prices: pd.DataFrame) -> pd.DataFrame:
        """
        Calculates the hedge ratio, spread, rolling z-score, and generates trading signals.
        
        :param df_prices: DataFrame containing historical adjusted close prices for both tickers.
        :return: DataFrame with added spread, z-score, and signal columns.
        """
        df = df_prices.copy()
        
        # 1. Calculate OLS Hedge Ratio (Y = beta * X + alpha)
        X = sm.add_constant(df[self.ticker_b])
        model = sm.OLS(df[self.ticker_a], X).fit()
        alpha, hedge_ratio = model.params['const'], model.params[self.ticker_b]
        
        logger.info(f"Fitted Model - Alpha: {alpha:.4f}, Hedge Ratio (Beta): {hedge_ratio:.4f}")

        # 2. Compute the spread
        df['spread'] = df[self.ticker_a] - (hedge_ratio * df[self.ticker_b])

        # 3. Compute rolling mean and standard deviation of the spread
        df['spread_mean'] = df['spread'].rolling(window=self.window).mean()
        df['spread_std'] = df['spread'].rolling(window=self.window).std()

        # 4. Compute Z-Score
        df['z_score'] = (df['spread'] - df['spread_mean']) / df['spread_std']

        # 5. Generate Trading Signals based on Z-Score Trigger Thresholds
        # Signal: +1 = Long Spread (Long Ticker A, Short Ticker B)
        # Signal: -1 = Short Spread (Short Ticker A, Long Ticker B)
        # Signal:  0 = Exit / Neutral
        df['signal'] = 0
        df.loc[df['z_score'] <= -self.z_threshold, 'signal'] = 1   # Spread is too low, expect upward reversion
        df.loc[df['z_score'] >= self.z_threshold, 'signal'] = -1  # Spread is too high, expect downward reversion
        
        # Mean reversion exit condition: Z-score crosses back towards zero (-0.5 to 0.5)
        df.loc[(df['z_score'] > -0.5) & (df['z_score'] < 0.5), 'signal'] = 0

        return df

if __name__ == "__main__":
    # Example execution simulation
    dates = pd.date_range(start="2025-01-01", periods=100, freq="B")
    np.random.seed(42)
    
    # Synthetic correlated price data
    pep_prices = 150 + np.cumsum(np.random.normal(0, 1, 100))
    ko_prices = 0.5 * pep_prices + 20 + np.cumsum(np.random.normal(0, 0.5, 100))
    
    sample_data = pd.DataFrame({'KO': ko_prices, 'PEP': pep_prices}, index=dates)
    
    strategy = CointegratedPairsStrategy(ticker_a='KO', ticker_b='PEP', window=30, z_threshold=2.0)
    
    # Check cointegration
    is_coint, p_val = strategy.test_cointegration(sample_data['KO'], sample_data['PEP'])
    
    if is_coint:
        results_df = strategy.generate_signals(sample_data)
        strategy.calculate_half_life(results_df['spread'])
        print(results_df[['KO', 'PEP', 'spread', 'z_score', 'signal']].tail(10))
    else:
        logger.warning("Pair is not cointegrated. Trading strategy aborted for this pair.")
