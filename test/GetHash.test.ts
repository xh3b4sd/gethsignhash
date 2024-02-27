import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { expect } from "chai";
import { ethers } from "hardhat";

describe("GetHash", () => {
  const deployContract = async () => {
    const sig = await ethers.getSigners();

    const con = await (await ethers.getContractFactory("GetHash")).deploy();

    return { sig, con };
  }

  it("should get message hash for 000", async () => {
    const { sig, con } = await loadFixture(deployContract);

    const has = await con.connect(sig[4]).getHash000(7);

    expect(has).to.equal("0xcbe76b843b0b2d4ee61809d1578b36268b773207965820ae20ebdaf978bf3afc");
  });

  it("should get message hash for 001", async () => {
    const { sig, con } = await loadFixture(deployContract);

    const has = await con.connect(sig[4]).getHash001("0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65");

    expect(has).to.equal("0x207cbaf2b538db65f2c33c1aaa3259d01c8caf9ed3d05efa3d4bdce10948cd76");
  });

  it("should get message hash for 002", async () => {
    const { sig, con } = await loadFixture(deployContract);

    const has = await con.connect(sig[4]).getHash002(7, 1);

    expect(has).to.equal("0xb92fe793e135cccdde4f009e010816d9470d47838f10c9f378379c7c6564fac1");
  });

  it("should get message hash for 003", async () => {
    const { sig, con } = await loadFixture(deployContract);

    const has = await con.connect(sig[4]).getHash003(7, 1, "0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65");

    expect(has).to.equal("0x0fcf785806e52fcea9fad680f96f13a8f4c5932ea6e248a2a05687a64190534b");
  });

  it("should get message hash for 004", async () => {
    const { sig, con } = await loadFixture(deployContract);

    const has = await con.connect(sig[4]).getHash004(7, 1, "0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65", 700);

    expect(has).to.equal("0xbc02743c0db4e578cea5caaf1a2f6705faa8e02ef58d14b4611555c22f6e4678");
  });

  it("should get message hash for 005", async () => {
    const { sig, con } = await loadFixture(deployContract);

    const has = await con.connect(sig[4]).getHash005(7, 1, "0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65", 700);

    expect(has).to.equal("0x4d84bf5afe05dff160c24b8f734d46a6260d39c21d4551bfa406baf9419aff53");
  });
});
