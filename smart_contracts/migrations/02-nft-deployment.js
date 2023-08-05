const CryptoCert = artifacts.require("CryptoCert")

module.exports = function(deployer){
    deployer.deploy(CryptoCert);
}