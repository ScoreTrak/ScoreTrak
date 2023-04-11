/// <reference types="vite/client" />


interface ImportMetaEnv {
  readonly VITE_API_SERVER_URL: string
  readonly VITE_API_ORY_URL: string
  // more env variables...
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}
