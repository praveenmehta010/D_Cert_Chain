// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Certificate {
    mapping(string => bytes32) public certificateHashes;

    event CertificateIssued(string certificateId, bytes32 hash);

    function storeCertificateHash(string memory certificateId, bytes32 hash) public {
        certificateHashes[certificateId] = hash;
        emit CertificateIssued(certificateId, hash);
    }

    function getCertificateHash(string memory certificateId) public view returns (bytes32) {
        return certificateHashes[certificateId];
    }
}
