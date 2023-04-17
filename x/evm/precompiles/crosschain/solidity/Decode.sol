// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

library Decode {
    function crossChain(bytes memory data) internal pure returns (bool) {
        bool result = abi.decode(data, (bool));
        return result;
    }

    function cancelSendToExternal(
        bytes memory data
    ) internal pure returns (bool) {
        bool result = abi.decode(data, (bool));
        return result;
    }

    function increaseBridgeFee(bytes memory data) internal pure returns (bool) {
        bool result = abi.decode(data, (bool));
        return result;
    }

    function bridgeCoin(
        bytes memory data
    ) internal pure returns (uint256) {
        uint256 amount = abi.decode(data, (uint256));
        return amount;
    }

    function ok(
        bool _result,
        bytes memory _data,
        string memory _msg
    ) internal pure {
        if (!_result) {
            string memory errMsg = abi.decode(_data, (string));
            if (bytes(_msg).length < 1) {
                revert(errMsg);
            }
            revert(string(abi.encodePacked(_msg, ": ", errMsg)));
        }
    }
}
