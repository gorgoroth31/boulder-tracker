/// <reference types="vite/client" />

interface ViteTypeOptions {
    // By adding this line, you can make the type of ImportMetaEnv strict
    // to disallow unknown keys.
    // strictImportMetaEnv: unknown
}

interface ImportMetaEnv {
    readonly VITE_SERVER_URL: string
    readonly VITE_LOGTO_APPID: string
    readonly VITE_LOGTO_ENDPOINT: string
    readonly VITE_APP_URI: string
}

interface ImportMeta {
    readonly env: ImportMetaEnv
}