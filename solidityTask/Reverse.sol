// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract Reverse {

    function tes1(string memory str) public pure returns(bytes memory){
        return bytes(str);
    }
    function tes2(string memory str) public pure returns(uint){
        return bytes(str).length;
    }

    function ReverseString(string memory str) public pure returns(string memory){
        bytes memory a = bytes(str);
        bytes memory b = new bytes(a.length);
        for (uint i = 0; i < a.length; i++) {
            b[i] = a[a.length - 1 - i];
        }
        return string(b);
    }
}