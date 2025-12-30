// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract BinarySearch{

    function test1(uint[] memory arr, uint target) public pure returns(int) {
        uint start = 0;
        uint end = arr.length-1;

        while (start <= end) {
            uint mid = (end-start)/2+start;
            if (arr[mid] == target){
                return int(mid);
            }
            if (arr[mid] > target){
                end = mid-1;
            } else {
                start = mid+1;
            }
        }
        return -1;
    }

    function test(uint[] memory arr, uint target) public returns(int){
        return test2(arr, target, 0, arr.length-1);
    }

    function test2(uint[] memory arr, uint target, uint start, uint end) private returns(int){
        if(start > end){
            return -1;
        }
        uint mid = (end-start)/2+start;
        if(arr[mid] == target){
            return int(mid);
        }
        if (arr[mid] > target){
            end = mid - 1;
        }else{
            start = mid + 1;
        }
        return test2(arr, target, start, end);
    }
}