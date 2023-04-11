require("@nomicfoundation/hardhat-toolbox");

// const ALCHEMY_API_KEY = "KEY";
// const GOERLI_PRIVATE_KEY = "YOUR GOERLI PRIVATE KEY";

module.exports = {
  solidity: "0.8.18",
  networks: {
    local: {
      url: "http://127.0.0.1:8545",
      accounts: ["9f199de615cc5c7c330ec0dd153b67f347387c2d01fa90def1f5417d3d6cf790"]
    },
    // goerli: {
    //   url: `https://eth-goerli.alchemyapi.io/v2/${ALCHEMY_API_KEY}`,
    //   accounts: [GOERLI_PRIVATE_KEY]
    // }
  },
  settings: {
    optimizer: {
      enabled: true,
      runs: 1000,
    },
  },
};
