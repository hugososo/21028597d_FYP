// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Snapshot.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/draft-ERC20Permit.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol";

contract LoreToken is
    ERC20,
    ERC20Burnable,
    ERC20Snapshot,
    Ownable,
    ERC20Permit,
    ERC20Votes
{
    /// @notice Address of the pool
    address public pool;

    constructor() ERC20("LoreToken", "LORE") ERC20Permit("LoreToken") {
        _mint(msg.sender, 3000000000 * 10 ** uint256(decimals()));
    }

    /// @notice Reverts if the caller is not a pool.
    modifier onlyOwnerOrPool() {
        require(
            msg.sender == pool || msg.sender == owner(),
            "LoreToken::OnlyPool: not calling from pool"
        );
        _;
    }

    function snapshot() public onlyOwnerOrPool {
        _snapshot();
    }

    function mint(address to, uint256 amount) public onlyOwnerOrPool {
        _mint(to, amount);
    }

    function setPool(address _pool) public onlyOwner {
        require(
            _pool != address(0),
            "LoreToken::constructor: address can't be zero"
        );
        pool = _pool;
    }

    // The following functions are overrides required by Solidity.

    function _beforeTokenTransfer(
        address from,
        address to,
        uint256 amount
    ) internal override(ERC20, ERC20Snapshot) {
        super._beforeTokenTransfer(from, to, amount);
    }

    function _afterTokenTransfer(
        address from,
        address to,
        uint256 amount
    ) internal override(ERC20, ERC20Votes) {
        super._afterTokenTransfer(from, to, amount);
    }

    function _mint(
        address to,
        uint256 amount
    ) internal override(ERC20, ERC20Votes) {
        super._mint(to, amount);
    }

    function _burn(
        address account,
        uint256 amount
    ) internal override(ERC20, ERC20Votes) {
        super._burn(account, amount);
    }
}
