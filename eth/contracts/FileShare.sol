pragma solidity ^0.4.24;

import "./Token.sol";

contract FileShareContract {

    struct Share {
        string fileHash;
        uint price;
    }

    mapping(address => Share[]) public shares;

    Token public token;

    constructor(Token _token) public {
        token = _token;
    }

    function publishFile(string memory _fileHash, uint _price) public {
        
        Share memory share = Share({
            fileHash: _fileHash,
            price: _price
        });
        
        shares[msg.sender].push(share);
    }

    function acceptPossible(address _recipient, uint _shareID) public view returns (bool) {
        address sharer = msg.sender;

        require(_shareID < shares[sharer].length);
        
        Share storage share = shares[sharer][_shareID];

        return token.allowance(_recipient, sharer) >= share.price;
    }

    function acceptRequest(address _recipient, uint _shareID) public {

        require(acceptPossible(_recipient, _shareID));

        Share storage share = shares[msg.sender][_shareID];

        // Send the funds through the contract to the sharer
        token.claimAllowance(_recipient, share.price);
        token.sendTo(msg.sender, share.price);    
    }

}
