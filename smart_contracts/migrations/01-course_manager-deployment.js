const CourseManager = artifacts.require("CourseManager")

module.exports = function(deployer){
    deployer.deploy(CourseManager);
}