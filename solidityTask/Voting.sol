// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract Voting {
    mapping (address => uint256) public votMap;

    function vote(address addr) public {
        votMap[addr] = votMap[addr]+uint256(1);
    }

    function getVotes(address addr) public view returns(uint256){
        return votMap[addr];
    }

    function resetVotes(address addr) public {
        delete votMap[addr];
    }
}