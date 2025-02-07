"""
This file was generated by the KCL auto-gen tool. DO NOT EDIT.
Editing this file might prove futile when you re-run the KCL auto-gen generate command.
"""
import k8s.apimachinery.pkg.apis.meta.v1


schema SAML:
    r"""
    SAML is the Schema for the SAMLs API.

    Attributes
    ----------
    apiVersion : str, default is "provider.authentik.crossplane.io/v1alpha1", required
        APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
    kind : str, default is "SAML", required
        Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
    metadata : v1.ObjectMeta, default is Undefined, optional
        metadata
    spec : ProviderAuthentikCrossplaneIoV1alpha1SAMLSpec, default is Undefined, required
        spec
    status : ProviderAuthentikCrossplaneIoV1alpha1SAMLStatus, default is Undefined, optional
        status
    """


    apiVersion: "provider.authentik.crossplane.io/v1alpha1" = "provider.authentik.crossplane.io/v1alpha1"

    kind: "SAML" = "SAML"

    metadata?: v1.ObjectMeta

    spec: ProviderAuthentikCrossplaneIoV1alpha1SAMLSpec

    status?: ProviderAuthentikCrossplaneIoV1alpha1SAMLStatus


schema ProviderAuthentikCrossplaneIoV1alpha1SAMLSpec:
    r"""
    SAMLSpec defines the desired state of SAML

    Attributes
    ----------
    deletionPolicy : str, default is "Delete", optional
        DeletionPolicy specifies what will happen to the underlying external
        when this managed resource is deleted - either "Delete" or "Orphan" the
        external resource.
        This field is planned to be deprecated in favor of the ManagementPolicies
        field in a future release. Currently, both could be set independently and
        non-default values would be honored if the feature flag is enabled.
        See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223
    forProvider : ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecForProvider, default is Undefined, required
        for provider
    initProvider : ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecInitProvider, default is Undefined, optional
        init provider
    managementPolicies : [str], default is ["*"], optional
        THIS IS A BETA FIELD. It is on by default but can be opted out
        through a Crossplane feature flag.
        ManagementPolicies specify the array of actions Crossplane is allowed to
        take on the managed and external resources.
        This field is planned to replace the DeletionPolicy field in a future
        release. Currently, both could be set independently and non-default
        values would be honored if the feature flag is enabled. If both are
        custom, the DeletionPolicy field will be ignored.
        See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223
        and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md
    providerConfigRef : ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecProviderConfigRef, default is Undefined, optional
        provider config ref
    publishConnectionDetailsTo : ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecPublishConnectionDetailsTo, default is Undefined, optional
        publish connection details to
    writeConnectionSecretToRef : ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecWriteConnectionSecretToRef, default is Undefined, optional
        write connection secret to ref
    """


    deletionPolicy?: "Orphan" | "Delete" = "Delete"

    forProvider: ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecForProvider

    initProvider?: ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecInitProvider

    managementPolicies?: [str] = ["*"]

    providerConfigRef?: ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecProviderConfigRef

    publishConnectionDetailsTo?: ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecPublishConnectionDetailsTo

    writeConnectionSecretToRef?: ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecWriteConnectionSecretToRef


schema ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecForProvider:
    r"""
    provider authentik crossplane io v1alpha1 s a m l spec for provider

    Attributes
    ----------
    acsUrl : str, default is Undefined, optional
        (String)
    assertionValidNotBefore : str, default is Undefined, optional
        5.
        Defaults to `minutes=-5`.
    assertionValidNotOnOrAfter : str, default is Undefined, optional
        (String) Defaults to minutes=5.
        Defaults to `minutes=5`.
    audience : str, default is Undefined, optional
        (String) Defaults to “.
        Defaults to “.
    authenticationFlow : str, default is Undefined, optional
        (String)
    authorizationFlow : str, default is Undefined, optional
        (String)
    authorizationFlowRef : ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecForProviderAuthorizationFlowRef, default is Undefined, optional
        authorization flow ref
    authorizationFlowSelector : ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecForProviderAuthorizationFlowSelector, default is Undefined, optional
        authorization flow selector
    defaultRelayState : str, default is Undefined, optional
        (String) Defaults to “.
        Defaults to “.
    digestAlgorithm : str, default is Undefined, optional
        (String) Allowed values:
        Allowed values:
        - `http://www.w3.org/2000/09/xmldsig#sha1`
        - `http://www.w3.org/2001/04/xmlenc#sha256`
        - `http://www.w3.org/2001/04/xmldsig-more#sha384`
        - `http://www.w3.org/2001/04/xmlenc#sha512`
        Defaults to `http://www.w3.org/2001/04/xmlenc#sha256`.
    encryptionKp : str, default is Undefined, optional
        (String)
    invalidationFlow : str, default is Undefined, optional
        (String)
    issuer : str, default is Undefined, optional
        (String) Defaults to authentik.
        Defaults to `authentik`.
    name : str, default is Undefined, optional
        (String)
    nameIdMapping : str, default is Undefined, optional
        (String)
    propertyMappings : [str], default is Undefined, optional
        (List of String)
    sessionValidNotOnOrAfter : str, default is Undefined, optional
        (String) Defaults to minutes=86400.
        Defaults to `minutes=86400`.
    signAssertion : bool, default is Undefined, optional
        (Boolean) Defaults to true.
        Defaults to `true`.
    signResponse : bool, default is Undefined, optional
        (Boolean) Defaults to false.
        Defaults to `false`.
    signatureAlgorithm : str, default is Undefined, optional
        (String) Allowed values:
        Allowed values:
        - `http://www.w3.org/2000/09/xmldsig#rsa-sha1`
        - `http://www.w3.org/2001/04/xmldsig-more#rsa-sha256`
        - `http://www.w3.org/2001/04/xmldsig-more#rsa-sha384`
        - `http://www.w3.org/2001/04/xmldsig-more#rsa-sha512`
        - `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha1`
        - `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha256`
        - `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha384`
        - `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha512`
        - `http://www.w3.org/2000/09/xmldsig#dsa-sha1`
        Defaults to `http://www.w3.org/2001/04/xmldsig-more#rsa-sha256`.
    signingKp : str, default is Undefined, optional
        (String)
    spBinding : str, default is Undefined, optional
        (String) Allowed values:
        Allowed values:
        - `redirect`
        - `post`
        Defaults to `redirect`.
    urlSloPost : str, default is Undefined, optional
        (String) Generated.
        Generated.
    urlSloRedirect : str, default is Undefined, optional
        (String) Generated.
        Generated.
    urlSsoInit : str, default is Undefined, optional
        (String) Generated.
        Generated.
    urlSsoPost : str, default is Undefined, optional
        (String) Generated.
        Generated.
    urlSsoRedirect : str, default is Undefined, optional
        (String) Generated.
        Generated.
    verificationKp : str, default is Undefined, optional
        (String)
    """


    acsUrl?: str

    assertionValidNotBefore?: str

    assertionValidNotOnOrAfter?: str

    audience?: str

    authenticationFlow?: str

    authorizationFlow?: str

    authorizationFlowRef?: ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecForProviderAuthorizationFlowRef

    authorizationFlowSelector?: ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecForProviderAuthorizationFlowSelector

    defaultRelayState?: str

    digestAlgorithm?: str

    encryptionKp?: str

    invalidationFlow?: str

    issuer?: str

    name?: str

    nameIdMapping?: str

    propertyMappings?: [str]

    sessionValidNotOnOrAfter?: str

    signAssertion?: bool

    signResponse?: bool

    signatureAlgorithm?: str

    signingKp?: str

    spBinding?: str

    urlSloPost?: str

    urlSloRedirect?: str

    urlSsoInit?: str

    urlSsoPost?: str

    urlSsoRedirect?: str

    verificationKp?: str


schema ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecForProviderAuthorizationFlowRef:
    r"""
    Reference to a Flow in authentik to populate authorizationFlow.

    Attributes
    ----------
    name : str, default is Undefined, required
        Name of the referenced object.
    policy : ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecForProviderAuthorizationFlowRefPolicy, default is Undefined, optional
        policy
    """


    name: str

    policy?: ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecForProviderAuthorizationFlowRefPolicy


schema ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecForProviderAuthorizationFlowRefPolicy:
    r"""
    Policies for referencing.

    Attributes
    ----------
    resolution : str, default is "Required", optional
        Resolution specifies whether resolution of this reference is required.
        The default is 'Required', which means the reconcile will fail if the
        reference cannot be resolved. 'Optional' means this reference will be
        a no-op if it cannot be resolved.
    resolve : str, default is Undefined, optional
        Resolve specifies when this reference should be resolved. The default
        is 'IfNotPresent', which will attempt to resolve the reference only when
        the corresponding field is not present. Use 'Always' to resolve the
        reference on every reconcile.
    """


    resolution?: "Required" | "Optional" = "Required"

    resolve?: "Always" | "IfNotPresent"


schema ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecForProviderAuthorizationFlowSelector:
    r"""
    Selector for a Flow in authentik to populate authorizationFlow.

    Attributes
    ----------
    matchControllerRef : bool, default is Undefined, optional
        MatchControllerRef ensures an object with the same controller reference
        as the selecting object is selected.
    matchLabels : {str:str}, default is Undefined, optional
        MatchLabels ensures an object with matching labels is selected.
    policy : ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecForProviderAuthorizationFlowSelectorPolicy, default is Undefined, optional
        policy
    """


    matchControllerRef?: bool

    matchLabels?: {str:str}

    policy?: ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecForProviderAuthorizationFlowSelectorPolicy


schema ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecForProviderAuthorizationFlowSelectorPolicy:
    r"""
    Policies for selection.

    Attributes
    ----------
    resolution : str, default is "Required", optional
        Resolution specifies whether resolution of this reference is required.
        The default is 'Required', which means the reconcile will fail if the
        reference cannot be resolved. 'Optional' means this reference will be
        a no-op if it cannot be resolved.
    resolve : str, default is Undefined, optional
        Resolve specifies when this reference should be resolved. The default
        is 'IfNotPresent', which will attempt to resolve the reference only when
        the corresponding field is not present. Use 'Always' to resolve the
        reference on every reconcile.
    """


    resolution?: "Required" | "Optional" = "Required"

    resolve?: "Always" | "IfNotPresent"


schema ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecInitProvider:
    r"""
    THIS IS A BETA FIELD. It will be honored
    unless the Management Policies feature flag is disabled.
    InitProvider holds the same fields as ForProvider, with the exception
    of Identifier and other resource reference fields. The fields that are
    in InitProvider are merged into ForProvider when the resource is created.
    The same fields are also added to the terraform ignore_changes hook, to
    avoid updating them after creation. This is useful for fields that are
    required on creation, but we do not desire to update them after creation,
    for example because of an external controller is managing them, like an
    autoscaler.

    Attributes
    ----------
    acsUrl : str, default is Undefined, optional
        (String)
    assertionValidNotBefore : str, default is Undefined, optional
        5.
        Defaults to `minutes=-5`.
    assertionValidNotOnOrAfter : str, default is Undefined, optional
        (String) Defaults to minutes=5.
        Defaults to `minutes=5`.
    audience : str, default is Undefined, optional
        (String) Defaults to “.
        Defaults to “.
    authenticationFlow : str, default is Undefined, optional
        (String)
    authorizationFlow : str, default is Undefined, optional
        (String)
    authorizationFlowRef : ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecInitProviderAuthorizationFlowRef, default is Undefined, optional
        authorization flow ref
    authorizationFlowSelector : ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecInitProviderAuthorizationFlowSelector, default is Undefined, optional
        authorization flow selector
    defaultRelayState : str, default is Undefined, optional
        (String) Defaults to “.
        Defaults to “.
    digestAlgorithm : str, default is Undefined, optional
        (String) Allowed values:
        Allowed values:
        - `http://www.w3.org/2000/09/xmldsig#sha1`
        - `http://www.w3.org/2001/04/xmlenc#sha256`
        - `http://www.w3.org/2001/04/xmldsig-more#sha384`
        - `http://www.w3.org/2001/04/xmlenc#sha512`
        Defaults to `http://www.w3.org/2001/04/xmlenc#sha256`.
    encryptionKp : str, default is Undefined, optional
        (String)
    invalidationFlow : str, default is Undefined, optional
        (String)
    issuer : str, default is Undefined, optional
        (String) Defaults to authentik.
        Defaults to `authentik`.
    name : str, default is Undefined, optional
        (String)
    nameIdMapping : str, default is Undefined, optional
        (String)
    propertyMappings : [str], default is Undefined, optional
        (List of String)
    sessionValidNotOnOrAfter : str, default is Undefined, optional
        (String) Defaults to minutes=86400.
        Defaults to `minutes=86400`.
    signAssertion : bool, default is Undefined, optional
        (Boolean) Defaults to true.
        Defaults to `true`.
    signResponse : bool, default is Undefined, optional
        (Boolean) Defaults to false.
        Defaults to `false`.
    signatureAlgorithm : str, default is Undefined, optional
        (String) Allowed values:
        Allowed values:
        - `http://www.w3.org/2000/09/xmldsig#rsa-sha1`
        - `http://www.w3.org/2001/04/xmldsig-more#rsa-sha256`
        - `http://www.w3.org/2001/04/xmldsig-more#rsa-sha384`
        - `http://www.w3.org/2001/04/xmldsig-more#rsa-sha512`
        - `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha1`
        - `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha256`
        - `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha384`
        - `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha512`
        - `http://www.w3.org/2000/09/xmldsig#dsa-sha1`
        Defaults to `http://www.w3.org/2001/04/xmldsig-more#rsa-sha256`.
    signingKp : str, default is Undefined, optional
        (String)
    spBinding : str, default is Undefined, optional
        (String) Allowed values:
        Allowed values:
        - `redirect`
        - `post`
        Defaults to `redirect`.
    urlSloPost : str, default is Undefined, optional
        (String) Generated.
        Generated.
    urlSloRedirect : str, default is Undefined, optional
        (String) Generated.
        Generated.
    urlSsoInit : str, default is Undefined, optional
        (String) Generated.
        Generated.
    urlSsoPost : str, default is Undefined, optional
        (String) Generated.
        Generated.
    urlSsoRedirect : str, default is Undefined, optional
        (String) Generated.
        Generated.
    verificationKp : str, default is Undefined, optional
        (String)
    """


    acsUrl?: str

    assertionValidNotBefore?: str

    assertionValidNotOnOrAfter?: str

    audience?: str

    authenticationFlow?: str

    authorizationFlow?: str

    authorizationFlowRef?: ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecInitProviderAuthorizationFlowRef

    authorizationFlowSelector?: ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecInitProviderAuthorizationFlowSelector

    defaultRelayState?: str

    digestAlgorithm?: str

    encryptionKp?: str

    invalidationFlow?: str

    issuer?: str

    name?: str

    nameIdMapping?: str

    propertyMappings?: [str]

    sessionValidNotOnOrAfter?: str

    signAssertion?: bool

    signResponse?: bool

    signatureAlgorithm?: str

    signingKp?: str

    spBinding?: str

    urlSloPost?: str

    urlSloRedirect?: str

    urlSsoInit?: str

    urlSsoPost?: str

    urlSsoRedirect?: str

    verificationKp?: str


schema ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecInitProviderAuthorizationFlowRef:
    r"""
    Reference to a Flow in authentik to populate authorizationFlow.

    Attributes
    ----------
    name : str, default is Undefined, required
        Name of the referenced object.
    policy : ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecInitProviderAuthorizationFlowRefPolicy, default is Undefined, optional
        policy
    """


    name: str

    policy?: ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecInitProviderAuthorizationFlowRefPolicy


schema ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecInitProviderAuthorizationFlowRefPolicy:
    r"""
    Policies for referencing.

    Attributes
    ----------
    resolution : str, default is "Required", optional
        Resolution specifies whether resolution of this reference is required.
        The default is 'Required', which means the reconcile will fail if the
        reference cannot be resolved. 'Optional' means this reference will be
        a no-op if it cannot be resolved.
    resolve : str, default is Undefined, optional
        Resolve specifies when this reference should be resolved. The default
        is 'IfNotPresent', which will attempt to resolve the reference only when
        the corresponding field is not present. Use 'Always' to resolve the
        reference on every reconcile.
    """


    resolution?: "Required" | "Optional" = "Required"

    resolve?: "Always" | "IfNotPresent"


schema ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecInitProviderAuthorizationFlowSelector:
    r"""
    Selector for a Flow in authentik to populate authorizationFlow.

    Attributes
    ----------
    matchControllerRef : bool, default is Undefined, optional
        MatchControllerRef ensures an object with the same controller reference
        as the selecting object is selected.
    matchLabels : {str:str}, default is Undefined, optional
        MatchLabels ensures an object with matching labels is selected.
    policy : ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecInitProviderAuthorizationFlowSelectorPolicy, default is Undefined, optional
        policy
    """


    matchControllerRef?: bool

    matchLabels?: {str:str}

    policy?: ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecInitProviderAuthorizationFlowSelectorPolicy


schema ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecInitProviderAuthorizationFlowSelectorPolicy:
    r"""
    Policies for selection.

    Attributes
    ----------
    resolution : str, default is "Required", optional
        Resolution specifies whether resolution of this reference is required.
        The default is 'Required', which means the reconcile will fail if the
        reference cannot be resolved. 'Optional' means this reference will be
        a no-op if it cannot be resolved.
    resolve : str, default is Undefined, optional
        Resolve specifies when this reference should be resolved. The default
        is 'IfNotPresent', which will attempt to resolve the reference only when
        the corresponding field is not present. Use 'Always' to resolve the
        reference on every reconcile.
    """


    resolution?: "Required" | "Optional" = "Required"

    resolve?: "Always" | "IfNotPresent"


schema ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecProviderConfigRef:
    r"""
    ProviderConfigReference specifies how the provider that will be used to
    create, observe, update, and delete this managed resource should be
    configured.

    Attributes
    ----------
    name : str, default is Undefined, required
        Name of the referenced object.
    policy : ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecProviderConfigRefPolicy, default is Undefined, optional
        policy
    """


    name: str

    policy?: ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecProviderConfigRefPolicy


schema ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecProviderConfigRefPolicy:
    r"""
    Policies for referencing.

    Attributes
    ----------
    resolution : str, default is "Required", optional
        Resolution specifies whether resolution of this reference is required.
        The default is 'Required', which means the reconcile will fail if the
        reference cannot be resolved. 'Optional' means this reference will be
        a no-op if it cannot be resolved.
    resolve : str, default is Undefined, optional
        Resolve specifies when this reference should be resolved. The default
        is 'IfNotPresent', which will attempt to resolve the reference only when
        the corresponding field is not present. Use 'Always' to resolve the
        reference on every reconcile.
    """


    resolution?: "Required" | "Optional" = "Required"

    resolve?: "Always" | "IfNotPresent"


schema ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecPublishConnectionDetailsTo:
    r"""
    PublishConnectionDetailsTo specifies the connection secret config which
    contains a name, metadata and a reference to secret store config to
    which any connection details for this managed resource should be written.
    Connection details frequently include the endpoint, username,
    and password required to connect to the managed resource.

    Attributes
    ----------
    configRef : ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecPublishConnectionDetailsToConfigRef, default is Undefined, optional
        config ref
    metadata : ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecPublishConnectionDetailsToMetadata, default is Undefined, optional
        metadata
    name : str, default is Undefined, required
        Name is the name of the connection secret.
    """


    configRef?: ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecPublishConnectionDetailsToConfigRef

    metadata?: ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecPublishConnectionDetailsToMetadata

    name: str


schema ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecPublishConnectionDetailsToConfigRef:
    r"""
    SecretStoreConfigRef specifies which secret store config should be used
    for this ConnectionSecret.

    Attributes
    ----------
    name : str, default is Undefined, required
        Name of the referenced object.
    policy : ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecPublishConnectionDetailsToConfigRefPolicy, default is Undefined, optional
        policy
    """


    name: str

    policy?: ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecPublishConnectionDetailsToConfigRefPolicy


schema ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecPublishConnectionDetailsToConfigRefPolicy:
    r"""
    Policies for referencing.

    Attributes
    ----------
    resolution : str, default is "Required", optional
        Resolution specifies whether resolution of this reference is required.
        The default is 'Required', which means the reconcile will fail if the
        reference cannot be resolved. 'Optional' means this reference will be
        a no-op if it cannot be resolved.
    resolve : str, default is Undefined, optional
        Resolve specifies when this reference should be resolved. The default
        is 'IfNotPresent', which will attempt to resolve the reference only when
        the corresponding field is not present. Use 'Always' to resolve the
        reference on every reconcile.
    """


    resolution?: "Required" | "Optional" = "Required"

    resolve?: "Always" | "IfNotPresent"


schema ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecPublishConnectionDetailsToMetadata:
    r"""
    Metadata is the metadata for connection secret.

    Attributes
    ----------
    annotations : {str:str}, default is Undefined, optional
        Annotations are the annotations to be added to connection secret.
        - For Kubernetes secrets, this will be used as "metadata.annotations".
        - It is up to Secret Store implementation for others store types.
    labels : {str:str}, default is Undefined, optional
        Labels are the labels/tags to be added to connection secret.
        - For Kubernetes secrets, this will be used as "metadata.labels".
        - It is up to Secret Store implementation for others store types.
    $type : str, default is Undefined, optional
        Type is the SecretType for the connection secret.
        - Only valid for Kubernetes Secret Stores.
    """


    annotations?: {str:str}

    labels?: {str:str}

    $type?: str


schema ProviderAuthentikCrossplaneIoV1alpha1SAMLSpecWriteConnectionSecretToRef:
    r"""
    WriteConnectionSecretToReference specifies the namespace and name of a
    Secret to which any connection details for this managed resource should
    be written. Connection details frequently include the endpoint, username,
    and password required to connect to the managed resource.
    This field is planned to be replaced in a future release in favor of
    PublishConnectionDetailsTo. Currently, both could be set independently
    and connection details would be published to both without affecting
    each other.

    Attributes
    ----------
    name : str, default is Undefined, required
        Name of the secret.
    namespace : str, default is Undefined, required
        Namespace of the secret.
    """


    name: str

    namespace: str


schema ProviderAuthentikCrossplaneIoV1alpha1SAMLStatus:
    r"""
    SAMLStatus defines the observed state of SAML.

    Attributes
    ----------
    atProvider : ProviderAuthentikCrossplaneIoV1alpha1SAMLStatusAtProvider, default is Undefined, optional
        at provider
    conditions : [ProviderAuthentikCrossplaneIoV1alpha1SAMLStatusConditionsItems0], default is Undefined, optional
        Conditions of the resource.
    observedGeneration : int, default is Undefined, optional
        ObservedGeneration is the latest metadata.generation
        which resulted in either a ready state, or stalled due to error
        it can not recover from without human intervention.
    """


    atProvider?: ProviderAuthentikCrossplaneIoV1alpha1SAMLStatusAtProvider

    conditions?: [ProviderAuthentikCrossplaneIoV1alpha1SAMLStatusConditionsItems0]

    observedGeneration?: int


schema ProviderAuthentikCrossplaneIoV1alpha1SAMLStatusAtProvider:
    r"""
    provider authentik crossplane io v1alpha1 s a m l status at provider

    Attributes
    ----------
    acsUrl : str, default is Undefined, optional
        (String)
    assertionValidNotBefore : str, default is Undefined, optional
        5.
        Defaults to `minutes=-5`.
    assertionValidNotOnOrAfter : str, default is Undefined, optional
        (String) Defaults to minutes=5.
        Defaults to `minutes=5`.
    audience : str, default is Undefined, optional
        (String) Defaults to “.
        Defaults to “.
    authenticationFlow : str, default is Undefined, optional
        (String)
    authorizationFlow : str, default is Undefined, optional
        (String)
    defaultRelayState : str, default is Undefined, optional
        (String) Defaults to “.
        Defaults to “.
    digestAlgorithm : str, default is Undefined, optional
        (String) Allowed values:
        Allowed values:
        - `http://www.w3.org/2000/09/xmldsig#sha1`
        - `http://www.w3.org/2001/04/xmlenc#sha256`
        - `http://www.w3.org/2001/04/xmldsig-more#sha384`
        - `http://www.w3.org/2001/04/xmlenc#sha512`
        Defaults to `http://www.w3.org/2001/04/xmlenc#sha256`.
    encryptionKp : str, default is Undefined, optional
        (String)
    id : str, default is Undefined, optional
        (String) The ID of this resource.
    invalidationFlow : str, default is Undefined, optional
        (String)
    issuer : str, default is Undefined, optional
        (String) Defaults to authentik.
        Defaults to `authentik`.
    name : str, default is Undefined, optional
        (String)
    nameIdMapping : str, default is Undefined, optional
        (String)
    propertyMappings : [str], default is Undefined, optional
        (List of String)
    sessionValidNotOnOrAfter : str, default is Undefined, optional
        (String) Defaults to minutes=86400.
        Defaults to `minutes=86400`.
    signAssertion : bool, default is Undefined, optional
        (Boolean) Defaults to true.
        Defaults to `true`.
    signResponse : bool, default is Undefined, optional
        (Boolean) Defaults to false.
        Defaults to `false`.
    signatureAlgorithm : str, default is Undefined, optional
        (String) Allowed values:
        Allowed values:
        - `http://www.w3.org/2000/09/xmldsig#rsa-sha1`
        - `http://www.w3.org/2001/04/xmldsig-more#rsa-sha256`
        - `http://www.w3.org/2001/04/xmldsig-more#rsa-sha384`
        - `http://www.w3.org/2001/04/xmldsig-more#rsa-sha512`
        - `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha1`
        - `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha256`
        - `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha384`
        - `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha512`
        - `http://www.w3.org/2000/09/xmldsig#dsa-sha1`
        Defaults to `http://www.w3.org/2001/04/xmldsig-more#rsa-sha256`.
    signingKp : str, default is Undefined, optional
        (String)
    spBinding : str, default is Undefined, optional
        (String) Allowed values:
        Allowed values:
        - `redirect`
        - `post`
        Defaults to `redirect`.
    urlSloPost : str, default is Undefined, optional
        (String) Generated.
        Generated.
    urlSloRedirect : str, default is Undefined, optional
        (String) Generated.
        Generated.
    urlSsoInit : str, default is Undefined, optional
        (String) Generated.
        Generated.
    urlSsoPost : str, default is Undefined, optional
        (String) Generated.
        Generated.
    urlSsoRedirect : str, default is Undefined, optional
        (String) Generated.
        Generated.
    verificationKp : str, default is Undefined, optional
        (String)
    """


    acsUrl?: str

    assertionValidNotBefore?: str

    assertionValidNotOnOrAfter?: str

    audience?: str

    authenticationFlow?: str

    authorizationFlow?: str

    defaultRelayState?: str

    digestAlgorithm?: str

    encryptionKp?: str

    id?: str

    invalidationFlow?: str

    issuer?: str

    name?: str

    nameIdMapping?: str

    propertyMappings?: [str]

    sessionValidNotOnOrAfter?: str

    signAssertion?: bool

    signResponse?: bool

    signatureAlgorithm?: str

    signingKp?: str

    spBinding?: str

    urlSloPost?: str

    urlSloRedirect?: str

    urlSsoInit?: str

    urlSsoPost?: str

    urlSsoRedirect?: str

    verificationKp?: str


schema ProviderAuthentikCrossplaneIoV1alpha1SAMLStatusConditionsItems0:
    r"""
    A Condition that may apply to a resource.

    Attributes
    ----------
    lastTransitionTime : str, default is Undefined, required
        LastTransitionTime is the last time this condition transitioned from one
        status to another.
    message : str, default is Undefined, optional
        A Message containing details about this condition's last transition from
        one status to another, if any.
    observedGeneration : int, default is Undefined, optional
        ObservedGeneration represents the .metadata.generation that the condition was set based upon.
        For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date
        with respect to the current state of the instance.
    reason : str, default is Undefined, required
        A Reason for this condition's last transition from one status to another.
    status : str, default is Undefined, required
        Status of this condition; is it currently True, False, or Unknown?
    $type : str, default is Undefined, required
        Type of this condition. At most one of each condition type may apply to
        a resource at any point in time.
    """


    lastTransitionTime: str

    message?: str

    observedGeneration?: int

    reason: str

    status: str

    $type: str


