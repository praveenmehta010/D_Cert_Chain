const hre = require("hardhat");

async function main() {
  const Certificate = await hre.ethers.getContractFactory("Certificate");
  const certificate = await Certificate.deploy(); // deploy the contract

  await certificate.waitForDeployment(); // instead of `.deployed()` (new in Ethers v6)

  console.log("Certificate deployed to:", await certificate.getAddress());
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
