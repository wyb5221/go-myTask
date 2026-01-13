// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";


contract MyNFT is ERC721, ERC721URIStorage, Ownable{
    // Token ID计数器 自增
    uint256 public _tokenIdCounter;
    // 最大供应量
    uint256 public constant MAX_SOUPPLY = 10000;
    // 铸造价格 
    uint256 public minPrice = 0.01 ether;
    /**
     * NFT铸造事件
     * minter 铸造者地址
     * tokenId 新创建的Token ID 
     * uri 元数据URI
     */
    event NFTMinted(address indexed minter, uint256 indexed tokenId, string uri);
    /**
     * 构造函数
     * 初始化NFT集合名称和符号，设置合约所有者
     */
    constructor() ERC721("MyNFT", "MNFT") Ownable(msg.sender){}

    function mint(string memory uri) public payable returns(uint256){
        //检查供应量限制，防止超量
        require(_tokenIdCounter < MAX_SOUPPLY, "ax supply reached");
        //检查支付金额，创造nft需要支付费用
        require(msg.value >= minPrice, "Insufficient payment");
        //创建tokenId, 递增计数器
        _tokenIdCounter++;
        uint256 newTokenId = _tokenIdCounter;
        //安全铸造NFT
        _safeMint(msg.sender, newTokenId);
        //设置元数据URI
        _setTokenURI(newTokenId, uri);
        emit NFTMinted(msg.sender, newTokenId, uri);
        return newTokenId;
    }
    /**
     * 重写tokenURI函数
     * tokenId Token ID
     * @return 元数据URI
     * @notice 需要重写以解决多重继承的冲突
     */
    function tokenURI(uint256 tokenId) public view override(ERC721, ERC721URIStorage) returns(string memory) {
        return super.tokenURI(tokenId);
    }
    /**
     * @dev 检查接口支持
     * @param interfaceId 接口ID
     * @return 是否支持该接口
     * @notice 实现ERC165标准，支持接口查询
     */
    function supportsInterface(bytes4 interfaceId) public view override (ERC721, ERC721URIStorage) returns (bool){
        return super.supportsInterface(interfaceId);
    }
    /**
     * @dev 查询总供应量
     * @return 已铸造的NFT数量
     */
    function totalSupply() public view returns(uint256){
        return _tokenIdCounter;
    }
    /**
     * @dev 提取铸造费用
     * @notice 只有合约所有者可以调用
     */
    function withdraw() public onlyOwner {
        //获取合约余额
        uint256 balance = address(this).balance;
        require(balance > 0, "No balance to withdraw");
        //payable(owner()).transfer(balance);
        (bool success, ) = payable(owner()).call{value: balance}("");
        require(success, "Failed to withdraw funds");
    }
    function setMintPrice(uint256 newPrice) public onlyOwner {
        minPrice = newPrice;
    }

}