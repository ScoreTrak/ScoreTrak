import { useAuth } from "../contexts/AuthContext";
import {Link, NavLink} from "react-router-dom";


export default function Nav() {
  const { user, logoutUrl } = useAuth()

  return (
    <>
      <header>
        <nav aria-label={"main-nav"} className={"container flex flex-row p-2"}>
          <div className={""}>
            <NavLink to={"/"} className={"text-3xl font-bold tracking-wide"}>ScoreTrak</NavLink>
          </div>
          <div className={"grow"}></div>
          <div className={""}>
            { !user &&
              <>
                <NavLink to={"/auth/login"} className={"secondary"} >Login</NavLink>
                <NavLink to={"/auth/register"} className={"secondary"} >Register</NavLink>
              </>
            }
            { user && logoutUrl &&
              <>
                <NavLink to={"/me"}>{user.traits?.username}</NavLink>
                <Link to={logoutUrl} className={"secondary"} >Logout</Link>
              </>
            }
          </div>
        </nav>
      </header>
    </>
  )
}
