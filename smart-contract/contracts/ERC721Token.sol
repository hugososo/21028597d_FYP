// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/security/Pausable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

contract ERC721Token is ERC721, ERC721Enumerable, Pausable, Ownable {
    using Counters for Counters.Counter;
    Counters.Counter private _tokenIdCounter;
    using SafeERC20 for IERC20;
    IERC20 public loreToken;

    uint256 public cost;
    uint256 public maxSupply;
    uint256 public maxMintAmountPerTx;
    string public customURI;

    constructor(
        string memory _tokenName,
        string memory _tokenSymbol,
        string memory _customURI,
        uint256 _cost,
        uint256 _maxSupply,
        uint256 _maxMintAmountPerTx
    ) ERC721(_tokenName, _tokenSymbol) {
        loreToken = IERC20(0x8dB4b10b718b097346e75720bBD2B36bB49d3723);
        customURI = _customURI;
        cost = _cost;
        maxSupply = _maxSupply;
        maxMintAmountPerTx = _maxMintAmountPerTx;
        transferOwnership(tx.origin);
    }

    modifier mintCompliance(uint256 _mintAmount) {
        require(
            _mintAmount > 0 && _mintAmount <= maxMintAmountPerTx,
            "Invalid mint amount!"
        );
        require(
            totalSupply() + _mintAmount <= maxSupply,
            "Max supply exceeded!"
        );
        _;
    }

    modifier mintPriceCompliance(uint256 _mintAmount) {
        require(
            loreToken.allowance(_msgSender(), address(this)) >=
                cost * _mintAmount,
            "Insufficient allowance!"
        );
        require(
            loreToken.balanceOf(_msgSender()) >= cost * _mintAmount,
            "You don't have enough LoreToken!"
        );
        _;
    }

    function mint() public mintCompliance(1) mintPriceCompliance(1) {
        require(!paused(), "Sale paused");
        uint256 supply = totalSupply();
        loreToken.transferFrom(msg.sender, owner(), cost);
        _safeMint(_msgSender(), supply + 1);
    }

    function pause() public onlyOwner {
        _pause();
    }

    function unpause() public onlyOwner {
        _unpause();
    }

    function safeMint(address to) public onlyOwner {
        uint256 tokenId = _tokenIdCounter.current();
        _tokenIdCounter.increment();
        _safeMint(to, tokenId);
    }

    function withdraw() public onlyOwner {
        (bool os, ) = payable(owner()).call{value: address(this).balance}("");
        require(os);
    }

    function _beforeTokenTransfer(
        address from,
        address to,
        uint256 tokenId,
        uint256 batchSize
    ) internal override(ERC721, ERC721Enumerable) whenNotPaused {
        super._beforeTokenTransfer(from, to, tokenId, batchSize);
    }

    // override methods
    function tokenURI(
        uint256 tokenId
    ) public view virtual override returns (string memory) {
        require(
            _exists(tokenId),
            "ERC721Metadata: URI query for nonexistent token"
        );
        string memory baseURI = _baseURI();
        return
            bytes(baseURI).length > 0 ? string(abi.encodePacked(baseURI)) : "";
    }

    function _baseURI() internal view virtual override returns (string memory) {
        return customURI;
    }

    // The following functions are overrides required by Solidity.
    function supportsInterface(
        bytes4 interfaceId
    ) public view override(ERC721, ERC721Enumerable) returns (bool) {
        return super.supportsInterface(interfaceId);
    }
}
