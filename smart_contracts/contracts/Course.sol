//SPDX-License-Identifier: MIT

pragma solidity ^0.8.9;

contract Course {
  address payable public owner;
    uint coursePriceInWei;
    uint courseId;
    mapping (string => bool) usersPurchased;

    constructor(address _owner,uint _coursePrice,uint _courseId){
        owner = payable(_owner);
        coursePriceInWei = _coursePrice;
        courseId = _courseId;
    }

    modifier alreadyPurchased(string memory username){
        require(!usersPurchased[username], "Already purchased course");
        _;
    }

    receive() external payable {
       revert("Invalid usage of receive function");
    }

     function purchaseCourse(string memory username) public payable alreadyPurchased(username) {
        require(msg.value >= coursePriceInWei, "Not enough funds for transaction.");

        uint256 amountToRefund = msg.value - coursePriceInWei;
        if (amountToRefund > 0) {
            payable(msg.sender).transfer(amountToRefund);
        }

        owner.transfer(coursePriceInWei);
        usersPurchased[username] = true;
    }

    function checkIfUserPurchased(string memory sender) public view returns(bool){
         return usersPurchased[sender];
    } 

      function getCoursePriceInWei() external view returns (uint){
          return coursePriceInWei;
      }
}