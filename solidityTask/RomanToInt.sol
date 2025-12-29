// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract RomanToInt{

    function t1(string memory str) public pure returns(string memory){
        //字符串转bytes
        bytes memory b = bytes(str);
        // uint len = b.length;
        //bytes转字符串
        bytes1 s = b[1];
        return string(abi.encodePacked(s));
    }

    mapping(string => int256) public map;
    constructor() {
        map["M"]=1000;
        map["D"]=500;
        map["C"]=100;
        map["L"]=50;
        map["X"]=10;
        map["V"]=5;
        map["I"]=1;
    }

    function t2(string memory str) external view returns(int256){
        int256 num=0;
        //字符串转bytes
        bytes memory b = bytes(str);
        for (uint i = 0; i < b.length; i++) {
            //获取字符串
            string memory s1 = string(abi.encodePacked(b[i]));
            int a = map[s1];
            if(i<b.length-1 && a<map[string(abi.encodePacked(b[i+1]))]){
                num-=a;
            }else{
                num+=a;
            }
        }
        return num;
    }
}