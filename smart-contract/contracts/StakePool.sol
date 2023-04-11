// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/security/Pausable.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./IStakePool.sol";

contract StakePool is IStakePool, Pausable, Ownable, ReentrancyGuard {
    using SafeERC20 for IERC20;

    IERC20 public loreToken;
    uint256 public rewardRate;
    mapping(address => uint256) public stakedBalances;
    mapping(address => uint256) public stakingStartTimes;
    uint256 public DAOThreshold = 2000;

    constructor(IERC20 _loreToken, uint256 _rewardRate) {
        loreToken = _loreToken;
        rewardRate = _rewardRate;
    }

    function stake(uint256 amount) external override whenNotPaused {
        require(amount > 0, "Cannot stake 0 tokens");
        require(
            loreToken.allowance(msg.sender, address(this)) > 0,
            "You need to approve token usage"
        );
        loreToken.safeTransferFrom(msg.sender, address(this), amount);
        stakedBalances[msg.sender] += amount;
        stakingStartTimes[msg.sender] = block.timestamp;
        emit Staked(msg.sender, amount);
    }

    function withdraw(uint256 amount) external override {
        require(amount > 0, "Cannot withdraw 0 tokens");
        require(
            stakedBalances[msg.sender] >= amount,
            "Not enough tokens staked"
        );
        claimReward();
        stakedBalances[msg.sender] -= amount;
        loreToken.safeTransfer(msg.sender, amount);
        emit Withdrawn(msg.sender, amount);
    }

    function claimReward() public override {
        uint256 reward = calculateReward(msg.sender);
        require(reward > 0, "No rewards available");
        stakingStartTimes[msg.sender] = block.timestamp;
        loreToken.safeTransfer(msg.sender, reward);
        emit RewardClaimed(msg.sender, reward);
    }

    function calculateReward(
        address user
    ) public view override returns (uint256) {
        uint256 stakingDuration = getStakingDuration(user);
        uint256 earn = (stakedBalances[user] * stakingDuration * rewardRate) /
            (1 days * 365) /
            100;
        return earn;
    }

    function isStaker(address wallet) public view returns (bool) {
        return stakedBalances[wallet] >= DAOThreshold ** 18;
    }

    function getStakingDuration(address user) public view returns (uint256) {
        return block.timestamp - stakingStartTimes[user];
    }
}
