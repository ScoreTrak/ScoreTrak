import { useAuth } from "../contexts/AuthContext";
import {Link} from "@tanstack/react-router";


export default function Nav() {
  const { session, logoutUrl } = useAuth()
  return (
    <>
      <Link to={"/"}>HOME</Link>
      {
        !session ?
          <>
            <Link to={"/auth/login"} >Login</Link>
            <Link to={"/auth/register"} >Register</Link>
          </>
          :
          <>
            <Link to={logoutUrl} >Logout</Link>
          </>

      }
      <Link to={"/competitions"} >Competitions</Link>
      <Link to={"/dashboard"} >Dashboard</Link>
    </>
  )
}
