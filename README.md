# terraform-pgp-provider
A Terraform pgp provider brought together and edited from two repos on Github.

This provider will provide two resources so far, pgp_encrypt_message and pgp_decrypt_message.

pgp_encrypt_message takes a plaintext message and a public key and outputs an encrypted message.

pgp_decrypt_message takes an encrypted message and a private key and outputs a plaintext message.
