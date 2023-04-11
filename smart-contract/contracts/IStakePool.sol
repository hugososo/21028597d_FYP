// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

interface IStakePool {
    event Staked(address indexed user, uint256 amount);
    event Withdrawn(address indexed user, uint256 amount);
    event RewardClaimed(address indexed user, uint256 amount);

    function stake(uint256 amount) external;

    function withdraw(uint256 amount) external;

    function claimReward() external;

    function calculateReward(address user) external view returns (uint256);

    function isStaker(address wallet) external view returns (bool);
}
