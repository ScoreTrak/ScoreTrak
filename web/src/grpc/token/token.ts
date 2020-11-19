import jwt_decode from "jwt-decode"

export enum Role {
  Red = "red",
  Black = "black",
  Blue = "blue"
}

export type JWTToken = {
  exp: number,
  jti: string,
  username: string,
  team_id: string,
  role: Role
}


const saveToken = (token : string) => {
  localStorage.setItem("token", JSON.stringify(token));
}

const logout = () => {
  localStorage.removeItem("token");
};

const getToken = ():string => {
  return JSON.parse(localStorage.getItem("token") as string);
};

const tokenExists = ():boolean =>{
  return !!localStorage.getItem("token")
}

const getDecodedJWT = (): JWTToken | null => {
  const item = localStorage.getItem("token")
  if (item !== null){
    return jwt_decode(item) as JWTToken
  }
  return null
}

const getCurrentRole = ():string | undefined => {
  if (isAValidToken()){
    return getDecodedJWT()?.role
  } else {
    return undefined
  }
};

const getCurrentTeamID = ():string | undefined => {
  return getDecodedJWT()?.team_id
};

const tokenExpired = ():boolean => {
  let current_time = new Date().getTime() / 1000;
  const exp = getDecodedJWT()?.exp
  if (exp){
    return current_time > exp
  }
  return false
}

const isAValidToken = () => {
  return (tokenExists() && !tokenExpired())
}
export const token = {
  saveToken,
  logout,
  getToken,
  getCurrentRole,
  getDecodedJWT,
  isAValidToken,
  getCurrentTeamID
};
