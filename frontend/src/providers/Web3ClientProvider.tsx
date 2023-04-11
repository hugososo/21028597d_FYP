import { PropsWithChildren } from "react";
import { WagmiConfig, createClient, configureChains } from "wagmi";
import { localhost } from "wagmi/chains";
import { jsonRpcProvider } from "wagmi/providers/jsonRpc";

const { provider, webSocketProvider } = configureChains(
  [localhost],
  [
    jsonRpcProvider({
      rpc: (chain) => ({
        http: `http://127.0.0.1:8545`,
      }),
    }),
  ]
);

const client = createClient({
  autoConnect: true,
  provider,
  webSocketProvider,
});

const Web3ClientProvider = ({ children }: PropsWithChildren) => {
  return <WagmiConfig client={client}>{children}</WagmiConfig>;
};

export default Web3ClientProvider;
