/// <reference types="vite/client" />

interface ViteTypeOptions {
    // By adding this line, you can make the type of ImportMetaEnv strict
    // to disallow unknown keys.
    // strictImportMetaEnv: unknown
}

interface ImportMetaEnv {
    readonly VITE_AUTH_DOMAIN: string
    readonly VITE_SERVER_URL: string
    readonly VITE_AUDIENCE: string
    readonly VITE_AUTH_CLIENT_ID: string
    readonly VITE_SCOPES: string
    // more env variables...
}

interface ImportMeta {
    readonly env: ImportMetaEnv
}