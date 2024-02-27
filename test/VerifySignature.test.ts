import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { expect } from "chai";
import { ethers } from "hardhat";

describe("VerifySignature", () => {
  const deployContract = async () => {
    const sig = await ethers.getSigners();

    // address for private key fe5e321af0d82edab25abed54f841e575953465e42ce7a51e2008f8807c5582b
    const con = await (await ethers.getContractFactory("VerifySignature")).deploy("0xe1DA85a0E14c901DCc005602aFA24266b35BBE30");

    return { sig, con };
  }

  it("should verify signature", async () => {
    const { sig, con } = await loadFixture(deployContract);

    const ver = await con.connect(sig[4]).verifySignature(7, 1, 700, "0x063e956331f6b03a51aef0e7ac4065ab9e7339b943535a0a1458953436478d995ebabbcfe29b6f93c02ee9fda5954a6290469b758e0df0a08c1c47a630cd6d391c");

    expect(ver).to.equal(true);
  });
});
