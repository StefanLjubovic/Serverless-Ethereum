// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";

contract MyToken is ERC721URIStorage  {
    constructor() ERC721("MyToken", "MTK") {}

    function safeMint(address to, uint256 tokenId, string memory tokenURI) public {
        _safeMint(to, tokenId);
         _setTokenURI(tokenId, tokenURI);
    }
}