import time
import requests
import logging
from web3 import Web3

logging.basicConfig(level=logging.INFO, format="%(asctime)s [%(levelname)s] [%(name)s] %(message)s")
logger = logging.getLogger("MonadRPCManager")

class MonadRPCManager:
    def __init__(self, rpc_endpoints: list[str] = None):
        """
        Manages high-performance RPC endpoints for the Monad EVM layer.
        """
        self.rpc_endpoints = rpc_endpoints or [
            "https://rpc.monad.xyz",
            "https://testnet-rpc.monad.xyz",
            "https://rpc-testnet.monadinfra.com"
        ]
        self.active_w3: Web3 | None = None
        self.current_url: str | None = None

    def get_fastest_rpc(self) -> str:
        """Benchmarks available Monad RPC endpoints and returns the lowest latency node."""
        best_rpc = None
        min_latency = float('inf')

        for rpc in self.rpc_endpoints:
            try:
                start_time = time.time()
                response = requests.post(
                    rpc,
                    json={"jsonrpc": "2.0", "method": "eth_blockNumber", "params": [], "id": 1},
                    timeout=2.0
                )
                if response.status_code == 200:
                    latency = time.time() - start_time
                    logger.debug(f"RPC {rpc} responded in {latency*1000:.2f}ms")
                    if latency < min_latency:
                        min_latency = latency
                        best_rpc = rpc
            except Exception as e:
                logger.warning(f"RPC endpoint {rpc} failed health check: {e}")

        if not best_rpc:
            raise ConnectionError("All configured Monad RPC nodes are unreachable.")
        
        logger.info(f"⚡ Selected optimal Monad RPC: {best_rpc} (Latency: {min_latency*1000:.2f}ms)")
        return best_rpc

    def connect(self) -> Web3:
        """Initializes and returns an active Web3 connection to the fastest Monad RPC."""
        self.current_url = self.get_fastest_rpc()
        self.active_w3 = Web3(Web3.HTTPProvider(self.current_url))
        
        if self.active_w3.is_connected():
            logger.info(f"Successfully established Web3 connection to chain ID: {self.active_w3.eth.chain_id}")
        else:
            raise ConnectionError(f"Failed to connect to Web3 provider at {self.current_url}")
            
        return self.active_w3

if __name__ == "__main__":
    manager = MonadRPCManager()
    w3_client = manager.connect()
    print(f"Latest Block Number: {w3_client.eth.block_number}")
