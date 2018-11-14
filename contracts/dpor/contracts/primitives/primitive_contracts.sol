pragma solidity ^0.4.24;

library PrimitiveContractsInterface {
    function getRank(address addr, uint256 blockNum) internal view returns (uint256 b) {
        assembly {
            let p := mload(0x40)
            mstore(p, addr)
            mstore(add(p, 0x14), blockNum)
            if iszero(staticcall(not(0), 0x64, p, 0x34, p, 0x20)) {
                revert(0, 0)
            }
            b := mload(p)
        }
    }

    function getMaintenance(address addr, uint256 blockNum) internal view returns (uint256 b) {
        assembly {
            let p := mload(0x40)
            mstore(p, addr)
            mstore(add(p, 0x14), blockNum)
            if iszero(staticcall(not(0), 0x65, p, 0x34, p, 0x20)) {
                revert(0, 0)
            }
            b := mload(p)
        }
    }

    function getProxyCount(address addr, uint256 blockNum) internal view returns (uint256 b) {
        assembly {
            let p := mload(0x40)
            mstore(p, addr)
            mstore(add(p, 0x14), blockNum)
            if iszero(staticcall(not(0), 0x66, p, 0x34, p, 0x20)) {
                revert(0, 0)
            }
            b := mload(p)
        }
    }

    function getUploadInfo(address addr, uint256 blockNum) internal view returns (uint256 b) {
        assembly {
            let p := mload(0x40)
            mstore(p, addr)
            mstore(add(p, 0x14), blockNum)
            if iszero(staticcall(not(0), 0x67, p, 0x34, p, 0x20)) {
                revert(0, 0)
            }
            b := mload(p)
        }
    }

    function getTxVolume(address addr, uint256 blockNum) internal view returns (uint256 b) {
        assembly {
            let p := mload(0x40)
            mstore(p, addr)
            mstore(add(p, 0x14), blockNum)
            if iszero(staticcall(not(0), 0x68, p, 0x34, p, 0x20)) {
                revert(0, 0)
            }
            b := mload(p)
        }
    }

    function isProxy(address addr, uint256 blockNum) internal view returns (uint256 b) {
        assembly {
            let p := mload(0x40)
            mstore(p, addr)
            mstore(add(p, 0x14), blockNum)
            if iszero(staticcall(not(0), 0x69, p, 0x34, p, 0x20)) {
                revert(0, 0)
            }
            b := mload(p)
        }
    }
}