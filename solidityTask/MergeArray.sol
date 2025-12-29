// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract MergeArray {

    function merge1(uint[] memory arr1, uint[] memory arr2) public pure returns(uint[] memory){
        uint[] memory arr = new uint[](arr1.length+arr2.length);
        for(uint i=0;i<arr1.length;i++){
            arr[i]=arr1[i];
        }
        for(uint i=0;i<arr2.length;i++){
            arr[arr1.length+i] = arr2[i];
        }
        
        return arr;
    }

    function merge(uint[] memory arr1, uint[] memory arr2) public pure returns(uint[] memory){
        uint[] memory arr = new uint[](arr1.length+arr2.length);
        
        uint i = 0;
        uint j = 0;
        uint k = 0;
        while (i< arr1.length && j<arr2.length){
            if(arr1[i]<arr2[j]){
                arr[k]=arr1[i];
                i++;
            }else{
                arr[k]=arr2[j];
                j++;
            }
            k++;
        }
        while(i<arr1.length){
            arr[k]=arr1[i];
            i++;
            k++;
        }
        while(j<arr2.length){
            arr[k]=arr2[j];
            j++;
            k++;
        }
        return arr;
        // arr1 = arr;
        // return arr1;
    }

}