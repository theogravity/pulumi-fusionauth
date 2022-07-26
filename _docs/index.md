---
title: FusionAuth
meta_desc: Provides configuration management for FusionAuth
layout: overview
---

FusionAuth for Pulumi can be used to configure FusionAuth instances.

## Example

{{< chooser language "typescript" >}}

{{% choosable language typescript %}}

```typescript
import { Provider } from 'pulumi-fusionauth';

const fusionAuthProvider = new Provider('fusion-auth', {
  host: process.env.FUSION_AUTH_HOST_URL,
  apiKey: process.env.FUSION_AUTH_API_KEY,
});

// Create a new signing key
const clientDataSigningKey = new FusionAuthKey(
  'sample-jwt-key',
  {
    algorithm: 'RS256',
    name: 'Sample jwt key',
    length: 2048,
  },
  { provider: fusionAuthProvider },
);

export const appClientDataJwtKeyId = clientDataSigningKey.id;
```

{{< /chooser >}}
