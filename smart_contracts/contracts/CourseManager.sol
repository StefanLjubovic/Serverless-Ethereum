//SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import "./Course.sol";

contract CourseManager {
    mapping(address => mapping(uint => address)) courses;

    constructor(){}

    function deployCourse(uint courseId, uint coursePrice) public {
        address newCourse = address(new Course(msg.sender,coursePrice, courseId));
        courses[msg.sender][courseId] = newCourse;
    }

    function getCourse(address owner, uint courseId) external  view returns (address) {
        return courses[owner][courseId];
    }
}