import * as pulumi from "@pulumi/pulumi";
import * as Authress from "@pulumi/authress";

// Create the provider wit the provider configuration.
// Or these values can also be set using `pulumi config set`:
//
// pulumi config set authress:customDomain https://login.company.com
// pulumi config set --secret authress:accessKey $AUTHRESS_KEY
//
// const authress = new Authress.Provider("authress", {
//     // Your Authress custom domain. [Configured a custom domain for Account](https://authress.io/app/#/settings?focus=domain)
//     // * Or use [provided domain](https://authress.io/app/#/api?route=overview).
//     customDomain: '',
    
//     // The access key for the Authress API. Should be [configured by your CI/CD](https://authress.io/knowledge-base/docs/category/cicd) for more information.
//     // * Or it can be overridden directly here. Do not commit this plaintext value to your source code.
//     accessKey: process.env.AUTHRESS_KEY || ''
// });

const authressRole = new Authress.Role("TestRole", {
    roleId: 'test-role',
    name: "Test Role",
    description: "An example description for this Role",
    permissions: {
        "documents:read": {
            allow: true
        }
    }
});

export let authressTestRoleId = authressRole.roleId;