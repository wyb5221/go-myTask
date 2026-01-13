// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/IERC721Metadata.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";



contract MyNFTAuction is
    IERC721Receiver,
    Initializable,
    UUPSUpgradeable,
    OwnableUpgradeable,
    ReentrancyGuardUpgradeable {
    /**
     * @dev 拍卖结构体
     */
    struct Auction {
        address seller;           // 卖家地址
        address nftContract;      // NFT合约地址
        uint256 tokenId;          // Token ID
        uint256 startPrice;       // 起拍价
        uint256 highestBid;       // 当前最高出价
        address highestBidder;    // 当前最高出价者
        uint256 endTime;          // 拍卖结束时间
        bool active;              // 是否激活
    }
    //拍卖映射
    mapping (uint256 => Auction) public auctions;
    // 待退款映射
    mapping(uint256 => mapping(address => uint256)) public pendingReturns;
    //拍卖id，自增器
    uint256 public auctionCounter;
    // 平台手续费（基点，10000 = 100%）
    uint256 public platformFee = 250; // 2.5%
    // 手续费接收地址
    address public feeRecipient;
    //拍卖创建事件
    event AuctionCreated(
        uint256 indexed auctionId,
        address indexed seller, 
        address nftContract, 
        uint256 tokenId, 
        uint256 startPrice,
        uint256 endTime);
    // 拍卖出价事件
    event BidPlaced(
        uint256 indexed auctionId,
        address indexed bidder,
        address indexed nftContract,
        uint256 tokenId,
        uint256 amount
    );
    //拍卖结束事件
    event AuctionEnd(
        uint256 indexed auctionId,
        address indexed winner,
        address indexed nftContract,
        uint256 tokenId,
        uint256 finalPrice
    );
    
   /**
     * @dev 初始化函数（替代构造函数）
     * @param _feeRecipient 手续费接收地址
     */
    function initialize(address _feeRecipient) public initializer {
        __Ownable_init(msg.sender);
        __UUPSUpgradeable_init();
        __ReentrancyGuard_init();
        
        require(_feeRecipient != address(0), "Invalid fee recipient");
        feeRecipient = _feeRecipient;
        auctionCounter = 0;
    }
    /**
     * @dev 创建拍卖
     * @param nftContract NFT合约地址
     * @param tokenId Token ID
     * @param startPrice 起拍价（wei）
     * @param duration 拍卖时长（天）
     * @return auctionId 拍卖ID
     */
    function createAuction(
        address nftContract,
        uint256 tokenId,
        uint256 startPrice,
        uint256 duration
    ) external returns(uint256){
        require(nftContract != address(0), "invalid nft");
        require(startPrice > 0, "Start price must be greater than 0");

        // 获取NFT合约实例
        IERC721 nft = IERC721(nftContract);
        require(nft.ownerOf(tokenId) == msg.sender, "not be NFT owner");
        // 验证授权
        require(
            nft.getApproved(tokenId) == address(this) ||
            nft.isApprovedForAll(msg.sender, address(this)),
            "Marketplace not approved"
        );
        auctionCounter++;
        auctions[auctionCounter] = Auction({
            seller: msg.sender,
            nftContract: nftContract,
            tokenId: tokenId,
            startPrice: startPrice,
            highestBid: 0,
            highestBidder: address(0),
            endTime: block.timestamp + (duration * 1 days),
            active: true
        });
        emit AuctionCreated(
            auctionCounter, 
            msg.sender, 
            nftContract, 
            tokenId, 
            startPrice, 
            auctions[auctionCounter].endTime);
            return auctionCounter;
    }
    //竞拍出价
    function placeBid(uint256 auctionId) external payable {
        //获取拍卖信息
        Auction storage auction = auctions[auctionId];
        require(auction.active, "Auction not active");
        require(block.timestamp<auction.endTime, "Auction ended");
        require(msg.sender!=auction.seller, "Seller cannot bid");
        // 计算最新出价
        uint256 newBid;
        if (auction.highestBid == 0) {
            newBid = auction.startPrice;
        } else {
            newBid = auction.highestBid + (auction.highestBid * 5 / 100); // 5% increment
        }
        require(msg.value >= newBid, "Bid too low");
        // 如果有之前的出价者，记录他们的待退款金额
        if (auction.highestBidder != address(0)) {
            pendingReturns[auctionId][auction.highestBidder] += auction.highestBid;
        }
        // 更新最高出价
        auction.highestBid=msg.value;
        auction.highestBidder=msg.sender;
        emit BidPlaced(
            auctionId,
            msg.sender,
            auction.nftContract,
            auction.tokenId,
            msg.value
        );
    }
    /**
     * @dev 提取出价退款
     * @param auctionId 拍卖ID
     * @notice 被超越的出价者可以提取他们的资金
     */
    function withdrawBid(uint256 auctionId) external {
         uint256 amount = pendingReturns[auctionId][msg.sender];
        require(amount > 0, "No pending return");
        
        pendingReturns[auctionId][msg.sender] = 0;
        
        (bool success, ) =  msg.sender.call{value: amount}("");
        require(success, "Transfer failed");
    }
    /**
     * @dev 结束拍卖
     * @param auctionId 拍卖ID
     * @notice 任何人都可以在拍卖结束后调用此函数进行结算
     */
    function endAuction(uint256 auctionId) external nonReentrant {
        Auction storage auction = auctions[auctionId];
        
        require(auction.active, "Auction not active");
        // require(block.timestamp >= auction.endTime, "Auction not ended");
        auction.active=false;
        // 获取NFT合约实例
        IERC721 nft = IERC721(auction.nftContract);
        if (auction.highestBidder != address(0)) {
            // 有人出价，获取手续费，进行结算
            uint256 fee = (auction.highestBid * platformFee) / 10000;
            uint256 sellerAmount = auction.highestBid - fee;
            // 转移NFT
            nft.safeTransferFrom(
                auction.seller,
                auction.highestBidder,
                auction.tokenId
            );
            //卖家提现
            (bool successSeller, ) = auction.seller.call{value: sellerAmount}("");
            require(successSeller, "Transfer to seller failed");
            
            (bool successFee, ) = feeRecipient.call{value: fee}("");
            require(successFee, "Transfer fee failed");
        }else{
            // 如果无人出价，则将NFT转回给卖家
            nft.safeTransferFrom(address(this), auction.seller, auction.tokenId);
        }
        emit AuctionEnd(
                auctionId,
                auction.highestBidder,
                auction.nftContract,
                auction.tokenId,
                auction.highestBid
            );
    }

    function onERC721Received(
        address operator,
        address from,
        uint256 tokenId,
        bytes calldata data
    ) external pure override returns (bytes4) {
        return this.onERC721Received.selector;
    }

    function _authorizeUpgrade(
        address newImplementation
    ) internal override onlyOwner {}
    
}