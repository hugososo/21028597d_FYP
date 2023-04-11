// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;
import "./ERC721Token.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract ERC721Factory {
    address[] public tokens;
    mapping(address => bool) public exists;

    // standard for DeBook detection
    event BookCreated(address owner, address tokenContract, string url);

    function deployERC721(
        string calldata _tokenName,
        string calldata _tokenSymbol,
        string calldata _customURI,
        uint256 _cost,
        uint256 _maxSupply,
        uint256 _maxMintAmountPerTx
    ) public returns (address) {
        ERC721Token newContract = new ERC721Token(
            _tokenName,
            _tokenSymbol,
            _customURI,
            _cost,
            _maxSupply,
            _maxMintAmountPerTx
        );
        address contractAddress = address(newContract);
        require(exists[contractAddress] == false, "Contract Existed");
        tokens.push(contractAddress);
        exists[contractAddress] = true;
        emit BookCreated(msg.sender, contractAddress, _customURI);
        return contractAddress;
    }

    function registryERC721Contract(
        address contractAddress,
        string calldata url
    ) public {
        require(exists[contractAddress] == false, "Contract Existed");
        tokens.push(contractAddress);
        exists[contractAddress] = true;
        emit BookCreated(msg.sender, contractAddress, url);
    }
}
