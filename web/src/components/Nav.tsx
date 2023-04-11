import { Link } from "@mui/material";
import { useAuth } from "../contexts/AuthContext";


export default function Nav() {
  const { session, logoutUrl } = useAuth()
  return (
    <>
      {
        !session ?
          <>
            <Link href="/auth/login">Login</Link>
            <Link href="/auth/register">Register</Link>
          </>
          :
          <>
            <Link href={logoutUrl}>Logout</Link>
          </>

      }
      <Link href="/competitions">Competitions</Link>
      <Link href="/dashboard">Dashboard</Link>
    </>
  )
}
