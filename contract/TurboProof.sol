// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/*
 * @dev Provides information about the current execution context, including the
 * sender of the transaction and its data. While these are generally available
 * via msg.sender and msg.data, they should not be accessed in such a direct
 * manner, since when dealing with meta-transactions the account sending and
 * paying for execution may not be the actual sender (as far as an application
 * is concerned).
 *
 * This contract is only required for intermediate, library-like contracts.
 */
abstract contract Context {
    function _msgSender() internal view virtual returns (address) {
        return msg.sender;
    }

    function _msgData() internal view virtual returns (bytes calldata) {
        this; // silence state mutability warning without generating bytecode - see https://github.com/ethereum/solidity/issues/2691
        return msg.data;
    }
}

/**
 * @dev Contract module which provides a basic access control mechanism, where
 * there is an account (an owner) that can be granted exclusive access to
 * specific functions.
 *
 * By default, the owner account will be the one that deploys the contract. This
 * can later be changed with {transferOwnership}.
 *
 * This module is used through inheritance. It will make available the modifier
 * `onlyOwner`, which can be applied to your functions to restrict their use to
 * the owner.
 */
abstract contract Ownable is Context {
    address private _owner;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    /**
     * @dev Initializes the contract setting the deployer as the initial owner.
     */
    constructor () {
        address msgSender = _msgSender();
        _owner = msgSender;
        emit OwnershipTransferred(address(0), msgSender);
    }

    /**
     * @dev Returns the address of the current owner.
     */
    function owner() public view virtual returns (address) {
        return _owner;
    }

    /**
     * @dev Throws if called by any account other than the owner.
     */
    modifier onlyOwner() {
        require(owner() == _msgSender(), "Ownable: caller is not the owner");
        _;
    }

    /**
     * @dev Leaves the contract without owner. It will not be possible to call
     * `onlyOwner` functions anymore. Can only be called by the current owner.
     *
     * NOTE: Renouncing ownership will leave the contract without an owner,
     * thereby removing any functionality that is only available to the owner.
     */
    function renounceOwnership() public virtual onlyOwner {
        emit OwnershipTransferred(_owner, address(0));
        _owner = address(0);
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     * Can only be called by the current owner.
     */
    function transferOwnership(address newOwner) public virtual onlyOwner {
        require(newOwner != address(0), "Ownable: new owner is the zero address");
        emit OwnershipTransferred(_owner, newOwner);
        _owner = newOwner;
    }
}

abstract contract IFactRegistry {
    function isValid(bytes32 fact) external view virtual returns(bool);
}

contract TurboProof is Ownable {
    mapping(uint256 => uint256) public    registriesHash;
    uint256                     public    registriesProgramHash;
    uint256                     public    identitiesProgramHash;
    IFactRegistry               public    CAIRO_VERIFIER;
    
    event UpdateRegistry(uint256, uint256, uint256);
    
    constructor(uint256 _registriesProgramHash, uint256 _identitiesProgramHash, address cairoVerifier) {
        identitiesProgramHash = _identitiesProgramHash;
        registriesProgramHash = _registriesProgramHash;
        CAIRO_VERIFIER = IFactRegistry(cairoVerifier);
    }

    function updateRegistry(uint256 registryKey, uint256 oldRegistryHash, uint256 newRegistryHash) onlyOwner public returns (bytes32) {
        bytes32 outputHash = keccak256(abi.encodePacked(oldRegistryHash, newRegistryHash));
        bytes32 fact = keccak256(abi.encodePacked(registriesProgramHash, outputHash));
        require(CAIRO_VERIFIER.isValid(fact), "INVALID_PROOF");
        require(registriesHash[registryKey] == oldRegistryHash);

        registriesHash[registryKey] = newRegistryHash;
        emit UpdateRegistry(registryKey, oldRegistryHash, newRegistryHash);
        return (fact);
    }

    function proveIdentity(uint256 registryKey, uint256 hash, uint256 registryHash) public view returns (bool) {
        bytes32 outputHash = keccak256(abi.encodePacked(hash, registryHash));
        bytes32 fact = keccak256(abi.encodePacked(identitiesProgramHash, outputHash));
        return registriesHash[registryKey] == registryHash && CAIRO_VERIFIER.isValid(fact);
    }
}