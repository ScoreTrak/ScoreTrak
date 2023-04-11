import { Identity, Session } from "@ory/client";
import { createContext, useContext, useEffect, useState } from "react";
import { ory } from "../lib/auth/ory";


type AuthContextType = {
  session: Session | undefined,
  logoutUrl: string | undefined,
  login: () => void,
}

const AuthContext = createContext<AuthContextType>({session: undefined, login:() => {}, logoutUrl: ""})

export const useAuth = () => {
  return useContext(AuthContext)
}

export default function AuthProvider({ children }) {
  const [session, setSession] = useState<Session | undefined>()
  const [logoutUrl, setLogoutUrl] = useState<string | undefined>()

  useEffect(() => {
    ory
      .toSession()
      .then(({ data: sessionData }) => {
        setSession(sessionData)
        ory.createBrowserLogoutFlow()
          .then(({ data: logoutFlowData }) => {
            setLogoutUrl(logoutFlowData.logout_url)
          })
      })
      .catch((err) => {
        console.error(err);
      })
  }, [])

  return (
    <>
      <AuthContext.Provider value={{session, logoutUrl, login: () => {}}}>
        {children}
      </AuthContext.Provider>
    </>
  )
}
