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



contract MyNFTAuctionV2 is
    IERC721Receiver,
    Initializable,
    UUPSUpgradeable,
    OwnableUpgradeable,
    ReentrancyGuardUpgradeable {
    /**
     *  拍卖结构体
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
        address paymentToken;    // 支付代币地址（address(0)表示使用ETH）
        uint256 highestBidInUSD;   // 以USD计价的最高出价(USD，精度8位)
    }

    //拍卖映射
    mapping (uint256 => Auction) public auctions;
    // 待退款映射（支持多代币）
    mapping(uint256 => mapping(address => mapping(address => uint256))) public pendingReturns;
    //拍卖id，自增器
    uint256 public auctionCounter;
    // 平台手续费（基点，10000 = 100%）
    uint256 public platformFee = 250; // 2.5%
    // 手续费接收地址
    address public feeRecipient;
    // 美元价格精度
    uint256 private constant USD_DECIMALS = 8;

    //拍卖创建事件
    event AuctionCreated(
        uint256 indexed auctionId,  //拍卖id
        address indexed seller,     //卖家地址
        address indexed nftContract, //NFT合约地址
        uint256 tokenId,            //Token ID    
        address paymentToken,       //支付代币地址
        uint256 amount,             //支付金额
        uint256 amountInUSD,        //以USD计价金额
        uint256 endTime             //拍卖结束时间
        );          
    // 拍卖出价事件
    event BidPlaced(
        uint256 indexed auctionId,  //拍卖id
        address indexed bidder,     //出价者地址
        address indexed nftContract,//NFT合约地址
        uint256 tokenId,            //Token ID
        address paymentToken,       //支付代币地址
        uint256 amount,             //支付金额
        uint256 amountInUSD         //以USD计价金额
    );
    //拍卖结束事件
    event AuctionEnd(
        uint256 indexed auctionId,  //拍卖id 
        address indexed winner,     //赢家地址
        address indexed nftContract,//NFT合约地址
        uint256 tokenId,            //Token ID
        address paymentToken,       //支付代币地址
        uint256 finalPrice,         //最终支付金额
        uint256 finalPriceUSD       //以USD计价最终金额
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
     * @dev 获取代币最新价格（USD，精度8位）
     * @param token 代币地址 (address(0) 表示 ETH)
     * @return price 价格（USD，精度8位）
     * @return decimals 价格源的小数位数
     */
    function getTokenPrice(address token) public view returns (uint256 price, uint8 decimals) {
        AggregatorV3Interface priceFeed = AggregatorV3Interface(token);
        require(address(priceFeed) != address(0), "Price feed not set");
        
        (, int256 answer, , , ) = priceFeed.latestRoundData();
        require(answer > 0, "Invalid price");
        
        decimals = priceFeed.decimals();
        price = uint256(answer);
    }
    
    /**
     * @dev 将代币金额转换为 USD
     * @param token 代币地址
     * @param amount 代币金额
     * @return usdAmount USD 金额（精度8位）
     */
    function convertToUSD(address token, uint256 amount) public view returns (uint256 usdAmount) {
        (uint256 price, uint8 priceDecimals) = getTokenPrice(token);
        uint8 tokenDecimals = token == address(0) ? 18 : IERC20Metadata(token).decimals();
        
        // 计算：amount * price / (10^(tokenDecimals + priceDecimals - USD_DECIMALS))
        usdAmount = (amount * price) / (10 ** (tokenDecimals + priceDecimals - USD_DECIMALS));
    }
    
    /**
     * @dev 创建拍卖
     * @param nftContract NFT合约地址
     * @param tokenId Token ID
     * @param startPrice 起拍价（USD，精度8位）
     * @param duration 拍卖时长（）
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
            active: true,
            paymentToken: address(0),  // 默认ETH支付
            highestBidInUSD: 0
        });
        //转移NFT到合约
        nft.safeTransferFrom(msg.sender, address(this), tokenId);
        
        emit AuctionCreated(
            auctionCounter, 
            msg.sender, 
            nftContract, 
            tokenId, 
            address(0),
            0, 
            0,
            auctions[auctionCounter].endTime);
        
        return auctionCounter;
    }
    //竞拍出价(支持ETH和USDC)
    function placeBid(uint256 auctionId, address payToken, uint256 amount) external payable {
        //获取拍卖信息
        Auction storage auction = auctions[auctionId];

        require(auction.active, "Auction not active");
        require(block.timestamp<auction.endTime, "Auction ended");
        require(msg.sender!=auction.seller, "Seller cannot bid");
        
        // 计算最新出价(USD)，用usd价格进行比较
        uint256 newBid;
        if (auction.highestBidInUSD == 0) {
            newBid = auction.startPrice;
        } else {
            newBid = auction.highestBidInUSD + (auction.highestBidInUSD * 5 / 100); // 5% increment
        }
        address payMentToken;
        if(payToken == address(0)){
            // ETH 出价
            payMentToken = "0x694AA1769357215DE4FAC081bf1f309aDC325306";
            // ETH 出价，当此函数成功执行后，msg.value对应的ETH会自动存入本合约余额中
            require(msg.value == amount, "ETH bid need value equal amount");
        }else{
            // ERC20 出价
            payMentToken = "0xA2F78ab2355fe2f984D808B5CeE7FD0A93D5270E";
            // 查询用户是否授权拍卖合约可以操作该ERC20代币金额大于等于此次支付金额
            uint256 allowance = IERC20(payToken).allowance(msg.sender, address(this));
            require(allowance >= amount, "ERC20 allowance not enough");
        }
        
        uint256 bidInUSD = convertToUSD(payMentToken, amount);
        require(bidInUSD >= newBid, "Bid too low");
        // 如果有之前的出价者，记录他们的待退款金额
        if (auction.highestBidder != address(0)) {
            pendingReturns[auctionId][payToken][auction.highestBidder] += auction.highestBid;
        }
        // 更新最高出价
        auction.highestBid=amount;
        auction.highestBidInUSD=bidInUSD;
        auction.highestBidder=msg.sender;

        // 竞拍出价成功，如果是ERC20出价，则需要把出价金额转到本合约
        if (payToken != address(0)) {
            bool transferSuccess = IERC20(payToken).transferFrom(msg.sender, address(this), amount);
            require(transferSuccess, "ERC20 transfer failed");
        }

        emit BidPlaced(
            auctionId,
            msg.sender,
            auction.nftContract,
            auction.tokenId,
            payToken,
            amount,
            bidInUSD
        );
    }
    /**
     * @dev 提取出价退款
     * @param auctionId 拍卖ID
     * @notice 被超越的出价者可以提取他们的资金
     */
    function withdrawBid(uint256 auctionId, address payToken) external {
         uint256 amount = pendingReturns[auctionId][payToken][msg.sender];
        require(amount > 0, "No pending return");
        
        pendingReturns[auctionId][payToken][msg.sender] = 0;
        
        if (payToken == address(0)) {
            // 退款ETH
            (bool success, ) = msg.sender.call{value: amount}("");
            require(success, "ETH transfer failed");
        } else {
            // 退款ERC20
            IERC20(payToken).transfer(msg.sender, amount);
        }
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
            if(auction.paymentToken == address(0)){
                //卖家提现ETH
                (bool successSeller, ) = auction.seller.call{value: sellerAmount}("");
                require(successSeller, "ETH Transfer to seller failed");
                
                (bool successFee, ) = feeRecipient.call{value: fee}("");
                require(successFee, "ETH  Transfer fee failed"); 

            }else{
                // ERC20支付（USDC）
                IERC20 token = IERC20(auction.paymentToken);
                require(token.transfer(auction.seller, sellerAmount), "Transfer to seller failed");
                require(token.transfer(feeRecipient, fee), "Transfer fee failed");
            }
        }else{
            // 如果无人出价，则将NFT转回给卖家
            nft.safeTransferFrom(address(this), auction.seller, auction.tokenId);
        }
        emit AuctionEnd(
                auctionId,
                auction.highestBidder,
                auction.nftContract,
                auction.tokenId,
                auction.paymentToken,
                auction.highestBid,
                auction.highestBidInUSD
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