// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

import "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";

contract GetHash {
    function getHash000(uint256 _game) public pure returns (bytes32) {
        return
            MessageHashUtils.toEthSignedMessageHash(
                keccak256(abi.encodePacked(_game))
            );
    }

    function getHash001(address _address) public pure returns (bytes32) {
        return
            MessageHashUtils.toEthSignedMessageHash(
                keccak256(abi.encodePacked(_address))
            );
    }

    function getHash002(
        uint256 _game,
        uint256 _tower
    ) public pure returns (bytes32) {
        return
            MessageHashUtils.toEthSignedMessageHash(
                keccak256(abi.encodePacked(_game, _tower))
            );
    }

    function getHash003(
        uint256 _game,
        uint256 _tower,
        address _address
    ) public pure returns (bytes32) {
        return
            MessageHashUtils.toEthSignedMessageHash(
                keccak256(abi.encodePacked(_game, _tower, _address))
            );
    }

    function getHash004(
        uint256 _game,
        uint256 _tower,
        address _address,
        uint256 _amount
    ) public pure returns (bytes32) {
        return
            MessageHashUtils.toEthSignedMessageHash(
                keccak256(abi.encodePacked(_game, _tower, _address, _amount))
            );
    }

    function getHash005(
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
}
