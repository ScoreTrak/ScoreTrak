import { FrontendApi, Configuration } from "@ory/client"


const basePath = import.meta.env.VITE_API_ORY_URL || "http://localhost:4000"
export const ory = new FrontendApi(
  new Configuration({
    basePath,
    baseOptions: {
      withCredentials: true,
    }
  })
)
