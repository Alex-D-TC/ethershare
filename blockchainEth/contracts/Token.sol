pragma solidity ^0.4.24;

interface Token {

    function totalSupply() external view returns (uint);

    function balanceOf(address who) external view returns (uint);
    function allowance(address from, address to) external view returns (uint);

    function sendTo(address to, uint256 value) external;
    function mint(address to, uint256 value) external;

    function changeOwner(address newOwner) external;

    function allow(address dest, uint value) external;
    function renounceAllowance(address to, uint value) external;
    function claimAllowance(address from, uint value) external;
}
