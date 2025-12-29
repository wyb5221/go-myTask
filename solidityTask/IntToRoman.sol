// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract IntToRoman{

    struct Tmp{
        uint key;
        string value;
    }
    Tmp[] public tmpArray;
    string public str;
    function t1(uint num) public returns(string memory){
        tmpArray.push(Tmp(1000, "M"));
        tmpArray.push(Tmp(900, "CM"));
        tmpArray.push(Tmp(500, "D"));
        tmpArray.push(Tmp(400, "CD"));
        tmpArray.push(Tmp(100, "C"));
        tmpArray.push(Tmp(90, "XC"));
        tmpArray.push(Tmp(50, "L"));
        tmpArray.push(Tmp(40, "XL"));
        tmpArray.push(Tmp(10, "X"));
        tmpArray.push(Tmp(9, "IX"));
        tmpArray.push(Tmp(5, "V"));
        tmpArray.push(Tmp(4, "IV"));
        tmpArray.push(Tmp(1, "I"));

        bytes memory b ;
        for (uint i=0;i<tmpArray.length;i++){
            Tmp memory t = tmpArray[i];
            uint key = t.key;
            while (num >= key){
                b = abi.encodePacked(b, t.value);
                num -=key;
            }
        }
        str = string(b);
        return string(b);
    }
}