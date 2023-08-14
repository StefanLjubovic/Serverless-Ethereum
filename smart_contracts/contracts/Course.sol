//SPDX-License-Identifier: MIT

pragma solidity ^0.8.9;

contract Course {
    address payable public owner;
    mapping (address =>bool) usersPurchased;
    uint coursePriceInWei;
    uint courseId;

    constructor(address _owner,uint _coursePrice,uint _courseId){
        owner = payable(_owner);
        coursePriceInWei = _coursePrice;
        courseId = _courseId;
    }

    modifier alreadyPurchased(){
        require(!usersPurchased[msg.sender], "Already purchased course");
        _;
    }

    receive() external payable alreadyPurchased {
        require(msg.value >= coursePriceInWei, "Not enough funds for transaction.");

         uint256 amountToRefund = msg.value - coursePriceInWei;
            if (amountToRefund > 0) {
            payable(msg.sender).transfer(amountToRefund);
         }

        owner.transfer(coursePriceInWei);
        usersPurchased[msg.sender] = true;
    }

    function checkIfUserPurchased(address sender) public view returns(bool){
         return usersPurchased[sender];
    } 

    function payInTokens(address sender) external  payable{
        usersPurchased[sender] = true;
    }

      function getCoursePriceInWei() external view returns (uint){
          return coursePriceInWei;
      }
}