// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

struct Token {
    string id;
    string name;
    uint64 price;
    string imageUrl;
}

contract NFT {
    address public minter;
    mapping(address => mapping(string => Token)) public tokens;
    Token[] public tokenList;

    constructor() {
        minter = msg.sender;
    }

    event Transfer(
        address indexed from,
        address indexed to,
        string indexed tokenId
    );
    event Mint(address indexed to, string indexed tokenId);

    function mint(
        string memory _id,
        string memory _name,
        uint64 _price,
        string memory _imageUrl
    ) public {
        require(msg.sender == minter, "Only minter can mint tokens");

        tokens[minter][_id] = Token(_id, _name, _price, _imageUrl);
        tokenList.push(tokens[minter][_id]);

        emit Mint(minter, _id);
    }

    function getTokens() public view returns (Token[] memory) {
        return tokenList;
    }

    function send(string memory _id, address _to) public {
        require(msg.sender == minter, "Only minter can transfer tokens");

        tokens[_to][_id] = tokens[minter][_id];
        delete tokens[minter][_id];

        emit Transfer(minter, _to, _id);
    }
}
