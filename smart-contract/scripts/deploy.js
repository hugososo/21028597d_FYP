// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// You can also run a script with `npx hardhat run <script>`. If you do that, Hardhat
// will compile your contracts, add the Hardhat Runtime Environment's members to the
// global scope, and execute the script.
const hre = require("hardhat");

async function main() {


  const LoreToken = await hre.ethers.getContractFactory("LoreToken");
  const loreToken = await LoreToken.deploy();
  await loreToken.deployed();
  console.log(
    `LoreToken deployed to ${loreToken.address}`
  );


  const StakePool = await hre.ethers.getContractFactory("StakePool");
  const stakePool = await StakePool.deploy(loreToken.address, "50000000000000000000");
  await stakePool.deployed();
  console.log(
    `StakePool deployed to ${stakePool.address}`
  );

  const DebookDAO = await hre.ethers.getContractFactory("DebookDAO");
  const debookDAO = await DebookDAO.deploy(stakePool.address);
  await debookDAO.deployed();
  console.log(
    `DebookDAO deployed to ${loreToken.address}`
  );

  const ERC721Factory = await hre.ethers.getContractFactory("ERC721Factory");
  const erc721Factory = await ERC721Factory.deploy();
  await erc721Factory.deployed();
  console.log(
    `DebookDAO deployed to ${erc721Factory.address}`
  );
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
