import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { expect } from "chai";
import { ethers } from "hardhat";

describe("GetPack", () => {
  const deployContract = async () => {
    const sig = await ethers.getSigners();

    const con = await (await ethers.getContractFactory("GetPack")).deploy();

    return { sig, con };
  }

  it("should get message pack for 000", async () => {
    const { sig, con } = await loadFixture(deployContract);

    const has = await con.connect(sig[4]).getPack000(7);

    expect(has).to.equal("0x0000000000000000000000000000000000000000000000000000000000000007");
  });

  it("should get message pack for 001", async () => {
    const { sig, con } = await loadFixture(deployContract);

    const has = await con.connect(sig[4]).getPack001("0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65");

    expect(has).to.equal("0x15d34aaf54267db7d7c367839aaf71a00a2c6a65");
  });

  it("should get message pack for 002", async () => {
    const { sig, con } = await loadFixture(deployContract);

    const has = await con.connect(sig[4]).getPack002(7, 1);

    expect(has).to.equal("0x00000000000000000000000000000000000000000000000000000000000000070000000000000000000000000000000000000000000000000000000000000001");
  });

  it("should get message pack for 003", async () => {
    const { sig, con } = await loadFixture(deployContract);

    const has = await con.connect(sig[4]).getPack003(7, 1, "0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65");

    expect(has).to.equal("0x0000000000000000000000000000000000000000000000000000000000000007000000000000000000000000000000000000000000000000000000000000000115d34aaf54267db7d7c367839aaf71a00a2c6a65");
  });

  it("should get message pack for 004", async () => {
    const { sig, con } = await loadFixture(deployContract);

    const has = await con.connect(sig[4]).getPack004(7, 1, "0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65", 700);

    expect(has).to.equal("0x0000000000000000000000000000000000000000000000000000000000000007000000000000000000000000000000000000000000000000000000000000000115d34aaf54267db7d7c367839aaf71a00a2c6a6500000000000000000000000000000000000000000000000000000000000002bc");
  });

  it("should get message pack for 005", async () => {
    const { sig, con } = await loadFixture(deployContract);

    const has = await con.connect(sig[4]).getPack005(7, 1, "0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65", 700);

    expect(has).to.equal("0x00000000000000000000000000000000000000000000000000000000000000073a00000000000000000000000000000000000000000000000000000000000000013a15d34aaf54267db7d7c367839aaf71a00a2c6a653a00000000000000000000000000000000000000000000000000000000000002bc");
  });
});
