// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract NFT {
    address public minter;
    mapping(address => string[]) public tokens;

    constructor() {
        minter = msg.sender;
    }

    event Transfer(
        address indexed from,
        address indexed to,
        string indexed tokenId
    );
    event Mint(address indexed to, string indexed tokenId);

    function mint(string memory _tokenId) public {
        require(msg.sender == minter, "Only minter can mint tokens");

        tokens[minter].push(_tokenId);

        emit Mint(minter, _tokenId);
    }

    function getTokens() public view returns (string[] memory) {
        return tokens[minter];
    }

    function send(string memory _tokenId, address _to) public {
        require(msg.sender == minter, "Only minter can transfer tokens");
        tokens[_to].push(_tokenId);

        emit Transfer(minter, _to, _tokenId);
    }
}
