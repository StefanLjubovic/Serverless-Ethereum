const CourseManager = artifacts.require("CourseManager");
const { BigNumber } = require("bignumber.js");
const Course = artifacts.require("Course");

contract("CourseManager", (accounts) => {
  it("Should create a new course and retrieve the address", async () => {
    const value = new BigNumber("1000000000000000000");
    const courseManagerInstance = await CourseManager.deployed();
    const owner = accounts[0];
    const courseId = 1;

    // Deploy the course
    await courseManagerInstance.deployCourse(courseId, value, { from: owner });

    // Get the course address
    const courseAddress = await courseManagerInstance.getCourse(owner, courseId);
    console.log(courseAddress);
    // Assert that the returned address is not zero
    assert.notEqual(courseAddress, "0x0000000000000000000000000000000000000000", "Invalid course address");
  });

//   it("should not have enough funds for transaction", async () => {
//     const value = new BigNumber("500000000000000000"); // Set a lower value than the course price
//     const courseManagerInstance = await CourseManager.deployed();
//     const owner = accounts[0];
//     const courseId = 1;
  
//     // Deploy the course
//     await courseManagerInstance.deployCourse(courseId, value, { from: owner });
  
//     // Get the course address
//     const courseAddress = await courseManagerInstance.getCourse(owner, courseId);
//     const lowerValue = new BigNumber("600000000000000000"); 
//     // Attempt to send a transaction with insufficient funds
//     try {
//         await web3.eth.sendTransaction({
//           from: accounts[1],
//           to: courseAddress,
//           value: lowerValue,
//         });
//         assert.fail("Transaction should have failed due to insufficient funds");
//       } catch (error) {
//         console.log(error.message)
//         assert.include(
//           error.message,
//           "Not enough funds for transaction.",
//           "Error message should contain 'Not enough funds for transaction.'"
//         );
//       }
//   });
it("Buy course", async () => {
    const courseManagerInstance = await CourseManager.deployed();
    const owner = accounts[0];
    const courseId = 1;
    const coursePriceInWei = new BigNumber("1000000000000000000"); // Set the course price
  
    // Deploy the course
    await courseManagerInstance.deployCourse(courseId, coursePriceInWei, { from: owner });
    const balanceBefore = await web3.eth.getBalance(accounts[1]);
  
    // Get the course address
    const courseAddress = await courseManagerInstance.getCourse(owner, courseId);
    const courseInstance = await Course.at(courseAddress);
  
    const fetchedCoursePriceInWei = await courseInstance.getCoursePriceInWei();
    console.log("Course price in Wei: " + fetchedCoursePriceInWei.toString());
  
    // Attempt to send a transaction with insufficient funds
    try {
      const transactionValue = coursePriceInWei;
      await web3.eth.sendTransaction({
        from: accounts[1],
        to: courseAddress,
        value: transactionValue,
      });
  
      const balanceAfter = await web3.eth.getBalance(accounts[1]);
      const difference = balanceBefore.sub(balanceAfter);
  
      console.log("Balance after:", balanceAfter.toString());
      console.log("Balance before:", balanceBefore.toString());
      console.log("Difference:", difference.toString());
  
      const isPurchased = await courseInstance.checkIfUserPurchased(accounts[1]);
      console.log("Is purchased:", isPurchased);
      assert.equal(isPurchased, true, "User should have purchased the course");
    } catch (error) {
      console.log(error);
    }
  });
  
  

});
