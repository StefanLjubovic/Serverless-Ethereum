// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

interface ICourse {
    function getCoursePriceInWei() external view returns (uint);
    function payInTokens(address) external  payable;
}

interface ICourseManager {
    function getCourse(address owner, uint courseId) external  view returns (address);
}

contract Token is ERC20 {

    uint tokenPriceInWeiPerUnit  = 0.001 ether;

    constructor() ERC20("Token", "TKN") {
    }

    event CoursePurchased(address indexed receiver,address indexed buyer);

    function rewardUser(address to, uint256 amount) public {
        _mint(to, amount* 10 ** decimals());
    }

    function buyCourse(address owner,uint courseId,uint amount,address courseManagerAddress) public{
        address course = ICourseManager(courseManagerAddress).getCourse(owner,courseId);
        uint courserPriceInWei = ICourse(course).getCoursePriceInWei();

        uint ammountInWei = tokenPriceInWeiPerUnit * amount;
        require(ammountInWei >=courserPriceInWei, "Not enough tokens");

        uint amountToReduce = courserPriceInWei / tokenPriceInWeiPerUnit;
        require(balanceOf(_msgSender()) >= amountToReduce * 10 ** decimals(), "Insufficient balance");

        _spendAllowance(owner, _msgSender(), amountToReduce* 10 ** decimals());
        _burn(_msgSender(), amountToReduce* 10 ** decimals());

        ICourse(course).payInTokens(_msgSender());
    }
}
