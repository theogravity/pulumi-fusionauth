// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace theogravity.Fusionauth.Inputs
{

    public sealed class FusionAuthApplicationSamlv2ConfigurationAssertionEncryptionConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The message digest algorithm to use when encrypting the symmetric key for transport. The possible values are: SHA1 - SHA-1 hashing algorithm, SHA256 - SHA-256 hashing algorithm, SHA384 - SHA-384 hashing algorithm or SHA512 - SHA-512 hashing algorithm. Using SHA256 or higher is recommended.
        /// </summary>
        [Input("digestAlgorithm")]
        public Input<string>? DigestAlgorithm { get; set; }

        /// <summary>
        /// Determines if SAML assertion encryption is enabled for this Application.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The symmetric key encryption algorithm that will be used to encrypt SAML assertions. A new symmetric key will be generated every time an assertion is encrypted. AES ciphers can operate in Cipher Block Chaining (CBC) or Galois/Counter Mode (GCM). The possible values are: AES128, AES192, AES256, AES128GCM, AES192GCM, AES256GCM or TripleDES.
        /// </summary>
        [Input("encryptionAlgorithm")]
        public Input<string>? EncryptionAlgorithm { get; set; }

        /// <summary>
        /// The location that the encrypted symmetric key information will be placed in the SAML response in relation to the EncryptedData element containing the encrypted assertion value. The possible values are: Child (The EncryptedKey element will be wrapped in a KeyInfo element and added inside the EncryptedData) or Sibling (The EncryptedKey element will be added to the document as a sibling of EncryptedData).
        /// </summary>
        [Input("keyLocation")]
        public Input<string>? KeyLocation { get; set; }

        /// <summary>
        /// The encryption algorithm used to encrypt the symmetric key for transport in the SAML response. The possible values are: RSAv15, RSA_OAEP or RSA_OAEP_MGF1P.
        /// </summary>
        [Input("keyTransportAlgorithm")]
        public Input<string>? KeyTransportAlgorithm { get; set; }

        /// <summary>
        /// The unique Id of the Key used to encrypt the symmetric key for transport in the SAML response. The selected Key must contain an RSA certificate. This parameter is required when application.samlv2Configuration.assertionEncryptionConfiguration.enabled is set to true.
        /// </summary>
        [Input("keyTransportEncryptionKeyId")]
        public Input<string>? KeyTransportEncryptionKeyId { get; set; }

        /// <summary>
        /// The mask generation function and hash function to use for the Optimal Asymmetric Encryption Padding when encrypting a symmetric key for transport. The possible values are: MGF1_SHA1, MGF1_SHA224, MGF1_SHA256, MGF1_SHA384 or MGF1_SHA512. This value is only used when the `application.samlv2Configuration.assertionEncryptionConfiguration.keyTransportAlgorithm` is set to RSA_OAEP. RSAv15 does not require a message digest function, and RSA_OAEP_MGF1P will always use MGF1_SHA1 regardless of this value.
        /// </summary>
        [Input("maskGenerationFunction")]
        public Input<string>? MaskGenerationFunction { get; set; }

        public FusionAuthApplicationSamlv2ConfigurationAssertionEncryptionConfigurationGetArgs()
        {
        }
        public static new FusionAuthApplicationSamlv2ConfigurationAssertionEncryptionConfigurationGetArgs Empty => new FusionAuthApplicationSamlv2ConfigurationAssertionEncryptionConfigurationGetArgs();
    }
}
