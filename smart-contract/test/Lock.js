const {
  time,
  loadFixture,
} = require("@nomicfoundation/hardhat-network-helpers");
const { anyValue } = require("@nomicfoundation/hardhat-chai-matchers/withArgs");
const { expect } = require("chai");

describe("LoreToken", function () {
  // We define a fixture to reuse the same setup in every test.
  // We use loadFixture to run this setup once, snapshot that state,
  // and reset Hardhat Network to that snapshot in every test.
  async function deployLoreToken() {
    // Contracts are deployed using the first signer/account by default
    const [owner, otherAccount] = await ethers.getSigners();

    const LoreToken = await ethers.getContractFactory("LoreToken");
    const loreToken = await LoreToken.deploy();

    return { loreToken, owner, otherAccount };
  }

  describe("Deployment", function () {
    it("Should set the right owner", async function () {
      const { loreToken, owner } = await loadFixture(deployLoreToken);
      expect(await loreToken.owner()).to.equal(owner.address);
    });
  });

  describe("Events", function () {
    it("Should emit an event on withdrawals", async function () {
      const { lock, unlockTime, lockedAmount } = await loadFixture(
        deployOneYearLockFixture
      );

      await time.increaseTo(unlockTime);

      await expect(lock.withdraw())
        .to.emit(lock, "Withdrawal")
        .withArgs(lockedAmount, anyValue); // We accept any value as `when` arg
    });
  });

  describe("Transfers", function () {
    it("Should transfer the funds to the owner", async function () {
      const { lock, unlockTime, lockedAmount, owner } = await loadFixture(
        deployOneYearLockFixture
      );

      await time.increaseTo(unlockTime);

      await expect(lock.withdraw()).to.changeEtherBalances(
        [owner, lock],
        [lockedAmount, -lockedAmount]
      );
    });
  });
});
