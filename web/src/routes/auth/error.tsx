import {useSearchParams} from "react-router-dom";


export function AuthErrorPage() {
  const [searchParams, setSearchParams] = useSearchParams();

  return (
    <>
      <p>Auth Error Page</p>
      <p>{searchParams.get("id")}</p>
    </>
  )
}
