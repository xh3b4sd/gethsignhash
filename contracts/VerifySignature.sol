// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";

contract VerifySignature {
    address private _signer;

    constructor(address _address) {
        _signer = _address;
    }

    function hashMessage(
        uint256 _game,
        uint256 _tower,
        address _address,
        uint256 _amount
    ) public pure returns (bytes32) {
        return
            MessageHashUtils.toEthSignedMessageHash(
                keccak256(
                    abi.encodePacked(
                        _game,
                        ":",
                        _tower,
                        ":",
                        _address,
                        ":",
                        _amount
                    )
                )
            );
    }

    function verifySignature(
        uint256 _game,
        uint256 _tower,
        uint256 _amount,
        bytes memory _signature
    ) public view returns (bool) {
        bytes32 digest = hashMessage(_game, _tower, msg.sender, _amount);
        address signer = ECDSA.recover(digest, _signature);

        return signer != address(0) && signer == _signer;
    }
}
