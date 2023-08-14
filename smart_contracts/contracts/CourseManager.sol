//SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "./Course.sol";

contract CourseManager {
    mapping(address => mapping(uint => address)) courses;
    mapping(uint => address) private courseOwners;
    constructor(){}

    function deployCourse(uint courseId, uint coursePrice) public {
        address newCourse = address(new Course(msg.sender,coursePrice, courseId));
        if (courseOwners[courseId] == address(0)) {
            // Only set the owner if it's not already set
            courseOwners[courseId] = msg.sender;
        }
        courses[msg.sender][courseId] = newCourse;
    }

    function getCourse(address owner, uint courseId) external  view returns (address) {
        return courses[owner][courseId];
    }

    function getOwnerByCourseId(uint courseId) external view returns (address) {
        return courseOwners[courseId];
    }
}