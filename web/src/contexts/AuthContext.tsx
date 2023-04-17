import { Identity, Session } from "@ory/client";
import { createContext, useContext, useEffect, useState } from "react";
import { ory } from "../lib/auth/ory";
import { PropsWithChildren } from "react"


type AuthContextType = {
  isLoading: boolean,
  session: Session | undefined,
  logoutUrl: string | undefined,
  refreshSession: () => void,
}

const AuthContext = createContext<AuthContextType>({isLoading: true, session: undefined, logoutUrl: "", refreshSession: () => {}})

export const useAuth = () => {
  return useContext(AuthContext)
}

export default function AuthProvider({ children }: PropsWithChildren) {
  const [isLoading, setIsLoading] = useState<boolean>(true)
  const [session, setSession] = useState<Session | undefined>()
  const [logoutUrl, setLogoutUrl] = useState<string | undefined>()

  const getSessionAndLogoutUrl = () => {
    setIsLoading(true)
    ory
      .toSession()
      .then(({ data: sessionData }) => {
        setSession(sessionData)
        ory.createBrowserLogoutFlow()
          .then(({ data: logoutFlowData }) => {
            setLogoutUrl(logoutFlowData.logout_url)
          })
          .catch((err) => {
          })
      })
      .catch((err) => {
      })
      .finally(() => setIsLoading(false))
  }

  const refresh = () => {
    getSessionAndLogoutUrl()
  }

  useEffect(() => {
    refresh()
  }, [])

  return (
    <>
      <AuthContext.Provider value={{session, logoutUrl, refreshSession: refresh}}>
        {children}
      </AuthContext.Provider>
    </>
  )
}
