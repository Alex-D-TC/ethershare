pragma solidity ^0.4.24;

import "./Token.sol";

contract TokenImpl is Token {

    address public owner;

    string public name;
    string public symbol;
    uint public maxSupply;
    uint public supply;

    mapping(address => uint) public balances;
    mapping(address => mapping(address => uint)) public allowances; 

    event Transfer(address indexed from, address indexed to, uint256 value);
    event Allowed(address indexed from, address indexed to, uint value);
    event RevokedAllowance(address indexed from, address indexed to, uint value);

    constructor(string memory _name, string memory _symbol, uint _maxSupply) public {
        name = _name;
        symbol = _symbol;
        maxSupply = _maxSupply;
        owner = msg.sender;
    }

    modifier isOwner(address addr) {
        require(addr == owner);
        _;
    }

    modifier canPay(address addr, uint toPay) {
        require(balances[addr] >= toPay, "Cannot send more coins than I own");
        _;
    }

    function totalSupply() public view returns (uint256) {
        return supply;
    }

    function balanceOf(address who) public view returns (uint256) {
        return balances[who];
    }

    function allowance(address from, address to) public view returns (uint256) {
        return allowances[from][to];
    }

    function sendTo(address to, uint256 value) public canPay(msg.sender, value) {
        balances[msg.sender] -= value;
        balances[to] += value;

        emit Transfer(msg.sender, to, value);
    }

    function mint(address to, uint256 value) public isOwner(msg.sender) {
        require(value + supply <= maxSupply, "Cannot mint more than the maximum allowed");
        balances[to] += value;
        supply += value;
    }

    function changeOwner(address newOwner) public isOwner(msg.sender) {
        owner = newOwner;
    }

    function allow(address dest, uint value) public canPay(msg.sender, value) {
        balances[msg.sender] -= value;
        allowances[msg.sender][dest] += value;

        emit Allowed(msg.sender, dest, value);
    }

    function renounceAllowance(address to, uint value) public {
        require(allowances[msg.sender][to] <= value, "Cannot renounce more than what I allowed to the wallet");

        balances[msg.sender] += value;
        allowances[msg.sender][to] -= value;

        emit RevokedAllowance(msg.sender, to, value);
    }

    function claimAllowance(address from, uint value) public {
        require(allowances[from][msg.sender] <= value, "Cannot claim more than what I am permitted");

        balances[msg.sender] += value;
        allowances[from][msg.sender] -= value;

        emit Transfer(from, msg.sender, value);
    }
}
