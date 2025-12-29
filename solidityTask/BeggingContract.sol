// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract BeggingContract{
    //所有者
    address public immutable owner;
    //期限，结束时间
    uint256 public immutable deadline;
    //收到的总金额
    uint256 public totalAmount;
    //捐款者地址数组
    address[] public addressArry;
    
    //支付用户的金额
    mapping (address => uint256) public payMap;
    //事件：收到捐款（包含捐款者、金额、最新总额）
    event ContributionReceived(address indexed form, uint256 amount, uint256 totalAmount);
    //事件：发起人提取金额
    event FundsWithdrawn(address indexed recipient, uint256 amount);
    //构造函数，初始化合约账号和期限
    constructor(uint256 _durtion){
        owner = msg.sender;
        deadline = block.timestamp + (_durtion * 1 days);
    }
    //捐款
    function donate() public payable returns(bool){
        //捐款金额大于0
        require(msg.value > 0, "payment amount is greater than zero");
        require(block.timestamp <= deadline , "The event has ended");
        //允许一个地址多次贡献
        payMap[msg.sender] += msg.value;
        totalAmount += msg.value;
        //记录捐款者地址
        addressArry.push(msg.sender);
        emit ContributionReceived(msg.sender, msg.value, totalAmount);
        return true;
    }
    //提取所有资金
    function withdraw() public {
        require(totalAmount > 0, "The total amount is greater than zero");
        require(msg.sender == owner, "Only the contract owner can make the call");
        uint256 amount = address(this).balance;
        // 在转账前清零余额，防止重入攻击
        totalAmount = 0;  
        (bool success, ) = payable (owner).call{value: amount}("");
        require(success, "Failed to withdraw funds");

        emit FundsWithdrawn(owner, amount);
    }
    //查询地址的捐款金额
    function getDonation(address addr) public view returns(uint256){
        return payMap[addr];
    }
    //查询当前合约的余额
    function getContractBalance() public view returns(uint256){ 
        return address(this).balance;
    }
    //获取发起人账户余额
    function getOwnerBalance() public view returns(uint256){ 
        return owner.balance;
    }

    struct AddressInfo {
        address addr;
        uint256 amount;
    }
    

    function getTopThree() public view returns(AddressInfo[3] memory){
        AddressInfo[3] memory topArry;
        uint256 len = addressArry.length;
        for (uint i=0;i<len;i++){
            address addr = addressArry[i];
            uint256 amount = payMap[addr];
            if(addr == topArry[0].addr || 
                addr == topArry[1].addr || 
                addr == topArry[2].addr){
                continue;
            }
            if(amount > topArry[0].amount){
                AddressInfo memory t1 = topArry[0];
                AddressInfo memory t2 = topArry[1];
                topArry[0] = AddressInfo({
                    addr: addr,
                    amount: amount
                });
                topArry[1] = t1;
                topArry[2] = t2;
            }else if(amount > topArry[1].amount){
                AddressInfo memory t1 = topArry[1];
                topArry[1] = AddressInfo({
                    addr: addr,
                    amount: amount
                });
                topArry[2] = t1;
            }else if(amount > topArry[2].amount){
                topArry[2] = AddressInfo({
                    addr: addr,
                    amount: amount
                });
            }
        }
        return topArry;
    }

}